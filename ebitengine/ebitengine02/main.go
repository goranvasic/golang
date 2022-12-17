package main

import (
	"fmt"
	"github.com/goranvasic/ebitengine02/assets"
	"github.com/hajimehoshi/ebiten/v2"
	image2 "image"
	"image/color"
	"log"
)

var (
	image *ebiten.Image
)

const (
	imageSize   = 32
	numOfFrames = 8
)

var (
	x       = 255
	flip    = true
	options = &ebiten.DrawImageOptions{}
)

type Game struct {
	tick uint64
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(image, nil)
	animationSpeed := 60 / assets.AnimationFPS
	frameNum := int(g.tick/animationSpeed) % numOfFrames
	frameX := frameNum * imageSize
	rect := image2.Rect(frameX, 0, frameX+imageSize, imageSize)
	subImg := assets.Sprites.SubImage(rect)
	screen.DrawImage(subImg.(*ebiten.Image), options)
}

func (g *Game) Update() error {
	if flip {
		x -= 3
	} else {
		x += 3
	}
	if x <= 1 {
		flip = false
		x = 1
	}
	if x >= 255 {
		flip = true
		x = 255
	}
	image = ebiten.NewImage(x, 100)
	image.Fill(color.RGBA{R: uint8(x), A: 0xff})
	g.tick++
	return nil
}

func main() {
	ebiten.SetWindowSize(640, 400)
	ebiten.SetWindowTitle("My 1st Game!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func init() {
	fmt.Println("Hello World")
}
