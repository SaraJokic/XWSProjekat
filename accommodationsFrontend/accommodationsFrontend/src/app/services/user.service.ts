import { ObserversModule } from '@angular/cdk/observers';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from '../model/user';
import { UpdateUserReq } from '../model/update-user-req';

const baseUrl = 'http://localhost:8000';
const url = 'http://localhost:8000/users/delete';
@Injectable({
  providedIn: 'root'
})
export class UserService {
  
  constructor(private http: HttpClient) { }
  getAll(): Observable<any>{
    return this.http.get<any>('http://localhost:8000/users'); 
  }
  updateUser(request: UpdateUserReq): Observable<any>{
    return this.http.put<any>('http://localhost:8000/users/update', request)
  }
  delete(id: any): Observable<any> {
    return this.http.delete(`${url}/${id}`);
  }
  getUserByUsername(username: any): Observable<any>{
    return this.http.get<any>(`${baseUrl}/users/getusername/${username}`); 
  }
  getUserById(id: any): Observable<any>{
    return this.http.get<any>(`${baseUrl}/users/get/${id}`); 
  }
}
