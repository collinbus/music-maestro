package main

import (
	"bufio"
	"fmt"
	"musicMaestro/internal/image"
	"musicMaestro/internal/music"
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
	fmt.Println("==========================")
	fmt.Println(" Welcome to Music Maestro ")
	fmt.Println("==========================")

	fmt.Println("What would you like to do?")
	fmt.Println("1) Update user")
	fmt.Println("2) Update music library")
	fmt.Println("3) Download images")
	fmt.Println("0) Exit")
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
	tokenService := token.NewService()
	trackService := music.NewTrackService(tokenService)
	trackService.FetchUserTracks()
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
