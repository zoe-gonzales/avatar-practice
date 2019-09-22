package main

import (
	"log"
	"os"

	"github.com/zoe-gonzales/avatar-practice/avatar"
	"github.com/zoe-gonzales/avatar-practice/hash"
)

// given a personal information
// such as an email address, IP address, or a public key,
// generate a unique avatar

func main() {
	if len(os.Args) != 2 {
		log.Fatal("\n Error: You must provide input to create an avatar. \n Please provide a single input without spaces: username, email, etc.")
	}
	// Take input (CL arg) and create a hash from this input
	h := hash.Email(os.Args[1])
	// Use hash to create an image
	avatar.Generate(h, "user-1")
}
