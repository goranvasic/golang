package main

import (
	player "ebitengine06/pkg"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"log"
	"math/rand"
	"time"
)

const (
	screenWidth, screenHeight = 640, 480
	playerWidth, playerHeight = 80, 80
	gamepadID                 = 1
)

var (
	p1 player.Player
)

type Game struct {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < ebiten.GamepadButtonCount(gamepadID); i++ {
		if inpututil.IsGamepadButtonJustPressed(gamepadID, ebiten.GamepadButton(i)) {
			fmt.Println(ebiten.GamepadButton(i))
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	player1 := ebiten.NewImage(p1.Width, p1.Height)
	player1.Fill(p1.Color)
	player1Options := ebiten.DrawImageOptions{}
	player1Options.GeoM.Translate(p1.PosX, p1.PosY)
	screen.DrawImage(player1, &player1Options)
}

func init() {
	p1 = player.Player{
		Width:  playerWidth,
		Height: playerHeight,
		PosX:   playerWidth,
		PosY:   screenHeight - playerHeight,
		Color:  color.RGBA{R: 0x88, G: 0x88, B: 0xff, A: 0xff},
		State:  player.Idle}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebitengine Experiment 08")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
