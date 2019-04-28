package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/ebitendrawutil"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

func TransitionGrowingBorder(screen *ebiten.Image, maxRect gfx.Rect, t float64, clr color.Color) {
	min := gfx.MathMin(maxRect.W(), maxRect.W())
	ebitendrawutil.DrawRect(screen, maxRect, colornames.Black, int(t*min/2.0))
}

func TransitionGrowingRect(screen *ebiten.Image, maxRect gfx.Rect, t float64, clr color.Color) {
	v := maxRect.Center().Lerp(maxRect.Min, t)
	w := gfx.Lerp(0, maxRect.W(), t)
	h := gfx.Lerp(0, maxRect.H(), t)
	ebitenutil.DrawRect(screen, v.X, v.Y, w, h, clr)
}

func TransitionShrinkingRect(screen *ebiten.Image, maxRect gfx.Rect, t float64, clr color.Color) {
	TransitionGrowingRect(screen, maxRect, 1-t, clr)
}
