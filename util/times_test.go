package util

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	type name struct {
		EventTime DateTime `json:"event_time"`
		Name      string   `json:"name"`
	}
	////////////////////////
	marshal, err := json.Marshal(&name{
		EventTime: DateTime{time.Now()},
		Name:      "zwz",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(marshal))
	///////////////////////
	var n name
	err = json.Unmarshal([]byte(marshal), &n)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(n)
}
