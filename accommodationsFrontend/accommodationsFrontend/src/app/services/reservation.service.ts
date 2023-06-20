import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Reservation } from '../model/reservation';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
}; 

@Injectable({
  providedIn: 'root'
})
export class ReservationService {
  private apiUrl = 'http://localhost:8000';
  constructor(private http: HttpClient) { }

  createReservation(reservation: any): Observable<any> {
    const url = `${this.apiUrl}/reservations/create`;
    return this.http.post<any>(url, 
      { guestId: reservation.guestId, accommodationId: reservation.accommodationId, startDate: reservation.startDate, endDate: reservation.endDate, numOfGuests: reservation.numOfGuests, status: reservation.status }, httpOptions);
  }
  getByAccommodationId(id: any): Observable<any>{
    return this.http.get<any>(`${this.apiUrl}/reservations/get/accommodationid/${id}`); 
  }
  changeStatus(id: string, status: number): Observable<any>{
    const url = `${this.apiUrl}/reservations/changestatus`;
    return this.http.put<any>(url, { id: id, status: status }, httpOptions);
  }
  delete(id: any): Observable<any> {
    return this.http.delete(`${this.apiUrl}/reservations/delete/${id}`);
  }
  getByUserId(id: any): Observable<any>{
    return this.http.get<any>(`${this.apiUrl}/reservations/get/userid/${id}`); 
  }
}
