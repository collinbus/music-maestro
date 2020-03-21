package main

import (
	"musicMaestro/internal/authorization"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		handleArgument()
	}
}

func handleArgument() {
	if os.Args[1] == "-code" {
		saveAccessCode(os.Args[2])
	}
}

func saveAccessCode(accessCode string) {
	accessCodeHandler := authorization.NewAccessCodeService(authorization.AccessCodeFileHandler{})
	accessCodeHandler.Save(accessCode)
}
