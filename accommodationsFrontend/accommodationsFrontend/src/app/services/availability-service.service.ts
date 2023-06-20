import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { AccommodationAvailability, AvailabilitySlot } from '../model/accommodation-availability';
import { Observable } from 'rxjs';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
}; 

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
  getAvailabilityByAccId(id: any): Observable<any>{
    return this.http.get<any>(`${this.apiUrl}/availability/getbyaccid/${id}`); 
  }
  updateAvailability(id: string, availabilbility: AccommodationAvailability): Observable<any>{
    const url = `${this.apiUrl}/availability/update`;
    return this.http.put<any>(url, { id: id, availability: availabilbility }, httpOptions);
  }
}
