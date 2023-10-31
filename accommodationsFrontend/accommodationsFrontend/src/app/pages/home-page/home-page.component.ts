import { Component, Input, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { Observable, map } from 'rxjs';
import { Accommodation } from 'src/app/model/accommodation';
import { SearchAccommodation } from 'src/app/model/search-accomodation';
import { AccommodationServiceService } from 'src/app/services/accommodation-service.service';
import { DialogService } from 'src/app/services/dialog.service';
import { UserService } from 'src/app/services/user.service';
import { MaterialModule } from 'src/app/material/material.module';
import { OwlOptions } from 'ngx-owl-carousel-o';



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
  showHint = false;

  PROMINENT: boolean = false;

  obojiProminent : boolean =false;

  accommodations: Accommodation[] = [];
  filteredAccommodations!: Accommodation[];
  
  customOptions: OwlOptions = {
    loop: true,  
    mouseDrag: true,  
     dots: false,  
    navSpeed: 700,  
     items:1,
     nav: true,
     navText: ['<', '>']
    
  }
  

  constructor(
    private accommodationService: AccommodationServiceService,private dialogService: DialogService, private userService:UserService) {
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
          id: accommodation.id || `accommodation-${index}` 
        }));
  
        this.accommodationPrices = this.accommodations.reduce((obj: { [id: string]: number }, accommodation, index) => {
          obj[accommodation.id!] = (index + 1) * 100; 
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

search(formData: any): void {
  this.filteredAccommodations = this.accommodations.filter(accommodation => {
    let odgovaraLokacija = true;
    let odgovaraBrojGostiju = true;
    let odgovaraDatum = true;

    // Provera lokacije
    if (formData.location) {
      odgovaraLokacija = accommodation.location.toLowerCase() === formData.location.toLowerCase();
    }

    // Provera broja gostiju
    if (formData.guests) {
      odgovaraBrojGostiju = accommodation.maxGuests >= formData.guests;
    }

    // Provera datuma (samo ako su oba datuma uneta)
   // if (formData.sdate && formData.edate) {
     // const dostupnoOd = new Date(accommodation.start_date);
      //const dostupnoDo = new Date(accommodation.end_date);

      //const trazeniPocetniDatum = new Date(formData.sdate);
      //const trazeniKrajnjiDatum = new Date(formData.edate);

      //odgovaraDatum = trazeniPocetniDatum >= dostupnoOd && trazeniKrajnjiDatum <= dostupnoDo;
    //}

   // return odgovaraLokacija && odgovaraBrojGostiju && odgovaraDatum;
    return odgovaraLokacija && odgovaraBrojGostiju;
  });
}

  /*search(searchArg: any) {
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
  }*/

  openDialog(accommodation: Accommodation): void {
    this.dialogService.openDialogReservation(accommodation);
  }
  openHostDialog(accommodation: Accommodation): void {
    this.dialogService.openHostDialog(accommodation);
  }
  openAccDialog(accommodation: Accommodation): void {
    this.dialogService.openAccDialog(accommodation);
  }
  openViewAccDialog(accommodation: Accommodation): void {
    //this.dialogService.openViewAccDialog(accommodation);
  }
  onlyTrue(benefits: any): boolean {
    if (!benefits) {
      return false;
    }
    return Object.values(benefits).includes(true);
  }

  Nadjen:any;

  
  applyFilters() {
    if (this.PROMINENT) {
      this.FindProminent();
    } else {
      this.filteredAccommodations = this.accommodations.filter(accommodation => {
        const price = this.accommodationPrices[accommodation.id || ''];
        const isWithinPriceRange = !this.selectedPriceValue || (price >= this.selectedPriceValue.min && price <= this.selectedPriceValue.max);
  
        const isWifiFiltered = !this.WIFI || accommodation.benefits?.wifi;
        const isKitchenFiltered = !this.KITCHEN || accommodation.benefits?.kitchen;
        const isFreeParkingFiltered = !this.FREEPARKING || accommodation.benefits?.freeParking;
  
        return isWithinPriceRange && isWifiFiltered && isKitchenFiltered && isFreeParkingFiltered;
      });
    }
  }
  
  FindProminent() {
    this.accommodationService.GetAllProminentAccommodation().subscribe(
      (accommodations) => {
        this.filteredAccommodations = accommodations.acc; // Assuming the accommodations are returned in the `acc` property
        console.log(this.filteredAccommodations);
        this.filteredAccommodations = this.filteredAccommodations.filter(accommodation => {
          const price = this.accommodationPrices[accommodation.id || ''];
          const isWithinPriceRange = !this.selectedPriceValue || (price >= this.selectedPriceValue.min && price <= this.selectedPriceValue.max);
  
          const isWifiFiltered = !this.WIFI || accommodation.benefits?.wifi;
          const isKitchenFiltered = !this.KITCHEN || accommodation.benefits?.kitchen;
          const isFreeParkingFiltered = !this.FREEPARKING || accommodation.benefits?.freeParking;
  
          return isWithinPriceRange && isWifiFiltered && isKitchenFiltered && isFreeParkingFiltered;
        });
      },
      (error: any) => {
        console.error(error);
      }
    );
  }

  isProminent(accommodation: any): boolean {
    let isProminent = false;
    this.accommodationService.GetAllProminentAccommodation().subscribe((prominentAccommodations: any[]) => {
      isProminent = prominentAccommodations.some((prominentAccommodation: any) => {
        return prominentAccommodation.id === accommodation.id;
      });
    });
    return isProminent;
  }
  
  
}