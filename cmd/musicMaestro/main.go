package main

import (
	"bufio"
	"fmt"
	"musicMaestro/internal/persistence"
	"musicMaestro/internal/token"
	"musicMaestro/internal/user"
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
	for {
		printMenu()
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		handleChoice(scanner.Text())
	}
}

func printMenu() {
	println("==========================")
	println(" Welcome to Music Maestro ")
	println("==========================")

	println("What would you like to do?")
	println("1) Update user")
	println("2) Update music library")
	println("0) Exit")
}

func handleChoice(choice string) {
	switch choice {
	case "0":
		exit()
	case "1":
		updateUser()
	case "2":
		updateMusicLibrary()
	default:
		fmt.Println("Invalid choice")
	}
}

func exit() {
	println("Quitting Music Maestro")
	os.Exit(0)
}

func updateUser() {
	appDataService := persistence.NewApplicationDataService()
	tokenService := token.NewService()
	userService := user.NewService(appDataService, tokenService)

	userService.UpdateCurrentUser()
}

func updateMusicLibrary() {
	println("Update Music Library")
	// TODO update music library
}
