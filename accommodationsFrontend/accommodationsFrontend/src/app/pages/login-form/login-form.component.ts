import { HttpErrorResponse } from '@angular/common/http';
import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { loginDto } from 'src/app/model/loginDto';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.css']
})
export class LoginFormComponent {

  
  constructor(private authService: AuthService, private router: Router) { }
  
  ngOnInit(): void {
    
  }
 
  ime:any;
  sifra:any;
  

    loginUser(user: any): void {
     
        let loginUser: loginDto = {
        username: this.ime,
        password: this.sifra
        }
      this.authService.login(loginUser)
        .subscribe(response => {
          this.authService.saveToken(response.token);
          alert("Welcome back! You have successfully logged in.")
          this.router.navigate(["/new/accommodation"]); 
        },
        (error: HttpErrorResponse) => {
          alert("Incorect username or password");
        });
    }
   
    
    redirect(){
      this.router.navigate(["/userRegister"]);
    }

  }


