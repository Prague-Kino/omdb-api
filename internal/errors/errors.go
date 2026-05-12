package errors

import "fmt"

type HttpGetError struct {
	Url string
	Err error
}

func (e *HttpGetError) Error() string {
	return fmt.Sprintf(
		"Error fetching data.\nURL: %s\nError: %s",
		e.Url,
		e.Err,
	)
}

type InvalidResponseError struct {
	Url string
	Err error
}

func (e *InvalidResponseError) Error() string {
	return fmt.Sprintf(
		"Invalid body response.\nURL: %s\nError: %s",
		e.Url,
		e.Err,
	)
}

type EnvLoadingError struct {
	Err error
}

func (e *EnvLoadingError) Error() string {
	return fmt.Sprintf("Error loading .env: %s", e.Err)
}

type ApiKeyNotSetError struct{}

func (e *ApiKeyNotSetError) Error() string {
	return "OMDB_API_KEY not set in .env file"
}

type InvalidAPIResponse struct {
	Err        error
	Body       []byte
	RequestURL string
}

func (e *InvalidAPIResponse) Error() string {
	return fmt.Sprintf(
		"Failed to parse API response.\n"+
			"Request URL: %s\n"+
			"Response: %s\n"+
			"Error: %s",
		e.RequestURL,
		e.Body,
		e.Err,
	)
}

type OmdbAPIError struct {
	Err        string
	RequestURL string
}

func (e OmdbAPIError) Error() string {
	return fmt.Sprintf(
		"OMDb API Error: %s\n"+
			"Request URL: %s\n",
		e.Err,
		e.RequestURL,
	)
}
