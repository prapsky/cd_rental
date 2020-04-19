package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCollection(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/collection/1", nil)
	response := httptest.NewRecorder()

	Collection(response, request)

	t.Run("Get collection: ", func(t *testing.T) {
		got := response.Body.Bytes()

		var want = []byte(`{"id":1,"dateTime":"2020-04-18T23:52:40.238858Z","title":"Star Wars","category":"Sci-Fi","quantity":20,"rate":15000}`)

		if !bytes.Equal(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
