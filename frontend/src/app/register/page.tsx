"use client";

import Image from "next/image";
import Link from "next/link";
import { useState, useEffect } from "react";
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

type PrenamOption = { id: string; th: string; en: string };
type OrgOption = { code: string; name_th: string; name_en: string };
type AccountTypeOption = { id: string; th: string; en: string };

const PRENAMES: PrenamOption[] = [
  { id: "MR",  th: "นาย",     en: "Mr."  },
  { id: "MRS", th: "นาง",     en: "Mrs." },
  { id: "MS",  th: "นางสาว",  en: "Miss" },
  { id: "OTH", th: "อื่นๆ",   en: "Other" },
];

const ORGANIZATIONS: OrgOption[] = [
  { code: "01",  name_th: "คณะมนุศย์ศาสตร์",                          name_en: "Faculty of Humanities" },
  { code: "02",  name_th: "คณะศึกษาศาสตร์",                            name_en: "Faculty of Education" },
  { code: "03",  name_th: "คณะวิจิตรศิลป์",                            name_en: "Faculty of Fine Art" },
  { code: "04",  name_th: "คณะสังคมศาสตร์",                            name_en: "Faculty of Social Sciences" },
  { code: "05",  name_th: "คณะวิทยาศาสตร์",                            name_en: "Faculty of Science" },
  { code: "06",  name_th: "คณะวิศวกรรมศาสตร์",                         name_en: "Faculty of Engineering" },
  { code: "07",  name_th: "คณะแพทยศาสตร์",                             name_en: "Faculty of Medicine" },
  { code: "08",  name_th: "คณะเกษตรศาสตร์",                            name_en: "Faculty of Agriculture" },
  { code: "09",  name_th: "คณะทันตแพทยศาสตร์",                         name_en: "Faculty of Dentistry" },
  { code: "10",  name_th: "คณะเภสัชศาสตร์",                            name_en: "Faculty of Pharmacy" },
  { code: "11",  name_th: "คณะเทคนิคการแพทย์",                         name_en: "Faculty of Medical Technology" },
  { code: "12",  name_th: "คณะพยาบาลศาสตร์",                           name_en: "Faculty of Nursing" },
  { code: "13",  name_th: "คณะอุตสาหกรรมเกษตร",                        name_en: "Faculty of Agro-Industry" },
  { code: "14",  name_th: "คณะสัตวแพทยศาสตร์",                         name_en: "Faculty of Veterinary Medicine" },
  { code: "15",  name_th: "คณะบริหารธุรกิจ",                           name_en: "Faculty of Business Administration" },
  { code: "16",  name_th: "คณะเศรษฐศาสตร์",                            name_en: "Faculty of Economics" },
  { code: "17",  name_th: "คณะสถาปัตยกรรมศาสตร์",                      name_en: "Faculty of Architecture" },
  { code: "18",  name_th: "คณะสื่อสารมวลชน",                           name_en: "Faculty of Mass Communication" },
  { code: "19",  name_th: "คณะรัฐศาสตร์และรัฐประศาสนศาสตร์",           name_en: "Faculty of Political Science and Public Administration" },
  { code: "20",  name_th: "คณะนิติศาสตร์",                             name_en: "Faculty of Law" },
  { code: "21",  name_th: "วิทยาลัยศิลปะ สื่อ และเทคโนโลยี",          name_en: "Faculty of Art, Media and Technology" },
  { code: "22",  name_th: "คณะสาธารณสุขศาสตร์",                        name_en: "Faculty of Public Health" },
  { code: "23",  name_th: "วิทยาลัยการศึกษาและการจัดการทะเล",          name_en: "Faculty of Education and Sea Management" },
  { code: "24",  name_th: "วิทยาลัยนานาชาตินวัตกรรมดิจิทัล",          name_en: "Faculty of International Digital Innovation" },
  { code: "25",  name_th: "สถาบันนโยบายสาธารณะ",                       name_en: "Faculty of Public Policy" },
  { code: "26",  name_th: "สถาบันวิศวกรรมชีวการแพทย์",                 name_en: "Faculty of Biomedical Engineering" },
  { code: "27",  name_th: "สถาบันวิจัยวิทยาศาสตร์สุขภาพ",              name_en: "Faculty of Health Science Research Institute" },
  { code: "28",  name_th: "วิทยาลัยพหุวิทยาการและสหวิทยาการ",          name_en: "Faculty of Interdisciplinary and Multidisciplinary Studies" },
  { code: "54",  name_th: "สำนักบริการวิชาการ",                        name_en: "Academic Service Center" },
  { code: "59",  name_th: "สำนักงานพัฒนาคุณภาพการศึกษา",               name_en: "Office of Education Quality Development" },
  { code: "61",  name_th: "สำนักงานมหาวิทยาลัย",                       name_en: "Office of the University" },
  { code: "64",  name_th: "ศูนย์วิจัยข้าวล้านนา",                      name_en: "Lanna Rice Research Center" },
  { code: "65",  name_th: "สถาบันภาษา",                                name_en: "Language Institute" },
  { code: "76",  name_th: "วิทยาลัยการศึกษาตลอดชีวิต",                 name_en: "School of Life Long Education" },
  { code: "107", name_th: "ศูนย์บริหารจัดการความปลอดภัยฯ SHE",        name_en: "Center of Safety, Occupational Health and Environment" },
  { code: "108", name_th: "ศูนย์นวัตกรรมการสอนและการเรียนรู้ TLIC",    name_en: "Teaching and Learning Innovation Center" },
];

