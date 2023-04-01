import { Injectable } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { BuyTicketDialogComponent } from 'src/app/components/buy-ticket-dialog/buy-ticket-dialog.component';
import { TicketDetailsDialogComponent } from 'src/app/components/ticket-details-dialog/ticket-details-dialog.component';
import { Flights } from 'src/models/flight.model';
import { Ticket } from 'src/models/ticket';

@Injectable({
  providedIn: 'root'
})
export class DialogService {

  constructor(private dialog: MatDialog) { }

  openDialogBuyingTicket(details: Flights): void {
    const dialogRef = this.dialog.open(BuyTicketDialogComponent, {
      width: '600px',
      height: '500px',
      data: details
    });

    dialogRef.afterClosed().subscribe((result) => {
      //console.log(`Dialog result: ${result}`);
    });
  }
  openDialogTicketDetails(details: Ticket): void {
    const dialogRef = this.dialog.open(TicketDetailsDialogComponent, {
      width: '600px',
      height: '500px',
      data: details
    });

    dialogRef.afterClosed().subscribe((result) => {
      //console.log(`Dialog result: ${result}`);
    });
  }
}
