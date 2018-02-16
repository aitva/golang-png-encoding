# Golang PNG Encoding

This repo test PNG compression parameters for Golang's `image/png` library.

## Usage

You need [dep](https://github.com/golang/dep) and [go](https://github.com/golang/go) install on your machine.

To download the dependencies: `dep ensure`
To build & run the program: `go run main.go apercu-slack.png`

Example output:

```
original image is 26297 bytes

size 1280x920px:
"DefaultCompression": image is 31368 bytes
"BestSpeed": image is 36923 bytes
"BestCompression": image is 27279 bytes

size 1024x736px:
"DefaultCompression": image is 60739 bytes
"BestSpeed": image is 62443 bytes
"BestCompression": image is 58774 bytes
```

## Todo

- test another PNG encoding, see https://github.com/foobaz/lossypng
