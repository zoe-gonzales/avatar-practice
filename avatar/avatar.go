package avatar

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// Generate function creates a png file from image content
func Generate(hsh []byte) {
	// define points using length
	length := 200
	upLeft := image.Point{0, 0}
	downRight := image.Point{length, length}
	// generate new image
	img := image.NewRGBA(image.Rectangle{upLeft, downRight})
	// define colors
	color1 := color.RGBA{hsh[0], hsh[1], hsh[2], 0xff}
	color2 := color.RGBA{hsh[3], hsh[4], hsh[5], 0xff}
	color3 := color.RGBA{hsh[5], hsh[7], hsh[8], 0xff}
	color4 := color.RGBA{hsh[9], hsh[10], hsh[11], 0xff}

	for j := 0; j < len(hsh); j++ {
		fmt.Printf("%v \n", hsh[j]) // uint8
		b := int(hsh[j])
		// fill image in with color
		for x := 0; x < length; x++ {
			for y := 0; y < length; y++ {
				switch {
				// upper left quadrant
				case x < b && y < b:
					img.Set(x, y, color1)
				// lower left quadrant
				case x < b && y > b:
					img.Set(x, y, color2)
				// lower right quadrant
				case x > b && y > b:
					img.Set(x, y, color3)
				// upper right quadrant
				case x > b && y < b:
					img.Set(x, y, color4)
				// default to black
				default:
					img.Set(x, y, color.Black)
				}
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
