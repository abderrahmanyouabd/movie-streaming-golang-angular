package service

import (
	"context"
	"movie-streaming-backend/internal/model"
	"movie-streaming-backend/internal/repository"
)

type MovieService interface {
	GetAllMovies(ctx context.Context) ([]model.Movie, error)
	AddMovie(ctx context.Context, movie *model.Movie) error
}

type movieService struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) MovieService {
	return &movieService{
		repo: repo,
	}
}

func (s *movieService) GetAllMovies(ctx context.Context) ([]model.Movie, error) {
	return s.repo.FindAllMovies(ctx)
}

func (s *movieService) AddMovie(ctx context.Context, movie *model.Movie) error {
	return s.repo.CreateMovie(ctx, movie)
}
