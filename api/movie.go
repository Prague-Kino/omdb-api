package api

import (
	"encoding/json"
	"fmt"
	"net/url"

	m "github.com/Prague-Kino/omdb-api/internal/models"
)

const (
	SUCCESSFUL_RESPONSE = "True"
	TYPE_MOVIE          = "movie"
)

func (o *OMDb) FetchMovie(title string) (*m.Movie, error) {
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
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if movie.Response != SUCCESSFUL_RESPONSE {
		return nil, fmt.Errorf("API error: %s", movie.Error)
	}

	return &movie, nil
}
