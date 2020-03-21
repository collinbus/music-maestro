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
	if os.Args[1] == "-clientId" {
		saveClientId(os.Args[2])
	}
	if os.Args[1] == "-clientSecret" {
		saveClientSecret(os.Args[2])
	}
}

func saveAccessCode(accessCode string) {
	accessCodeHandler := authorization.NewAccessCodeService(authorization.ApplicationDataFileHandler{})
	accessCodeHandler.SaveAccessCode(accessCode)
}

func saveClientId(id string) {
	accessCodeHandler := authorization.NewAccessCodeService(authorization.ApplicationDataFileHandler{})
	accessCodeHandler.SaveClientId(id)
}

func saveClientSecret(secret string) {
	accessCodeHandler := authorization.NewAccessCodeService(authorization.ApplicationDataFileHandler{})
	accessCodeHandler.SaveClientSecret(secret)
}
