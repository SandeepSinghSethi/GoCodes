package auth

import (
	"errors"
	"net/http"
	"strings"
)

// api header :
// Authorization : api_key <64 digit value>
func GetAPIKey(header http.Header) (string, error) {
	value := header.Get("Authorization")
	if value == "" {
		return "", errors.New("no authorization header provided for authentication")
	}
	apival := strings.Split(value, " ")
	if len(apival) != 2 {
		return "", errors.New("malformed auth header")
	}

	if apival[0] != "api_key" {
		return "", errors.New("invalid first part of auth header")
	}
	return apival[1], nil
}
