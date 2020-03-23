package main

import (
	"musicMaestro/internal/network"
	"musicMaestro/internal/persistence"
	"os"
)

func main() {
	arguments, err := ParseArguments(os.Args[1:])
	if err != nil {
		println(err.Error())
	}

	if len(arguments) > 0 {
		HandleArguments(arguments)
	}

	startMusicMaestro()
}

func startMusicMaestro() {
	applicationData := persistence.RetrieveApplicationData()

	if applicationData.RefreshToken == "" {
		applicationData = network.RequestApiToken(applicationData)
		persistence.SaveApplicationData(applicationData)
		println("Updated Api Token")
	}

	//TODO application menu
}
