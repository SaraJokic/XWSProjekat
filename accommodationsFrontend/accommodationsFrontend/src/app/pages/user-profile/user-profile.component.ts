import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/model/user';
import { UserService } from 'src/app/services/user.service';
import { UpdateUserReq } from 'src/app/model/update-user-req';
import { Router } from '@angular/router';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent implements OnInit{
  constructor(private userService: UserService,  private router: Router){}
  user: User = {
    Name: 'John',
    LastName: 'Doe',
    City: 'New York',
    Country: 'USA',
    Username: 'johndoe123',
    Password: 'password123',
    Role: 0,
    Email: '',
    id: ''
  };
  ngOnInit(): void {
    this.getUser();
  }
  getUser(){
    this.userService.getAll().subscribe(
      (data) => {
        this.user = data.users[0]
        console.log(this.user)
        //this.router.navigate(["/login"]);
      },
      (error: HttpErrorResponse) => {
        //alert("Username or email are already taken");
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
        //this.router.navigate(["/login"]);
      },
      (error: HttpErrorResponse) => {
        alert("Error");
      }
    );
  }
  deleteProfile(id: any){
    this.userService.delete(id).subscribe(
      (data) => {
        console.log(data)
        //this.router.navigate(["/login"]);
        this.router.navigate(["/register"]);
        
      },
      (error: HttpErrorResponse) => {
        alert("Error");
      }
    );
  }
}
