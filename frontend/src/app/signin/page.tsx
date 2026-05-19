"use client";

import Image from "next/image";
import Link from "next/link";
import { useState } from "react";
import { useRouter } from "next/navigation";
import axios, { AxiosError } from "axios";

type SignInResponse = { ok: true } | { ok: false; message: string };

const base = process.env.NEXT_PUBLIC_BASE_PATH || "";

export default function SignInPage() {
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setLoading(true);
    try {
      const res = await axios.post<SignInResponse>(`${base}/api/signIn`, {
        email,
        password,
      });
      if (res.data.ok) {
        router.push("/main");
      }
    } catch (err: unknown) {
      const axiosErr = err as AxiosError<SignInResponse>;
      const msg =
        axiosErr.response?.data && !axiosErr.response.data.ok
          ? axiosErr.response.data.message
          : "เกิดข้อผิดพลาด กรุณาลองใหม่อีกครั้ง";
      setError(msg);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-[#dcc0f2] px-4">
      <div className="bg-white rounded-3xl shadow-xl px-10 py-12 w-full max-w-md space-y-8">
        {/* Logo */}
        <div className="flex justify-center">
          <Image
            src={`${base}/logo_le.png`}
            alt="Logo"
            width={200}
            height={200}
            priority
          />
        </div>

        {/* Form */}
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Email
            </label>
            <input
              type="email"
              required
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              placeholder="your@email.com"
              className="w-full px-4 py-2.5 border border-gray-300 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-purple-400"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Password
            </label>
            <input
              type="password"
              required
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="••••••••"
              className="w-full px-4 py-2.5 border border-gray-300 rounded-xl text-sm focus:outline-none focus:ring-2 focus:ring-purple-400"
            />
          </div>

          {error && (
            <p className="text-red-500 text-sm text-center">{error}</p>
          )}

          <button
            type="submit"
            disabled={loading}
            className="w-full bg-gradient-to-r from-purple-800 to-purple-400 text-white font-semibold py-3 rounded-full shadow-md text-base disabled:opacity-60 disabled:cursor-not-allowed transition"
          >
            {loading ? "กำลังเข้าสู่ระบบ..." : "เข้าสู่ระบบ"}
          </button>
        </form>

        <p className="text-center text-sm text-gray-500">
          ยังไม่มีบัญชี?{" "}
          <Link
            href="/register"
            className="text-purple-700 font-medium hover:underline"
          >
            ลงทะเบียน
          </Link>
        </p>
      </div>
    </div>
  );
}
