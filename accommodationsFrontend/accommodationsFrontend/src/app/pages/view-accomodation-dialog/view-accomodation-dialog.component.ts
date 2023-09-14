import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';
import { OwlOptions } from 'ngx-owl-carousel-o';
import { Accommodation } from 'src/app/model/accommodation';

@Component({
  selector: 'app-view-accomodation-dialog',
  templateUrl: './view-accomodation-dialog.component.html',
  styleUrls: ['./view-accomodation-dialog.component.css']
})
export class ViewAccomodationDialogComponent {

  constructor(@Inject(MAT_DIALOG_DATA) public data: Accommodation)
  {}
  showHint = false;
  customOptions: OwlOptions = {
    loop: true,  
    mouseDrag: true,  
     dots: false,  
    navSpeed: 700,  
     items:1,
     nav: true,
     navText: ['<', '>']
    
  }

}
