package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Prague-Kino/omdb-api/internal/enums/searchtype"
	"github.com/Prague-Kino/omdb-api/internal/errors"
	"github.com/Prague-Kino/omdb-api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MovieTestCase struct {
	name     string
	title    string
	year     string
	expected *models.Movie
}

func newMockServer_Success(t *testing.T, title, year string) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		assert.Equal(t, title, r.URL.Query().Get("t"))
		assert.Equal(t, year, r.URL.Query().Get("y"))
		assert.Equal(t, searchtype.Movie.String(), r.URL.Query().Get("type"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte(fmt.Sprintf(`{
			"Title": %q,
			"Year": %q,
			"Response": "True"
		}`, title, year)))

		require.NoError(t, err)
	}))
}

func newMockServer_Error(t *testing.T, response string) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte(response))
		require.NoError(t, err)
	}))
}

func newMockServer_OmdbAPIError(t *testing.T) *httptest.Server {
	t.Helper()
	return newMockServer_Error(
		t,
		`{"Response": "False", "Error": "Movie not found!"}`,
	)
}

func newMockServer_JsonError(t *testing.T) *httptest.Server {
	t.Helper()
	return newMockServer_Error(
		t,
		`{"malformed json zf;v'';';aa"}`,
	)
}

func TestFetchMovie_Success(t *testing.T) {
	tests := []MovieTestCase{
		{
			name:  "fetch movie by name only",
			title: "The Straight Story",
			expected: &models.Movie{
				Title: "The Straight Story",
			},
		},
		{
			name:  "fetch movie by name and year",
			title: "Obsession",
			year:  "2026",
			expected: &models.Movie{
				Title: "Obsession",
				Year:  "2026",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, testFetchMovie(tc))
	}
}

func testFetchMovie(tc MovieTestCase) func(t *testing.T) {
	return func(t *testing.T) {
		server := newMockServer_Success(t, tc.title, tc.year)
		defer server.Close()

		omdb := &OMDb{
			apiKey:  "test-key",
			baseURL: server.URL + "/",
		}
		movie, err := omdb.FetchMovie(tc.title, tc.year)

		require.NoError(t, err)
		require.NotNil(t, movie)
		assert.Equal(t, tc.expected.Title, movie.Title)
		assert.Equal(t, tc.expected.Year, movie.Year)
	}
}

func TestFetchMovie_OmdbError(t *testing.T) {
	server := newMockServer_OmdbAPIError(t)
	defer server.Close()

	omdb := &OMDb{
		apiKey:  "test-key",
		baseURL: server.URL + "/",
	}
	movie, err := omdb.FetchMovie("Garmonbozia", "3000")
	require.Error(t, err)
	var apiErr *errors.OmdbAPIError
	require.ErrorAs(t, err, &apiErr)
	require.Nil(t, movie)
}

func TestFetchMovie_JsonError(t *testing.T) {
	server := newMockServer_JsonError(t)
	defer server.Close()

	omdb := &OMDb{
		apiKey:  "test-key",
		baseURL: server.URL + "/",
	}
	movie, err := omdb.FetchMovie("Garmonbozia", "3000")
	require.Error(t, err)
	var apiErr *errors.InvalidAPIResponse
	require.ErrorAs(t, err, &apiErr)
	require.Nil(t, movie)
}
