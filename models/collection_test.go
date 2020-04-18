package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNewCollectionResponse(t *testing.T) {
	gotStruct := NewCollectionResponse(1, time.Date(2019, 01, 01, 16, 00, 0, 0, time.UTC), "Star Wars", "Sci-Fi", 20, 15000)
	gotJson, _ := json.Marshal(gotStruct)
	got := string(gotJson)

	var want = `{"id":1,"dateTime":"2019-01-01T16:00:00Z","title":"Star Wars","category":"Sci-Fi","quantity":20,"rate":15000}`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
