package models

import "github.com/Prague-Kino/cast/cast"

type Movie struct {
	cast.Film
	APIResponse
}

func (m *Movie) ToFilm() cast.Film {
	return m.Film
}
