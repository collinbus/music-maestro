package main

import (
	"bufio"
	"fmt"
	"musicMaestro/internal/image"
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
	println("3) Download images")
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
	case "3":
		downloadImages()
	default:
		fmt.Println("Invalid choice")
	}
}

func exit() {
	println("Quitting Music Maestro")
	os.Exit(0)
}

func updateUser() {
	tokenService := token.NewService()
	userService := user.NewService(tokenService)

	userService.UpdateCurrentUserFromServer()
}

func updateMusicLibrary() {
	println("Update Music Library")
	// TODO update music library
}

func downloadImages() {
	tokenService := token.NewService()
	userService := user.NewService(tokenService)
	imageService := image.NewService()

	fetchedUser := userService.FetchUser()

	imgData := imageService.DownloadImage(fetchedUser.Image.Url)
	fetchedUser.Image.Data = imgData

	userService.UpdateUser(fetchedUser)
}
