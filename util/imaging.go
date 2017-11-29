package util

import ("image"
	"image/png"
	"os")

func SavePNG(name string, img image.Image) {
	name += ".png"
	f, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
