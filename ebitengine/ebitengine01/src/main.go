package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	gopher "goranvasic/ebitengine/experiment/pkg"
	"image/color"
	"log"
)

const (
	screenWidth, screenHeight = 1024, 768
	maxSpeed                  = 10
	gamepadID                 = 1
)

var (
	player  gopher.Gopher
	green   = color.RGBA{G: 255, A: 99}
	bullets []gopher.Bullet
)

type MyGame struct {
	pressedKeys []ebiten.Key
}

func init() {
	var err error
	player.Sprite, _, err = ebitenutil.NewImageFromFile("assets/gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	player.Width, player.Height = player.Sprite.Size()
	player.PosX = float64(screenWidth/2 - player.Width/2)
	player.PosY = float64(screenHeight - player.Height)
	player.Speed = maxSpeed
	player.InitSfx()
}

func (g *MyGame) Update() error {
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])
	if len(g.pressedKeys) > 0 {
		key := g.pressedKeys[0]
		if key.String() == "ArrowRight" {
			if player.PosX <= float64(screenWidth-player.Width) {
				player.PosX += player.Speed
			}
		}
		if key.String() == "ArrowLeft" {
			if player.PosX >= 0 {
				player.PosX -= player.Speed
			}
		}
	}

	for i := 0; i < ebiten.GamepadButtonCount(gamepadID); i++ {
		if inpututil.IsGamepadButtonJustPressed(gamepadID, ebiten.GamepadButton(i)) {
			fmt.Println(ebiten.GamepadButton(i))
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsGamepadButtonJustPressed(gamepadID, ebiten.GamepadButton0) {
		var bullet gopher.Bullet
		bullet.Width = 20
		bullet.Height = 100
		bullet.Sprite = ebiten.NewImage(bullet.Width, bullet.Height)
		bullet.Sprite.Fill(green)
		bullet.PosX = player.PosX + float64(player.Width/2-bullet.Width/2)
		bullet.PosY = player.PosY
		bullet.Speed = maxSpeed
		bullets = append(bullets, bullet)
		err := player.Fire(bullets)
		if err != nil {
			return err
		}
	}

	deadBullets := 0
	for index, bullet := range bullets {
		if bullet.PosY <= float64(-bullet.Height) {
			deadBullets++
		} else {
			bullets[index].PosY -= bullets[index].Speed
		}
	}
	bullets = bullets[deadBullets:]

	return nil
}

func (g *MyGame) Draw(screen *ebiten.Image) {
	imageOptions := &ebiten.DrawImageOptions{}
	imageOptions.GeoM.Translate(player.PosX, player.PosY)
	screen.DrawImage(player.Sprite, imageOptions)
	for _, bullet := range bullets {
		bulletImageOptions := &ebiten.DrawImageOptions{}
		bulletImageOptions.GeoM.Translate(bullet.PosX, bullet.PosY)
		screen.DrawImage(bullet.Sprite, bulletImageOptions)
	}
}

func (g *MyGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("My 1st Game!")
	if err := ebiten.RunGame(&MyGame{}); err != nil {
		log.Fatal(err)
	}
}
