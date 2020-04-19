package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNewRentResponse(t *testing.T) {
	gotStruct := NewRentResponse(1, time.Date(2019, 01, 01, 16, 00, 0, 0, time.UTC), 1, 1, 1, 1)
	gotJson, _ := json.Marshal(gotStruct)
	got := string(gotJson)

	var want = `{"id":1,"dateTime":"2019-01-01T16:00:00Z","queueNumber":1,"userId":1,"cdId":1,"rentQuantity":1}`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
