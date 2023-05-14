import { Benefits } from "./benefits";

export interface Accommodation {
    name: string;
    location: string;
    benefits: Benefits;
    minGuests: number;
    maxGuests: number;
}
