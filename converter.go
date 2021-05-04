package main


import (
	"fmt"
	"image"
	_ "image/png"
	_ "image/jpeg"
	"os"
	"log"
)

func main() {
	file, _ := os.Open("./sample.jpeg")
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	fmt.Println(width,height)

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row,convToPixel(img.At(x,y).RGBA()))
		}
		pixels = append(pixels,row)
	}

	var avg_array [][]string
	dest := image.NewGray16(bounds)
	for y := 0; y < height; y++ {
		var row []string
		for x := 0; x < width; x++ {
			row = append(row,convToAVG(pixels[y][x]))
		}
		avg_array = append(avg_array,row)
	}

	for y := 0; y < height; y++ {
		fmt.Println(avg_array[y][0:width])
	}
}

func convToPixel(r uint32,g uint32,b uint32,a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

func convToAVG(target Pixel) string {
	avg_data := (float64(target.R*0.3))+(float64(target.G*0.59))+(float64(target.B*0.11)) / 3
	if avg_data > 128 {
		return "."
	} else if avg_data == 0 {
		return " "
	} else {
		return ":"
	}
}

type Pixel struct {
	R int
	G int
	B int
	A int
}