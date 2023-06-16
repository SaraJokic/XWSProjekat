import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/model/user';
import { UserService } from 'src/app/services/user.service';
import { UpdateUserReq } from 'src/app/model/update-user-req';
import { Router } from '@angular/router';
import { LoggedUserInfo } from 'src/app/model/logged-user-info';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent implements OnInit{
  constructor(private userService: UserService,  private router: Router, private authService: AuthService){}
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
      }
    );
  }
  onSubmit(){
    const updateUser: UpdateUserReq = {
      user: this.user,
      UserId: this.user.id,
      id: " "
    }
    console.log("update user: ", this.user)
    this.userService.updateUser(updateUser).subscribe(
      (data) => {
        console.log(data)
      },
      (error: HttpErrorResponse) => {
        alert("Error");
      }
    );
  }
  deleteProfile(id: any){
    console.log(this.user)
    this.userService.delete(this.user.id).subscribe(
      (data) => {
        console.log(data)
        this.authService.removeToken();
        this.router.navigate(["/register"]);
        
      },
      (error: HttpErrorResponse) => {
        alert("Error");
      }
    );
  }
}
