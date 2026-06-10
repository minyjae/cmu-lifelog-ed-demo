import { cookies } from "next/headers";
import { NextResponse } from "next/server";

type SuccessResponse = { ok: true };
type ErrorResponse = { ok: false };
export type SignOutResponse = SuccessResponse | ErrorResponse;

export async function POST(): Promise<NextResponse<SignOutResponse>> {
  (await cookies()).delete("backend-api-token");
  return NextResponse.json({ ok: true });
}
