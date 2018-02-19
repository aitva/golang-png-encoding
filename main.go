package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/disintegration/imaging"
	"github.com/foobaz/lossypng/lossypng"
)

const folder = "output"

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
	fmt.Println("decoded image:", describeImage(img))
	paletted := img.(*image.Paletted)

	res := imaging.Resize(img, 1024, 0, imaging.NearestNeighbor)
	fmt.Println("resized image:", describeImage(res))
	err = saveImage(res, folder, "reduced.png")
	if err != nil {
		fmt.Println("fail to save image:", err)
		os.Exit(1)
	}

	tmp := image.NewPaletted(res.Bounds(), paletted.Palette)
	draw.Src.Draw(tmp, tmp.Bounds(), res, res.Bounds().Min)
	fmt.Println("indexed image:", describeImage(tmp))
	err = saveImage(tmp, folder, "indexed.png")
	if err != nil {
		fmt.Println("fail to save image:", err)
		os.Exit(1)
	}

	lossy := lossypng.Compress(res, lossypng.NoConversion, 20)
	fmt.Println("compressed image:", describeImage(lossy))
	err = saveImage(tmp, folder, "lossy.png")
	if err != nil {
		fmt.Println("fail to save image:", err)
		os.Exit(1)
	}

	tmp = image.NewPaletted(lossy.Bounds(), paletted.Palette)
	draw.Src.Draw(tmp, tmp.Bounds(), lossy, lossy.Bounds().Min)
	fmt.Println("indexed image:", describeImage(tmp))
	err = saveImage(tmp, folder, "lossy-indexed.png")
	if err != nil {
		fmt.Println("fail to save image:", err)
		os.Exit(1)
	}
}

func describeImage(img image.Image) string {
	var buf bytes.Buffer
	png.Encode(&buf, img)
	point := img.Bounds().Size()
	return fmt.Sprintf("%T, %dx%dpx and %d bytes", img, point.X, point.Y, buf.Len())
}

func saveImage(img image.Image, folder, filename string) error {
	_, err := os.Lstat(folder)
	if err != nil {
		err = os.Mkdir(folder, 0775)
		if err != nil {
			return err
		}
	}

	path := folder + string(os.PathSeparator) + filename
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	err = png.Encode(f, img)
	f.Close()
	return err
}
