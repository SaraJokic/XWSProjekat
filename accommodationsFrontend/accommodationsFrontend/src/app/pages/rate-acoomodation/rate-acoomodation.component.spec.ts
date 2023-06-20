import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RateAcoomodationComponent } from './rate-acoomodation.component';

describe('RateAcoomodationComponent', () => {
  let component: RateAcoomodationComponent;
  let fixture: ComponentFixture<RateAcoomodationComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [RateAcoomodationComponent]
    });
    fixture = TestBed.createComponent(RateAcoomodationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
