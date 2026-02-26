import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Movie } from '../models/movie.model';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class MovieService {
  private apiUrl = `${environment.apiUrl}/movies`;
  private http: HttpClient = inject(HttpClient);

    getMovies(): Observable<Movie[]> {
        return this.http.get<Movie[]>(this.apiUrl);
    }
}
