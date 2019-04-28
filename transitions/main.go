package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	shovelknightresources "github.com/kyeett/animex/resources/shovelknight"
	"github.com/kyeett/ebitendrawutil"
	"github.com/oakmound/shiny/materialdesign/colornames"
	"github.com/peterhellberg/gfx"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth, screenHeight = 400, 300
)

var (
	backgroundImages []*ebiten.Image
	scoreboard       *ebiten.Image
	i                int
)

func update(screen *ebiten.Image) error {
	screen.DrawImage(backgroundImages[1], &ebiten.DrawImageOptions{})

	// Approximations
	scoreboardHeight := 22.0
	textX := int(screenWidth/2 - 30)

	switch {
	case i < 150:
		// Shrink rectangle
		ebitendrawutil.DrawRect(screen, gfx.R(0, scoreboardHeight, screenWidth-1, screenHeight-1), colornames.Black, i)
	case i < 300:
		// Print text
		screen.Fill(color.Black)

		ebitenutil.DebugPrintAt(screen, "Next level", textX, int((screenHeight-scoreboardHeight)/2))
	case i < 450:
		ebitendrawutil.DrawRect(screen, gfx.R(0, scoreboardHeight, screenWidth-1, screenHeight-1), colornames.Black, 450-i)

	default:
		i = 0
	}

	i += 5
	screen.DrawImage(scoreboard, &ebiten.DrawImageOptions{})
	return nil
}

func main() {

	scene1Data, err := shovelknightresources.Asset("scene_1.png")
	if err != nil {
		log.Fatal(err)
	}
	scene2Data, err := shovelknightresources.Asset("scene_2.png")
	if err != nil {
		log.Fatal(err)
	}
	scoreboardData, err := shovelknightresources.Asset("scoreboard.png")
	if err != nil {
		log.Fatal(err)
	}

	backgroundImages = append(backgroundImages, ebitendrawutil.ImgFromBytes(scene1Data))
	backgroundImages = append(backgroundImages, ebitendrawutil.ImgFromBytes(scene2Data))

	scoreboard = ebitendrawutil.ImgFromBytes(scoreboardData)

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "my application"); err != nil {
		log.Fatal(err)
	}
}
