package main

import (
	"cyber-shield/lib"
	"fmt"
)

func main() {
	auth := lib.AuthOptions{
		Hostname: "127.0.0.1",
		Port:     17234,

		License: "<LICENSE>",
	}

	license, err := auth.Auth()

	if err != nil {
		panic(err)
	}

	fmt.Println("Succesfully authenticated!", license.Rank)
}
