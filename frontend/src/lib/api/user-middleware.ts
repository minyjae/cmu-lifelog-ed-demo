//// ============================================================
////                         Middleware
//// ============================================================
/** GET /user/me
 *  ดึงข้อมูล user ตัวเอง (สำหรับ Middleware)
 */

export async function getUserEdge(token?: string) {
  if (!token) return undefined;

  const base = process.env.INTERNAL_API_URL || process.env.NEXT_PUBLIC_API_URL; // ✅ ใช้ INTERNAL ก่อน
  const url = `${base}/user/me`;

  const ctrl = new AbortController();
  const timeout = setTimeout(() => ctrl.abort(), 2000);

  try {
    const res = await fetch(url, {
      headers: { Authorization: `Bearer ${token}` },
      cache: "no-store",
      signal: ctrl.signal,
    });
    if (!res.ok) return undefined;
    // backend ตอบ { message, data: {...user} } ต้องแกะ .data ออกก่อน
    const json = await res.json();
    return (json?.data ?? json) as { role?: string };
  } catch {
    return undefined;
  } finally {
    clearTimeout(timeout);
  }
}
