package player

import (
	"fmt"
	"image/color"
)

type Player struct {
	Width  int
	Height int
	PosX   float64
	PosY   float64
	Color  color.RGBA
	State  State
}

func (p *Player) Jump() {
	p.State = Jumping
	fmt.Println("Jump!")
}
