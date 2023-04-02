import { Component, Input, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Flights } from 'src/models/flight.model';
import { FlightService } from 'src/services/flight.service';

@Component({
  selector: 'app-update-flights',
  templateUrl: './update-flights.component.html',
  styleUrls: ['./update-flights.component.css']
})
export class UpdateFlightsComponent implements OnInit {

  @Input() viewMode = false;

  //starttime?:Date;
  //endtime?:Date;


  @Input() currentFlight: Flights = {
    fromplace: '',
    toplace: '',
    starttime:  new Date,
    endtime: new  Date, 
    numofseats: 0,
    ticketprice:0,
  };

  tp:number=0;
  ns:number=0;

  message = '';
  starttime?:Date;

  totalSum:any=this.ns*this.tp;

  constructor(private flightService: FlightService, private router: Router, private route: ActivatedRoute,) { }

  ngOnInit(): void {
    if (!this.viewMode) {
      this.message = '';
      this.getFlightID(this.route.snapshot.params["id"]);
      this.racunaj(this.ns,this.tp );
    }

  }



racunaj(a: any, b:any){
  return this.totalSum = a*b;
}


  getFlightID(id: string): void {
    this.flightService.getById(id)
      .subscribe({
        next: (data) => {
          this.currentFlight = data;
          console.log(data);
        },
        error: (e) => console.error(e)
      });
  }


  updateFlight(): void {
    this.message = '';

    if(confirm("Are you sure to update?"))
      {
      this.flightService.update(this.currentFlight.id, this.currentFlight)
        .subscribe({
          next: (res) => {
            console.log(res);
            this.message = res.message ? res.message : 'This flight was updated successfully!';
          },
          error: (e) => {
            console.error(e);
            this.message = "Validation not good!";
          }
              });
    }
  }



}
