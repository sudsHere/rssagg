package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts api key from request headers from a HTTP request
// example:
// Authorization: ApiKey <insert API key here>

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("authorization header not found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("Malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed first part of auth header")
	}

	return vals[1], nil
}
