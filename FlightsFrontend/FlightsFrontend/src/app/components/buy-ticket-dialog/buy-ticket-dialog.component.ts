import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit, Inject } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Flights } from 'src/models/flight.model';
import { Ticket } from 'src/models/ticket';
import { TicketService } from 'src/services/ticket.service';

@Component({
  selector: 'app-buy-ticket-dialog',
  templateUrl: './buy-ticket-dialog.component.html',
  styleUrls: ['./buy-ticket-dialog.component.css']
})
export class BuyTicketDialogComponent implements OnInit {

  constructor(@Inject(MAT_DIALOG_DATA) public data: Flights, private ticketService: TicketService) { }
  
  numTickets: number = 1;
  flight: Flights ={
    id: this.data.id,
    fromplace: this.data.fromplace,
    toplace:this.data.toplace,
    starttime: this.data.starttime ? new Date(this.data.starttime) : undefined,
    endtime: this.data.endtime ? new Date(this.data.endtime) : undefined,
    ticketprice: this.data.ticketprice,
    numofseats:this.data.numofseats
  };
  
  ngOnInit(): void {
    console.log(this.data)
  }
  buyTicket(){
    const newTicket: Ticket = {
      userid: "6426f65971b16d7d27fe5bb8",
      flightid: this.data.id!,
      expired: false,
      quantity: this.numTickets,
    };
    this.ticketService.add(newTicket).subscribe(
      (data) => {
        alert("Success!");
      },
      (error: HttpErrorResponse) => {
        alert(error.message);
      }
    )
  }

}
