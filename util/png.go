package util

import ("image"
	"image/color"
	"image/png"
	"os")

func SavePNG(name string, img image.Image) {
name += ".png"
f, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0600)
defer f.Close()
png.Encode(f, img)
}

const WIDTH, HEIGHT int = 512, 256

/*
func main() {
img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))

//Loop through whole Image
for x := 0; x < WIDTH ; x++ {
	for y := 0; y < HEIGHT ; y++ {
		X := float32(x)/float32(WIDTH)
		Y := float32(y)/float32(HEIGHT)
		
		//Point Operations here
		col := RGBFloat1(X*Y)
		img.Set(x,y, col)
	}
}

SavePNG("test", img)

}
*/

func RGBFloat3(r,g,b float32) color.RGBA {
return color.RGBA{uint8(r*255),uint8(g*255),uint8(b*255),255}
}

func RGBFloat1(v float32) color.RGBA {
if v > 0 {
	return color.RGBA{uint8(v*255),uint8(v*255),uint8(v*255),255}
	}
	return color.RGBA{0,0,0,255}
}

func GrayFloat(v float32) color.Gray {
return color.Gray{uint8(v*255)}
}