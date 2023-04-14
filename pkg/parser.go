package pkg

import (
	"errors"
	"io"
	"net/http"
)

// In the ParseApi function here we retrieve data from json with the get method
// And we return an array of bytes and an error

func ParseAPI(url, id string) ([]byte, error) {
	// Get  make HTTP (or HTTPS) requests
	resp, err := http.Get(url + id)
	// Status check
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, errors.New("Can't find api address")
	}

	// ReadAll reads from r until an error or EOF and returns the data it read. ​
	bytes, err := io.ReadAll(resp.Body)
	// Checking error
	if err != nil {
		return nil, errors.New("Can't translate to bytes")
	}

	// return massiv byte and nil
	return bytes, nil
}
