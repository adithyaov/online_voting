import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BallotBtnComponent } from './ballot-btn.component';

describe('BallotBtnComponent', () => {
  let component: BallotBtnComponent;
  let fixture: ComponentFixture<BallotBtnComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BallotBtnComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BallotBtnComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
