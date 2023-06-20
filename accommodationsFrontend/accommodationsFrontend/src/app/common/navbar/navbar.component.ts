import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit{
  userRole: any
  role: any 
  constructor(private authService: AuthService, private router: Router) {
    this.router.events.subscribe(() => {
      if(localStorage['authToken'] != null){
        this.userRole = this.authService.getUserRole();
        
        console.log("Navbar je namesten za ", this.userRole);
      }
    });
    }
  ngOnInit(): void {
    this.userRole = 'ROLE_NOTAUTH';
    if(localStorage['authToken'] != null){
      this.userRole = this.authService.getUserRole();
      //console.log("Navbar je namesten za ", this.userRole);
      this.role = this.userRole;
      //console.log(this.role === 1)
    }
  }
   
    logout(){
      this.authService.removeToken();
      this.userRole = 'ROLE_NOTAUTH';
      this.router.navigate(['/login'])
    }

 
}



