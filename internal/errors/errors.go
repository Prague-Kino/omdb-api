package errors

import "fmt"

type MovieNotFoundError struct {
	Title string
}

func (e *MovieNotFoundError) Error() string {
	return fmt.Sprintf("Movie not found: [%s]\n", e.Title)
}

type InvalidKeyError struct {
	Key string
}

func (e *InvalidKeyError) Error() string {
	return fmt.Sprintf("Invalid API key: %s\n", e.Key)
}
