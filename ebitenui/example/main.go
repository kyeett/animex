package main

import (
	"fmt"
	"log"

	"github.com/kyeett/animex/ebitenui"
	"github.com/kyeett/civ/mousehandler"
	"golang.org/x/image/colornames"

	"github.com/peterhellberg/gfx"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth, screenHeight = 640, 480
)

type UI struct {
	ebitenui.Component2
}

func NewUI() UI {
	ui := UI{}

	r := gfx.R(0, 0, 32, 32).Moved(gfx.V(screenWidth-32, 0))
	openMenu := ebitenui.NewButton(nil, r)
	openMenu.OnClick = func() { fmt.Println("\nbutton\n") }

	r = gfx.R(0, 0, screenWidth-100, screenHeight-100).Moved(gfx.V(50, 50))
	vert := ebitenui.NewVerticalLayout(r)

	vert.OnClick = func() { fmt.Println("\nvert\n") }
	// vert.Visible = false
	openMenu2 := ebitenui.NewButton(vert, gfx.R(0, 0, 200, 32))
	openMenu3 := ebitenui.NewButton(vert, gfx.R(0, 0, 200, 32))
	openMenu4 := ebitenui.NewButton(vert, gfx.R(0, 0, 200, 32))
	openMenu2.OnClick = func() {
		fmt.Println("CLICKED!!1")
	}
	openMenu3.OnClick = func() {
		fmt.Println("CLICKED!!2")
	}
	openMenu4.OnClick = func() {
		fmt.Println("CLICKED!!3")
	}

	vert.Add(openMenu2)
	// log.Fatal(openMenu2)
	vert.Add(openMenu3)
	vert.Add(openMenu4)
	// vert.Add(&openMenu2)

	ui.Add(vert)
	ui.Add(openMenu)

	return ui
}

func ComponentAt(pos gfx.Vec, root *ebitenui.Component2) (*ebitenui.Component2, bool) {
	fmt.Println(pos)
	if !root.Contains(pos) {
		return &ebitenui.Component2{}, false
	}

	for _, c := range root.Children {
		// Check children
		if w, found := ComponentAt(pos.Sub(root.Rect.Min), c); found {
			return w, true
		}
	}

	return root, true
}

func RenderComponents(screen *ebiten.Image, root *ebitenui.Component2, offset gfx.Vec) {
	// fmt.Println("offset", offset, root.Rect)
	if !root.Visible {
		return
	}

	root.Render(screen, offset)
	for _, c := range root.Children {
		RenderComponents(screen, c, root.Min.Add(offset))
	}
}

func ClickComponent(w ebitenui.Componenter) {
	if w == nil {
		fmt.Println("No click parent found")
		return
	}
	c := w.GetComponent()

	if c.OnClick != nil {
		fmt.Println("Click component")
		c.OnClick()
		return
	}
	fmt.Println("click parent")
	ClickComponent(c.Parent)
}

func (ui *UI) Clicked(v gfx.Vec) bool {
	for _, c := range ui.Children {
		if w, found := ComponentAt(v, c); found {
			fmt.Printf("%v, %T\n", w, w)
			ClickComponent(w)
			return true
		}
	}
	return false
}

func (ui *UI) Render(screen *ebiten.Image) {
	for _, c := range ui.Children {
		RenderComponents(screen, c, gfx.V(0, 0))
	}
}

func (ui *UI) Add(c ebitenui.Componenter) {

	ui.Children = append(ui.Children, c.GetComponent())

	// if v, ok := val.(ebitenui.Clickable); ok {
	// 	ui.clickable = append(ui.clickable, v)
	// 	log.Printf("clickable", v)
	// }
	// if v, ok := val.(ebitenui.Hoverable); ok {
	// 	ui.hoverable = append(ui.hoverable, v)
	// }
	// if v, ok := val.(ebitenui.Renderable); ok {
	// 	ui.renderable = append(ui.renderable, v)
	// 	log.Printf("renderable", v)
	// }
}

func update(screen *ebiten.Image) error {
	screen.Fill(colornames.Green)
	mouseHandler.Update()
	ui.Render(screen)
	return nil
}

var (
	ui           UI
	mouseHandler mousehandler.MouseHandler
)

func main() {
	ui = NewUI()
	mouseHandler = mousehandler.New(func(v gfx.Vec) { ui.Clicked(v) }, nil)

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "UI"); err != nil {
		log.Fatal(err)
	}
}