const ACCOUNT_TYPES: AccountTypeOption[] = [
  { id: "MISEmpAcc", th: "บุคลากร",            en: "MIS Employee" },
  { id: "StdAcc",    th: "นักศึกษาปัจจุบัน",   en: "Student Account" },
  { id: "AlumAcc",   th: "ศิษย์เก่า",           en: "Alumni Account" },
];

function pickRandom<T>(arr: T[]): T {
  return arr[Math.floor(Math.random() * arr.length)];
}

function buildInitialForm(): FormData {
  const org = pickRandom(ORGANIZATIONS);
  const acc = pickRandom(ACCOUNT_TYPES);
  return {
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
    organization_code:    org.code,
    organization_name_th: org.name_th,
    organization_name_en: org.name_en,
    itaccounttype_id:  acc.id,
    itaccounttype_th:  acc.th,
    itaccounttype_en:  acc.en,
  };
}

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
      <label htmlFor={id} className="block text-xs font-medium text-gray-600 mb-1">
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

function ReadonlyField({ label, value }: { label: string; value: string }) {
  return (
    <div>
      <label className="block text-xs font-medium text-gray-600 mb-1">{label}</label>
      <div className="w-full px-3 py-2 border border-gray-200 rounded-lg text-sm bg-gray-50 text-gray-500 select-none">
        {value || "—"}
      </div>
    </div>
  );
}

