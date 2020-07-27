package main

import "testing"

func TestUpdateName(t *testing.T) {
	Init()
	UpdateName("アンチうしbot")
}

func TestUpdateDescription(t *testing.T) {
	Init()
	UpdateDescription("私はうしbotです")
}
