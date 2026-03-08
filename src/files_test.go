package src_test

import (
	"testing"

	"github.com/VladiTNT/imagic/src"
)

func TestDecodeConfig(t *testing.T) {
	cfg, format, err := src.DecodeConfigFromFile("./assets/base/test.png")
	if err != nil {
		t.Errorf("Failed to decode config: %v\n", err)
		return
	}

	t.Logf("Image format: %s\n", format)
	t.Logf("Width: %d and Height: %d\n", cfg.Width, cfg.Height)
}
