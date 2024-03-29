import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AllFlightsComponent } from './components/all-flights/all-flights.component';
import { AddFlightsComponent } from './components/add-flights/add-flights.component';
import { RegistrationFormComponent } from './registration/components/registration-form/registration-form.component';
import { UpdateFlightsComponent } from './components/update-flights/update-flights.component';
import { MyTicketsComponent } from './components/my-tickets/my-tickets.component';
import { LoginFormComponent } from './registration/components/login-form/login-form.component';
import { WelcomePageComponent } from './components/welcome-page/welcome-page.component';
import { AuthGuardService } from './registration/services/auth-guard.service';
import { LogoutComponent } from './components/logout/logout.component';
import { AuthGuardAdminService } from './registration/services/auth-guard-admin.service';


const routes: Routes = [
  {path: '', component: WelcomePageComponent },
  {path: 'flights', component: AllFlightsComponent},
  {path: 'flights/add', component: AddFlightsComponent ,canActivate :[AuthGuardAdminService] },
  { path: 'userRegister', component: RegistrationFormComponent },
  { path: 'userLogin', component:  LoginFormComponent },
  { path: 'userLogout', component:  LogoutComponent, canActivate :[AuthGuardService] },
  {path: 'mytickets', component: MyTicketsComponent, canActivate :[AuthGuardService]},
  { path: 'flights/:id', component: UpdateFlightsComponent ,canActivate :[AuthGuardAdminService] },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
