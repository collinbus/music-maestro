package main

import (
	"errors"
	"musicMaestro/internal/persistence"
)

func ParseArguments(args []string) (map[string]string, error) {
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

func HandleArguments(args map[string]string) {
	applicationData := parseApplicationData(args)
	appDataService := persistence.NewApplicationDataService()
	appDataService.SaveApplicationData(applicationData)
}

func parseApplicationData(args map[string]string) *persistence.ApplicationData {
	accessCode := ""
	clientId := ""
	clientSecret := ""
	if argumentsContains("-accessCode", args) {
		accessCode = args["-accessCode"]
	}
	if argumentsContains("-clientId", args) {
		clientId = args["-clientId"]
	}
	if argumentsContains("-clientSecret", args) {
		clientSecret = args["-clientSecret"]
	}
	return persistence.NewApplicationData(accessCode, clientId, clientSecret)
}

func argumentsContains(key string, args map[string]string) bool {
	_, hasValue := args[key]
	return hasValue
}
