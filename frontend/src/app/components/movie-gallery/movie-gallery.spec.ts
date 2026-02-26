import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MovieGallery } from './movie-gallery';

describe('MovieGallery', () => {
  let component: MovieGallery;
  let fixture: ComponentFixture<MovieGallery>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MovieGallery],
    }).compileComponents();

    fixture = TestBed.createComponent(MovieGallery);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
