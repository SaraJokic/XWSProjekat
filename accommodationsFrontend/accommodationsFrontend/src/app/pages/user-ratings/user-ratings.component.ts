import { Component, OnInit } from '@angular/core';
import { CreateNewHostRatingRequest } from 'src/app/model/createNewHostRatingRequest';
import { LoggedUserInfo } from 'src/app/model/logged-user-info';
import { User } from 'src/app/model/user';
import { AuthService } from 'src/app/services/auth.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-user-ratings',
  templateUrl: './user-ratings.component.html',
  styleUrls: ['./user-ratings.component.css']
})
export class UserRatingsComponent implements OnInit {
  constructor( private userServce: UserService, private authService: AuthService){}
  ratings:  any[] =  [];
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
      this.getInfo();
    }
    getUser(){
      this.userServce.getUserByUsername(this.logedUser.username).subscribe(
        (data) => {
          this.user = data.user
        }
      );
    }
   
public getInfo(){
  
  console.log("userid", this.user.id)
    this.userServce.getRatingsUserById(this.user.id).subscribe(
      (accommodations) => {
        this.ratings = accommodations;
        console.log(accommodations);
      },
      (error: any) => {
        console.error(error);
      }
    );
}
}
