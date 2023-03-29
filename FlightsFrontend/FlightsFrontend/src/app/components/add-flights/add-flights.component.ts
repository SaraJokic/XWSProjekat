import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FlightService } from 'src/services/flight.service';

@Component({
  selector: 'app-add-flights',
  templateUrl: './add-flights.component.html',
  styleUrls: ['./add-flights.component.css']
})

export class AddFlightsComponent  {

    fromPlace:string='';
    toPlace:string='';
    startTime:any;
    endTime:any;
    ticketPrice?:number=0;
    numOfSeats:number=0;


    constructor(private flightService: FlightService, private router: Router) { }

    saveFlight(): void {
      const data = {
        fromPlace: this.fromPlace,
        toPlace: this.toPlace,
        startTime: this.startTime,
        endTime:this.endTime,
        ticketPrice: this.ticketPrice,
        numOfSeats: this.numOfSeats,
          
      };
    
      this.flightService.create(data)
        .subscribe({
          next: (res) => {
            console.log(res);
            console.log("kreirao");
            //this.router.navigate(["/flights"]); 
          },
          error: (e) => console.error(e)
        });
    }


}
