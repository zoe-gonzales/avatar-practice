package main

import (
	"crypto/sha1"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	generateAvatar()
}

// Given a personal information
// such as an email address, IP address, or a public key,
// generate a unique avatar

// Take input and create a hash from this input
// Use hash to create an image

// Creating an image
// make rectangle
// define a color palette to use in generating the image
// use NewPaletted function to create an image from the rectangle and chosen color palette

// use NewRGBA function to create image from a rectangle

// 2 ways to generate hash value using sha1

// 1st way
func hashText1(t string) []byte {
	// 1) create new hash
	h := sha1.New()
	// 2) use io package to write value to it (added to byte slice)
	io.WriteString(h, t)
	fmt.Printf("Method 1: \n")
	fmt.Printf("% x", h.Sum(nil))
	fmt.Printf("\n")
	return h.Sum(nil)
}

// 2nd way
func hashText2(t string) [20]byte {
	// 1) convert string to byte slice
	txt := []byte(t)
	fmt.Printf("Method 2: \n")
	// 2) run slice through Sum fxn
	fmt.Printf("% x", sha1.Sum(txt))
	fmt.Printf("\n")
	return sha1.Sum(txt)
}

// Create square png using image pkg
func generateAvatar() {
	// define points using length
	length := 200
	upLeft := image.Point{0, 0}
	downRight := image.Point{length, length}
	// generate new image
	img := image.NewRGBA(image.Rectangle{upLeft, downRight})
	// define colors
	plum := color.RGBA{145, 61, 136, 0xff}
	black := color.RGBA{0, 0, 0, 0xff}
	purple := color.RGBA{140, 20, 252, 0xff}
	white := color.RGBA{255, 255, 255, 0xff}

	// fill image in with color
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			switch {
			// upper left quadrant
			case x < length/2 && y < length/2:
				img.Set(x, y, plum)
			// lower left quadrant
			case x < length/2 && y > length/2:
				img.Set(x, y, purple)
			// lower right quadrant
			case x >= length/2 && y >= length/2:
				img.Set(x, y, black)
			// upper right quadrant
			case x > length/2 && y < length/2:
				img.Set(x, y, white)
			// default to black
			default:
				img.Set(x, y, black)
			}
		}
	}
	// generate png file
	avatar, _ := os.Create("avatar.png")
	// encode png with img content
	err := png.Encode(avatar, img)
	if err != nil {
		log.Fatal("Error encoding image to a png", err)
	}
}