export default function RegisterPage() {
  const router = useRouter();
  const [form, setForm] = useState<FormData>(() => buildInitialForm());
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const [mounted, setMounted] = useState(false);

  useEffect(() => { setMounted(true); }, []);

  const setField = (id: keyof FormData, val: string) =>
    setForm((prev) => ({ ...prev, [id]: val }));

  const handlePrename = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selected = PRENAMES.find((p) => p.id === e.target.value);
    if (!selected) {
      setForm((prev) => ({ ...prev, prename_id: "", prename_th: "", prename_en: "" }));
      return;
    }
    setForm((prev) => ({
      ...prev,
      prename_id: selected.id,
      prename_th: selected.th,
      prename_en: selected.en,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    if (!form.prename_id) {
      setError("กรุณาเลือกคำนำหน้าชื่อ");
      return;
    }
    if (form.password !== form.confirmPassword) {
      setError("รหัสผ่านไม่ตรงกัน");
      return;
    }

    setLoading(true);
    try {
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      const { confirmPassword, ...payload } = form;
      const res = await axios.post<RegisterResponse>(`${base}/api/register`, payload);
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

  if (!mounted) return null;

  return (
    <div className="min-h-screen bg-[#dcc0f2] px-4 py-10 flex items-start justify-center">
      <div className="bg-white rounded-3xl shadow-xl px-8 py-10 w-full max-w-2xl">
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
              <Field label="ชื่อที่แสดง (Name)" id="name" required placeholder="ชื่อ-นามสกุล หรือชื่อเล่น" value={form.name} onChange={setField} />
              <Field label="Email" id="email" type="email" required placeholder="your@email.com" value={form.email} onChange={setField} />
              <Field label="Password" id="password" type="password" required placeholder="อย่างน้อย 8 ตัวอักษร" value={form.password} onChange={setField} />
              <Field label="ยืนยัน Password" id="confirmPassword" type="password" required placeholder="••••••••" value={form.confirmPassword} onChange={setField} />
            </div>
          </div>

          {/* ── คำนำหน้า ── */}
          <div>
            <SectionHeader title="คำนำหน้าชื่อ" />
            <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
              <div className="sm:col-span-1">
                <label htmlFor="prename_select" className="block text-xs font-medium text-gray-600 mb-1">
                  คำนำหน้า <span className="text-red-500">*</span>
                </label>
                <select
                  id="prename_select"
                  value={form.prename_id}
                  onChange={handlePrename}
                  className="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-purple-400 bg-white"
                >
                  <option value="">-- เลือกคำนำหน้า --</option>
                  {PRENAMES.map((p) => (
                    <option key={p.id} value={p.id}>
                      {p.th} ({p.en})
                    </option>
                  ))}
                </select>
              </div>
              <ReadonlyField label="คำนำหน้า (ไทย)" value={form.prename_th} />
              <ReadonlyField label="Prename (EN)" value={form.prename_en} />
            </div>
          </div>

          {/* ── ชื่อ-นามสกุล ── */}
          <div>
            <SectionHeader title="ชื่อ - นามสกุล" />
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <Field label="ชื่อ (ไทย)" id="firstname_th" placeholder="ชื่อภาษาไทย" value={form.firstname_th} onChange={setField} />
              <Field label="First Name (EN)" id="firstname_en" placeholder="First name" value={form.firstname_en} onChange={setField} />
              <Field label="นามสกุล (ไทย)" id="lastname_th" placeholder="นามสกุลภาษาไทย" value={form.lastname_th} onChange={setField} />
              <Field label="Last Name (EN)" id="lastname_en" placeholder="Last name" value={form.lastname_en} onChange={setField} />
            </div>
          </div>

          {/* ── องค์กร (สุ่มอัตโนมัติ) ── */}
          <div>
            <SectionHeader title="ข้อมูลองค์กร" />
            <p className="text-xs text-gray-400 mb-3">กำหนดอัตโนมัติตามระบบ</p>
            <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
              <ReadonlyField label="Organization Code" value={form.organization_code} />
              <ReadonlyField label="ชื่อองค์กร (ไทย)" value={form.organization_name_th} />
              <ReadonlyField label="Organization Name (EN)" value={form.organization_name_en} />
            </div>
          </div>

          {/* ── ประเภทบัญชี (สุ่มอัตโนมัติ) ── */}
          <div>
            <SectionHeader title="ประเภทบัญชี (IT Account Type)" />
            <p className="text-xs text-gray-400 mb-3">กำหนดอัตโนมัติตามระบบ</p>
            <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
              <ReadonlyField label="Account Type ID" value={form.itaccounttype_id} />
              <ReadonlyField label="ประเภทบัญชี (ไทย)" value={form.itaccounttype_th} />
              <ReadonlyField label="Account Type (EN)" value={form.itaccounttype_en} />
            </div>
          </div>

          {error && <p className="text-red-500 text-sm text-center">{error}</p>}

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
          <Link href="/signin" className="text-purple-700 font-medium hover:underline">
            เข้าสู่ระบบ
          </Link>
        </p>
      </div>
    </div>
  );
}
