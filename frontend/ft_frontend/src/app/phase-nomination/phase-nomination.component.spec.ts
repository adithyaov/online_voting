import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PhaseNominationComponent } from './phase-nomination.component';

describe('PhaseNominationComponent', () => {
  let component: PhaseNominationComponent;
  let fixture: ComponentFixture<PhaseNominationComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PhaseNominationComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PhaseNominationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
