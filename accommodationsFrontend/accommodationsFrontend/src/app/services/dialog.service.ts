import { Injectable } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Accommodation } from '../model/accommodation';
import { MakeReservationDialogComponent } from '../pages/make-reservation-dialog/make-reservation-dialog.component';


@Injectable({
  providedIn: 'root'
})
export class DialogService {

  constructor(private dialog: MatDialog) { }

  openDialogReservation(details: Accommodation): void {
    const dialogRef = this.dialog.open(MakeReservationDialogComponent, {
      width: '600px',
      height: '550px',
      data: details
    });

    dialogRef.afterClosed().subscribe((result) => {
      //console.log(`Dialog result: ${result}`);
    });
  }
}
