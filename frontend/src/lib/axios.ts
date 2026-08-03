import axios from "axios";

export const authHeader = (token?: string) => {
  return token ? { Authorization: `Bearer ${token}` } : undefined;
};

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  withCredentials: true,
  timeout: 15000,
  headers: { "Content-Type": "application/json" },
});

// backend ตอบในรูปแบบมาตรฐาน { message, code?, data? } เสมอ
// interceptor นี้แกะ envelope ออกให้ res.data เป็น payload จริง (array/object)
// เพื่อให้ทุก api function อ่าน res.data ได้ตรงๆ เหมือนเดิม
api.interceptors.response.use(
  (res) => {
    const body = res.data;
    if (
      body &&
      typeof body === "object" &&
      "message" in body &&
      "data" in body
    ) {
      res.data = body.data;
    }
    return res;
  },
  (err) => {
    // ใช้ message จาก backend ({ message, code }) ถ้ามี ไม่งั้น fallback เป็น message ของ axios
    const backendMessage = err?.response?.data?.message;
    return Promise.reject(new Error(backendMessage ?? err.message));
  }
);

export default api;
