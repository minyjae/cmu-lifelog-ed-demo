import axios from "axios";
import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

type SuccessResponse = { ok: true };
type ErrorResponse = { ok: false; message: string };
export type SignInResponse = SuccessResponse | ErrorResponse;

const http = axios.create({ proxy: false, timeout: 15000 });

export async function POST(req: NextRequest) {
  try {
    const { email, password } = await req.json();

    if (!email || !password) {
      return NextResponse.json(
        { ok: false, message: "Email and password are required" },
        { status: 400 }
      );
    }

    const backendUrl =
      process.env.INTERNAL_API_URL ?? process.env.NEXT_PUBLIC_API_URL!;

    let backendToken: string | null = null;
    try {
      const res = await http.post(`${backendUrl}/auth`, { email, password });
      backendToken = res.data.token;
    } catch (error) {
      if (axios.isAxiosError(error)) {
        const status = error.response?.status ?? 500;
        const message =
          status === 401
            ? "Invalid email or password"
            : "Authentication failed";
        return NextResponse.json({ ok: false, message }, { status });
      }
      return NextResponse.json(
        { ok: false, message: "Unexpected error" },
        { status: 500 }
      );
    }

    if (!backendToken) {
      return NextResponse.json(
        { ok: false, message: "No token received from backend" },
        { status: 500 }
      );
    }

    const cookieStore = await cookies();
    cookieStore.set({
      name: "backend-api-token",
      value: backendToken,
      expires: new Date(Date.now() + 72 * 3600 * 1000),
      httpOnly: false,
      sameSite: "lax",
      secure: process.env.NODE_ENV === "production",
      path: "/",
    });

    return NextResponse.json({ ok: true });
  } catch (error) {
    console.error("Unexpected error in sign-in handler:", error);
    return NextResponse.json(
      { ok: false, message: "Internal server error" },
      { status: 500 }
    );
  }
}
