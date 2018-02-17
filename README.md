# Golang PNG Encoding

This repo test PNG compression parameters for Golang's `image/png` library.

## Usage

You need [dep](https://github.com/golang/dep) and [go](https://github.com/golang/go) install on your machine.

- To download the dependencies: `dep ensure`
- To build & run the program: `go run main.go apercu-slack.png`

Example output:

```
original image is 26297 bytes
decoded image: *image.Paletted, 1280x920px and 31368 bytes
resized image: *image.NRGBA, 1024x736px and 47339 bytes
indexed image: *image.Paletted, 1024x736px and 23617 bytes
compressed image: *image.NRGBA, 1024x736px and 80984 bytes
indexed image: *image.Paletted, 1024x736px and 31352 bytes
```

## Summary

- `png.DefaultCompression` gives the best result
- original image comes with a custom color palette
- reduced image lose the color palette
- image reduced with nearest neighbor have a smaller size
- `lossypng` does not help to get back to the original size
- applying the original color palette to the reduced image is the best 💖
