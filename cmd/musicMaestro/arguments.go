package main

import "errors"

func HandleArguments(args []string) (map[string]string, error) {
	numberOfArguments := len(args)
	if numberOfArguments%2 != 0 {
		return nil, errors.New("No value found for argument: " + args[numberOfArguments-1])
	}
	return nil, nil
}
