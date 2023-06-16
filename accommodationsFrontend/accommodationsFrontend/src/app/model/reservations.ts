import { Accommodation } from "./accommodation";
import { Reservation } from "./reservation";
import { User } from "./user";

export interface Reservations {
    r: Reservation,
    u?: User,
    a?: Accommodation
}
