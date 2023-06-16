import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { RegisterUserRequest } from 'src/app/model/register-user-request';
import { User } from 'src/app/model/user';
import { RegistrationServiceService } from 'src/app/services/registration-service.service';

@Component({
  selector: 'app-registration-form',
  templateUrl: './registration-form.component.html',
  styleUrls: ['./registration-form.component.css']
})

export class RegistrationFormComponent {
  
  constructor(private registrationService: RegistrationServiceService,  private router: Router) { }
  
  ngOnInit(): void {
   // throw new Error('Method not implemented.');
  }

  korisnickoIme:any;
  ime:any;
  prezime:any;
  sifra:any;
  email:any;
  city:any;
  country:any;
  isHost: boolean = false;


  registerUser(user: any){
    let newUser: User = {
      id: "",
      Username: this.korisnickoIme,
      Password: this.sifra,
      Email: this.email,
      Name: this.ime,
      LastName: this.prezime,
      Role: 0,
      City: this.city,
      Country: this.country,
      timesCancelled: 0
    }
    console.log(typeof this.isHost)
    if(this.isHost == true){
      newUser.Role = 1;
    }
    let request: RegisterUserRequest = {
      user: newUser
    }
    console.log(newUser)

    this.registrationService.add(request).subscribe(
      (data) => {
        alert(data.message);
        this.router.navigate(["/login"]);
      },
      (error: HttpErrorResponse) => {
        alert("Username or email are already taken");
      }
    );
  }

  redirect(){
    this.router.navigate(["/login"]);
  }

}
