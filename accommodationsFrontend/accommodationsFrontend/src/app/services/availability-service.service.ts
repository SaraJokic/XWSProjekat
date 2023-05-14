import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AccommodationAvailability } from '../model/accommodation-availability';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AvailabilityServiceService {
  private apiUrl = 'http://localhost:8000';
  constructor(private http: HttpClient) { }
  createNewAvailability(request: AccommodationAvailability): Observable<any> {
    const url = `${this.apiUrl}/availability/add`;
    return this.http.post<any>(url, request);
  }
}
