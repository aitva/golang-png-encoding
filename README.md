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
- image reduced with nearest neighbor have a smaller size:
    - indexed image with box: 26174 bytes
    - indexed image with nearest neighbor: 23617 bytes
    - compressed image with box: 98563 bytes
    - compressed image with nearest neighbor: 80984 bytes
- `lossypng` does not help to get back to the original size
- applying the original color palette to the reduced image is the best :purple_heart:

## Images

Encoding | Image
---------|------
Reduced | ![Gryzzly Slack Reduced](output/reduced.png)
Reduced & Indexed | ![Grzzly Slack Indexed](output/indexed.png)
Reduced & Lossy | ![Gryzzly Slack Lossy](output/lossy.png)
Reduced, Lossy & Indexed | ![Gryzzly Slack Lossy Indexed](output/lossy-indexed.png)

