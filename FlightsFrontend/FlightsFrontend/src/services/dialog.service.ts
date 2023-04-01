import { Injectable } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { BuyTicketDialogComponent } from 'src/app/components/buy-ticket-dialog/buy-ticket-dialog.component';

@Injectable({
  providedIn: 'root'
})
export class DialogService {

  constructor(private dialog: MatDialog) { }

  openDialog(details: any): void {
    const dialogRef = this.dialog.open(BuyTicketDialogComponent, {
      width: '400px',
      height: '500px',
      data: details
    });

    dialogRef.afterClosed().subscribe((result) => {
      console.log(`Dialog result: ${result}`);
    });
  }
}
