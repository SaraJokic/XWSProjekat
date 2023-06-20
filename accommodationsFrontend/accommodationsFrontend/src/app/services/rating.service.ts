import { Injectable } from '@angular/core';
import { CreateNewHostRatingRequest } from '../model/createNewHostRatingRequest';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { CreateAccommodationDto } from '../model/create-accommodation-dto';
import { CreateNewAccommodationRatingRequest } from '../model/createNewAccommodationRatingRequest';

@Injectable({
  providedIn: 'root'
})
export class RatingService {
  private apiUrl = 'http://localhost:8000';
  
  constructor(private http: HttpClient) { }

  rateHost(rating: CreateNewHostRatingRequest): Observable<any> {
    const url = `${this.apiUrl}/ratings/addhost`;
    return this.http.post<any>(url,rating );
  }

  rateAcc(rating: CreateNewAccommodationRatingRequest): Observable<any> {
    const url = `${this.apiUrl}/ratings/addacc`;
    return this.http.post<any>(url,rating );
  }
}


