import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReservationRequestsHostComponent } from './reservation-requests-host.component';

describe('ReservationRequestsHostComponent', () => {
  let component: ReservationRequestsHostComponent;
  let fixture: ComponentFixture<ReservationRequestsHostComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ReservationRequestsHostComponent]
    });
    fixture = TestBed.createComponent(ReservationRequestsHostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
