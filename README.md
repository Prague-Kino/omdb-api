# omdb-api

This is a very simple GO API for querying [OMDb](https://www.omdbapi.com/), a free movie database.

## Instructions

### Constructor

To use this API, create a new `OMDb` object using an API key. All API methods come from this object.

```go
import "github.com/Prague-Kino/omdb-api/api"

apiKey := "..."
omdb := api.NewOMDb(key)
```

---

### Methods

As of now, all you can do in this API is search for a movie. What else would you do? I guess you could also want to search for TV shows but in this case I only care about movies, so that's all it does.

#### FetchMovie

Search for a movie by its Title and optionally also its Release Year.
If found, returns a [`Movie`](#movie) model, otherwise returns an error.

```go
movie, err := omdb.FetchMovie("Die my love", "2025")
```

## Models

The only model exported by this API is [`Movie`](#movie), which stores all the JSON data returned by OMDb.

### Movie

```go
type Movie struct {
    Title      string
    Year       string
    Rated      string
    Released   string
    Runtime    string
    Genre      string
    Director   string
    Plot       string
    Poster     string
    IMDbRating string
    Response   string
    Error      string
}
```

---

That's all for now!
