package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
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

func NewOMDb(apiKey string) *OMDb {
	return &OMDb{
		apiKey:  apiKey,
		baseURL: OMDB_API_URL,
	}
}

func (o *OMDb) formatRequestURL(params url.Values) string {
	return o.baseURL + "?" + params.Encode()
}

func (o *OMDb) get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return body, nil
}
