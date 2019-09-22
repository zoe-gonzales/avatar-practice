package avatar

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// Generate function creates a png file from image content
func Generate(hsh []byte, filename string) {
	// define points using length
	length := 200
	upLeft := image.Point{0, 0}
	downRight := image.Point{length, length}
	// generate new image from rectangle
	img := image.NewRGBA(image.Rectangle{upLeft, downRight})
	// fill picture using hash values
	colorPicture(hsh, length, img)
	// generate png file
	avatar, _ := os.Create(filename + ".png")
	// encode png with img content
	err := png.Encode(avatar, img)
	if err != nil {
		log.Fatal("Error encoding image to a png", err)
	}
}

// colorPicture function takes a hash, defines the image's colors, and fills it
func colorPicture(hsh []byte, length int, img *image.RGBA) {
	// define colors
	color1 := color.RGBA{hsh[0], hsh[1], hsh[2], 0xff}
	color2 := color.RGBA{hsh[3], hsh[4], hsh[5], 0xff}
	color3 := color.RGBA{hsh[5], hsh[7], hsh[8], 0xff}
	color4 := color.RGBA{hsh[9], hsh[10], hsh[11], 0xff}
	// fill image in with color
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			switch {
			// upper left quadrant
			case x < int(hsh[12]) && y < int(hsh[13]):
				img.Set(x, y, color1)
			// lower left quadrant
			case x < int(hsh[14]) && y > int(hsh[15]):
				img.Set(x, y, color2)
			// lower right quadrant
			case x > int(hsh[16]) && y > int(hsh[17]):
				img.Set(x, y, color3)
			// upper right quadrant
			case x > int(hsh[18]) && y < int(hsh[19]):
				img.Set(x, y, color4)
			// default to black
			default:
				img.Set(x, y, color.Black)
			}
		}
	}
}
