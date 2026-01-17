package api

import (
	"io"
	"net/http"
	"net/url"

	e "github.com/Prague-Kino/omdb-api/internal/errors"
	su "github.com/Prague-Kino/omdb-api/internal/stringutils"
)

const (
	OMDB_API_URL = "http://www.omdbapi.com/"
	PARAM_KEY    = "apikey"
	PARAM_TITLE  = "t"
	PARAM_TYPE   = "type"
)

type OMDb struct {
	apiKey  string
	baseURL string
}

func NewOMDb(apiKey string) (*OMDb, error) {
	if su.IsEmpty(apiKey) {
		return nil, &e.InvalidKeyError{Key: apiKey}
	}
	return &OMDb{
		apiKey:  apiKey,
		baseURL: OMDB_API_URL,
	}, nil
}

func (o *OMDb) formatRequestURL(params url.Values) string {
	return o.baseURL + "?" + params.Encode()
}

func (o *OMDb) get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
