import { Injectable } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { Flights } from "src/models/flight.model";
import { Observable } from "rxjs";

const baseUrl = 'http://localhost:4200';

const novoUrl = 'http://localhost:4200/filter';

@Injectable({
  providedIn: 'root'
})
export class FlightService {

  constructor(private http: HttpClient) {}

  getById(id: any): Observable<Flights> {
    return this.http.get<Flights>(`${baseUrl}/${id}`);
  }

  getAll(): Observable<Flights[]> {
    return this.http.get<Flights[]>(baseUrl);
  }

  
  create(data: Flights): Observable<any> {
    return this.http.post(baseUrl, data);
  }


  delete(id: any): Observable<any> {
    return this.http.delete(`${baseUrl}/${id}`);
  }

  /*
  update(id: any, data: any): Observable<any> {
    return this.http.put(`${baseUrl}/${id}`, data);
  }
*/

  findByName(title: any): Observable<Flights[]> {
    return this.http.get<Flights[]>(`${baseUrl}?name=${name}`);
  }



  Search(nesto:any): Observable<Flights[]> {
    console.log(nesto);
    return this.http.get<Flights[]>(`${novoUrl}/${nesto}`);

    

  }
}