import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { NewAccommodationComponent } from './pages/new-accommodation/new-accommodation.component';
import { RegistrationFormComponent } from './pages/registration-form/registration-form.component';
import { LoginFormComponent } from './pages/login-form/login-form.component';
import { UserProfileComponent } from './pages/user-profile/user-profile.component';

const routes: Routes = [
  {path: 'new/accommodation', component: NewAccommodationComponent },
  {path: 'register', component: RegistrationFormComponent },
  {path: 'login', component: LoginFormComponent },
  {path: 'myprofile', component: UserProfileComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
