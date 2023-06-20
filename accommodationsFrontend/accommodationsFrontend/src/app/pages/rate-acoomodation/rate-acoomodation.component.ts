import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation';
import { CreateNewAccommodationRatingRequest } from 'src/app/model/createNewAccommodationRatingRequest';
import { LoggedUserInfo } from 'src/app/model/logged-user-info';
import { User } from 'src/app/model/user';
import { AuthService } from 'src/app/services/auth.service';
import { AvailabilityServiceService } from 'src/app/services/availability-service.service';
import { RatingService } from 'src/app/services/rating.service';
import { ReservationService } from 'src/app/services/reservation.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-rate-acoomodation',
  templateUrl: './rate-acoomodation.component.html',
  styleUrls: ['./rate-acoomodation.component.css']
})
export class RateAcoomodationComponent {

  constructor( @Inject(MAT_DIALOG_DATA) public data: Accommodation,private router: Router, private dialog:MatDialog,
    private availabilityService: AvailabilityServiceService,
   private userService: UserService,
    private reservationService: ReservationService, private authService: AuthService,
    private rating: RatingService){};

    reateNewAccommodationRatingRequest: CreateNewAccommodationRatingRequest = {
      guestId: "",
      dateRating: "",
      rating: 0,
      accommodationId: ""
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
    accomodation(vrednost : any){
    this.reateNewAccommodationRatingRequest.guestId=this.user.id;
    this.reateNewAccommodationRatingRequest.dateRating=new Date().toISOString();
    this.reateNewAccommodationRatingRequest.rating=vrednost.grade;
     
  if (this.data && this.data.id) {
    this.reateNewAccommodationRatingRequest.accommodationId = this.data.id;
  }
   
    console.log( "rating reqiest",this.reateNewAccommodationRatingRequest);
    this.rating.rateAcc(this.reateNewAccommodationRatingRequest).subscribe(
      () => {
        console.log("Zahtev uspešno poslat!"); 
      },
      (error) => {
        console.log("Došlo je do greške prilikom slanja zahteva:", error); 
      }
    );
  };
}
