package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Ranking struct {
	Name  string `json:"ranking_name" bson:"ranking_name"`
	Value int    `json:"ranking_value" bson:"ranking_value"`
}

type Movie struct {
	ID          bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ImdbID      string        `json:"imdb_id" bson:"imdb_id"`
	Title       string        `json:"title" bson:"title"`
	PosterPath  string        `json:"poster_path" bson:"poster_path"`
	YoutubeID   string        `json:"youtube_id" bson:"youtube_id"`
	Genres      []string      `json:"genres" bson:"genres"`
	AdminReview string        `json:"admin_review" bson:"admin_review"`
	Ranking     Ranking       `json:"ranking" bson:"ranking"`
}
