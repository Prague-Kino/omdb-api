package models

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Country    string `json:"Country"`
	Language   string `json:"Language"`
	Plot       string `json:"Plot"`
	Poster     string `json:"Poster"`
	IMDbRating string `json:"imdbRating"`
	Response   string `json:"Response"`
	Error      string `json:"Error,omitempty"`
}

func (m *Movie) String() string {
	s, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return fmt.Sprintf(m.Title, m.Year)
	}
	return string(s)
}
