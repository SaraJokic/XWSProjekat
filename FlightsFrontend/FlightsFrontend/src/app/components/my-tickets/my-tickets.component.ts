import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Ticket } from 'src/models/ticket';
import { TicketService } from 'src/services/ticket.service';
import { MatCardModule } from '@angular/material/card';
import { Flights } from 'src/models/flight.model';
import { FlightService } from 'src/services/flight.service';
import { DialogService } from 'src/services/dialog.service';
import { logedUserInfo } from 'src/app/registration/model/logedUserInfo';
import { AuthService } from 'src/app/registration/services/auth.service';

@Component({
  selector: 'app-my-tickets',
  templateUrl: './my-tickets.component.html',
  styleUrls: ['./my-tickets.component.css']
})
export class MyTicketsComponent implements OnInit{
  
  constructor(private ticketService: TicketService, private flightservice: FlightService,
    private dialogService: DialogService, private authService: AuthService) { }

  public tickets : Ticket[] = [];
  public flights : Flights[] = [];

  logedUser: logedUserInfo = {
    id: "",
    username: "",
    role: "",
    name: ''
  };
  
  ngOnInit(): void {
    this.logedUser = this.authService.getLogedUserInfo() ?? {username: "", role: "", id: "", name: ""};
    this.getMyTickets();
  }
  getMyTickets(): void{
    this.ticketService.findByUserId(this.logedUser.id).subscribe((data) => {
      for (const ticket of data) {
        this.flightservice.getById(ticket.flightid).subscribe(flight => {
          ticket.flight = flight;
        });
      }
      this.tickets = data;
    },
    (error: HttpErrorResponse) => {
      console.log("greskicaa")
      alert(error.message);
    });
  }
  getAllFlights(): void{
    this.flightservice.getAll().subscribe((data) => {
      this.flights = data;
    })
  }

  openTicketDetailsDialog(ticket: Ticket): void{
    this.dialogService.openDialogTicketDetails(ticket);
  }
}
