package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	shovelknightresources "github.com/kyeett/animex/resources/shovelknight"
	"github.com/peterhellberg/gfx"

	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/ebitendrawutil"
)

const (
	screenWidth, screenHeight = 400, 300
)

var (
	scene1, tmpScreen *ebiten.Image
	t                 float64
)

func simpleScreenShake(screen *ebiten.Image, t float64) {
	// Draw scene onto tmpScreen
	tmpScreen.DrawImage(scene1, &ebiten.DrawImageOptions{})

	// Draw help text
	ebitenutil.DrawRect(tmpScreen, 5+115, 20, 150, 18, color.Black)
	ebitenutil.DebugPrintAt(tmpScreen, "Press X to shake screen", 10+115, 20)

	// Draw tmpScreen on an offset (screen shake) applied
	maxAmplitude := 10.0
	op := &ebiten.DrawImageOptions{}
	if t < 1 {
		amplitude := maxAmplitude * gfx.Lerp(1, 0, t)
		dx := amplitude * (2*rand.Float64() - 1)
		dy := amplitude * (2*rand.Float64() - 1)
		op.GeoM.Translate(dx, dy)
	}
	screen.DrawImage(tmpScreen, op)
}

func update(screen *ebiten.Image) error {
	// touchScreenTouched := len(inpututil.JustPressedTouchIDs()) > 0
	// if touchScreenTouched || inpututil.IsKeyJustPressed(ebiten.KeyX) {
	// 	t = 0
	// }

	// // Shake every 5 seconds
	// if t > 5 {
	// 	t = 0
	// }

	simpleScreenShake(screen, t)
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

	// Create temp screen
	tmpScreen, err = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Screen shake"); err != nil {
		log.Fatal(err)
	}
}
