import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Ticket } from 'src/models/ticket';
import { TicketService } from 'src/services/ticket.service';
import { MatCardModule } from '@angular/material/card';
import { Flights } from 'src/models/flight.model';
import { FlightService } from 'src/services/flight.service';
import { DialogService } from 'src/services/dialog.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-my-tickets',
  templateUrl: './my-tickets.component.html',
  styleUrls: ['./my-tickets.component.css']
})
export class MyTicketsComponent implements OnInit{
  
  constructor(private ticketService: TicketService, private flightservice: FlightService,
    private dialogService: DialogService, private router: Router) { }

  public tickets : Ticket[] = [];
  public flights : Flights[] = [];
  
  ngOnInit(): void {
    this.getMyTickets();
  }
  getMyTickets(): void{
    this.ticketService.findByUserId("6428e5416833a3ee718b4af0").subscribe((data) => {
      for (const ticket of data) {
        this.flightservice.getById(ticket.flightid).subscribe(flight => {
          ticket.flight = flight;
        });
      }
      this.tickets = data;
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
  GoToAllFlightsPage(): void{
    this.router.navigate(["/flights"]); 
  }
}
