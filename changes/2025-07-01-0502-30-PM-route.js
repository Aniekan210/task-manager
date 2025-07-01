import { createSupabaseServerClient } from "@/lib/supabase/server";
import { NextResponse } from "next/server";

export async function GET(request) {
  const { searchParams } = new URL(request.url);
  const filter = searchParams.get("filter") || "today";
  const page = parseInt(searchParams.get("page") || "1");
  const itemsPerPage = parseInt(searchParams.get("itemsPerPage") || "5");
  const customDate = searchParams.get("customDate");

  try {
    const supabase = await createSupabaseServerClient();
    const {
      data: { user },
    } = await supabase.auth.getUser();

    if (!user) {
      return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
    }

    const today = new Date();
    let query = supabase
      .from("bookings")
      .select("*")
      .eq("business_id", user.id)
      .order("booking_date", { ascending: true })
      .order("booking_time", { ascending: true });

    switch (filter) {
      case "today":
        const todayStr = today.toISOString().split("T")[0];
        query = query
          .gte("booking_date", todayStr)
          .lte("booking_date", todayStr);
        break;
      case "tomorrow":
        const tomorrow = new Date(today);
        tomorrow.setDate(today.getDate() + 1);
        const tomorrowStr = tomorrow.toISOString().split("T")[0];
        query = query
          .gte("booking_date", tomorrowStr)
          .lte("booking_date", tomorrowStr);
        break;
      case "next7":
        const nextWeek = new Date(today);
        nextWeek.setDate(today.getDate() + 7);
        query = query
          .gte("booking_date", today.toISOString().split("T")[0])
          .lte("booking_date", nextWeek.toISOString().split("T")[0]);
        break;
... (truncated for brevity)