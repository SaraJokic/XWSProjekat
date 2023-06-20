export interface Reservation {
    id?: string,
    guestId: string,
    accommodationId: string,
    numOfGuests: number,
    status: number,
    startDate: string,
    endDate: string,
}
