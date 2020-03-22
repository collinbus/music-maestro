package main

import "testing"

func TestInvalidNumberOfArguments(t *testing.T) {
	args := []string{"first"}

	_, err := HandleArguments(args)

	if err == nil {
		t.Error("There should always be an even amount of arguments")
	}
}
