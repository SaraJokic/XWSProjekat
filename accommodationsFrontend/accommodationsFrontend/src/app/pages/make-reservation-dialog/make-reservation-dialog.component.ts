import { HttpErrorResponse } from '@angular/common/http';
import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation';
import { AccommodationAvailability, AvailabilitySlot } from 'src/app/model/accommodation-availability';
import { AvailabilityServiceService } from 'src/app/services/availability-service.service';
import { MatDatepickerInputEvent } from '@angular/material/datepicker';
import { Reservation } from 'src/app/model/reservation';
import { User } from 'src/app/model/user';
import { UserService } from 'src/app/services/user.service';
import { ReservationService } from 'src/app/services/reservation.service';
import { DatePipe } from '@angular/common';
import { LoggedUserInfo } from 'src/app/model/logged-user-info';
import { AuthService } from 'src/app/services/auth.service';


@Component({
  selector: 'app-make-reservation-dialog',
  templateUrl: './make-reservation-dialog.component.html',
  styleUrls: ['./make-reservation-dialog.component.css']
})
export class MakeReservationDialogComponent implements OnInit{
  availability: AccommodationAvailability | undefined
  availableSlots: any[] = []
  numberOfGuests: number = 0;
  startDate: Date | undefined;
  endDate: Date | undefined;
  datePipe: DatePipe = new DatePipe('en-US');
  constructor(@Inject(MAT_DIALOG_DATA) public data: Accommodation, 
  private router: Router, private dialog:MatDialog, private availabilityService: AvailabilityServiceService,
  private userService: UserService, private reservationService: ReservationService, private authService: AuthService){}
  user: User = {
    Name: '',
    LastName: '',
    City: '',
    Country: '',
    Username: '',
    Password: '',
    Role: 0,
    Email: '',
    id: '',
    timesCancelled: 0
  };
  logedUser: LoggedUserInfo = {
    id: "",
    username: "",
    role: "",
    name: '',
    email:''
  };
  ngOnInit(): void {
    this.logedUser = this.authService.getLogedUserInfo() ?? {username: "", role: "", id: "", name: "", email:""};
    this.getAvailabilityByAccId()
    this.getUser();
  }
  getUser(){
    this.userService.getUserByUsername(this.logedUser.username).subscribe(
      (data) => {
        this.user = data.user
      }
    );
  }
  getAvailabilityByAccId(){
    this.availabilityService.getAvailabilityByAccId(this.data.id).subscribe(
      (response) => {
        this.availability = response.availability;
        console.log("Availability:", this.availability)
        this.availableSlots = this.availability?.availableSlots || [];
        console.log(this.availableSlots)
      },
      (error: HttpErrorResponse) => {
        //alert("Username or email are already taken");
      }
    );
  }
  makeReservation(){
    const convertedStartDate = this.datePipe.transform(this.startDate, 'yyyy-MM-dd\'T\'HH:mm:ss\'Z\'');
    const convertedEndDate = this.datePipe.transform(this.endDate, 'yyyy-MM-dd\'T\'HH:mm:ss\'Z\'');
    const reservation: Reservation = {
      guestId: this.user.id,
      accommodationId: this.data.id || "",
      numOfGuests: this.numberOfGuests,
      status: 0,
      startDate: convertedStartDate ||"",
      endDate: convertedEndDate ||"",
    }
    console.log("Rezervacija: ", reservation)
    this.reservationService.createReservation(reservation).subscribe(
      (responseres) => {
        alert("Success!")
        if(this.data.automaticApprove){
          this.approveReservation(responseres.id);
        }
      },
      (error: HttpErrorResponse) => {
        alert(error.error.message);
      }
    );
  }
  approveReservation(id: string){
    this.reservationService.changeStatus(id ||"", 1).subscribe(
      (data) => {
        alert(data.message)
      },
      (error: HttpErrorResponse) => {
        alert("Greska");
      });

  }
}
