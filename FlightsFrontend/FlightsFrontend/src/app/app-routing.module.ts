import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AllFlightsComponent } from './components/all-flights/all-flights.component';
import { AddFlightsComponent } from './components/add-flights/add-flights.component';
import { RegistrationFormComponent } from './registration/components/registration-form/registration-form.component';
import { UpdateFlightsComponent } from './components/update-flights/update-flights.component';

const routes: Routes = [
  {path: 'userRegister', component: RegistrationFormComponent },
  {path: 'flights', component: AllFlightsComponent },
  {path: 'flights/add', component: AddFlightsComponent },
  { path: 'flights/:id', component: UpdateFlightsComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
