package main

import (
	"fmt"
	"os"

	"github.com/zoe-gonzales/avatar-practice/avatar"
	"github.com/zoe-gonzales/avatar-practice/hash"
)

func main() {
	fmt.Println()
	e := os.Args[1]
	hash.Email(e)
	avatar.Generate()
}

// Given a personal information
// such as an email address, IP address, or a public key,
// generate a unique avatar

// Take input and create a hash from this input
// Use hash to create an image

// Creating an image
// make rectangle
// define a color palette to use in generating the image
// use NewRGBA function to create image from a rectangle
