import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit, Inject, ViewChild } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { Router } from '@angular/router';
import { Flights } from 'src/models/flight.model';
import { Ticket } from 'src/models/ticket';
import { FlightService } from 'src/services/flight.service';
import { TicketService } from 'src/services/ticket.service';

@Component({
  selector: 'app-buy-ticket-dialog',
  templateUrl: './buy-ticket-dialog.component.html',
  styleUrls: ['./buy-ticket-dialog.component.css']
})

export class BuyTicketDialogComponent implements OnInit {


  constructor(@Inject(MAT_DIALOG_DATA) public data: Flights, private ticketService: TicketService, private router: Router, private dialog:MatDialog) { }
  
  flights = new MatTableDataSource<Flights[]>;


  
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

  @ViewChild(MatSort)
  sort: MatSort = new MatSort;

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
        this.dialog.closeAll();
        this.reloadCurrentRoute();

      },
      (error: HttpErrorResponse) => {
        alert(error.message);
      }
    )
  }


  reloadCurrentRoute() {
    let currentUrl = this.router.url;
    this.router.navigateByUrl('/', {skipLocationChange: true}).then(() => {
        this.router.navigate([currentUrl]);
    });
}


}
