package main

import (
	"fmt"
	"testing"
)

func TestEscapeProcessing(t *testing.T) {
	Init()
	a := map[string]interface{}{"latte": "\"aaa\""}
	PostDMToMe(fmt.Sprintf("%#v", a))
}

func TestGetDetailedTime(t *testing.T) {
	fmt.Println(GetDetailedTime())
}
