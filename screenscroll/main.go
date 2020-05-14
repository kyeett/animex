package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/kyeett/ebitendrawutil"
	"golang.org/x/image/colornames"

	"github.com/peterhellberg/gfx"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth, screenHeight = 1280, 512
)

var (
	playerImage            *ebiten.Image
	backgroundImage        *ebiten.Image
	player, camera         gfx.Vec
	previousPlayer, target gfx.Vec
	cameraBounds           = gfx.R(0, 0, 512, screenHeight)
	directionLeft          bool
)

func init() {
	img := gfx.MustOpenImage("map.png")
	backgroundImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img = gfx.MustOpenImage("player.png")
	playerImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

func update(screen *ebiten.Image) error {

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		previousPlayer = player
		player.X += -5
		directionLeft = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		previousPlayer = player
		player.X += 5
		directionLeft = false
	}
	pBounds := gfx.BoundsToRect(playerImage.Bounds())
	player.X = gfx.Clamp(player.X, pBounds.W()/2, screenWidth-pBounds.W()/2)

	offset := 100.0
	speed := 10.0
	target.X = player.X + offset
	if player.X < previousPlayer.X {
		target.X = player.X - offset
	}
	// camera.X = gfx.Clamp(player.X, camera.X-3, camera.X+3) // Limit speed
	camera.X = gfx.Clamp(target.X-cameraBounds.W()/2, camera.X-speed, camera.X+speed) // Limit speed
	camera.X = gfx.Clamp(camera.X, 0, screenWidth-cameraBounds.W())

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(backgroundImage, op)

	// Draw player
	bounds := gfx.BoundsToRect(playerImage.Bounds())
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-bounds.W()/2, bounds.H()/2)
	if directionLeft {
		op.GeoM.Scale(-1, 1)
	}
	op.GeoM.Translate(player.X, player.Y)
	screen.DrawImage(playerImage, op)

	ebitendrawutil.DrawRect(screen, cameraBounds.Moved(camera), colornames.Red)
	ebitenutil.DrawRect(screen, 0, 0, camera.X, cameraBounds.H(), colornames.Black)
	ebitenutil.DrawRect(screen, camera.X+cameraBounds.W(), 0, screenWidth-(camera.X+cameraBounds.W()), cameraBounds.H(), colornames.Black)
	ebitenutil.DrawLine(screen, target.X, 0, target.X, screenHeight, colornames.Yellow)

	return nil
}

func main() {
	player = gfx.V(100, 360)

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "my application"); err != nil {
		log.Fatal(err)
	}
}
