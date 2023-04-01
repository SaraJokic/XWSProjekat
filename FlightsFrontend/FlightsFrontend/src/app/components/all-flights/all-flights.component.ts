import { AfterViewInit, Component, ViewChild } from '@angular/core';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { Flights } from 'src/models/flight.model';
import { FlightService } from 'src/services/flight.service';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { Ticket } from 'src/models/ticket';
import { TicketService } from 'src/services/ticket.service';
import { HttpErrorResponse } from '@angular/common/http';
import { DialogService } from 'src/services/dialog.service';


@Component({
  selector: 'app-all-flights',
  templateUrl: './all-flights.component.html',
  styleUrls: ['./all-flights.component.css'],
})

export class AllFlightsComponent implements AfterViewInit {

  displayedColumns:string[] = ['fromplace','toplace', 'starttime','endtime','ticketprice','numofseats', 'Delete', 'Buy'];
  flights = new MatTableDataSource<Flights[]>;
  
  


  currentIndex = -1;
  public izabran : any ;
  fromPlace='';
  toPlace='';



  @ViewChild(MatSort)
  sort: MatSort = new MatSort;
  

  ngAfterViewInit(): void {
    this.retrieveFlights();
    this.flights.sort = this.sort;
    
  }

  constructor(private flightService: FlightService, private ticketservice: TicketService,
    private dialogService: DialogService) { }

  retrieveFlights(): void {
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
          this.flights = new MatTableDataSource(<any>data.filter(flights => flights.fromplace?.toLowerCase().includes(this.fromPlace.toLowerCase()) && flights.toplace?.toLowerCase().includes(this.toPlace.toLowerCase())));
        },
        error: (e) => console.error(e)
      });
  }

  buyTicket(flight: any){
    const newTicket: Ticket = {
      userid: "6426f65971b16d7d27fe5bb8",
      flightid: flight.id,
      expired: false,
      quantity: 4,
    };
    console.log(flight.id)
    this.ticketservice.add(newTicket).subscribe(
      (data) => {
        alert("Success!");
        this.retrieveFlights()
      },
      (error: HttpErrorResponse) => {
        alert(error.message);
      }
    )
  }
  deleteFlight(deleting : any){

    this.izabran = deleting.id;
    this.flightService.delete(this.izabran).subscribe(
      (resp) =>{
        this.retrieveFlights();
        return console.log("Deleted!");
      }, err=>{
        return console.error("Neuspesno");
      });
  }
  openDialog(flight: Flights): void {
    this.dialogService.openDialog(flight);
  }
    



/*
Search(nesto:any){
  this.flightService.Search(nesto)
  .subscribe((data)=>{
    this.flights.data = data.push;
  } )
    
}
*/
/*
onSubmitSearch(inputSearch : any){
  
  this.flightService.Search(inputSearch.searchinput).subscribe((data) => {
    this.flights = data;
    
  });
}

*/



}