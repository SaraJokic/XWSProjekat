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
        console.log(loginUser)
      this.authService.login(loginUser)
        .subscribe(response => {
          this.authService.saveToken(response.token);
          alert("Welcome back! You have successfully logged in.")
          this.router.navigate([""]); 
        },
        (error: HttpErrorResponse) => {
          alert(error.error.message);
        });
    }
   
    
    redirect(){
      this.router.navigate(["/userRegister"]);
    }

  }


