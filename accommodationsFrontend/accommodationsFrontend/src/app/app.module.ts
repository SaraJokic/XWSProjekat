import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './common/navbar/navbar.component';
import { MaterialModule } from './material/material.module';
import { NewAccommodationComponent } from './pages/new-accommodation/new-accommodation.component';
import { HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { RegistrationFormComponent } from './pages/registration-form/registration-form.component';
import { UserProfileComponent } from './pages/user-profile/user-profile.component';
import { LoginFormComponent } from './pages/login-form/login-form.component';
import { MyAccommodationsComponent } from './pages/my-accommodations/my-accommodations.component';



@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    NewAccommodationComponent,
    RegistrationFormComponent,
    UserProfileComponent,
    LoginFormComponent,
    MyAccommodationsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MaterialModule,
    HttpClientModule,
    ReactiveFormsModule
    
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
