import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EnqueuedComponent } from './enqueued.component';

describe('EnqueuedComponent', () => {
  let component: EnqueuedComponent;
  let fixture: ComponentFixture<EnqueuedComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EnqueuedComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EnqueuedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
