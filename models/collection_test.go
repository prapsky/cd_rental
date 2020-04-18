package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNewCollectionResponse(t *testing.T) {
	gotStruct := NewCollectionResponse(time.Date(2019, 01, 01, 16, 00, 0, 0, time.UTC), 1, "Star Wars", "Sci-Fi", 20, 15000)
	gotJson, _ := json.Marshal(gotStruct)
	got := string(gotJson)

	var want = `{"dateTime":"2019-01-01T16:00:00Z","id":1,"title":"Star Wars","category":"Sci-Fi","quantity":20,"rate":15000}`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
