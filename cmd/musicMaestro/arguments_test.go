package main

import "testing"

func TestInvalidNumberOfArguments(t *testing.T) {
	args := []string{"first"}

	_, err := ParseArguments(args)

	if err == nil {
		t.Error("There should always be an even amount of arguments")
	}
}

func TestParseKeyValuesOfProvidedArguments(t *testing.T) {
	args := []string{"firstKey", "firstValue"}

	arguments, _ := ParseArguments(args)

	if arguments["firstKey"] != "firstValue" {
		t.Error("Arguments should contain [firstKey]={firstValue}")
	}
}

func TestWithThreeArguments(t *testing.T) {
	args := []string{"first", "second", "third"}

	arguments, err := ParseArguments(args)

	if err == nil {
		t.Error("There should always be an even amount of arguments")
	}

	if arguments != nil {
		t.Error("No arguments should be returned")
	}
}
