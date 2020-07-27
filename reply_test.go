package main

import (
	"encoding/json"
	"testing"
)

func TestReply(t *testing.T) {
	Init()
	str := `{"latte":"malta"}`
	var val interface{}
	json.Unmarshal([]byte(str), &val)
	if err := Reply(val); err != nil {
		t.Fatal(err)
	}
}
