package main

import "testing"

func TestInvalidNumberOfArguments(t *testing.T) {
	args := []string{"first"}

	_, err := HandleArguments(args)

	if err == nil {
		t.Error("There should always be an even amount of arguments")
	}
}

func TestParseKeyValuesOfProvidedArguments(t *testing.T) {
	args := []string{"firstKey", "firstValue"}

	arguments, _ := HandleArguments(args)

	if arguments["firstKey"] != "firstValue" {
		t.Error("Arguments should contain [firstKey]={firstValue}")
	}
}
