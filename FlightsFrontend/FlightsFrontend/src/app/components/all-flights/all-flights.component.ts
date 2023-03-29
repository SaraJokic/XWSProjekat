import { AfterViewInit, Component, ViewChild } from '@angular/core';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { Flights } from 'src/models/flight.model';
import { FlightService } from 'src/services/flight.service';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-all-flights',
  templateUrl: './all-flights.component.html',
  styleUrls: ['./all-flights.component.css']
})

export class AllFlightsComponent implements AfterViewInit {

  displayedColumns:string[] = ['fromPlace','toPlace', 'startTime','endTime','ticketPrice','numOfSeats'];
  flights = new MatTableDataSource<Flights[]>;
  currentCentre: Flights = {};
  currentIndex = -1;

  fromPlace='';
  toPlace='';

  @ViewChild(MatSort)
  sort: MatSort = new MatSort;

  ngAfterViewInit(): void {
    this.retrieveCentres();
    this.flights.sort = this.sort;
    
  }

  constructor(private flightService: FlightService) { }

  retrieveCentres(): void {
    this.flightService.getAll()
      .subscribe({
        next: (data) => {
          this.flights = new MatTableDataSource(<Flights[][]>data);
          this.flights.sort = this.sort;
          console.log(data);
        },
        error: (e) => console.error(e)
      });
  }


  filterByNameAndSurname(): void {
    this.flightService.getAll()
      .subscribe({
        next: (data) => {
          this.flights = new MatTableDataSource(<any>data.filter(flights => flights.fromPlace?.includes(this.fromPlace) && flights.toPlace?.includes(this.toPlace)));
        },
        error: (e) => console.error(e)
      });
  }


}
