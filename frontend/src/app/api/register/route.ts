import axios from "axios";
import { NextRequest, NextResponse } from "next/server";

type SuccessResponse = { ok: true; message: string };
type ErrorResponse = { ok: false; message: string };
export type RegisterResponse = SuccessResponse | ErrorResponse;

const http = axios.create({ proxy: false, timeout: 15000 });

export async function POST(req: NextRequest) {
  try {
    const body = await req.json();

    if (!body.name || !body.email || !body.password) {
      return NextResponse.json(
        { ok: false, message: "Name, email, and password are required" },
        { status: 400 }
      );
    }

    const backendUrl =
      process.env.INTERNAL_API_URL ?? process.env.NEXT_PUBLIC_API_URL!;

    try {
      await http.post(`${backendUrl}/auth/register`, body);
    } catch (error) {
      if (axios.isAxiosError(error)) {
        const status = error.response?.status ?? 500;
        const message = error.response?.data?.error ?? "Registration failed";
        return NextResponse.json({ ok: false, message }, { status });
      }
      return NextResponse.json(
        { ok: false, message: "Unexpected error" },
        { status: 500 }
      );
    }

    return NextResponse.json(
      { ok: true, message: "Registration successful" },
      { status: 201 }
    );
  } catch {
    return NextResponse.json(
      { ok: false, message: "Internal server error" },
      { status: 500 }
    );
  }
}
