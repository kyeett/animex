package ebitenui

import (
	"github.com/peterhellberg/gfx"
	"github.com/sirupsen/logrus"
)

type BaseRect struct {
	gfx.Rect
	Visible bool
	Hovered bool
	OnClick func()
	OnHover func()
}

func (r *BaseRect) IsHovered(v gfx.Vec) bool {
	return r.Rect.Contains(v)
}

func (r *BaseRect) Size() gfx.Vec {
	return gfx.V(r.W(), r.H())
}

func (r *BaseRect) Hover(gfx.Vec) {
	logrus.New()
	if r.OnHover == nil {
		return
	}
	r.OnHover()
}

func (r *BaseRect) IsClicked(v gfx.Vec) bool {
	return r.Rect.Contains(v)
}

func (r *BaseRect) Click(gfx.Vec) {
	if r.OnClick == nil {
		return
	}
	r.OnClick()
}
