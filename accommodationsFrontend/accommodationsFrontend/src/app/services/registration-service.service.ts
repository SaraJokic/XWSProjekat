import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { User } from '../model/user';
import { RegisterUserRequest } from '../model/register-user-request';
import { Observable } from 'rxjs';

const baseUrl = 'http://localhost:8000';
@Injectable({
  providedIn: 'root'
})
export class RegistrationServiceService {

  constructor(private http: HttpClient) { }
  
  add(newUser: RegisterUserRequest): Observable<any>{
    return this.http.post<User>(baseUrl + '/users/register', newUser);
  }
}
