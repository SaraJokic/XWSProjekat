import { Injectable } from '@angular/core';
import { AuthService } from './auth.service';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthGuardAdminService {

  constructor(private authService: AuthService, private router: Router) {}

  canActivate(): boolean {
    const userRole = this.authService.getUserRole();
    console.log("ovo je uloga u authadminservice",userRole)
    if (!this.authService.getToken() || userRole != '1') {
      console.log("niste admin")
      this.router.navigate(['/userLogout']);
      return false;
    }
  
    return true;
  }
}
