import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BallotDetailsComponent } from './ballot-details.component';

describe('BallotDetailsComponent', () => {
  let component: BallotDetailsComponent;
  let fixture: ComponentFixture<BallotDetailsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BallotDetailsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BallotDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
