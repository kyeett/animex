package ebitenui

import (
	"fmt"
	"image/color"

	"github.com/kyeett/ebitendrawutil"
	"github.com/oakmound/shiny/materialdesign/colornames"

	"github.com/hajimehoshi/ebiten"

	"github.com/peterhellberg/gfx"
)

type BaseContainer struct {
	gfx.Rect
	Visible    bool
	Background color.Color
	components []Component
}

type Component struct {
	Clickable
	Hoverable
	Renderable
}

func (ct *BaseContainer) Click(v gfx.Vec) bool {
	for _, c := range ct.components {
		if c.IsClicked(v) {
			c.Click(v)
			return true
		}
		fmt.Println(c)
	}
	return false
}

func (ct *BaseContainer) Hovered(v gfx.Vec) bool {
	// fmt.Println("hover")
	// for _, h := range ct.components {
	// 	fmt.Println("hover2")
	// 	if h.IsHovered(v) {
	// 		h.Hover()
	// 		return true
	// 	}
	// }
	return false
}

func (ct *BaseContainer) Add(val interface{}) {

	component := Component{}
	if v, ok := val.(Renderable); ok {
		component.Renderable = v
	}
	if v, ok := val.(Clickable); ok {
		component.Clickable = v
	}
	if v, ok := val.(Hoverable); ok {
		component.Hoverable = v
	}

	ct.components = append(ct.components, component)
}

type HorizontalContainer struct {
	BaseContainer
	contentHeight float64
}

func NewHorizontalContainer(r gfx.Rect) HorizontalContainer {
	return HorizontalContainer{
		BaseContainer{
			Rect: r,
		},
		0.0,
	}
}

func (ct *HorizontalContainer) Render(screen *ebiten.Image, offset gfx.Vec) {
	if !ct.Visible {
		return
	}

	ebitendrawutil.DrawRect(screen, ct.Rect, colornames.Black, 10)

	offsetStepY := (ct.Rect.H() - ct.contentHeight) / float64((len(ct.components) + 1))

	// Render components
	offsetY := 0.0
	for _, r := range ct.components {
		x := (ct.Rect.W() - r.Size().X) / 2
		y := offsetY + offsetStepY
		r.Render(screen, ct.Rect.Min.Add(offset).AddXY(x, y))
		offsetY = y + r.Size().Y
	}
}

func (ct *HorizontalContainer) Click(v gfx.Vec) {
	offsetStepY := (ct.Rect.H() - ct.contentHeight) / float64((len(ct.components) + 1))

	// Render components
	offsetY := 0.0
	for _, c := range ct.components {
		x := (ct.Rect.W() - c.Size().X) / 2
		y := offsetY + offsetStepY

		relativeV := v.Sub(ct.Rect.Min).Sub(gfx.V(x, y))
		// c.Render(screen, ct.Rect.Min.Add(offset).AddXY(x, y))
		if c.IsClicked(relativeV) {
			c.Click(relativeV)
		}
		offsetY = y + c.Size().Y
	}
}

func (ct *HorizontalContainer) IsClicked(v gfx.Vec) bool {
	return ct.Contains(v)
}

func (ct *HorizontalContainer) Add(val interface{}) {
	ct.BaseContainer.Add(val)

	contentHeight := 0.0
	for _, r := range ct.components {
		contentHeight += r.Size().Y
	}
	ct.contentHeight = contentHeight
}
