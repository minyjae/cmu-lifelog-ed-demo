import { redirect } from "next/navigation";

export default function CmuEntraIDCallback() {
  redirect("/signin");
}
