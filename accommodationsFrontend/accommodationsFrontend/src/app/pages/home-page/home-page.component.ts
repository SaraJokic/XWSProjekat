import { Component } from '@angular/core';
import { Accommodation } from 'src/app/model/accommodation';
import { SearchAccommodation } from 'src/app/model/search-accomodation';
import { AccommodationServiceService } from 'src/app/services/accommodation-service.service';
import { DialogService } from 'src/app/services/dialog.service';
import { MaterialModule } from 'src/app/material/material.module';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent {

search(searchArg: any) {
  const searchQuery: SearchAccommodation = {
    location: searchArg.location,
    guests: searchArg.guests,
    start_date: searchArg.sdate,
    end_date: searchArg.sdate
  };
  console.log("ovako izgleda search dto ",searchQuery)
  this.accommodationService.Search(searchQuery).subscribe(
    (accommodations) => {
      this.accommodations = accommodations.acc;
      console.log(accommodations)
    },
    (error: any) => {
      console.error(error);
    }
  );
  
}
  accommodations: Accommodation[] = [];

  constructor(private accommodationService: AccommodationServiceService, 
    private dialogService: DialogService) { }

  ngOnInit() {
    this.getAccommodations();
  }

  getAccommodations() {
    this.accommodationService.getAll().subscribe(
      (accommodations) => {
        this.accommodations = []
        this.accommodations = accommodations.acc;
        console.log(accommodations)
      },
      (error: any) => {
        console.error(error);
      }
    );
  }
  openDialog(accommodation: Accommodation): void {
    this.dialogService.openDialogReservation(accommodation);
  }
}
