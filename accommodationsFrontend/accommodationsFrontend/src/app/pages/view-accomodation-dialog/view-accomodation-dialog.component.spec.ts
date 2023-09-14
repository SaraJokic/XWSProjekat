import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewAccomodationDialogComponent } from './view-accomodation-dialog.component';

describe('ViewAccomodationDialogComponent', () => {
  let component: ViewAccomodationDialogComponent;
  let fixture: ComponentFixture<ViewAccomodationDialogComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ViewAccomodationDialogComponent]
    });
    fixture = TestBed.createComponent(ViewAccomodationDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
