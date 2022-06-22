import type { Guest } from "./Guest";

export type ReservationCreate = {
  id: string;
  first_name: string;
  last_name: string;
  phone_number: string;
  remark: string;
  email: string;
  guest_count: number;
  guest_persona: Guest[];
  is_cancelled: boolean;
  is_rescheduled: boolean;
  dts_id: string;
  restaurant_id: string;
};
