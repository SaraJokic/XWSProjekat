import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RegistrationFormComponent } from './components/registration-form/registration-form.component';
import { MaterialModule } from '../material/material.module';




@NgModule({
  declarations: [
    RegistrationFormComponent,
  ],
  imports: [
    CommonModule,
    MaterialModule
    
    
  ]
})
export class RegistrationModule { }
