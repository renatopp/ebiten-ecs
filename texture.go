package sk

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Texture struct {
	Image *ebiten.Image
}

func NewTextureFromImage(image *ebiten.Image) *Texture {
	return &Texture{Image: image}
}

func NewTextureFromBytes(data []byte) (*Texture, error) {
	imgData, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return &Texture{Image: nil}, err
	}

	img := ebiten.NewImageFromImage(imgData)
	return &Texture{Image: img}, nil
}

func NewEmptyTexture(width, height int) *Texture {
	return &Texture{Image: ebiten.NewImage(width, height)}
}
