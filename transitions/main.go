package main

import (
	"log"

	shovelknightresources "github.com/kyeett/animex/resources/shovelknight"
	"github.com/kyeett/ebitendrawutil"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth, screenHeight = 400, 300
)

var (
	scene1, scene2, scene3, scoreboard *ebiten.Image
	t                                  float64
)

func update(screen *ebiten.Image) error {
	// Draw the scenes
	switch int(t / 2) {
	case 0:
		screen.DrawImage(scene1, &ebiten.DrawImageOptions{})
	case 1:
		screen.DrawImage(scene2, &ebiten.DrawImageOptions{})
	case 2:
		screen.DrawImage(scene3, &ebiten.DrawImageOptions{})
	}

	// Draw transition
	scoreboardHeight := 44.0
	maxRect := gfx.R(0, scoreboardHeight, screenWidth, screenHeight)
	switch int(t) {
	case 0:
		TransitionShrinkingRect(screen, maxRect, t, colornames.Black)
	case 1:
		offset := 1.0
		TransitionGrowingRect(screen, maxRect, t-offset, colornames.Black)
	case 2:
		offset := 2.0
		TransitionShrinkingBorder(screen, maxRect, t-offset, colornames.Black)
	case 3:
		offset := 3.0
		TransitionGrowingBorder(screen, maxRect, t-offset, colornames.Black)
	case 4:
		offset := 4.0
		TransitionBlinds(screen, maxRect, 6, 1-(t-offset), colornames.Black)
	case 5:
		offset := 5.0
		TransitionBlinds(screen, maxRect, 6, t-offset, colornames.Black)
	}

	// Draw scoreboard
	screen.DrawImage(scoreboard, &ebiten.DrawImageOptions{})

	t += 0.02
	if t >= 6 {
		t = 0
	}
	return nil
}

func main() {
	// Load resources
	scene1Data, err := shovelknightresources.Asset("scene_1.png")
	if err != nil {
		log.Fatal(err)
	}
	scene2Data, err := shovelknightresources.Asset("scene_2.png")
	if err != nil {
		log.Fatal(err)
	}
	scene3Data, err := shovelknightresources.Asset("scene_3.png")
	if err != nil {
		log.Fatal(err)
	}
	scoreboardData, err := shovelknightresources.Asset("scoreboard.png")
	if err != nil {
		log.Fatal(err)
	}
	scene1 = ebitendrawutil.ImageFromBytes(scene1Data)
	scene2 = ebitendrawutil.ImageFromBytes(scene2Data)
	scene3 = ebitendrawutil.ImageFromBytes(scene3Data)
	scoreboard = ebitendrawutil.ImageFromBytes(scoreboardData)

	// Start the Ebiten update loop
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "my application"); err != nil {
		log.Fatal(err)
	}
}
