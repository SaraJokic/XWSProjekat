import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Flights } from 'src/models/flight.model';
import { FlightService } from 'src/services/flight.service';

@Component({
  selector: 'app-add-flights',
  templateUrl: './add-flights.component.html',
  styleUrls: ['./add-flights.component.css']
})

export class AddFlightsComponent  {

    fromplace:string='';
    toplace:string='';
    starttime?:Date;
    endtime?:Date;
    ticketprice:number=0;
    numofseats:number=0;
    totalSum:any=this.numofseats*this.ticketprice;

    


    constructor(private flightService: FlightService, private router: Router) { }

    saveFlight(): void {
      const data: Flights = {
        fromplace: this.fromplace,
        toplace: this.toplace,
        starttime: this.starttime ? new Date(this.starttime) : undefined,
        endtime:this.endtime ? new Date(this.endtime) : undefined,
        ticketprice: this.ticketprice,
        numofseats: this.numofseats,
      };
      data.starttime?.setHours(data.starttime.getHours() + 2)
      data.endtime?.setHours(data.endtime.getHours() + 2)
      
      this.flightService.create(data)
        .subscribe({
          next: (res) => {
            console.log(res);
            //console.log("kreirao");
            
            this.router.navigate(["/flights"]); 
          },
          error: (e) => console.error(e)
        });
    }


}
