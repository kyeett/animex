package main

import (
	"log"
	"math/rand"

	shovelknightresources "github.com/kyeett/animex/resources/shovelknight"
	"github.com/peterhellberg/gfx"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/ebitendrawutil"
)

const (
	screenWidth, screenHeight = 400, 300
	padding                   = 20.0
)

var (
	scene1, tmpScreen, biggerTmpScreen *ebiten.Image
	t                                  float64
	maxAmplitude                       = 10.0
)

func simpleScreenShake(screen *ebiten.Image, t float64) {
	// Draw scene onto tmpScreen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-padding, -padding)
	tmpScreen.DrawImage(scene1, op)

	// Draw tmpScreen on an offset (screen shake) applied

	op = &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(padding, -padding)
	if t < 1 {
		amplitude := maxAmplitude * gfx.Lerp(1, 0, t)
		dx := amplitude * (2*rand.Float64() - 1)
		dy := amplitude * (2*rand.Float64() - 1)
		op.GeoM.Translate(dx, dy)
	}
	screen.DrawImage(tmpScreen, op)
}

func nicerScreenShake(screen *ebiten.Image, t float64) {
	// Draw scene onto tmpScreen
	biggerTmpScreen.DrawImage(scene1, &ebiten.DrawImageOptions{})

	// Draw tmpScreen on an offset (screen shake) applied
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-padding, -padding)
	if t < 1 {
		amplitude := maxAmplitude * gfx.Lerp(1, 0, t)
		dx := amplitude * (2*rand.Float64() - 1)
		dy := amplitude * (2*rand.Float64() - 1)
		op.GeoM.Translate(-dx, -dy)
	}

	screen.DrawImage(biggerTmpScreen, op)
}

var shakeNicely = true

func update(screen *ebiten.Image) error {
	nicerScreenShake(screen, t) // Previously simpleScreenShake(screen, t)
	t += 1 / 60.0
	return nil
}

func main() {
	// Load resources
	scene1Data, err := shovelknightresources.Asset("scene_1.png")
	if err != nil {
		log.Fatal(err)
	}
	scene1 = ebitendrawutil.ImageFromBytes(scene1Data)

	// Create temp screens
	tmpScreen, err = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	// Create temp screen
	biggerTmpScreen, err = ebiten.NewImage(screenWidth+40, screenHeight+40, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Screen shake"); err != nil {
		log.Fatal(err)
	}
}
