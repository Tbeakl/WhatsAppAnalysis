import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BasicSummaryComponent } from './basic-summary.component';

describe('BasicSummaryComponent', () => {
  let component: BasicSummaryComponent;
  let fixture: ComponentFixture<BasicSummaryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BasicSummaryComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BasicSummaryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
