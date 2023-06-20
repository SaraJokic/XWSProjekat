import { Injectable } from '@angular/core';
import { CreateNewHostRatingRequest } from '../model/createNewHostRatingRequest';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RatingService {
  private apiUrl = 'http://localhost:8000';

  constructor(private http: HttpClient) { }

  createAccommodation(rating: CreateNewHostRatingRequest): Observable<any> {
    const url = `${this.apiUrl}/ratings/addhost`;
    return this.http.post<any>(url,rating );
  }
}
