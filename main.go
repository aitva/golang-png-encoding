package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/disintegration/imaging"
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
	fmt.Printf("\nsize %s:\n", imageSizeToString(img))
	testEncodings(img)

	img = imaging.Resize(img, 1024, 0, imaging.Box)

	fmt.Printf("\nsize %s:\n", imageSizeToString(img))
	testEncodings(img)
}

func imageSizeToString(img image.Image) string {
	point := img.Bounds().Size()
	return fmt.Sprintf("%dx%dpx", point.X, point.Y)
}

func compToString(c png.CompressionLevel) string {
	str := "InvalidCompression"
	switch c {
	case png.DefaultCompression:
		str = "DefaultCompression"
	case png.BestSpeed:
		str = "BestSpeed"
	case png.BestCompression:
		str = "BestCompression"
	}
	return str
}

func testEncodings(img image.Image) {
	compressions := []png.CompressionLevel{
		png.DefaultCompression,
		png.BestSpeed,
		png.BestCompression,
	}
	var buf bytes.Buffer
	var enc png.Encoder
	for _, c := range compressions {
		enc.CompressionLevel = c
		err := enc.Encode(&buf, img)
		if err != nil {
			fmt.Printf("%q: %v\n", compToString(c), err)
			continue
		}
		fmt.Printf("%q: image is %d bytes\n", compToString(c), buf.Len())
		buf.Reset()
	}
}
