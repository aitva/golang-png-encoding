# Golang PNG Encoding

This repo test PNG compression parameters for Golang's `image/png` library.

## Usage

Build & run the program with: `go run main.go apercu-slack.png`

Example output:

```
original image is 26297 bytes
"DefaultCompression": image is 31368 bytes
"BestSpeed": image is 36923 bytes
"BestCompression": image is 27279 bytes
```
