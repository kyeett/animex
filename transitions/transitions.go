package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/ebitendrawutil"
	"github.com/peterhellberg/gfx"
)

// TransitionGrowingBorder draws a border that expands inwards until it fills maxRect
func TransitionGrowingBorder(screen *ebiten.Image, maxRect gfx.Rect, t float64, clr color.Color) {
	min := gfx.MathMin(maxRect.W(), maxRect.H())
	ebitendrawutil.DrawRect(screen, maxRect, clr, int(t*min/2.0))
}

// TransitionShrinkingBorder is TransitionGrowingBorder run backwards
func TransitionShrinkingBorder(screen *ebiten.Image, maxRect gfx.Rect, t float64, clr color.Color) {
	TransitionGrowingBorder(screen, maxRect, 1-t, clr)
}

// TransitionGrowingRect draws a rectangle that grows from the center to fill maxRect as t goes from 0 to 1
func TransitionGrowingRect(screen *ebiten.Image, maxRect gfx.Rect, t float64, clr color.Color) {
	v := maxRect.Center().Lerp(maxRect.Min, t)
	w := gfx.Lerp(0, maxRect.W(), t)
	h := gfx.Lerp(0, maxRect.H(), t)
	ebitenutil.DrawRect(screen, v.X, v.Y, w, h, clr)
}

// TransitionShrinkingRect draws a rectangle that shrinks towards the center as to goes from 0 to 1
func TransitionShrinkingRect(screen *ebiten.Image, maxRect gfx.Rect, t float64, clr color.Color) {
	TransitionGrowingRect(screen, maxRect, 1-t, clr)
}

// TransitionBlinds draws n rectangles that expands up to cover maxRect  as t goes from 0 to 1
func TransitionBlinds(screen *ebiten.Image, maxRect gfx.Rect, nRectangles int, t float64, clr color.Color) {
	blindMaxHeight := maxRect.H() / float64(nRectangles)
	for i := 0; i < nRectangles; i++ {
		x := maxRect.Min.X
		y := maxRect.Min.Y + float64(i)*blindMaxHeight
		height := gfx.Lerp(0, blindMaxHeight, t)
		ebitenutil.DrawRect(screen, x, y, maxRect.W(), height, clr)
	}
}
