package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNewUserResponse(t *testing.T) {
	gotStruct := NewUserResponse(1, time.Date(2019, 01, 01, 16, 00, 0, 0, time.UTC), "Jeffrey", "085624136123", "Jalan A no.1 Jakarta Selatan")
	gotJson, _ := json.Marshal(gotStruct)
	got := string(gotJson)

	var want = `{"id":1,"dateTime":"2019-01-01T16:00:00Z","name":"Jeffrey","phoneNumber":"085624136123","address":"Jalan A no.1 Jakarta Selatan"}`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
