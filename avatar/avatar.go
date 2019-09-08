package avatar

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// Generate function creates a png file from image content
func Generate() {
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
