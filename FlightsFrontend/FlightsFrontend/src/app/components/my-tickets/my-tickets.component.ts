import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit, ViewChild } from '@angular/core';
import { Ticket } from 'src/models/ticket';
import { TicketService } from 'src/services/ticket.service';
import { MatCardModule } from '@angular/material/card';
import { Flights } from 'src/models/flight.model';
import { FlightService } from 'src/services/flight.service';
import { DialogService } from 'src/services/dialog.service';
import { MatTableDataSource } from '@angular/material/table';
import { MatSort } from '@angular/material/sort';
import { Router } from '@angular/router';

@Component({
  selector: 'app-my-tickets',
  templateUrl: './my-tickets.component.html',
  styleUrls: ['./my-tickets.component.css']
})
export class MyTicketsComponent implements OnInit{
  
  constructor(private ticketService: TicketService, private flightservice: FlightService,

    private dialogService: DialogService,  private router: Router) { }


  public tickets : Ticket[] = [];
  public flights : Flights[] = [];

  public izabran : any ;
  ticKets = new MatTableDataSource<Ticket[]>;
  message?:any;

  @ViewChild(MatSort)
  sort: MatSort = new MatSort;


  ngOnInit(): void {
    this.getMyTickets();
    this.retrieveTickets(this.izabran);
  }
  getMyTickets(): void{
    this.ticketService.findByUserId("6428e5416833a3ee718b4af0").subscribe((data) => {
      for (const ticket of data) {
        this.flightservice.getById(ticket.flightid).subscribe(flight => {
          ticket.flight = flight;
          this.CheckIfTicketExpired(flight.starttime, ticket)
        },
        error => {
          // Delete from tickets and db
          const index = this.tickets.indexOf(ticket);
          if (index !== -1) {
            this.tickets.splice(index, 1);
          }
          this.DeleteTicket(ticket.id ?? "")
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


/*
  deleteTicket(deleting : any){
    this.message = '';
    this.izabran = deleting.id;
    if(confirm("Are you sure?")){
    this.ticketService.delete(this.izabran).subscribe(
      (resp) =>{
        this.retrieveTickets(deleting);
        this.reloadCurrentRoute();
        return console.log("Ticket Deleted!");
         
      }, err=>{
         console.error("Neuspesno");
      });
    }
  }


  


  retrieveTickets(id:any): void {
    this.ticketService.findByUserId(id)
      .subscribe({
        next: (data) => {
          this.ticKets = new MatTableDataSource(<Ticket[][]><unknown>data);
          this.ticKets.sort = this.sort;
          console.log(data);
        },
        error: (e) => console.error(e)
      });
  }

  reloadCurrentRoute() {
    let currentUrl = this.router.url;
    this.router.navigateByUrl('/', {skipLocationChange: true}).then(() => {
        this.router.navigate([currentUrl]);
    });
}

*/


  GoToAllFlightsPage(): void{
    this.router.navigate(["/flights"]); 
  }
  DeleteTicket(id: string): void{
    this.ticketService.delete(id ?? "").subscribe((resp) =>{
      console.log("Deleted!");
    }, err=>{
      return console.error("Not deleted");
    });
  }
  //if the date of departure passed
  CheckIfTicketExpired(date: Date, ticket: Ticket){
    const currentDate = new Date();
    if(date < currentDate){
      this.ticketService.update(ticket, (ticket.id ?? "")).subscribe((resp) =>{
        console.log("Ticket set to expired");
      }, err =>{
        console.error("Error")
      });
    }
  }

}
