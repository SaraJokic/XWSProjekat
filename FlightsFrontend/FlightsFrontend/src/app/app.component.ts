import { Component, OnInit } from '@angular/core';
import { AuthService } from './registration/services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{

  title = 'FlightsFrontend';
  userRole: any
  role: any 
  constructor(private authService: AuthService, private router: Router) {
    this.router.events.subscribe(() => {
      if(localStorage['authToken'] != null){
        this.userRole = this.authService.getUserRole();
        
        //console.log("Navbar je namesten za ", this.userRole);
      }
    });
   }
  ngOnInit(): void {
    this.userRole = 'ROLE_NOTAUTH';
    if(localStorage['authToken'] != null){
      this.userRole = this.authService.getUserRole();
      //console.log("Navbar je namesten za ", this.userRole);
      console.log(typeof this.userRole)
      //console.log(this.userRole === '1')
      this.role = this.userRole;
      console.log(this.role === 1)
    }
  }

  logout(){
    this.authService.removeToken();
    this.userRole = 'ROLE_NOTAUTH';
    this.router.navigate(['/userLogin'])
  }
}
