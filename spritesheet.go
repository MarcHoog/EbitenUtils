package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type SpriteSheet struct {
	sprites []*BasicSprite
}

func NewBasicSpriteSheet(sheet *ebiten.Image, width, height int, pos image.Point, blanks bool) *SpriteSheet {

	sprites := make([]*ebiten.Image, 0)

	b := sheet.Bounds()
	w, h := b.Dx(), b.Dy()
	for y := 0; y < h; y += height {
		for x := 0; x < w; x += width {
			rect := image.Rect(x, y, x+width, y+height)
			sub := sheet.SubImage(rect).(*ebiten.Image)
			sprite := ebiten.NewImageFromImage(sub)
			sprites = append(sprites, sprite)
		}

	}

	return &SpriteSheet{}
}
