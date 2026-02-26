package repository

import (
	"context"
	"movie-streaming-backend/internal/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MovieRepository interface {
	CreateMovie(ctx context.Context, movie *model.Movie) error
	FindAllMovies(ctx context.Context) ([]model.Movie, error)
}

type mongoMovieRepository struct {
	collection *mongo.Collection
}

func NewMovieRepository(db *mongo.Database) MovieRepository {
	return &mongoMovieRepository{collection: db.Collection("movies")}
}

func (r *mongoMovieRepository) CreateMovie(ctx context.Context, movie *model.Movie) error {
	// Generate an ID before inserting so our Go struct has it too!
	movie.ID = bson.NewObjectID()
	_, err := r.collection.InsertOne(ctx, movie)
	return err
}

func (r *mongoMovieRepository) FindAllMovies(ctx context.Context) ([]model.Movie, error) {
	var movies []model.Movie

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &movies); err != nil {
		return nil, err
	}

	return movies, nil
}
