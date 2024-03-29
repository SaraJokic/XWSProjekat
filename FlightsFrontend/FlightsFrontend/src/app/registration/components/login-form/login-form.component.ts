import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { HttpErrorResponse } from '@angular/common/http';
import { Router } from '@angular/router';
import { loginDto } from '../../model/loginDto';

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
          this.router.navigate(["/flights"]); 
        },
        (error: HttpErrorResponse) => {
          alert("Incorect username or password");
        });
    }
   
    
    redirect(){
      this.router.navigate(["/userRegister"]);
    }

  }

