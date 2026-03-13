package process

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

func GrayScale(img image.Image) image.Image {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newImg := image.NewGray16(img.Bounds())

	for i := range width {
		for j := range height {
			newImg.SetGray16(i, j, averageGray(img.At(i, j)))
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
			r, g, b, a := (bitShiftCorrection(c.RGBA()))

			newR := factor*(float64(r)+brightness-128) + 128
			newG := factor*(float64(g)+brightness-128) + 128
			newB := factor*(float64(b)+brightness-128) + 128

			newImg.Set(i, j, color.RGBA{
				R: clamp(newR),
				G: clamp(newG),
				B: clamp(newB),
				A: uint8(a),
			})
		}
	}

	return newImg
}

// Blends the image using a color and the draw package.
func Blend(img image.Image, c color.Color) image.Image {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newImg := image.NewRGBA64(img.Bounds())

	for i := range width {
		for j := range height {
			oldR, oldG, oldB, oldA := bitShiftCorrection(img.At(i, j).RGBA())
			r, g, b, _ := bitShiftCorrection(c.RGBA())

			newImg.Set(i, j, color.RGBA{
				R: clamp(float64((oldR + r) / 2)),
				G: clamp(float64((oldG + g) / 2)),
				B: clamp(float64((oldB + b) / 2)),
				A: uint8(oldA),
			})
		}
	}

	return newImg
}

// Blurs the image by averaging out the pixels.
func Blur(img image.Image) image.Image {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newImg := image.NewRGBA64(img.Bounds())

	for i := range width {
		for j := range height {
			var colors []color.Color

			if i-1 >= 0 && j-1 >= 0 {
				colors = append(colors, img.At(i-1, j-1))
			}

			if j-1 >= 0 {
				colors = append(colors, img.At(i, j-1))
			}

			if i+1 < width && j-1 >= 0 {
				colors = append(colors, img.At(i+1, j-1))
			}

			if i-1 >= 0 {
				colors = append(colors, img.At(i-1, j))
			}

			colors = append(colors, img.At(i, j))

			if i+1 <= width {
				colors = append(colors, img.At(i+1, j))
			}

			if i-1 >= 0 && j+1 < height {
				colors = append(colors, img.At(i-1, j+1))
			}

			if j+1 < height {
				colors = append(colors, img.At(i, j+1))
			}

			if i+1 < width && j+1 < height {
				colors = append(colors, img.At(i+1, j+1))
			}

			newImg.Set(i, j, averageColor(colors))
		}
	}

	return newImg
}

// Convert to ASCII for ASCII Art.
func ToAscii(img image.Image) []byte {
	size := img.Bounds().Size()
	capacity := size.X*size.Y + size.Y
	buf := make([]byte, 0, capacity)

	for i := range size.X {
		for j := range size.Y {
			y := averageGray(img.At(j, i)).Y

			switch {
			case y < 16000:
				buf = append(buf, '@')
			case y >= 16000 && y < 32000:
				buf = append(buf, '#')
			case y >= 32000 && y < 40000:
				buf = append(buf, '8')
			case y >= 40000 && y < 50000:
				buf = append(buf, '-')
			case y >= 50000 && y < 60000:
				buf = append(buf, ',')
			case y >= 60000 && y < 65535:
				buf = append(buf, '.')
			}
		}
		buf = append(buf, '\n')
	}

	return buf
}

// DrawImage draws the subject image into the target image using Go's standard image drawing
// tools.
//
// TopLeft represents the point inside target from which subject will start being drawn inside
// of target.
func DrawImage(target, subject image.Image, TopLeft image.Point) image.Image {
	newImg := image.NewRGBA64(target.Bounds())

	for i := range newImg.Bounds().Dx() {
		for j := range newImg.Bounds().Dx() {
			newImg.Set(i, j, target.At(i, j))
		}
	}

	// Width and height of the subject image.
	width := subject.Bounds().Dx()
	height := subject.Bounds().Dy()
	// Area in the target where the subject will be drawn.
	area := image.Rect(TopLeft.X, TopLeft.Y, TopLeft.X+width, TopLeft.Y+height)
	// Draw function.
	draw.Over.Draw(newImg, area, subject, image.Pt(0, 0))

	return newImg
}
