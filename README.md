# omdb-api

This is an API client to fetch film data from OMDb.

All it does is fetch movies by title.

---

## Constructor

The OMDb constructor takes in one parameter:

- `key` - API key string

If the key is invalid, it will throw the following error:

- `InvalidKeyError`

### Usage

```go
omdb, err := api.NewOMDb(key)
if err != nil {
  log.Fatal(err)
}
```

---

## Fetch Movie

Queries the API for movie data by the movie title (case insensitive).

### Usage

```go
movie, err := omdb.FetchMovie("Die my love")
if err != nil {
  log.Fatal(err)
}
prettyPrint(movie)
```

**Output:**

```go
{
    "Title": "Die My Love",
    "Year": "2025",
    "Rated": "R",
    "Released": "07 Nov 2025",
    "Runtime": "119 min",
    "Genre": "Drama, Thriller",
    "Director": "Lynne Ramsay",
    "Plot": "Grace, a writer and young mother, is slowly slipping into madness...",
    "Poster": "https://m.media-amazon.com/images/M/MV5BYjc5OWZlZmYtMTg3Yy00YzFmLTg0YTgtNjVjN2M2ZWJjOWM1XkEyXkFqcGc@._V1_SX300.jpg",
    "imdbRating": "6.1",
    "Response": "True"
}
```

---
