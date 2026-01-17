package models

type Movie struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Plot       string `json:"Plot"`
	Poster     string `json:"Poster"`
	IMDbRating string `json:"imdbRating"`
	Response   string `json:"Response"`
	Error      string `json:"Error,omitempty"`
}
