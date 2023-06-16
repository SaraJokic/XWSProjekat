import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { ActivatedRoute, Router } from '@angular/router';
import { Observable, forkJoin, map } from 'rxjs';
import { Reservation } from 'src/app/model/reservation';
import { Reservations } from 'src/app/model/reservations';
import { User } from 'src/app/model/user';
import { ReservationService } from 'src/app/services/reservation.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-reservation-requests-host',
  templateUrl: './reservation-requests-host.component.html',
  styleUrls: ['./reservation-requests-host.component.css']
})
export class ReservationRequestsHostComponent implements OnInit{
  reservations: Reservation[] = []
  pendingReservations: Reservation[] = []
  reservationsMapped: Reservations[] = []
  pending = new MatTableDataSource<Reservations>;
  displayedColumns: string[] = ['name', 'startDate', 'endDate', 'numOfGuests', 'timesCancelled', 'actions'];
  constructor(private route: ActivatedRoute, private reservationService: ReservationService,
    private userService: UserService){
  }
  ngOnInit(): void {
    const id = this.route.snapshot.params['id'];
    this.fillTablePending(id);
  }
  fillTablePending(id: string) {
    this.reservationService.getByAccommodationId(id).subscribe(
      (data) => {
        this.reservations = data.reservations;
        this.pendingReservations = []
        this.reservationsMapped = []
        for (let i = 0; i < this.reservations.length; i++) {
          const item = this.reservations[i];
          const st: any = item.status;
          if (st == 'Pending') {
            this.pendingReservations.push(item);
          }
        }
  
        const observables = this.pendingReservations.map((reservation) =>
          this.getReservationUser1(reservation.guestId)
        );
  
        forkJoin(observables).subscribe((users: User[]) => {
          for (let i = 0; i < this.pendingReservations.length; i++) {
            const newRes: Reservations = {
              r: this.pendingReservations[i],
              u: users[i],
            };
            this.reservationsMapped.push(newRes);
          }
          this.pending.data = this.reservationsMapped;
          console.log(this.reservationsMapped);
        });
      },
      (error: HttpErrorResponse) => {
        // Handle error
      }
    );
  }
  getReservationUser1(id: string): Observable<User> {
    return this.userService.getUserById(id).pipe(
      map((data) => data.user)
    );
  }
  approveReservation(reservation: Reservation){
    this.reservationService.changeStatus(reservation.id ||"", 1).subscribe(
      (data) => {
        alert(data.message)
        this.fillTablePending(this.route.snapshot.params['id'])
      },
      (error: HttpErrorResponse) => {
        alert("Greska");
      });

  }
  denyReservation(reservation: Reservation){
    this.reservationService.changeStatus(reservation.id ||"", 2).subscribe(
      (data) => {
        alert(data.message)
        this.fillTablePending(this.route.snapshot.params['id'])
      });
  }
  

}
