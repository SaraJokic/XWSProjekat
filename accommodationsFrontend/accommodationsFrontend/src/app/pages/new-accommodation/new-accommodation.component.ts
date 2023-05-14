import { Component, OnInit } from '@angular/core';
import { FormArray, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Accommodation } from 'src/app/model/accommodation';
import { AccommodationAvailability, AvailabilitySlot, PriceChange } from 'src/app/model/accommodation-availability';
import { CreateAccommodationDto } from 'src/app/model/create-accommodation-dto';
import { AccommodationServiceService } from 'src/app/services/accommodation-service.service';
import { AvailabilityServiceService } from 'src/app/services/availability-service.service';

@Component({
  selector: 'app-new-accommodation',
  templateUrl: './new-accommodation.component.html',
  styleUrls: ['./new-accommodation.component.css']
})
export class NewAccommodationComponent implements OnInit {
  accommodationForm: FormGroup;
  newaccommodationID: string | undefined;
  
  constructor(private accommodationService: AccommodationServiceService, private fb: FormBuilder
    , private availabilityService: AvailabilityServiceService) {
    this.accommodationForm = this.fb.group({
      name: ['', Validators.required],
      location: ['', Validators.required],
      benefits: this.fb.group({
        wifi: false,
        freeParking: false,
        kitchen: false,
      }),
      minguests: [0, Validators.required],
      maxguests: [0, Validators.required],
      availability: this.fb.group({
        accommodationId: ['', Validators.required],
        availableSlots: this.fb.array([]),
        price: [0, Validators.required],
        isPricePerGuest: [false, Validators.required],
        changePrice: this.fb.array([]),
      }),
    });
   }

  ngOnInit(): void {
  }
  onSubmit() {
    const accommodation: Accommodation = {
      name: this.accommodationForm?.get('name')?.value,
      location: this.accommodationForm?.get('location')?.value,
      benefits: this.accommodationForm?.get('benefits')?.value,
      minGuests: this.accommodationForm?.get('minguests')?.value, 
      maxGuests: this.accommodationForm?.get('maxguests')?.value, 
    };
    const availabilityFormData: AccommodationAvailability = {
      accommodationId: "",
      availableSlots: this.accommodationForm?.get('availability.availableSlots')?.value.map((slot: any) => ({
        start_date: slot.startDate + 'T10:00:00Z',
        end_date: slot.endDate+ 'T10:00:00Z',
      })),
      price: this.accommodationForm?.get('availability.price')?.value,
      isPricePerPerson: this.accommodationForm?.get('availability.isPricePerGuest')?.value,
      changePrice: this.accommodationForm?.get('availability.changePrice')?.value.map((change: any) => ({
        startdate: change.startDate+ 'T10:00:00Z',
        enddate: change.endDate+ 'T10:00:00Z',
        change: change.change,
      })),
    };
    console.log(accommodation)
    console.log(availabilityFormData)
    console.log(this.accommodationForm.get('availability.changePrice')?.value);
      this.accommodationService.createAccommodation(accommodation).subscribe(
        response => {
          console.log(response);
          this.newaccommodationID = response.id;
        },
        error => {
          console.log(error);
        }
    );
    availabilityFormData.accommodationId = this.newaccommodationID || " ";
    this.availabilityService.createNewAvailability(availabilityFormData).subscribe(
      response => {
        console.log(response);
      },
      error => {
        console.log(error);
      }
  );
  }
  get availabilitySlots() {
    return this.accommodationForm.get('availability.availableSlots') as FormArray;
  }
  get changePrice(){
    return this.accommodationForm.get('availability.changePrice') as FormArray
  }

  addAvailabilitySlot() {
    const availableSlots = this.accommodationForm.get('availability.availableSlots') as FormArray;
    availableSlots.push(
      this.fb.group({
        startDate: ['', Validators.required],
        endDate: ['', Validators.required],
      })
    );
  }

  removeAvailabilitySlot(index: number) {
    const availableSlots = this.accommodationForm.get('availability.availableSlots') as FormArray;
    availableSlots.removeAt(index);
  }
    
  removePriceChange(index: number) {
    const changePrice = this.accommodationForm.get('availability.changePrice') as FormArray;
    changePrice.removeAt(index);
  }
  
  addPriceChange() {
    const changePrice = this.accommodationForm.get('availability.changePrice') as FormArray;
    changePrice.push(this.fb.group({
      startDate: ['', Validators.required],
      endDate: ['', Validators.required],
      change: ['', Validators.required]
    }));
  }
    

}
