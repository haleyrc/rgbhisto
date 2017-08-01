package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
)

func main() {
	r, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("open:", err)
	}
	defer r.Close()

	img, _, err := image.Decode(r)
	if err != nil {
		log.Fatalln("decode:", err)
	}

	bounds := img.Bounds()

	red := make([]int64, 256)
	blue := make([]int64, 256)
	green := make([]int64, 256)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			red[r>>8]++
			blue[b>>8]++
			green[g>>8]++
		}
	}

	f, err := os.Create("histo.csv")
	if err != nil {
		log.Fatalln("create:", err)
	}

	io.WriteString(f, "red,blue,green\n")
	for i := 0; i < 256; i++ {
		fmt.Fprintf(f, "%d,%d,%d\n", red[i], blue[i], green[i])
	}
}
