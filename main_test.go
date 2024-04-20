package main

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	var headers http.Header

	_, error := auth.GetAPIKey(headers)

	if error == nil {
		t.FailNow()
		t.Error(error)
	}
}
