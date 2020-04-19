package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRent(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/rent/1", nil)
	response := httptest.NewRecorder()

	Rent(response, request)

	t.Run("Get rent: ", func(t *testing.T) {
		got := response.Body.Bytes()

		var want = []byte(`{"id":1,"dateTime":"2020-04-19T22:38:40.12395Z","queueNumber":1,"userId":1,"cdId":1,"rentQuantity":1}`)

		if !bytes.Equal(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
