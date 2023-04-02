import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, tap } from 'rxjs';
import { loginUser } from '../model/loginUser';
import jwt_decode from 'jwt-decode';
import { Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';




@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private baseUrl = 'http://localhost:8080';

  constructor( private http: HttpClient
   ) { }
  
      private tokenKey = 'authToken'; // ključ za čuvanje tokena u local storage-u
    
      
      login(user: loginUser): Observable<any> {
        return this.http.post<any>( `${this.baseUrl}/login`,user);
      }
    
      saveToken(token: string): void {
        localStorage.setItem(this.tokenKey, token);
      }
    
      getToken(): string | null {
        return localStorage.getItem(this.tokenKey);
      }
      removeToken(): void {
        localStorage.removeItem(this.tokenKey);
      }

      getUserRole(): string {
        const token= this.getToken();
        if (token!=null) {
          const payload: any = jwt_decode(token);
          const userRole: string = payload.role;
          //console.log(userRole); 
          return userRole;
        } else {
          return '';
        }
      }
    }
    





