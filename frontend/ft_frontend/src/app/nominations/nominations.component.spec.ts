import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { NominationsComponent } from './nominations.component';

describe('NominationsComponent', () => {
  let component: NominationsComponent;
  let fixture: ComponentFixture<NominationsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NominationsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(NominationsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
