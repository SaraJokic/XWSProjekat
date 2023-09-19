import { Injectable } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Accommodation } from '../model/accommodation';
import { MakeReservationDialogComponent } from '../pages/make-reservation-dialog/make-reservation-dialog.component';
import { RatingComponent } from '../pages/rating/rating.component';
import { RateAcoomodationComponent } from '../pages/rate-acoomodation/rate-acoomodation.component';
import { ViewAccomodationDialogComponent } from '../pages/view-accomodation-dialog/view-accomodation-dialog.component';


@Injectable({
  providedIn: 'root'
})
export class DialogService {

  constructor(private dialog: MatDialog) { }

  openDialogReservation(details: Accommodation): void {
    const dialogRef = this.dialog.open(MakeReservationDialogComponent, {
      width: '1900px',
      height: '1000px',
      data: details
    });

    dialogRef.afterClosed().subscribe((result) => {
      //console.log(`Dialog result: ${result}`);
    });
  }

openHostDialog(details: Accommodation): void {
  const dialogRef = this.dialog.open(RatingComponent, {
    width: '600px',
      height: '450px',
      data: details
    });
    dialogRef.afterClosed().subscribe((result) => {
      //console.log(`Dialog result: ${result}`);
    });
  }
  openViewAccDialog(details: Accommodation): void {
    const dialogRef = this.dialog.open(ViewAccomodationDialogComponent, {
      width: '1600px',
        height: '1000px',
        data: details
      });
      dialogRef.afterClosed().subscribe((result) => {
        //console.log(`Dialog result: ${result}`);
      });
    }
openAccDialog(details: Accommodation): void {
  const dialogRef = this.dialog.open(RateAcoomodationComponent, {
    width: '600px',
      height: '450px',
      data: details
    });
    dialogRef.afterClosed().subscribe((result) => {
      //console.log(`Dialog result: ${result}`);
    });
  }
}

