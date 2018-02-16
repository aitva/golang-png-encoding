package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
)

func usage() {
	fmt.Printf("usage: %s image.png\n", os.Args[0])
	fmt.Printf("\nRead a PNG image and try different compression algorithm on it.\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("fail to open file:", err)
		os.Exit(1)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		fmt.Println("fail to read file info:", err)
		os.Exit(1)
	}
	fmt.Printf("original image is %d bytes\n", stat.Size())

	img, err := png.Decode(f)
	if err != nil {
		fmt.Println("fail to open image:", err)
		os.Exit(1)
	}
	testEncodings(img)
}

func testEncodings(img image.Image) {
	compressions := map[string]png.CompressionLevel{
		"DefaultCompression": png.DefaultCompression,
		"BestSpeed":          png.BestSpeed,
		"BestCompression":    png.BestCompression,
	}
	var buf bytes.Buffer
	var enc png.Encoder
	for k, v := range compressions {
		enc.CompressionLevel = v
		err := enc.Encode(&buf, img)
		if err != nil {
			fmt.Printf("%q: %v\n", k, err)
			continue
		}
		fmt.Printf("%q: image is %d bytes\n", k, buf.Len())
		buf.Reset()
	}
}
