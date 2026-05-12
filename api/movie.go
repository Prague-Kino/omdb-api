package api

import (
	"encoding/json"

	"github.com/Prague-Kino/omdb-api/internal/errors"
	"github.com/Prague-Kino/omdb-api/models"
)

const (
	SuccessfulResponse = "True"
	ParamTypeMovie     = "movie"
	ParamYear          = "y"
)

// Fetch a movie's data by its title and release year.
//
// title - the name of the movie
//
// year (optional) - year of release
func (o *OMDb) FetchMovie(title string, year ...string) (*models.Movie, error) {
	params := o.setupSearchParams(title, year...)

	requestURL := o.formatRequestURL(params)
	body, err := o.get(requestURL)
	if err != nil {
		return nil, err
	}

	var movie models.Movie
	if err := json.Unmarshal(body, &movie); err != nil {
		return nil, &errors.InvalidAPIResponse{
			Err:  err,
			Body: body,
		}
	}

	if movie.Response != SuccessfulResponse {
		return nil, &errors.OmdbAPIError{
			Err:        movie.Error,
			RequestURL: requestURL,
		}
	}

	return &movie, nil
}
