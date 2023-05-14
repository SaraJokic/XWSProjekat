import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { NewAccommodationComponent } from './pages/new-accommodation/new-accommodation.component';

const routes: Routes = [
  {path: 'new/accommodation', component: NewAccommodationComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
