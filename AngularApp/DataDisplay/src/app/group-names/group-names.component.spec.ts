import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupNamesComponent } from './group-names.component';

describe('GroupNamesComponent', () => {
  let component: GroupNamesComponent;
  let fixture: ComponentFixture<GroupNamesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GroupNamesComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(GroupNamesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
