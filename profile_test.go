package main

import "testing"

func TestUpdateName(t *testing.T) {
	Init()
	if err := UpdateName("アンチうしbot"); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateDescription(t *testing.T) {
	Init()
	if err := UpdateDescription("私はうしbotです"); err != nil {
		t.Fatal(err)
	}
}
