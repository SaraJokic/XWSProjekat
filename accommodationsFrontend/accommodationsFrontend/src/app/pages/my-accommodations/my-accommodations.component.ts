import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { Router } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation';
import { LoggedUserInfo } from 'src/app/model/logged-user-info';
import { User } from 'src/app/model/user';
import { AccommodationServiceService } from 'src/app/services/accommodation-service.service';
import { AuthService } from 'src/app/services/auth.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-my-accommodations',
  templateUrl: './my-accommodations.component.html',
  styleUrls: ['./my-accommodations.component.css']
})
export class MyAccommodationsComponent implements OnInit{
  displayedColumns: string[] = ['name', 'location', 'view', 'req'];
  accommodations = new MatTableDataSource<Accommodation>;
  logedUser: LoggedUserInfo = {
    id: "",
    username: "",
    role: "",
    name: '',
    email:'',
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
  constructor(private accommodationService: AccommodationServiceService, 
    private userService: UserService, private router: Router, private authService: AuthService){
  }
  ngOnInit(): void {
    this.getUser()
  }
  getAccommodations() {
    this.accommodationService.getMyAccommodations(this.user.id).subscribe(
      (data) => {
        this.accommodations.data = data.acc;
      },
      (error: HttpErrorResponse) => {
        alert('No accommodations');
      }
    );
  }
  getUser(){
    this.logedUser = this.authService.getLogedUserInfo() ?? {username: "", role: "", id: "", name: "",email:""};
    this.userService.getUserByUsername(this.logedUser.username).subscribe(
      (data) => {
        this.user = data.user
        this.getAccommodations()
      },
      (error: HttpErrorResponse) => {
      }
    );
  }
  viewAccommodation(id: string) {
    this.router.navigate(['/accommodationview', id]);
  }
  viewRequests(id: string){
    this.router.navigate(['/accommodationrequests', id]);
  }
  

}
