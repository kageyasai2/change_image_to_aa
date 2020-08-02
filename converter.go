package main


import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"log"
)

func main() {
	file, _ := os.Open("./test.png")
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

	var avg_array [][]int
	for y := 0; y < height; y++ {
		var row []int
		for x := 0; x < width; x++ {
			row = append(row,convToAVG(pixels[y][x]))
		}
		avg_array = append(avg_array,row)
	}

	fmt.Println(avg_array[0][0])
}

func convToPixel(r uint32,g uint32,b uint32,a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

func convToAVG(x Pixel) int {
	return 1
}

type Pixel struct {
	R int
	G int
	B int
	A int
}