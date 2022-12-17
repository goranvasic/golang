package player

import (
	"bytes"
	"fmt"
	sfx "github.com/goranvasic/ebitengine07/assets/audio"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"log"
)

var (
	playerColorIdle  = color.RGBA{B: 0xff, A: 0xff}
	playerColorError = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}

	sfxPlayer    *audio.Player
	audioContext *audio.Context
	blipSelect   = sfx.BlipselectOgg
)

type Player struct {
	Name           string
	GamepadId      ebiten.GamepadID
	Width, Height  float64
	PosX, PosY     float64
	XBound, YBound float64
	State          State
	Color          color.RGBA
}

func (p *Player) MoveUp() {
	p.PosY -= p.Height
	fmt.Println("Move Up")
}

func (p *Player) MoveDown() {
	p.PosY += p.Height
	fmt.Println("Move Down")
}

func (p *Player) MoveLeft() {
	p.PosX -= p.Width
	fmt.Println("Move Left")
}

func (p *Player) MoveRight() {
	p.PosX += p.Width
	fmt.Println("Move Right")
}

func (p *Player) Move() {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		if p.PosY-p.Height >= 0 {
			p.MoveUp()
		} else {
			p.State = Error
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		if p.PosY+p.Height <= p.YBound-p.Height {
			p.MoveDown()
		} else {
			p.State = Error
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if p.PosX-p.Width >= 0 {
			p.MoveLeft()
		} else {
			p.State = Error
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if p.PosX+p.Width <= p.XBound-p.Width {
			p.MoveRight()
		} else {
			p.State = Error
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	sprite := ebiten.NewImage(int(p.Width), int(p.Height))
	if p.State == Error {
		sprite.Fill(playerColorError)
		if err := sfxPlayer.Rewind(); err != nil {
			log.Fatal(err)
		}
		sfxPlayer.Play()
		p.State = Idle
	} else {
		sprite.Fill(p.Color)
	}
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.PosX, p.PosY)
	screen.DrawImage(sprite, &options)
}

func CreatePlayer(name string, gamepadId ebiten.GamepadID, screenWidth float64, screenHeight float64) Player {
	audioContext = audio.NewContext(44100)
	blip, err := vorbis.DecodeWithoutResampling(bytes.NewReader(blipSelect))
	if err != nil {
		log.Fatal(err)
	}
	sfxPlayer, err = audioContext.NewPlayer(blip)
	if err != nil {
		log.Fatal(err)
	}
	if err := sfxPlayer.Rewind(); err != nil {
		log.Fatal(err)
	}
	return Player{
		Name:      name,
		GamepadId: gamepadId,
		Width:     40,
		Height:    40,
		PosX:      screenWidth/2 - 20,
		PosY:      screenHeight/2 - 20,
		XBound:    screenWidth,
		YBound:    screenHeight,
		State:     Idle,
		Color:     playerColorIdle,
	}
}
