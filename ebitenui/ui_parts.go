package ebitenui

import (
	"fmt"

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
	Component2
}

func NewButton(parent Componenter, r gfx.Rect) *Button {
	b := Button{}
	b.Component2 = NewComponent2(r, b.Render)
	b.Parent = parent
	fmt.Println("new button")
	return &b
}

func (b *Button) Render(screen *ebiten.Image, offset gfx.Vec) {
	c := colornames.White
	// fmt.Println("Base:", b.Rect)
	// fmt.Println("Moved", b.Rect.Moved(offset), offset)
	ebitendrawutil.DrawRect(screen, b.Rect.Moved(offset), c, 3)
}
