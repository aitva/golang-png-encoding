# Golang PNG Encoding

This repo test PNG compression parameters for Golang's `image/png` library.

## Usage

You need [dep](https://github.com/golang/dep) and [go](https://github.com/golang/go) install on your machine.

To download the dependencies: `dep ensure`
To build & run the program: `go run main.go apercu-slack.png`

Example output:

```
original image is 26297 bytes

*image.Paletted of 1280x920px:
"DefaultCompression": image is 31368 bytes
"BestSpeed": image is 36923 bytes
"BestCompression": image is 27279 bytes

*image.NRGBA of 1024x736px:
"DefaultCompression": image is 86268 bytes
"BestSpeed": image is 91306 bytes
"BestCompression": image is 84924 bytes

lossypng *image.NRGBA of 1024x736px:
"DefaultCompression": image is 39282 bytes
"BestSpeed": image is 44420 bytes
"BestCompression": image is 37317 bytes
```

## Summary

- The original image comes with a color palette, which greatly reduce its size.
- Reducing the image with `imaging` removes the color palette and double the image size.
- Passing the reduced image to `lossypng` helps but we don't even get close to the original size.
