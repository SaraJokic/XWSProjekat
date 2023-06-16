import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReservationRequestsGuestComponent } from './reservation-requests-guest.component';

describe('ReservationRequestsGuestComponent', () => {
  let component: ReservationRequestsGuestComponent;
  let fixture: ComponentFixture<ReservationRequestsGuestComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ReservationRequestsGuestComponent]
    });
    fixture = TestBed.createComponent(ReservationRequestsGuestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
