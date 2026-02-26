export interface Ranking {
  ranking_name: string;
  ranking_value: number;
}

export interface Movie {
  id?: string;
  imdb_id: string;
  title: string;
  poster_path: string;
  youtube_id: string;
  genres: string[];
  admin_review: string;
  ranking: Ranking;
}
