package src

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

func GetImageFromFile(path string) (img image.Image, format string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	img, format, err = image.Decode(f)
	return
}

func GetGifFromFile(path string) (g *gif.GIF, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	g, err = gif.DecodeAll(f)
	return
}

func WriteImageToFile(img image.Image, format, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	switch format {
	case "png":
		return png.Encode(f, img)
	case "jpeg":
		return jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
	default:
		return fmt.Errorf("Unknown image format: %s", format)
	}
}

func WriteGifToFile(g *gif.GIF, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return gif.EncodeAll(f, g)
}

func DecodeConfigFromFile(path string) (image.Config, string, error) {
	f, err := os.Open(path)
	if err != nil {
		return image.Config{}, "", err
	}
	defer f.Close()

	return image.DecodeConfig(f)
}
