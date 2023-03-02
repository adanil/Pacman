package utility

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"os"
)

func GetFont(baseFont *opentype.Font, size, dpi float64) (font.Face, error) {
	return opentype.NewFace(baseFont, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func ReadImage(imagePath string) (*ebiten.Image, error) {
	f, err := os.Open(imagePath)
	//nolint:staticcheck
	defer f.Close()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return ebiten.NewImageFromImage(img), nil
}
