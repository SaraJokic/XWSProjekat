import { TestBed } from '@angular/core/testing';

import { AccommodationServiceService } from './accommodation-service.service';

describe('AccommodationServiceService', () => {
  let service: AccommodationServiceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AccommodationServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
