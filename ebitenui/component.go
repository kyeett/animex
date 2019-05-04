package ebitenui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"
)

type Component2 struct {
	gfx.Rect
	Visible  bool
	Parent   Componenter
	Children []*Component2
	OnClick  func()
	onRender func(*ebiten.Image, gfx.Vec) // Should be implemented by component creator, not by user of component
}

func (c Component2) Render(screen *ebiten.Image, v gfx.Vec) {
	c.onRender(screen, v)
}

func NewComponent2(r gfx.Rect, renderFunction func(*ebiten.Image, gfx.Vec)) Component2 {
	return Component2{
		Rect:     r,
		Visible:  true,
		onRender: renderFunction,
	}
}

func (c *Component2) GetComponent() *Component2 {
	return c
}

type Componenter interface {
	GetComponent() *Component2
}
