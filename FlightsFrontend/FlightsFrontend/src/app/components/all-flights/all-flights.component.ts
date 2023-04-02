import { AfterViewInit, Component, OnInit, ViewChild } from '@angular/core';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { Flights } from 'src/models/flight.model';
import { FlightService } from 'src/services/flight.service';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { SearchFlightsDTO } from 'src/models/flightDTO.model';
import { Ticket } from 'src/models/ticket';
import { TicketService } from 'src/services/ticket.service';
import { HttpErrorResponse } from '@angular/common/http';
import { DialogService } from 'src/services/dialog.service';
import { AuthService } from 'src/app/registration/services/auth.service';
import { logedUserInfo } from 'src/app/registration/model/logedUserInfo';
import { Pipe, PipeTransform } from '@angular/core';


@Component({
  selector: 'app-all-flights',
  templateUrl: './all-flights.component.html',
  styleUrls: ['./all-flights.component.css'],
})

export class AllFlightsComponent implements OnInit, AfterViewInit {

  
  
  flights = new MatTableDataSource<Flights[]>;
  role: string = 'ROLE_NOTAUTH'
  logedUser: logedUserInfo = {
    id: "",
    username: "",
    role: "",
    name: ''
  };
  
  
  probavam:any;
  displayedColumns!: string[];
  currentIndex = -1;
  public izabran : any ;
  fromPlace='';
  toPlace='';
  random='';

  tp:number=0;
  ns:number=0;
  fromDate='';
  toDate:any;

  totalSum:any=this.tp*this.ns;


  allFlights : Array<Flights> = new Array
  isAdmin = false ;
  startPlace: string = "";
  endPlace: string = "";
  startDate: Date | undefined;
  endDate : Date | undefined;

  @ViewChild(MatSort)
  sort: MatSort = new MatSort;
  
  ngOnInit(): void {
    this.displayedColumns = this.getdisplayedColumns()
    this.logedUser = this.authService.getLogedUserInfo() ?? {username: "", role: "", id: "", name: ""};
    this.racunaj(this.tp,this.ns);
    //console.log(this.logedUser)
  }
  racunaj(a: any, b:any){
    return this.totalSum = a*b;
  }

  ngAfterViewInit(): void {
    this.retrieveFlights();
    this.flights.sort = this.sort;
  }

  getdisplayedColumns(): string[] {
    //console.log(typeof this.role)
    if (this.role.toString() === '1') {
      //console.log("admina bc number is ", this.role)
      return ['fromplace', 'toplace', 'starttime', 'endtime', 'ticketprice', 'numofseats', 'totalsum','Edit', 'Delete', 'Buy'];
      
    } else if (this.role.toString() === '0'){
      //console.log("user bc number is ", this.role)
      return ['fromplace', 'toplace', 'starttime', 'endtime', 'ticketprice', 'numofseats', 'totalsum', 'Buy'];
      
    }
    else{
      //console.log("not auth bc number is ", this.role)
      return ['fromplace', 'toplace', 'starttime', 'endtime', 'ticketprice', 'numofseats', 'totalsum'];
    }
  }

  constructor(private flightService: FlightService, private ticketservice: TicketService,
    private dialogService: DialogService, private authService: AuthService) {
      this.GetUserRole();
      //console.log("ROLA USERA KOD FLIGHTS JE ", this.role)
     }
  

  retrieveFlights(): void {
    this.flightService.getAll()
      .subscribe({
        next: (data) => {
          this.flights = new MatTableDataSource(<Flights[][]>data);
          this.flights.sort = this.sort;
          //console.log(data);
        },
        error: (e) => console.error(e)
      });
  }


  FilterByNameAndSurname(smor:any,smor2:any,smor3:any){
    this.flightService.SearchByAll(smor,smor2,smor3)
    .subscribe((data:any) =>{
      this.flights=data;
    })
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
    this.dialogService.openDialogBuyingTicket(flight);
  }
  
  GetUserRole(): void{
    if(localStorage['authToken'] != null){
      this.role = this.authService.getUserRole();
    }
  }

 forDate(): void {
  this.flightService.getAll()
    .subscribe({
      next: (data) => {
        this.flights = new MatTableDataSource(<any>data.filter(flights => flights.fromplace?.includes(this.fromDate)));
        console.log("Nadji " + this.fromDate);
      },
      error: (e) => console.error(e)
    });
}


Search(nesto:any){
  this.flightService.Search(nesto)
  .subscribe((data:any)=>{
    this.flights.data = data.push;
  } )
    
}

filterByPlaces(): void {
  this.flightService.getAll()
    .subscribe({
      next: (data) => {
        this.flights = new MatTableDataSource(<any>data.filter(flights => flights.fromplace?.toLowerCase().includes(this.fromPlace.toLowerCase()) && flights.toplace?.toLowerCase().includes(this.toPlace.toLowerCase())));
        console.log("Nadji " + this.fromPlace + " " + this.toPlace);
      },
      error: (e) => console.error(e)
    });
}



/*
Nadji(nesto:any){
  this.flightService.Proba(nesto)
  .subscribe((data:any)=>{
    this.flights.data = data.push; //iLI data.push
  } )
    */

}






/*
onSubmitSearch(inputSearch : any){
  
  this.flightService.Search(inputSearch.searchinput).subscribe((data) => {
    this.flights = data;
    
  });
}

*/



