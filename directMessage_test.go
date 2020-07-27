package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPostDMToMe(t *testing.T) {
	Init()

	blob := []byte(`{"latte":"malta"}`)
	var val interface{}
	json.Unmarshal(blob, &val)
	PostDMToMe(fmt.Sprintf("%#v", val))
}
