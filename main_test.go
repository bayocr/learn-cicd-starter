package main

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

type test struct {
	name     string
	headers  http.Header
	exp      string
	experror error
}

func TestGetAPIKey(t *testing.T) {
	tests := []test{
		{
			name:     "Missing Authorization header",
			headers:  http.Header{},
			exp:      "",
			experror: auth.ErrNoAuthHeaderIncluded,
		},
		{
			name:     "Correcct ApiKey",
			headers:  http.Header{"Authorization": {"ApiKey supersecretkey"}},
			exp:      "supersecretkey",
			experror: nil,
		},
		{
			name:     "Malformed Authorization header",
			headers:  http.Header{"Authorization": {"BasicAuth badvalue"}},
			exp:      "",
			experror: errors.New("malformed authorization header"),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			key, _error := auth.GetAPIKey(testCase.headers)

			if key != testCase.exp {
				t.Errorf("Expected: %v, got: %v", testCase.exp, key)
			}

			if _error != testCase.experror {
				t.Errorf("Expected: %v, got: %v", testCase.experror, _error)
			}
		})
	}

}
