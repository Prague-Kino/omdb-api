package api

import (
	"encoding/json"
	"net/url"

	e "github.com/Prague-Kino/omdb-api/internal/errors"
	m "github.com/Prague-Kino/omdb-api/internal/models"
	su "github.com/Prague-Kino/omdb-api/internal/stringutils"
)

const (
	SUCCESSFUL_RESPONSE = "True"
	TYPE_MOVIE          = "movie"
)

func (o *OMDb) FetchMovie(title string) (*m.Movie, error) {
	if su.IsEmpty(title) {
		return nil, &e.InvalidMovieNameError{Title: title}
	}

	params := url.Values{}
	params.Add(PARAM_KEY, o.apiKey)
	params.Add(PARAM_TITLE, title)
	params.Add(PARAM_TYPE, TYPE_MOVIE)

	requestURL := o.formatRequestURL(params)
	body, err := o.get(requestURL)
	if err != nil {
		return nil, err
	}

	var movie m.Movie
	if err := json.Unmarshal(body, &movie); err != nil {
		return nil, err
	}

	if movie.Response != SUCCESSFUL_RESPONSE {
		return nil, &e.MovieNotFoundError{Title: title}
	}

	return &movie, nil
}
