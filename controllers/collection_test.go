package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRealtimeData(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/collection", nil)
	response := httptest.NewRecorder()

	GetCollection(response, request)

	t.Run("Get collection: ", func(t *testing.T) {
		got := response.Body.Bytes()

		var want = []byte(`{"dateTime":"2019-01-01T16:00:00Z","id":1,"title":"Star Wars","category":"Sci-Fi","quantity":20,"rate":15000}`)

		if !bytes.Equal(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
