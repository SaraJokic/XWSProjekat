import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Flights } from 'src/models/flight.model';
import { Ticket } from 'src/models/ticket';

@Component({
  selector: 'app-ticket-details-dialog',
  templateUrl: './ticket-details-dialog.component.html',
  styleUrls: ['./ticket-details-dialog.component.css']
})
export class TicketDetailsDialogComponent implements OnInit{

  constructor(@Inject(MAT_DIALOG_DATA) public data: Ticket) { }

  flight: Flights ={
    id: this.data.id,
    fromplace: this.data.flight?.fromplace,
    toplace:this.data.flight?.toplace,
    starttime: this.data.flight?.starttime ? new Date(this.data.flight?.starttime) : undefined,
    endtime: this.data.flight?.endtime ? new Date(this.data.flight?.endtime) : undefined,
    ticketprice: this.data.flight?.ticketprice,
    numofseats:this.data.flight?.numofseats
  };
  
  ngOnInit(): void {
    //console.log(typeof this.flight.starttime)
  }



}
