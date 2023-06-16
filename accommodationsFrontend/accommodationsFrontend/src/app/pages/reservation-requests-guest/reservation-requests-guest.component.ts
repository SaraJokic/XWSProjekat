import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { Observable, forkJoin, map } from 'rxjs';
import { Accommodation } from 'src/app/model/accommodation';
import { Benefits } from 'src/app/model/benefits';
import { LoggedUserInfo } from 'src/app/model/logged-user-info';
import { Reservation } from 'src/app/model/reservation';
import { Reservations } from 'src/app/model/reservations';
import { UpdateUserReq } from 'src/app/model/update-user-req';
import { User } from 'src/app/model/user';
import { AccommodationServiceService } from 'src/app/services/accommodation-service.service';
import { AuthService } from 'src/app/services/auth.service';
import { ReservationService } from 'src/app/services/reservation.service';
import { UserService } from 'src/app/services/user.service';


@Component({
  selector: 'app-reservation-requests-guest',
  templateUrl: './reservation-requests-guest.component.html',
  styleUrls: ['./reservation-requests-guest.component.css']
})
export class ReservationRequestsGuestComponent implements OnInit{
  reservations: Reservation[] = []
  mappedReservations: Reservations[] = []
  reservationsource = new MatTableDataSource<Reservations>;
  displayedColumns: string[] = ['name', 'location', 'startDate', 'endDate', 'status', 'cancel'];
  constructor(private authService: AuthService, private userService: UserService, private accommodationService: AccommodationServiceService,
    private reservationService: ReservationService){}
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
    name: ''
  };


  ngOnInit(): void {
    this.logedUser = this.authService.getLogedUserInfo() ?? {username: "", role: "", id: "", name: ""};
    this.getUser();
  }
  getUser(){
    this.userService.getUserByUsername(this.logedUser.username).subscribe(
      (data) => {
        this.user = data.user
        this.getReservations();
      }
    );
  }
  getReservations(){
    this.reservationService.getByUserId(this.user.id).subscribe(
      (data) => {
        this.reservations = data.reservations;
        const observables = this.reservations.map((reservation) =>
          this.getReservationAccommodation1(reservation.accommodationId)
        );
        forkJoin(observables).subscribe((accommodations: Accommodation[]) => {
          for (let i = 0; i < this.reservations.length; i++) {
            const newRes: Reservations = {
              r: this.reservations[i],
              a: accommodations[i],
            };
            this.mappedReservations.push(newRes);
          }
          this.reservationsource.data = this.mappedReservations;
          console.log(this.mappedReservations);
        });
      })
  }
  
  getReservationAccommodation1(id: string): Observable<Accommodation> {
    return this.accommodationService.getAccommodationById(id).pipe(
      map((data) => data.acc)
    );
  }
  cancelReservation(r: Reservation){
    if (this.isBeforeStartDate(new Date(r.startDate))){
      alert("You can't cancel your reservation. It's too late now.")
      return;
    }
    if(r.status.toString() == "Pending"){
      alert("Your reservation is still pending.")
      return;
    }
    this.reservationService.delete(r.id).subscribe(
      (data) => {
        this.user.timesCancelled++;
        const request: UpdateUserReq = {
          user: this.user,
          id: '',
          UserId: this.user.id
        }
        this.userService.updateUser(request).subscribe(
          (data) =>{
            console.log("poruka:", data.message)
          } )
        alert("Sucessfully deleted your reservation.")
      })
  }
  isBeforeStartDate(startDate: Date): boolean {
    const today = new Date();
    const start = new Date(startDate);
    const oneDayBeforeStartDate = new Date(startDate);
    oneDayBeforeStartDate.setDate(start.getDate() - 1);
    console.log(!(today < oneDayBeforeStartDate))
    
    return !(today < oneDayBeforeStartDate);
  }
}
