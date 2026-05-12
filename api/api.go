package api

import (
	"io"
	"net/http"
	"net/url"

	"github.com/Prague-Kino/omdb-api/internal/enums/searchparams"
	"github.com/Prague-Kino/omdb-api/internal/enums/searchtype"
	"github.com/Prague-Kino/omdb-api/internal/errors"
)

const OmdbApiUrl = "http://www.omdbapi.com/"

type OMDb struct {
	apiKey  string
	baseURL string
}

func NewOMDb(apiKey string) *OMDb {
	return &OMDb{
		apiKey:  apiKey,
		baseURL: OmdbApiUrl,
	}
}

func (o *OMDb) formatRequestURL(params url.Values) string {
	return o.baseURL + "?" + params.Encode()
}

func (o *OMDb) get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, &errors.HttpGetError{
			Url: url,
			Err: err,
		}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &errors.InvalidResponseError{
			Url: url,
			Err: err,
		}
	}

	return body, nil
}

func (o *OMDb) setupSearchParams(title string, year ...string) url.Values {
	params := url.Values{}
	params.Add(searchparams.Key, o.apiKey)
	params.Add(searchparams.Title, title)
	params.Add(searchparams.Type, searchtype.Movie.String())

	if len(year) > 0 {
		params.Add(searchparams.Year, year[0])
	}

	return params
}
