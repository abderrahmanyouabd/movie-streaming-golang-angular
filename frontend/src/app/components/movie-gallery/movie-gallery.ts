import { Component, inject } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { MovieService } from '../../services/movie.service';

@Component({
  selector: 'app-movie-gallery',
  standalone: true,
  imports: [],
  templateUrl: './movie-gallery.html',
  styleUrl: './movie-gallery.scss',
})
export class MovieGallery {
  private movieService = inject(MovieService);

  movies = toSignal(this.movieService.getMovies(), { initialValue: [] });
}
