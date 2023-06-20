import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation';
import { CreateNewAccommodationRatingRequest } from 'src/app/model/createNewAccommodationRatingRequest';
import { LoggedUserInfo } from 'src/app/model/logged-user-info';
import { AuthService } from 'src/app/services/auth.service';
import { AvailabilityServiceService } from 'src/app/services/availability-service.service';
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
    private reservationService: ReservationService, private authService: AuthService){};

    createNewHostRatingRequest: CreateNewAccommodationRatingRequest = {
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
    ngOnInit(): void {
      this.logedUser = this.authService.getLogedUserInfo() ?? {username: "", role: "", id: "", name: "", email:""};
  
    }


    accomodation(vrednost : any){
    this.createNewHostRatingRequest.guestId=this.logedUser.id;
    this.createNewHostRatingRequest.dateRating=new Date().toISOString();
    this.createNewHostRatingRequest.rating=vrednost.grade;
     
  if (this.data && this.data.id) {
    this.createNewHostRatingRequest.accommodationId = this.data.id;
  }
   
    console.log( "rating reqiest",this.createNewHostRatingRequest);
  };
}
