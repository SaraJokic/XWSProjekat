import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Flights } from 'src/models/flight.model';
import { FlightService } from 'src/services/flight.service';

@Component({
  selector: 'app-add-flights',
  templateUrl: './add-flights.component.html',
  styleUrls: ['./add-flights.component.css']
})

export class AddFlightsComponent implements OnInit  {


    fromplace:string='';
    toplace:string='';
    starttime?:Date;
    endtime?:Date;
    ticketprice:number=0;
    numofseats:string='';


    tp:number=0;
    ns:number=0;
    
    totalSum:any=this.tp*this.ns;

    

racunaj(a: any, b:any){
  return this.totalSum = a*b;
}

    constructor(private flightService: FlightService, private router: Router) { }
  ngOnInit(): void {
    console.log(this.getCurrentDateTime());
    this.racunaj(this.tp,this.ns );
  }

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

    getCurrentDateTime() {
      return new Date().toISOString().slice(0, 16);
    }


}
