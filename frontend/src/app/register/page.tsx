"use client";

import Image from "next/image";
import Link from "next/link";
import { useState } from "react";
import { useRouter } from "next/navigation";
import axios, { AxiosError } from "axios";

type RegisterResponse = { ok: true; message: string } | { ok: false; message: string };

const base = process.env.NEXT_PUBLIC_BASE_PATH || "";

type FormData = {
  name: string;
  email: string;
  password: string;
  confirmPassword: string;
  prename_id: string;
  prename_th: string;
  prename_en: string;
  firstname_th: string;
  firstname_en: string;
  lastname_th: string;
  lastname_en: string;
  organization_code: string;
  organization_name_th: string;
  organization_name_en: string;
  itaccounttype_id: string;
  itaccounttype_th: string;
  itaccounttype_en: string;
};

const initialForm: FormData = {
  name: "",
  email: "",
  password: "",
  confirmPassword: "",
  prename_id: "",
  prename_th: "",
  prename_en: "",
  firstname_th: "",
  firstname_en: "",
  lastname_th: "",
  lastname_en: "",
  organization_code: "",
  organization_name_th: "",
  organization_name_en: "",
  itaccounttype_id: "",
  itaccounttype_th: "",
  itaccounttype_en: "",
};

function SectionHeader({ title }: { title: string }) {
  return (
    <h3 className="text-sm font-semibold text-purple-700 uppercase tracking-wide border-b border-purple-100 pb-1 mb-3">
      {title}
    </h3>
  );
}

function Field({
  label,
  id,
  type = "text",
  required = false,
  placeholder,
  value,
  onChange,
}: {
  label: string;
  id: keyof FormData;
  type?: string;
  required?: boolean;
  placeholder?: string;
  value: string;
  onChange: (id: keyof FormData, val: string) => void;
}) {
  return (
    <div>
      <label
        htmlFor={id}
        className="block text-xs font-medium text-gray-600 mb-1"
      >
        {label}
        {required && <span className="text-red-500 ml-0.5">*</span>}
      </label>
      <input
        id={id}
        type={type}
        required={required}
        value={value}
        placeholder={placeholder}
        onChange={(e) => onChange(id, e.target.value)}
        className="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-purple-400"
      />
    </div>
  );
}

