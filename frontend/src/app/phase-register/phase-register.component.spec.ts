import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PhaseRegisterComponent } from './phase-register.component';

describe('PhaseRegisterComponent', () => {
  let component: PhaseRegisterComponent;
  let fixture: ComponentFixture<PhaseRegisterComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PhaseRegisterComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PhaseRegisterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
