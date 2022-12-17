package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var (
	Sprites      *ebiten.Image
	err          error
	AnimationFPS uint64
)

func init() {
	Sprites, _, err = ebitenutil.NewImageFromFile("assets/sprites.png")
	if err != nil {
		log.Fatal(err)
	}
	AnimationFPS = 6
}
