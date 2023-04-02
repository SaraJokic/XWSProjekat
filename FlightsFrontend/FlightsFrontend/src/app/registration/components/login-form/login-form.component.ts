import { Component } from '@angular/core';
import { loginUser } from '../../model/loginUser';
import { AuthService } from '../../services/auth.service';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.css']
})
export class LoginFormComponent {
  constructor(private authService: AuthService) { }
  
  ngOnInit(): void {
    
  }
 
    loginUser(user: any): void {
     
        let loginUser: loginUser = {
        username: user.username,
        password: user.password1
        }
      this.authService.login(loginUser)
        .subscribe(response => {
          this.authService.saveToken(response.token);
        });
    }
   
    
  }

