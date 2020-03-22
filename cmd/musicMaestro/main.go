package main

import (
	"os"
)

func main() {
	_, err := ParseArguments(os.Args[1:])
	if err != nil {
		println(err.Error())
	}
}

/*func handleArgument() {
	if os.Args[1] == "-code" {
		saveAccessCode(os.Args[2])
	}
	if os.Args[1] == "-clientId" {
		saveClientId(os.Args[2])
	}
	if os.Args[1] == "-clientSecret" {
		saveClientSecret(os.Args[2])
	}
}*/
