import { Injectable } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { BuyTicketDialogComponent } from 'src/app/components/buy-ticket-dialog/buy-ticket-dialog.component';
import { Flights } from 'src/models/flight.model';

@Injectable({
  providedIn: 'root'
})
export class DialogService {

  constructor(private dialog: MatDialog) { }

  openDialog(details: Flights): void {
    const dialogRef = this.dialog.open(BuyTicketDialogComponent, {
      width: '600px',
      height: '500px',
      data: details
    });

    dialogRef.afterClosed().subscribe((result) => {
      //console.log(`Dialog result: ${result}`);
    });
  }
}
