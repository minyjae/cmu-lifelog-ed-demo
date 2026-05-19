import { cookies } from "next/headers";
import { NextResponse } from "next/server";
import { User } from "@/types/api/user";

type SuccessResponse = { ok: true; cmuBasicInfo: User[] };
type ErrorResponse = { ok: false; message: string };
export type WhoAmIResponse = SuccessResponse | ErrorResponse;

export async function GET(): Promise<NextResponse<WhoAmIResponse>> {
  const cookieStore = await cookies();
  const token = cookieStore.get("backend-api-token")?.value;

  if (!token) {
    return NextResponse.json(
      { ok: false, message: "Not authenticated" },
      { status: 401 }
    );
  }

  const base =
    process.env.INTERNAL_API_URL || process.env.NEXT_PUBLIC_API_URL;

  try {
    const res = await fetch(`${base}/user/me`, {
      headers: { Authorization: `Bearer ${token}` },
      cache: "no-store",
    });

    if (!res.ok) {
      return NextResponse.json(
        { ok: false, message: "User not found" },
        { status: 401 }
      );
    }

    const user = (await res.json()) as User;
    return NextResponse.json({ ok: true, cmuBasicInfo: [user] });
  } catch {
    return NextResponse.json(
      { ok: false, message: "Failed to fetch user" },
      { status: 500 }
    );
  }
}
