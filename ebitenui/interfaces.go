package ebitenui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
)

type Clickable interface {
	IsClicked(gfx.Vec) bool
	Click(gfx.Vec)
}

type Hoverable interface {
	IsHovered(gfx.Vec) bool
	Hover(gfx.Vec)
}

type Renderable interface {
	Render(*ebiten.Image, gfx.Vec)
	Size() gfx.Vec
}
