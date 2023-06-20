import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MakeReservationDialogComponent } from './make-reservation-dialog.component';

describe('MakeReservationDialogComponent', () => {
  let component: MakeReservationDialogComponent;
  let fixture: ComponentFixture<MakeReservationDialogComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [MakeReservationDialogComponent]
    });
    fixture = TestBed.createComponent(MakeReservationDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
