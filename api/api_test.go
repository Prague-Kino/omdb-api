package api

import (
	"fmt"
	"math/rand/v2"
	"net/url"
	"strings"
	"testing"

	"github.com/Prague-Kino/omdb-api/internal/enums/searchparams"
	"github.com/stretchr/testify/assert"
)

type ParamsTestCase struct {
	name           string
	title          string
	year           string
	expectedURL    string
	expectedParams url.Values
}

func (t ParamsTestCase) createUrlParams(omdb *OMDb) url.Values {
	if t.year != "" {
		return omdb.setupSearchParams(t.title, t.year)
	}

	return omdb.setupSearchParams(t.title)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
func generateRandomString(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.IntN(len(charset))])
	}
	return sb.String()
}

// -------------------------------------------------

func TestNewOMDB(t *testing.T) {
	key := generateRandomString(10)
	o := NewOMDb(key)

	assert.Equal(t, o.apiKey, key)
	assert.Equal(t, o.baseURL, OmdbApiUrl)
}

func TestSetupSearchParams(t *testing.T) {
	key := generateRandomString(6)
	o := NewOMDb(key)

	tests := []ParamsTestCase{
		{
			name:  "just title - single world",
			title: "Psycho",
			expectedParams: url.Values{
				searchparams.Key:   {key},
				searchparams.Title: {"Psycho"},
				searchparams.Type:  {"movie"},
			},
		},
		{
			name:  "just title - multiple words",
			title: "There Will Be Blood",
			expectedParams: url.Values{
				searchparams.Key:   {key},
				searchparams.Title: {"There Will Be Blood"},
				searchparams.Type:  {"movie"},
			},
		},
		{
			name:  "title and year",
			title: "The Invite",
			year:  "2026",
			expectedParams: url.Values{
				searchparams.Key:   {key},
				searchparams.Title: {"The Invite"},
				searchparams.Type:  {"movie"},
				searchparams.Year:  {"2026"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, testSetupSearchParams(tc, o))
	}
}

func testSetupSearchParams(tc ParamsTestCase, omdb *OMDb) func(t *testing.T) {
	return func(t *testing.T) {
		params := tc.createUrlParams(omdb)
		assert.Equal(t, tc.expectedParams, params)
	}
}

func TestFormatRequestURL(t *testing.T) {
	key := generateRandomString(6)
	o := NewOMDb(key)

	tests := []ParamsTestCase{
		{
			name:        "just title - single word",
			title:       "M",
			expectedURL: fmt.Sprintf(
				"http://www.omdbapi.com/?apiKey=%s&t=M&type=movie",
				key,
			),
		},
		{
			name:        "just title - multiple words",
			title:       "The Night of the Hunter",
			expectedURL: fmt.Sprintf(
				"http://www.omdbapi.com/?apiKey=%s&t=The+Night+of+the+Hunter&type=movie",
				key,
			),
		},
		{
			name:        "title and year",
			title:       "Fire Walk With Me",
			year:        "1992",
			expectedURL: fmt.Sprintf(
				"http://www.omdbapi.com/?apiKey=%s&t=Fire+Walk+With+Me&type=movie&y=1992",
				key,
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, testFormatRequestURL(tc, o))
	}
}

func testFormatRequestURL(tc ParamsTestCase, omdb *OMDb) func(t *testing.T) {
	return func(t *testing.T) {
		params := tc.createUrlParams(omdb)
		requestURL := omdb.formatRequestURL(params)
		assert.Equal(t, tc.expectedURL, requestURL)
	}
}