export default function RegisterPage() {
  const router = useRouter();
  const [form, setForm] = useState<FormData>(initialForm);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const setField = (id: keyof FormData, val: string) =>
    setForm((prev) => ({ ...prev, [id]: val }));

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    if (form.password !== form.confirmPassword) {
      setError("รหัสผ่านไม่ตรงกัน");
      return;
    }

    setLoading(true);
    try {
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      const { confirmPassword, ...payload } = form;
      const res = await axios.post<RegisterResponse>(
        `${base}/api/register`,
        payload
      );
      if (res.data.ok) {
        router.push("/signin");
      }
    } catch (err: unknown) {
      const axiosErr = err as AxiosError<RegisterResponse>;
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
    <div className="min-h-screen bg-[#dcc0f2] px-4 py-10 flex items-start justify-center">
      <div className="bg-white rounded-3xl shadow-xl px-8 py-10 w-full max-w-2xl">
        {/* Logo */}
        <div className="flex justify-center mb-6">
          <Image
            src={`${base}/logo_le.png`}
            alt="Logo"
            width={140}
            height={140}
            priority
          />
        </div>

        <h2 className="text-xl font-bold text-center text-purple-800 mb-8">
          ลงทะเบียนผู้ใช้งาน
        </h2>

        <form onSubmit={handleSubmit} className="space-y-6">
          {/* ── ข้อมูลบัญชี ── */}
          <div>
            <SectionHeader title="ข้อมูลบัญชี" />
            <div className="grid grid-cols-1 gap-4">
              <Field
                label="ชื่อที่แสดง (Name)"
                id="name"
                required
                placeholder="ชื่อ-นามสกุล หรือชื่อเล่น"
                value={form.name}
                onChange={setField}
              />
              <Field
                label="Email"
                id="email"
                type="email"
                required
                placeholder="your@email.com"
                value={form.email}
                onChange={setField}
              />
              <Field
                label="Password"
                id="password"
                type="password"
                required
                placeholder="อย่างน้อย 8 ตัวอักษร"
                value={form.password}
                onChange={setField}
              />
              <Field
                label="ยืนยัน Password"
                id="confirmPassword"
                type="password"
                required
                placeholder="••••••••"
                value={form.confirmPassword}
                onChange={setField}
              />
            </div>
          </div>

          {/* ── คำนำหน้า ── */}
          <div>
            <SectionHeader title="คำนำหน้าชื่อ" />
            <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
              <Field
                label="Prename ID"
                id="prename_id"
                placeholder="เช่น 1, 2, 3"
                value={form.prename_id}
                onChange={setField}
              />
              <Field
                label="คำนำหน้า (ไทย)"
                id="prename_th"
                placeholder="นาย / นาง / นางสาว"
                value={form.prename_th}
                onChange={setField}
              />
              <Field
                label="คำนำหน้า (EN)"
                id="prename_en"
                placeholder="Mr. / Mrs. / Miss"
                value={form.prename_en}
                onChange={setField}
              />
            </div>
          </div>

          {/* ── ชื่อ-นามสกุล ── */}
          <div>
            <SectionHeader title="ชื่อ - นามสกุล" />
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <Field
                label="ชื่อ (ไทย)"
                id="firstname_th"
                placeholder="ชื่อภาษาไทย"
                value={form.firstname_th}
                onChange={setField}
              />
              <Field
                label="First Name (EN)"
                id="firstname_en"
                placeholder="First name"
                value={form.firstname_en}
                onChange={setField}
              />
              <Field
                label="นามสกุล (ไทย)"
                id="lastname_th"
                placeholder="นามสกุลภาษาไทย"
                value={form.lastname_th}
                onChange={setField}
              />
              <Field
                label="Last Name (EN)"
                id="lastname_en"
                placeholder="Last name"
                value={form.lastname_en}
                onChange={setField}
              />
            </div>
          </div>

          {/* ── องค์กร ── */}
          <div>
            <SectionHeader title="ข้อมูลองค์กร" />
            <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
              <Field
                label="Organization Code"
                id="organization_code"
                placeholder="รหัสองค์กร"
                value={form.organization_code}
                onChange={setField}
              />
              <Field
                label="ชื่อองค์กร (ไทย)"
                id="organization_name_th"
                placeholder="ชื่อองค์กรภาษาไทย"
                value={form.organization_name_th}
                onChange={setField}
              />
              <Field
                label="Organization Name (EN)"
                id="organization_name_en"
                placeholder="Organization name"
                value={form.organization_name_en}
                onChange={setField}
              />
            </div>
          </div>

          {/* ── ประเภทบัญชี ── */}
          <div>
            <SectionHeader title="ประเภทบัญชี (IT Account Type)" />
            <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
              <Field
                label="Account Type ID"
                id="itaccounttype_id"
                placeholder="เช่น MISEmpAcc"
                value={form.itaccounttype_id}
                onChange={setField}
              />
              <Field
                label="ประเภทบัญชี (ไทย)"
                id="itaccounttype_th"
                placeholder="ประเภทบัญชีภาษาไทย"
                value={form.itaccounttype_th}
                onChange={setField}
              />
              <Field
                label="Account Type (EN)"
                id="itaccounttype_en"
                placeholder="Account type"
                value={form.itaccounttype_en}
                onChange={setField}
              />
            </div>
          </div>

          {error && (
            <p className="text-red-500 text-sm text-center">{error}</p>
          )}

          <button
            type="submit"
            disabled={loading}
            className="w-full bg-gradient-to-r from-purple-800 to-purple-400 text-white font-semibold py-3 rounded-full shadow-md text-base disabled:opacity-60 disabled:cursor-not-allowed transition"
          >
            {loading ? "กำลังลงทะเบียน..." : "ลงทะเบียน"}
          </button>
        </form>

        <p className="text-center text-sm text-gray-500 mt-6">
          มีบัญชีแล้ว?{" "}
          <Link
            href="/signin"
            className="text-purple-700 font-medium hover:underline"
          >
            เข้าสู่ระบบ
          </Link>
        </p>
      </div>
    </div>
  );
}
