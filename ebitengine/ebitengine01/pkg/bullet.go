package gopher

import "github.com/hajimehoshi/ebiten/v2"

type Bullet struct {
	PosX   float64
	PosY   float64
	Speed  float64
	Sprite *ebiten.Image
	Width  int
	Height int
}
