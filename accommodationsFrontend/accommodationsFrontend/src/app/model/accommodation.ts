import { Benefits } from "./benefits";

export interface Accommodation {
    id?: string
    name: string;
    location: string;
    benefits: Benefits;
    minGuests: number;
    maxGuests: number;
    description: string;
    pictures: string[];
    hostId: string;
    automaticApprove: boolean;
}
