package process_test

import (
	"image"
	"image/color"
	"testing"

	"github.com/VladiTNT/imagic/src"
	"github.com/VladiTNT/imagic/src/process"
)

func TestGrayScale(t *testing.T) {
	img, format, err := src.GetImageFromFile("../assets/base/test.png")
	if err != nil {
		t.Errorf("Failed to load image: %v\n", err)
		return
	}

	err = src.WriteImageToFile(process.GrayScale(img), format, "../assets/test/grayscale.png")
	if err != nil {
		t.Errorf("Failed to write image: %v\n", err)
	}
}

func TestRotation90(t *testing.T) {
	img, format, err := src.GetImageFromFile("../assets/base/test.png")
	if err != nil {
		t.Errorf("Failed to load image: %v\n", err)
		return
	}

	err = src.WriteImageToFile(process.Rotate90(img), format, "../assets/test/rotate90.png")
	if err != nil {
		t.Errorf("Failed to write image: %v\n", err)
	}
}

func TestRotation180(t *testing.T) {
	img, format, err := src.GetImageFromFile("../assets/base/test.png")
	if err != nil {
		t.Errorf("Failed to load image: %v\n", err)
		return
	}

	err = src.WriteImageToFile(process.Rotate180(img), format, "../assets/test/rotate180.png")
	if err != nil {
		t.Errorf("Failed to write image: %v\n", err)
	}
}

func TestFlip(t *testing.T) {
	img, format, err := src.GetImageFromFile("../assets/base/test.png")
	if err != nil {
		t.Errorf("Failed to load image: %v\n", err)
		return
	}

	err = src.WriteImageToFile(process.Flip(img), format, "../assets/test/flip.png")
	if err != nil {
		t.Errorf("Failed to write image: %v\n", err)
	}
}

func TestCrop(t *testing.T) {
	img, format, err := src.GetImageFromFile("../assets/base/test.png")
	if err != nil {
		t.Errorf("Failed to load image: %v\n", err)
		return
	}

	newImg, err := process.Crop(img, image.Point{42, 42}, image.Point{85, 85})
	if err != nil {
		t.Errorf("Failed to crop image: %v\n", err)
		return
	}

	err = src.WriteImageToFile(newImg, format, "../assets/test/crop.png")
	if err != nil {
		t.Errorf("Failed to write image: %v\n", err)
	}
}

func TestContrast(t *testing.T) {
	img, format, err := src.GetImageFromFile("../assets/base/test.png")
	if err != nil {
		t.Errorf("Failed to load image: %v\n", err)
		return
	}

	err = src.WriteImageToFile(process.ContrastAndBrightness(img, 50, 20), format, "../assets/test/contrast.png")
	if err != nil {
		t.Errorf("Failed to write image: %v\n", err)
		return
	}
}

func TestBlend(t *testing.T) {
	img, format, err := src.GetImageFromFile("../assets/base/test.png")
	if err != nil {
		t.Errorf("Failed to load image: %v\n", err)
		return
	}

	err = src.WriteImageToFile(process.Blend(img, color.RGBA{R: 0, G: 0, B: 255, A: 255}), format, "../assets/test/blend.png")
	if err != nil {
		t.Errorf("Failed to write image: %v\n", err)
		return
	}
}
