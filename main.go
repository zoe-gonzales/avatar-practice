package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	// 2 ways to generate hash value using sha1

	// 1st way
	// 1) create new hash
	// 2) use io package to write value to it (added to byte slice)
	em := "bobsmith123@gmail.com"
	h := sha1.New()
	io.WriteString(h, em)
	fmt.Printf("Method 1: \n")
	fmt.Printf("% x", h.Sum(nil))
	fmt.Printf("\n")

	// 2nd way
	// 1) convert string to byte slice
	// 2) run slice through Sum fxn
	email := []byte("bobsmith123@gmail.com")
	fmt.Printf("Method 2: \n")
	fmt.Printf("% x", sha1.Sum(email))
	fmt.Printf("\n")
}

// Given a personal information
// such as an email address, IP address, or a public key,
// generate a unique avatar

// Take input and create a hash from this input
// Use hash to create an image
