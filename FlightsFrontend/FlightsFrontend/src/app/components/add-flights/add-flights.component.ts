import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
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
    ticketprice?:number=0;
    numofseats:number=0;

    


    constructor(private flightService: FlightService, private router: Router) { }

    saveFlight(): void {
      const data = {
        fromplace: this.fromplace,
        toplace: this.toplace,
        starttime: this.starttime?.toISOString,
        endtime:this.endtime?.toDateString,
        ticketprice: this.ticketprice,
        numofseats: this.numofseats,
          
      };
    
      
      this.flightService.create(data)
        .subscribe({
          next: (res) => {
            console.log(res);
            console.log(this.starttime);
            //console.log("kreirao");
            
            this.router.navigate(["/flights"]); 
          },
          error: (e) => console.error(e)
        });
    }


}
