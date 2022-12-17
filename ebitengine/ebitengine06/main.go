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
	pylonWidth, pylonHeight   = 80, 150
	playerWidth, playerHeight = 80, 80
	gamepadID                 = 1
)

var (
	pylon      Pylon
	p1         player.Player
	speed      float64
	jumpHeight float64
)

type Game struct {
}

type Pylon struct {
	Width  int
	Height int
	PosX   float64
	PosY   float64
	Color  color.RGBA
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	rand.Seed(time.Now().UnixNano())
	randomLength := rand.Intn(screenWidth)
	if pylon.PosX < float64(-pylon.Width) {
		pylon.PosX = float64(screenWidth + randomLength)
	} else {
		pylon.PosX -= speed
	}
	button := ebiten.GamepadButton0
	if inpututil.IsGamepadButtonJustPressed(gamepadID, button) && p1.State == player.Idle {
		p1.Jump()
	}

	for i := 0; i < ebiten.GamepadButtonCount(gamepadID); i++ {
		if inpututil.IsGamepadButtonJustPressed(gamepadID, ebiten.GamepadButton(i)) {
			fmt.Println(ebiten.GamepadButton(i))
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && p1.State == player.Idle {
		p1.Jump()
	}
	switch p1.State {
	case player.Jumping:
		if jumpHeight > 1 {
			jumpHeight /= 1.25
			p1.PosY -= jumpHeight / 3
		} else {
			p1.State = player.Falling
		}
	case player.Falling:
		if jumpHeight < playerHeight*2 {
			jumpHeight *= 1.25
		} else {
			jumpHeight = playerHeight * 2
		}
		if p1.PosY < screenHeight-playerHeight {
			p1.PosY += jumpHeight / 3
		} else {
			p1.PosY = screenHeight - playerHeight
			p1.State = player.Idle
		}
	case player.Idle:

	}
	speed += 0.001
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	upperPylon := ebiten.NewImage(pylon.Width, pylon.Height)
	upperPylon.Fill(pylon.Color)
	upperPylonOptions := ebiten.DrawImageOptions{}
	upperPylonOptions.GeoM.Translate(pylon.PosX, 0)
	screen.DrawImage(upperPylon, &upperPylonOptions)

	bottomPylon := ebiten.NewImage(pylon.Width, pylon.Height)
	bottomPylon.Fill(pylon.Color)
	bottomPylonOptions := ebiten.DrawImageOptions{}
	bottomPylonOptions.GeoM.Translate(pylon.PosX, pylon.PosY)
	screen.DrawImage(bottomPylon, &bottomPylonOptions)

	player1 := ebiten.NewImage(p1.Width, p1.Height)
	player1.Fill(p1.Color)
	player1Options := ebiten.DrawImageOptions{}
	player1Options.GeoM.Translate(p1.PosX, p1.PosY)
	screen.DrawImage(player1, &player1Options)
}

func init() {
	speed = 6
	pylon = Pylon{
		Width:  pylonWidth,
		Height: pylonHeight,
		PosX:   screenWidth - pylonWidth,
		PosY:   screenHeight - pylonHeight,
		Color:  color.RGBA{R: 0x88, G: 0xff, B: 0xff, A: 0xff}}
	p1 = player.Player{
		Width:  playerWidth,
		Height: playerHeight,
		PosX:   playerWidth,
		PosY:   screenHeight - playerHeight,
		Color:  color.RGBA{R: 0x88, G: 0x88, B: 0xff, A: 0xff},
		State:  player.Idle}
	jumpHeight = playerHeight * 2
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebitengine Experiment 06")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
