import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MatLegacyTableModule as MatTableModule } from '@angular/material/legacy-table';
import { MatSortModule} from '@angular/material/sort';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { AllFlightsComponent } from './components/all-flights/all-flights.component';
import { AddFlightsComponent } from './components/add-flights/add-flights.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import {MatTooltipModule} from '@angular/material/tooltip';
import { MatIconModule} from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { NavbarComponent } from './common/navbar/navbar.component';
import { RegistrationModule } from './registration/registration.module';
import { MaterialModule } from './material/material.module';
import { UpdateFlightsComponent } from './components/update-flights/update-flights.component';
import { MyTicketsComponent } from './components/my-tickets/my-tickets.component';
import { MatCardModule } from '@angular/material/card';
import { BuyTicketDialogComponent } from './components/buy-ticket-dialog/buy-ticket-dialog.component';
import { MatDialogModule } from '@angular/material/dialog';
import { TicketDetailsDialogComponent } from './components/ticket-details-dialog/ticket-details-dialog.component';
import { AuthInterceptorService } from './registration/services/auth-interceptor.service';
import { WelcomePageComponent } from './components/welcome-page/welcome-page.component';
import { LogoutComponent } from './components/logout/logout.component';

@NgModule({
  declarations: [
    AppComponent,
    AllFlightsComponent,
    AddFlightsComponent,
    NavbarComponent,
    UpdateFlightsComponent,
    MyTicketsComponent,
    BuyTicketDialogComponent,
    TicketDetailsDialogComponent,
    WelcomePageComponent,
    LogoutComponent,
  ],
  imports: [
    BrowserModule,
    MaterialModule,
    AppRoutingModule,
    MatSortModule,
    HttpClientModule,
    FormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatDatepickerModule,
    MatNativeDateModule,
    MatTooltipModule,
    MatIconModule,
    MatButtonModule,
    RegistrationModule,
    BrowserAnimationsModule,
    MatCardModule,
    MatDialogModule,
  ],
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: AuthInterceptorService, multi: true }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
