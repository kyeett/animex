package ebitenui

import (
	"fmt"

	"github.com/kyeett/ebitendrawutil"
	"github.com/oakmound/shiny/materialdesign/colornames"

	"github.com/hajimehoshi/ebiten"

	"github.com/peterhellberg/gfx"
)

type VerticalLayout struct {
	// BaseContainer
	Component2
}

func NewVerticalLayout(r gfx.Rect) *VerticalLayout {
	l := VerticalLayout{}
	l.Component2 = NewComponent2(r, l.Render)
	return &l
}

func (l *VerticalLayout) Render(screen *ebiten.Image, offset gfx.Vec) {
	ebitendrawutil.DrawRect(screen, l.Rect.Moved(offset), colornames.Black, 10)
}

func (l *VerticalLayout) Click(v gfx.Vec) {
	// offsetStepY := (l.Rect.H() - l.contentHeight) / float64((len(l.components) + 1))

	// // Render components
	// offsetY := 0.0
	// for _, c := range l.components {
	// 	x := (l.Rect.W() - c.Size().X) / 2
	// 	y := offsetY + offsetStepY

	// 	relativeV := v.Sub(l.Rect.Min).Sub(gfx.V(x, y))
	// 	// c.Render(screen, l.Rel.Min.Add(offset).AddXY(x, y))
	// 	if c.IsClicked(relativeV) {
	// 		c.Click(relativeV)
	// 	}
	// 	offsetY = y + c.Size().Y
	// }
}

func (l *VerticalLayout) IsClicked(v gfx.Vec) bool {
	return l.Contains(v)
}

func (l *VerticalLayout) Add(c Componenter) {
	l.Children = append(l.Children, c.GetComponent())
	l.UpdatePositions()
	fmt.Println("Add to vertical", len(l.Children))
}

func (l *VerticalLayout) UpdatePositions() {
	contentHeight := 0.0
	for _, c := range l.Children {
		contentHeight += c.H()
	}

	offsetStepY := (l.H() - contentHeight) / float64((len(l.Children) + 1))

	offsetY := 0.0
	for i, c := range l.Children {
		x := (l.W() - c.W()) / 2
		y := offsetY + offsetStepY

		relativeV := gfx.V(x, y)
		l.Children[i].Rect = gfx.R(0, 0, c.Rect.W(), c.Rect.H()).Moved(relativeV)

		// c.Render(screen, l.Rel.Min.Add(offset).AddXY(x, y))
		// if c.IsClicked(relativeV) {
		// 	c.Click(relativeV)
		// }
		offsetY = y + c.H()
	}
	fmt.Println(l.Children)
}
