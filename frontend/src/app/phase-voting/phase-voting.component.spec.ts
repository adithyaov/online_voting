import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PhaseVotingComponent } from './phase-voting.component';

describe('PhaseVotingComponent', () => {
  let component: PhaseVotingComponent;
  let fixture: ComponentFixture<PhaseVotingComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PhaseVotingComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PhaseVotingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
