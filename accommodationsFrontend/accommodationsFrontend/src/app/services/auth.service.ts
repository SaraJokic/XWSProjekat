import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { loginDto } from '../model/loginDto';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private baseUrl = 'http://localhost:8000';

  constructor( private http: HttpClient
   ) { }
  
      private tokenKey = 'authToken'; 
    
      
      login(user: loginDto): Observable<any> {
        return this.http.post<any>( `${this.baseUrl}/auth/login`,user);
      }
    
      saveToken(token: string): void {
        localStorage.setItem(this.tokenKey, token);
      }

      removeToken(): void {
        localStorage.removeItem(this.tokenKey);
      }

}
