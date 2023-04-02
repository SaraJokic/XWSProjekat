import { Component, OnInit } from '@angular/core';
import { RegistrationService } from '../../services/registration.service';
import { User } from '../../model/user';
import { HttpErrorResponse } from '@angular/common/http';


@Component({
  selector: 'app-registration-form',
  templateUrl: './registration-form.component.html',
  styleUrls: ['./registration-form.component.css']
})
export class RegistrationFormComponent implements OnInit{

  constructor(private registrationService: RegistrationService) { }
  
  ngOnInit(): void {
    throw new Error('Method not implemented.');
  }
  registerUser(user: any){
    let newUser: User = {
    username: user.username,
    password: user.password1,
    email: user.email,
    name: user.name,
    lastname: user.surname,
    role: 0
    }
    console.log(newUser)

    this.registrationService.add(newUser).subscribe(
      (data) => {
        alert("Success!");
      },
      (error: HttpErrorResponse) => {
        alert("Username or email are already taken");
      }
    );
  }
}
