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
		sfx.Cu01_ogg,
		sfx.Cu02_ogg,
		sfx.Cu03_ogg,
		sfx.Cu04_ogg,
		sfx.Cu05_ogg,
		sfx.Cu06_ogg,
		sfx.Cu07_ogg,
		sfx.Cu08_ogg,
		sfx.Cu09_ogg,
		sfx.Cu10_ogg}
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
