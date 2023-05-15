import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { CreateAccommodationDto } from '../model/create-accommodation-dto';
import { Observable } from 'rxjs';
import { Accommodation } from '../model/accommodation';
import { AccommodationAvailability } from '../model/accommodation-availability';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
}; 

@Injectable({
  providedIn: 'root'
})

export class AccommodationServiceService {
  private apiUrl = 'http://localhost:8000';
  constructor(private http: HttpClient) { }

  createAccommodation(accommodation: Accommodation): Observable<any> {
    const url = `${this.apiUrl}/accommodations/create`;
    return this.http.post<any>(url, { acc: accommodation }, httpOptions);
  }
}
