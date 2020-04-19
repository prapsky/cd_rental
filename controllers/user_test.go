package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/user/1", nil)
	response := httptest.NewRecorder()

	GetUser(response, request)

	t.Run("Get user: ", func(t *testing.T) {
		got := response.Body.Bytes()

		var want = []byte(`{"id":1,"dateTime":"2020-04-19T17:09:26.710061Z","name":"Jeffrey","phoneNumber":"085624136123","address":"Jalan A no.1 Jakarta Selatan"}`)

		if !bytes.Equal(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
