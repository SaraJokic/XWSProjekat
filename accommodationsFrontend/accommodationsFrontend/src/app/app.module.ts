import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './common/navbar/navbar.component';

import { NewAccommodationComponent } from './pages/new-accommodation/new-accommodation.component';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { RegistrationFormComponent } from './pages/registration-form/registration-form.component';
import { UserProfileComponent } from './pages/user-profile/user-profile.component';
import { LoginFormComponent } from './pages/login-form/login-form.component';
import { MyAccommodationsComponent } from './pages/my-accommodations/my-accommodations.component';
import { AccommodationViewComponent } from './pages/accommodation-view/accommodation-view.component';
import { HomePageComponent } from './pages/home-page/home-page.component';
import { MatDialogModule } from '@angular/material/dialog';
import { MakeReservationDialogComponent } from './pages/make-reservation-dialog/make-reservation-dialog.component';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { ReservationRequestsHostComponent } from './pages/reservation-requests-host/reservation-requests-host.component';
import { ReservationRequestsGuestComponent } from './pages/reservation-requests-guest/reservation-requests-guest.component';
import { AuthInterceptorService } from './services/auth-interceptor.service';
import { MatButtonModule } from '@angular/material/button';
import { MaterialModule } from './material/material.module';
import { RatingDialogComponent } from './pages/rating-dialog/rating-dialog.component';
import { RatingComponent } from './pages/rating/rating.component';
import { RateAcoomodationComponent } from './pages/rate-acoomodation/rate-acoomodation.component';
import { UserRatingsComponent } from './pages/user-ratings/user-ratings.component';
import { ViewAccomodationDialogComponent } from './pages/view-accomodation-dialog/view-accomodation-dialog.component';





@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    NewAccommodationComponent,
    RegistrationFormComponent,
    UserProfileComponent,
    LoginFormComponent,
    MyAccommodationsComponent,
    AccommodationViewComponent,
    HomePageComponent,
    MakeReservationDialogComponent,
    ReservationRequestsHostComponent,
    ReservationRequestsGuestComponent,
    RatingDialogComponent,
    RatingComponent,
    RateAcoomodationComponent,
    UserRatingsComponent,
    ViewAccomodationDialogComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MaterialModule,
    HttpClientModule,
    ReactiveFormsModule,
    MatDialogModule,
    MatDatepickerModule,
    MatNativeDateModule,
    BrowserAnimationsModule,
    NoopAnimationsModule
  ],
  providers: [ { provide: HTTP_INTERCEPTORS, useClass: AuthInterceptorService, multi: true }
],
  bootstrap: [AppComponent]
})
export class AppModule { }
