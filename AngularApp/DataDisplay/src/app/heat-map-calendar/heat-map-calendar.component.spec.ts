import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HeatMapCalendarComponent } from './heat-map-calendar.component';

describe('HeatMapCalendarComponent', () => {
  let component: HeatMapCalendarComponent;
  let fixture: ComponentFixture<HeatMapCalendarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HeatMapCalendarComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(HeatMapCalendarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
