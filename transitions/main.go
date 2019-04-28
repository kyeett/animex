package main

import (
	"log"

	"golang.org/x/image/colornames"

	"github.com/peterhellberg/gfx"

	shovelknightresources "github.com/kyeett/animex/resources/shovelknight"
	"github.com/kyeett/ebitendrawutil"

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
	switch step {
	case 0, 1, 2:
		screen.DrawImage(scene1, &ebiten.DrawImageOptions{})
	case 3, 4, 5:
		screen.DrawImage(scene2, &ebiten.DrawImageOptions{})
	}

	// Draw the transition
	scoreboardHeight := 44.0
	maxRect := gfx.R(0, scoreboardHeight, screenWidth, screenHeight)
	switch step {
	case 0:
		TransitionShrinkingRect(screen, maxRect, t, colornames.Black)
	case 1:
		// Increase faster during idle time
		t += 0.01
	case 2:
		TransitionGrowingBorder(screen, maxRect, t, colornames.Black)
	case 3:
		// TransitionShrinkingRect(screen, maxRect, t, colornames.Black)
		TransitionGrowingBorder(screen, maxRect, 1-t, colornames.Black)
	case 4:
		// Increase faster during idle time
		t += 0.01
	case 5:
		TransitionGrowingBorder(screen, maxRect, t, colornames.Black)
	}

	t += 0.01
	if t > 1 {
		t = 0
		step = (step + 1) % 6
	}

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

	scene1 = ebitendrawutil.ImageFromBytes(scene1Data)
	scene2 = ebitendrawutil.ImageFromBytes(scene2Data)
	scoreboard = ebitendrawutil.ImageFromBytes(scoreboardData)

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "my application"); err != nil {
		log.Fatal(err)
	}
}
