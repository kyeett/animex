package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

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

	touchScreenTouched := len(inpututil.JustPressedTouchIDs()) > 0
	if touchScreenTouched || inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		t = 0
		if inpututil.IsKeyJustPressed(ebiten.KeyW) {
			shakeNicely = false
		} else {
			shakeNicely = true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		maxAmplitude = gfx.Clamp(maxAmplitude-1.0, 2, 20)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		maxAmplitude = gfx.Clamp(maxAmplitude+1.0, 2, 20)
	}

	// Shake every 5 seconds
	if t > 5 {
		t = 0
	}

	if shakeNicely {
		nicerScreenShake(screen, t)
	} else {
		simpleScreenShake(screen, t)
	}

	// Draw help text
	ebitenutil.DrawRect(screen, 0, 0, 200, 70, color.Black)
	ebitenutil.DebugPrintAt(screen, "Press W to shake screen", 7, 7)
	ebitenutil.DebugPrintAt(screen, "Press S to shake screen nicely", 7, 27)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Change amplitude (%2.0f) with A/D", maxAmplitude), 7, 47)

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
