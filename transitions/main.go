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
	scene1, scene2, scoreboard *ebiten.Image
	t                          float64
	step                       int
)

func update(screen *ebiten.Image) error {
	// Draw the scenes
	screen.DrawImage(scene1, &ebiten.DrawImageOptions{})

	// Draw transition
	scoreboardHeight := 44.0
	maxRect := gfx.R(0, scoreboardHeight, screenWidth, screenHeight)
	switch step {
	case 0:
		TransitionGrowingRect(screen, maxRect, t, colornames.Black)
	}

	// Draw scoreboard
	screen.DrawImage(scoreboard, &ebiten.DrawImageOptions{})

	if t < 1 {
		t += 0.01
	}
	return nil
}

func main() {
	// Load resources
	scene1Data, err := shovelknightresources.Asset("scene_1.png")
	if err != nil {
		log.Fatal(err)
	}
	scoreboardData, err := shovelknightresources.Asset("scoreboard.png")
	if err != nil {
		log.Fatal(err)
	}
	scene1 = ebitendrawutil.ImageFromBytes(scene1Data)
	scoreboard = ebitendrawutil.ImageFromBytes(scoreboardData)

	// Start the Ebiten update loop
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "my application"); err != nil {
		log.Fatal(err)
	}
}
