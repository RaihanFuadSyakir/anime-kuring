package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AnimeSeason struct {
	Season string  `json:"season" bson:"season"`
	Year   int     `json:"year" bson:"year"`
	Animes []Anime `json:"animes" bson:"animes"`
}

type License struct {
	Name string `json:"name" bson:"name"`
	URL  string `json:"url" bson:"url"`
}

type Anime struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Sources     []string           `json:"sources" bson:"sources"`
	Title       string             `json:"title" bson:"title"`
	Type        string             `json:"type" bson:"type"`
	Episodes    int                `json:"episodes" bson:"episodes"`
	Status      string             `json:"status" bson:"status"`
	AnimeSeason AnimeSeason        `json:"animeSeason" bson:"animeSeason"`
	Picture     string             `json:"picture" bson:"picture"`
	Thumbnail   string             `json:"thumbnail" bson:"thumbnail"`
	Synonyms    []string           `json:"synonyms" bson:"synonyms"`
	Relations   []string           `json:"relations" bson:"relations"`
	Tags        []string           `json:"tags" bson:"tags"`
}

type AnimeData struct {
	License    License `json:"license" bson:"license"`
	Repository string  `json:"repository" bson:"repository"`
	LastUpdate string  `json:"lastUpdate" bson:"lastUpdate"`
	Data       []Anime `json:"data" bson:"data"`
}
