import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { NewAccommodationComponent } from './pages/new-accommodation/new-accommodation.component';
import { RegistrationFormComponent } from './pages/registration-form/registration-form.component';
import { LoginFormComponent } from './pages/login-form/login-form.component';
import { UserProfileComponent } from './pages/user-profile/user-profile.component';
import { MyAccommodationsComponent } from './pages/my-accommodations/my-accommodations.component';
import { AccommodationViewComponent } from './pages/accommodation-view/accommodation-view.component';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { ReservationRequestsHostComponent } from './pages/reservation-requests-host/reservation-requests-host.component';
import { ReservationRequestsGuestComponent } from './pages/reservation-requests-guest/reservation-requests-guest.component';

const routes: Routes = [
  {path: 'new/accommodation', component: NewAccommodationComponent },
  {path: 'register', component: RegistrationFormComponent },
  {path: 'login', component: LoginFormComponent },
  {path: 'myprofile', component: UserProfileComponent },
  {path: 'myaccommodations', component: MyAccommodationsComponent },
  {path: 'accommodationview/:id', component: AccommodationViewComponent },
  {path: '', component: HomePageComponent },
  {path: 'accommodationrequests/:id', component: ReservationRequestsHostComponent },
  {path: 'myreservationrequests', component: ReservationRequestsGuestComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
