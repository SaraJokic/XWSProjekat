import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, tap } from 'rxjs';
import jwt_decode from 'jwt-decode';
import { Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import { loginDto } from '../model/loginDto';
import { logedUserInfo } from '../model/logedUserInfo';
import { NonNullAssert } from '@angular/compiler';




@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private baseUrl = 'http://localhost:8080';

  constructor( private http: HttpClient
   ) { }
  
      private tokenKey = 'authToken'; 
    
      
      login(user: loginDto): Observable<any> {
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
          return userRole;
        } else {
          return '';
        }
      }
      getLogedUserInfo(): logedUserInfo | null {
        const token= this.getToken();
        
        if (token!=null) {
          const payload: any = jwt_decode(token);
          const Name : string = payload.name;
          const Username: string = payload.username;
          const Role: string = payload.role;
          let logedUserInfo: logedUserInfo = {
            name: Name,
            username: Username,
            role: Role
            }
            console.log(logedUserInfo)
          return logedUserInfo;
        } else {
          return null;
        }
      }
    }
    





