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
  private getApiUrl = 'http://localhost:8000/accommodations';

  private baseUrl='http://localhost:8000/GetAllProminentAccommodation';

  constructor(private http: HttpClient) { }

  Search(accommodation: any) {
    const url = `${this.apiUrl}/accommodations/search`;
    return this.http.post<any>(url, accommodation , httpOptions);
  }
  createAccommodation(accommodation: Accommodation): Observable<any> {
    const url = `${this.apiUrl}/accommodations/create`;
    return this.http.post<any>(url, { acc: accommodation }, httpOptions);
  }
  getMyAccommodations(id: any): Observable<any>{
    return this.http.get<any>(`${this.apiUrl}/accommodations/getbyhostid/${id}`); 
  }
  getAccommodationById(id: any): Observable<any>{
    return this.http.get<any>(`${this.apiUrl}/accommodations/get/${id}`); 
  }
  updateAccommodation(id: string, accommodation: Accommodation): Observable<any>{
    const url = `${this.apiUrl}/accommodations/update`;
    return this.http.put<any>(url, { id: id, accommodation: accommodation }, httpOptions);
  }
  getAll(): Observable<any> {
    return this.http.get<any>(this.getApiUrl);
  }

  GetAllProminentAccommodation(): Observable<any> {
    return this.http.get<any>(this.baseUrl);
  }

  
}
