import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation';
import { CreateNewHostRatingRequest } from 'src/app/model/createNewHostRatingRequest';
import { LoggedUserInfo } from 'src/app/model/logged-user-info';
import { User } from 'src/app/model/user';
import { AuthService } from 'src/app/services/auth.service';
import { AvailabilityServiceService } from 'src/app/services/availability-service.service';
import { RatingService } from 'src/app/services/rating.service';
import { ReservationService } from 'src/app/services/reservation.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-rating',
  templateUrl: './rating.component.html',
  styleUrls: ['./rating.component.css']
})
export class RatingComponent {


 constructor(@Inject(MAT_DIALOG_DATA) public data: Accommodation, private router: Router, private dialog:MatDialog,
  private availabilityService: AvailabilityServiceService,
 private userService: UserService,
  private reservationService: ReservationService, private authService: AuthService, private rating : RatingService){};

  createNewHostRatingRequest: CreateNewHostRatingRequest = {
    guestId: "",
    dateRating: "",
    rating: 0,
    hostId: ""
  };
  logedUser: LoggedUserInfo = {
    id: "",
    username: "",
    role: "",
    name: '',
    email:''
  };
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
  ngOnInit(): void {
    this.logedUser = this.authService.getLogedUserInfo() ?? {username: "", role: "", id: "", name: "", email:""};
    this.getUser();
  }
  getUser(){
    this.userService.getUserByUsername(this.logedUser.username).subscribe(
      (data) => {
        this.user = data.user
      }
    );
  }
  host(vrednost : any){
  this.createNewHostRatingRequest.guestId=this.user.id;
  const date = new Date();
const year = date.getUTCFullYear();
const month = String(date.getUTCMonth() + 1).padStart(2, '0');
const day = String(date.getUTCDate()).padStart(2, '0');
const hours = String(date.getUTCHours()).padStart(2, '0');
const minutes = String(date.getUTCMinutes()).padStart(2, '0');
const seconds = String(date.getUTCSeconds()).padStart(2, '0');

const formattedDate = `${year}-${month}-${day}T${hours}:${minutes}:${seconds}Z`;

  this.createNewHostRatingRequest.dateRating=formattedDate;
  this.createNewHostRatingRequest.rating=vrednost.grade;
  if (this.data && this.data.id) {
    this.createNewHostRatingRequest.hostId = this.data.hostId;
  }
  console.log("forma",this.createNewHostRatingRequest);

  this.rating.rateHost(this.createNewHostRatingRequest).subscribe(
    () => {
      console.log("Zahtev uspešno poslat!"); 
    },
    (error) => {
      console.log("Došlo je do greške prilikom slanja zahteva:", error); 
    }
  );
}
}
