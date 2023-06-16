export interface AccommodationAvailability {
    id?: string
    accommodationId: string;
    availableSlots: AvailabilitySlot[];
    price: number;
    isPricePerPerson: boolean;
    changePrice: PriceChange[];
}
export interface AvailabilitySlot {
    start_date: Date;
    end_date: Date;
  }
  
  export interface PriceChange {
    startdate: Date;
    enddate: Date;
    change: number;
  }
