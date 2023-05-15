import { Component, OnInit } from '@angular/core';
import { RegistrationService } from '../../services/registration.service';
import { User } from '../../model/user';
import { HttpErrorResponse } from '@angular/common/http';
import { Router } from '@angular/router';


@Component({
  selector: 'app-registration-form',
  templateUrl: './registration-form.component.html',
  styleUrls: ['./registration-form.component.css']
})
export class RegistrationFormComponent implements OnInit{

  constructor(private registrationService: RegistrationService,  private router: Router) { }
  
  ngOnInit(): void {
   // throw new Error('Method not implemented.');
  }

  korisnickoIme:any;
  ime:any;
  prezime:any;
  sifra:any;
  email:any;


  registerUser(user: any){
    let newUser: User = {
    username: this.korisnickoIme,
    password: this.sifra,
    email: this.email,
    name: this.ime,
    lastname: this.prezime,
    role: 0
    }
    console.log(newUser)

    this.registrationService.add(newUser).subscribe(
      (data) => {
        alert("Success!");
        this.router.navigate(["/userLogin"]);
      },
      (error: HttpErrorResponse) => {
        alert("Username or email are already taken");
      }
    );
  }

  redirect(){
    this.router.navigate(["/userLogin"]);
  }


}
