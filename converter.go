package main


import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"log"
	"strconv"
)

func main() {
	file, _ := os.Open("./test.png")
	defer file.Close()

	bounds := file.Bounds()
	fmt.Println(bounds)
	config, format, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("image format:" + format)
	fmt.Println("width:" + strconv.Itoa(config.Width) + ", height:" + strconv.Itoa(config.Height))
}