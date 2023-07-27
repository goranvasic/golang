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

	blipAudio    *audio.Player
	audioContext *audio.Context
	blipRaw      = sfx.BlipOgg
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

func (p *Player) Draw(screen *ebiten.Image) {
	sprite := ebiten.NewImage(int(p.Width), int(p.Height))
	if p.State == Error {
		sprite.Fill(playerColorError)
		playErrorSound()
		p.State = Idle
	} else {
		sprite.Fill(p.Color)
	}
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(p.PosX, p.PosY)
	screen.DrawImage(sprite, &options)
}

func playErrorSound() {
	if err := blipAudio.Rewind(); err != nil {
		log.Fatal(err)
	}
	blipAudio.Play()
}

func CreatePlayer(name string, gamepadId ebiten.GamepadID, screenWidth float64, screenHeight float64) Player {
	enableSound()
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

func enableSound() {
	audioContext = audio.NewContext(44100)
	blip, err := vorbis.DecodeWithoutResampling(bytes.NewReader(blipRaw))
	if err != nil {
		log.Fatal(err)
	}
	blipAudio, err = audioContext.NewPlayer(blip)
	if err != nil {
		log.Fatal(err)
	}
	if err := blipAudio.Rewind(); err != nil {
		log.Fatal(err)
	}
}

func (p *Player) Move() {
	upKeyPressed := inpututil.IsKeyJustPressed(ebiten.KeyUp)
	downKeyPressed := inpututil.IsKeyJustPressed(ebiten.KeyDown)
	leftKeyPressed := inpututil.IsKeyJustPressed(ebiten.KeyLeft)
	rightKeyPressed := inpututil.IsKeyJustPressed(ebiten.KeyRight)
	anyKeyPressed := upKeyPressed || downKeyPressed || leftKeyPressed || rightKeyPressed

	if upKeyPressed && playerCanMoveUp(p) {
		p.MoveUp()
	} else if downKeyPressed && playerCanMoveDown(p) {
		p.MoveDown()
	} else if leftKeyPressed && playerCanMoveLeft(p) {
		p.MoveLeft()
	} else if rightKeyPressed && playerCanMoveRight(p) {
		p.MoveRight()
	} else if anyKeyPressed {
		p.State = Error
	}
}

func playerCanMoveUp(p *Player) bool {
	return p.PosY-p.Height >= 0
}

func playerCanMoveDown(p *Player) bool {
	return p.PosY+p.Height <= p.YBound-p.Height
}

func playerCanMoveLeft(p *Player) bool {
	return p.PosX-p.Width >= 0
}

func playerCanMoveRight(p *Player) bool {
	return p.PosX+p.Width <= p.XBound-p.Width
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
