package gopher

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	sfx "goranvasic/ebitengine/experiment/assets/audio"
	"log"
	"math/rand"
	"time"
)

var (
	sfxPlayer    *audio.Player
	audioContext *audio.Context
	sounds       = [][]byte{
		sfx.Cu01Ogg,
		sfx.Cu02Ogg,
		sfx.Cu03Ogg,
		sfx.Cu04Ogg,
		sfx.Cu05Ogg,
		sfx.Cu06Ogg,
		sfx.Cu07Ogg,
		sfx.Cu08Ogg,
		sfx.Cu09Ogg,
		sfx.Cu10Ogg}
	sfxPlayers []*audio.Player
)

type Gopher struct {
	PosX   float64
	PosY   float64
	Speed  float64
	Sprite *ebiten.Image
	Width  int
	Height int
}

func (g *Gopher) InitSfx() {
	audioContext = audio.NewContext(44100)
	for _, sound := range sounds {
		cu, err := vorbis.DecodeWithoutResampling(bytes.NewReader(sound))
		if err != nil {
			log.Fatal(err)
		}
		sfxPlayer, err = audioContext.NewPlayer(cu)
		if err != nil {
			log.Fatal(err)
		}
		sfxPlayers = append(sfxPlayers, sfxPlayer)
	}
}

func (g *Gopher) Fire(bullets []Bullet) error {
	rand.Seed(time.Now().UnixNano())
	randomSfxIndex := rand.Intn(len(sfxPlayers))
	sfxPlayer = sfxPlayers[randomSfxIndex]
	if err := sfxPlayer.Rewind(); err != nil {
		return err
	}
	sfxPlayer.Play()
	fmt.Println("FIRE!", len(bullets), randomSfxIndex)
	return nil
}

func (g *Gopher) ThreadSleep() {
	time.Sleep(5 * time.Second)
}
