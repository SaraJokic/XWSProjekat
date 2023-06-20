import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { loginDto } from '../model/loginDto';
import jwt_decode from 'jwt-decode';
import { LoggedUserInfo } from '../model/logged-user-info';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private baseUrl = 'http://localhost:8000';

  constructor( private http: HttpClient
   ) { }
  
      private tokenKey = 'authToken'; 
    
      
      login(user: loginDto): Observable<any> {
        return this.http.post<any>( `${this.baseUrl}/auth/login`,{id: "", username: user.username, password: user.password});
      }
    
      saveToken(token: string): void {
        localStorage.setItem(this.tokenKey, token);
      }

      removeToken(): void {
        localStorage.removeItem(this.tokenKey);
      }
      getLogedUserInfo(): LoggedUserInfo | null {
        const token= this.getToken();
        if (token!=null) {
          const payload: any = jwt_decode(token);
          const Name : string = payload.name;
          const Email : string = payload.email;
          const Id : string = payload.id;
          const Username: string = payload.username;
          const Role: string = payload.role;
          let logedUserInfo: LoggedUserInfo = {
            name: Name,
            username: Username,
            role: Role,
            id: Id,
            email:Email,
            }
            console.log(logedUserInfo)
          return logedUserInfo;
        } else {
          return null;
        }
      }
      getToken(): string | null {
        return localStorage.getItem(this.tokenKey);
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

}
