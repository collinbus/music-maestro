package main

import "errors"

func HandleArguments(args []string) (map[string]string, error) {
	numberOfArguments := len(args)
	arguments := make(map[string]string)
	if numberOfArguments%2 != 0 {
		return nil, errors.New("No value found for argument: " + args[numberOfArguments-1])
	}

	for i := 0; i < numberOfArguments; i += 2 {
		arguments[args[i]] = args[i+1]
	}

	return arguments, nil
}
