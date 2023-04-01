import { Flights } from "./flight.model"

export interface Ticket {
    id?: string,
    userid: string,
    flightid: string,
    quantity: number,
    expired:boolean,
    flight?: Flights
}
