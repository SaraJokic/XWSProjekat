import { Component, Input, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { Accommodation } from 'src/app/model/accommodation';
import { SearchAccommodation } from 'src/app/model/search-accomodation';
import { AccommodationServiceService } from 'src/app/services/accommodation-service.service';
import { DialogService } from 'src/app/services/dialog.service';
@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})

export class HomePageComponent implements OnInit {

  selectedPriceValue: any | undefined;
  WIFI: boolean = false;
  KITCHEN: boolean = false;
  FREEPARKING: boolean = false;

  accommodations: Accommodation[] = [];
  filteredAccommodations!: Accommodation[];

  constructor(
    private accommodationService: AccommodationServiceService,private dialogService: DialogService) {
      this.selectedPriceValue = { min: 1, max: 1000000 };
    }

  accommodationPrices: { [id: string]: number } = {};

  ngOnInit() {
    this.getAccommodations();
  }


  getAccommodations() {
    this.accommodationService.getAll().subscribe(
      (accommodations) => {
        this.accommodations = accommodations.acc.map((accommodation: Accommodation, index: number) => ({
          ...accommodation,
          id: accommodation.id || `accommodation-${index}` // Assign a default id if undefined
        }));
  
        this.accommodationPrices = this.accommodations.reduce((obj: { [id: string]: number }, accommodation, index) => {
          obj[accommodation.id!] = (index + 1) * 100; // Generate a new value for selectedPrice
          return obj;
        }, {});
  
        this.filteredAccommodations = this.accommodations;
        this.applyFilters();
        console.log(accommodations);
      },
      (error: any) => {
        console.error(error);
      }
    );
  }

  search(searchArg: any) {
    const searchQuery: SearchAccommodation = {
      location: searchArg.location,
      guests: searchArg.guests,
      start_date: searchArg.sdate,
      end_date: searchArg.edate 
    };
    console.log("ovako izgleda search dto ", searchQuery);
    this.accommodationService.Search(searchQuery).subscribe(
      (accommodations) => {
        this.accommodations = accommodations.acc;
        console.log(accommodations);
      },
      (error: any) => {
        console.error(error);
      }
    );
  }

  openDialog(accommodation: Accommodation): void {
    this.dialogService.openDialogReservation(accommodation);
  }

  onlyTrue(benefits: any): boolean {
    if (!benefits) {
      return false;
    }
    return Object.values(benefits).includes(true);
  }

  applyFilters() {
    this.filteredAccommodations = this.accommodations.filter(accommodation => {
      const price = this.accommodationPrices[accommodation.id || ''];
      const isWithinPriceRange =!this.selectedPriceValue || (price >= this.selectedPriceValue.min && price <= this.selectedPriceValue.max);

      console.log("Cena od " + this.selectedPriceValue.min + " do " + this.selectedPriceValue.max)
      
      return (
        isWithinPriceRange &&
        (!this.WIFI || accommodation.benefits?.wifi) &&
        (!this.KITCHEN || accommodation.benefits?.kitchen) &&
        (!this.FREEPARKING || accommodation.benefits?.freeParking)
      );
    });
  }

  
}
