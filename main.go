package main

import (
	"os"

	"github.com/zoe-gonzales/avatar-practice/avatar"
	"github.com/zoe-gonzales/avatar-practice/hash"
)

// given a personal information
// such as an email address, IP address, or a public key,
// generate a unique avatar

func main() {
	e := os.Args[1]
	// Take input (CL arg) and create a hash from this input
	h := hash.Email(e)
	// Use hash to create an image
	avatar.Generate(h)
}
