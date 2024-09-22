package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Collections struct {
	Movies string
}

var Collection = &Collections{
	Movies: "movies",
}

type Movies struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Plot             string             `bson:"plot" json:"plot"`
	Genres           []string           `bson:"genres" json:"genres"`
	Runtime          int                `bson:"runtime" json:"runtime"`
	Cast             []string           `bson:"cast" json:"cast"`
	NumMflixComments int                `bson:"num_mflix_comments" json:"num_mflix_comments"`
	Poster           string             `bson:"poster" json:"poster"`
	Title            string             `bson:"title" json:"title"`
	Fullplot         string             `bson:"fullplot" json:"fullplot"`
	Languages        []string           `bson:"languages" json:"languages"`
	Released         time.Time          `bson:"released" json:"released"`
	Directors        []string           `bson:"directors" json:"directors"`
	Writers          []string           `bson:"writers" json:"writers"`
	Awards           struct {
		Wins        int    `bson:"wins" json:"wins"`
		Nominations int    `bson:"nominations" json:"nominations"`
		Text        string `bson:"text" json:"text"`
	} `bson:"awards" json:"awards"`
	Lastupdated string `bson:"lastupdated" json:"lastupdated"`
	Year        int    `bson:"year" json:"year"`
	Imdb        struct {
		Rating float64 `bson:"rating" json:"rating"`
		Votes  int     `bson:"votes" json:"votes"`
		ID     int     `bson:"id" json:"id"`
	} `bson:"imdb" json:"imdb"`
	Countries []string `bson:"countries" json:"countries"`
	Type      string   `bson:"type" json:"type"`
	Tomatoes  struct {
		Viewer struct {
			Rating     float64 `bson:"rating" json:"rating"`
			NumReviews int     `bson:"numReviews" json:"numReviews"`
		} `bson:"viewer" json:"viewer"`
		Production  string    `bson:"production" json:"production"`
		LastUpdated time.Time `bson:"lastUpdated" json:"lastUpdated"`
	} `bson:"tomatoes" json:"tomatoes"`
	PlotEmbedding []float64 `bson:"plot_embedding" json:"plot_embedding"`
}
