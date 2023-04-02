import { Injectable } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { Flights } from "src/models/flight.model";
import { Observable } from "rxjs";
import { SearchFlightsDTO } from "src/models/flightDTO.model";

const baseUrl = 'http://localhost:8080';
const novoUrl = 'http://localhost:8080/a/filter';
const probaUrl = 'http://localhost:8080/flight/filter/seats';


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

  
  update(id: any, data: any): Observable<any> {
    return this.http.patch(`${baseUrl}/${id}`, data);
  }


  findByName(title: any): Observable<Flights[]> {
    return this.http.get<Flights[]>(`${baseUrl}?name=${name}`);
  }


  Proba(proba: any): Observable<Flights[]> {
    return this.http.get<Flights[]>(`${novoUrl}?fromplace=${proba}`);
  }

  Search(nesto:any): Observable<Flights[]> {
    console.log(nesto);
    return this.http.get<Flights[]>(`${novoUrl}/${nesto}`);
  }


  
  SearchByAll(prva:any, druga:any, treca:any): Observable<Flights[]> {
    return this.http.get<Flights[]>(`${probaUrl}?fromplace=${prva}&toplace=${druga}numofseats=${treca}`);
  }



/*
 SearchFlights(dto: SearchFlightsDTO){
  let flights:any;

  this.http.put(baseUrl + "/flight/search", dto).subscribe((response:any) => {
    flights = response.data });

  return flights;
}

AAAA(startPlace: string, endPlace: String, startDateString: string, endDateString: string): Observable<any[]> {
  return this.http.get<Flights[]>(`${novoUrl}?startPlace=${startPlace}`+ '&endPlace=' + endPlace + '&startDate=' + startDateString + '&endDate=' + endDateString);
}
*/
  
}