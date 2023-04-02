import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Ticket } from 'src/models/ticket';

const baseUrl = 'http://localhost:8080';

@Injectable({
  providedIn: 'root'
})
export class TicketService {

  constructor(private http: HttpClient) { }

  findAllTickets(){
    return this.http.get<Ticket[]>(baseUrl + '/tickets/all');
  }
  findById(input: number){
    return this.http.get<Ticket>(baseUrl + `/tickets/get/${input}`);
  }
  findByUserId(input: string){
    return this.http.get<Ticket[]>(baseUrl + `/tickets/getbyuser/${input}`);
  }
  delete(input: string){
    return this.http.delete(baseUrl + `/tickets/delete/${input}`)
  }
  add(newTicket: Ticket){
    return this.http.post<Ticket>(baseUrl + '/tickets/buy', newTicket);
  }
  update(newTicket: Ticket, input : string){
    return this.http.patch<Ticket>(baseUrl + `/tickets/update/${input}`, newTicket);
  }
}
