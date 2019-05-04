package ebitenui

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyeett/ebitendrawutil"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

var Logger logger

type logger interface {
	Debugf(string, ...interface{})
	Printf(string, ...interface{})
}

type Button struct {
	BaseRect
}

func NewButton(r gfx.Rect) Button {
	return Button{
		BaseRect{
			Rect:    r,
			Visible: true,
		},
	}
}

func (b *Button) Render(screen *ebiten.Image, offset gfx.Vec) {
	if b.Visible {
		c := colornames.White
		if b.Hovered {
			c = colornames.Yellow
		}
		ebitendrawutil.DrawRect(screen, b.Rect.Moved(offset), c, 3)
	}
}
