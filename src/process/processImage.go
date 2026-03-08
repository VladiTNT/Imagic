package process

import (
	"fmt"
	"image"
	"image/color"
)

func GrayScale(img image.Image) image.Image {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newImg := image.NewGray16(img.Bounds())

	for i := range width {
		for j := range height {
			// Getting RGB values.
			r, g, b, _ := img.At(i, j).RGBA()
			// Getting the max value to use for the gray scale.
			t := max(r, g, b)
			newImg.SetGray16(i, j, color.Gray16{Y: uint16(t)})
		}
	}

	return newImg
}

// Rotate the image 90 degrees.
func Rotate90(img image.Image) image.Image {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newImg := image.NewRGBA64(img.Bounds())

	for i := range width {
		for j := range height {
			newImg.Set(i, j, img.At(j, width-i))
		}
	}

	return newImg
}

// Rotates the image 180 degrees.
func Rotate180(img image.Image) image.Image {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newImg := image.NewRGBA64(img.Bounds())

	for i := range width {
		for j := range height {
			newImg.Set(i, j, img.At(width-i, height-j))
		}
	}

	return newImg
}

// Flips the image.
func Flip(img image.Image) image.Image {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newImg := image.NewRGBA64(img.Bounds())

	for i := range width {
		for j := range height {
			newImg.Set(i, j, img.At(width-i, j))
		}
	}

	return newImg
}

// Crop the image, using subimage.
func Crop(img image.Image, p1, p2 image.Point) (image.Image, error) {
	if p1.X > p2.X {
		return nil, fmt.Errorf("Invalid crop: p1.X = %d > p2.X = %d", p1.X, p2.X)
	}

	if p1.Y > p2.Y {
		return nil, fmt.Errorf("Invalid crop: p1.Y = %d > p2.Y = %d", p1.X, p2.X)
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newImg := image.NewRGBA64(img.Bounds())

	for i := range width {
		for j := range height {
			newImg.Set(i, j, img.At(i, j))
		}
	}

	return newImg.SubImage(image.Rectangle{Min: p1, Max: p2}), nil
}

// Contrast & Brightness
func ContrastAndBrightness(img image.Image, contrast, brightness float64) image.Image {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newImg := image.NewRGBA64(img.Bounds())

	factor := (259 * (contrast + 255)) / (255 * (259 - contrast))

	for i := range width {
		for j := range height {
			c := img.At(i, j)
			r32, g32, b32, a32 := c.RGBA()

			r := float64(r32 >> 8)
			g := float64(g32 >> 8)
			b := float64(b32 >> 8)

			r = factor*(r+brightness-128) + 128
			g = factor*(g+brightness-128) + 128
			b = factor*(b+brightness-128) + 128

			newImg.Set(i, j, color.RGBA{
				R: clamp(r),
				G: clamp(g),
				B: clamp(b),
				A: uint8(a32 >> 8),
			})
		}
	}

	return newImg
}

// Blends the image using a color and the draw package.
func Blend() {

}

// Blurs the image by averaging out the pixels.
func Blur() {

}

// Change format JPEG, PNG and GIF.
func FormatImage() {

}

// Convert to ASCII for ASCII Art.
func ToAscii() {

}

// Use the draw.Draw function to overlay an image like a logo onto the original.
func DrawImage() {

}
