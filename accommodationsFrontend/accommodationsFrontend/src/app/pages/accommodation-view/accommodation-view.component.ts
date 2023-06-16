import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormArray, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation';
import { AccommodationAvailability, AvailabilitySlot, PriceChange } from 'src/app/model/accommodation-availability';
import { AccommodationServiceService } from 'src/app/services/accommodation-service.service';
import { AvailabilityServiceService } from 'src/app/services/availability-service.service';

@Component({
  selector: 'app-accommodation-view',
  templateUrl: './accommodation-view.component.html',
  styleUrls: ['./accommodation-view.component.css']
})
export class AccommodationViewComponent implements OnInit {
  accommodation: Accommodation | undefined;
  accommodationForm: FormGroup;
  availability: AccommodationAvailability | undefined
  selectedImages: File[] = [];
  accImages: string[] = [];
  slots: AvailabilitySlot[] = [];
  pricechanges: PriceChange[] = [];

  constructor(private formBuilder: FormBuilder, private route: ActivatedRoute, 
    private accommodationService: AccommodationServiceService, private availabilityService: AvailabilityServiceService) {
      this.accommodationForm = this.formBuilder.group({
        name: ['', Validators.required],
        location: ['', Validators.required],
        description: ['', Validators.required],
        automaticApprove: false,
        benefits: this.formBuilder.group({
          wifi: false,
          freeParking: false,
          kitchen: false,
        }),
        pictures: [[]],
        minguests: [0, Validators.required],
        maxguests: [0, Validators.required],
        availability: this.formBuilder.group({
          accommodationId: ['', Validators.required],
          availableSlots: this.formBuilder.array([]),
          price: [0, Validators.required],
          isPricePerGuest: [false, Validators.required],
          changePrice: this.formBuilder.array([]),
        }),
      });
     }

  ngOnInit() {
    const id = this.route.snapshot.params['id'];
    console.log(id)
    this.accommodationService.getAccommodationById(id).subscribe(
      (data) => {
        this.accommodation = data.acc;
        console.log("Accommodation:", this.accommodation)
        this.getAvailabilityByAccId()
        
      },
      (error: HttpErrorResponse) => {
        //alert("Username or email are already taken");
      }
    );
  }
  getAvailabilityByAccId(){
    this.availabilityService.getAvailabilityByAccId(this.route.snapshot.params['id']).subscribe(
      (data) => {
        this.availability = data.availability;
        console.log("Availability:", this.availability)
        this.populateForm()
      },
      (error: HttpErrorResponse) => {
        //alert("Username or email are already taken");
      }
    );
  }
  populateForm(): void {
    this.accommodationForm?.patchValue({
      name: this.accommodation?.name,
      location: this.accommodation?.location,
      description: this.accommodation?.description,
      minguests: this.accommodation?.minGuests,
      maxguests: this.accommodation?.maxGuests,
      pictures: this.accommodation?.pictures,
      automaticApprove: this.accommodation?.automaticApprove,
      benefits: {
        wifi: this.accommodation?.benefits?.wifi,
        freeParking: this.accommodation?.benefits?.freeParking,
        kitchen: this.accommodation?.benefits?.kitchen,
      },
      availability: {
        accommodationId: this.route.snapshot.params['id'],
        price: this.availability?.price,
        isPricePerGuest: this.availability?.isPricePerPerson,
      },
    });
    this.populateAvailableSlots();
    this.populateChangePrice();
  }
  populateAvailableSlots(): void {
    const availableSlotsArray = this.accommodationForm?.get('availability.availableSlots') as FormArray;
    const slots = this.availability?.availableSlots as any[];
    if(slots){
      slots.forEach((slot) => {
        const slotGroup = this.formBuilder.group({
          startDate: slot.startDate.substring(0, 10),
          endDate: slot.endDate.substring(0, 10),
        });
        availableSlotsArray.push(slotGroup);
      });
    }
    
  }
  
  populateChangePrice(): void {
    const changePriceArray = this.accommodationForm?.get('availability.changePrice') as FormArray;
    const changes = this.availability?.changePrice as any[]
    if(changes){
      changes.forEach((change) => {
        const changeGroup = this.formBuilder.group({
          startDate: change.startdate.substring(0, 10),
          endDate: change.enddate.substring(0, 10),
          change: change.change,
        });
        changePriceArray.push(changeGroup);
      });
    }
    
  }
  get availabilitySlots() {
    return this.accommodationForm?.get('availability.availableSlots') as FormArray;
  }
  get changePrice(){
    return this.accommodationForm?.get('availability.changePrice') as FormArray
  }
  addAvailabilitySlot() {
    const availableSlots = this.accommodationForm?.get('availability.availableSlots') as FormArray;
    availableSlots.push(
      this.formBuilder.group({
        startDate: ['', Validators.required],
        endDate: ['', Validators.required],
      })
    );
  }

  removeAvailabilitySlot(index: number) {
    const availableSlots = this.accommodationForm?.get('availability.availableSlots') as FormArray;
    availableSlots.removeAt(index);
  }
    
  removePriceChange(index: number) {
    const changePrice = this.accommodationForm?.get('availability.changePrice') as FormArray;
    changePrice.removeAt(index);
  }
  
  addPriceChange() {
    const changePrice = this.accommodationForm?.get('availability.changePrice') as FormArray;
    changePrice.push(this.formBuilder.group({
      startDate: ['', Validators.required],
      endDate: ['', Validators.required],
      change: ['', Validators.required]
    }));
  }
  onSubmit(){
    const accommodationNew: Accommodation = {
      id: this.accommodation?.id,
      name: this.accommodationForm?.get('name')?.value,
      location: this.accommodationForm?.get('location')?.value,
      benefits: this.accommodationForm?.get('benefits')?.value,
      minGuests: this.accommodationForm?.get('minguests')?.value,
      maxGuests: this.accommodationForm?.get('maxguests')?.value,
      description: this.accommodationForm?.get('description')?.value,
      automaticApprove: this.accommodationForm.get('automaticApprove')?.value,
      pictures: this.accImages,
      hostId: this.accommodation?.hostId || "",
    };
    const availabilityFormData: AccommodationAvailability = {
      id: this.availability?.id,
      accommodationId: this.accommodation?.id || "",
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
    if (this.accommodationForm?.get('availability.availableSlots')?.value.length === 0){
      availabilityFormData.availableSlots = this.slots;
    }
    if (this.accommodationForm?.get('availability.changePrice')?.value.length === 0){
      availabilityFormData.changePrice = this.pricechanges;
    }
    console.log("OVO JE STA SE UPDATE-UJE: ", accommodationNew)
    console.log("OVO JE STA SE UPDATE-UJE: ", availabilityFormData)
    this.accommodationService.updateAccommodation(this.route.snapshot.params['id'], accommodationNew).subscribe(
      (data) => {
        console.log("posle accommodation update-a: ", data)
        this.availabilityService.updateAvailability(this.availability?.id ||"", availabilityFormData).subscribe(
          (response) => {
            console.log("Posle availability update-a ", response)
          },
          (error: HttpErrorResponse) => {
            //alert("Username or email are already taken");
          }
        );
      },
      (error: HttpErrorResponse) => {
        //alert("Username or email are already taken");
      }
    );
  }
  handleImageUpload(event: Event): void {
    const inputElement = event.target as HTMLInputElement;
    this.selectedImages = Array.from(inputElement.files || []);
    for(let i = 0; i < this.selectedImages.length; i++){
      this.accImages.push(this.selectedImages[i].name)
    }
    console.log("selektovane slike: ", this.accImages)
  }
}

