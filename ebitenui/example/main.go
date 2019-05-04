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
	clickable  []ebitenui.Clickable
	hoverable  []ebitenui.Hoverable
	renderable []ebitenui.Renderable
}

func NewUI() UI {
	ui := UI{}

	openMenu := ebitenui.NewButton(gfx.R(0, 0, 32, 32).Moved(gfx.V(screenWidth-32, 0)))

	ui.Add(&openMenu)
	// ui.Add(&menu)
	hor := ebitenui.NewHorizontalContainer(gfx.R(0, 0, screenWidth-100, screenHeight-100).Moved(gfx.V(50, 50)))
	hor.Visible = false
	openMenu2 := ebitenui.NewButton(gfx.R(0, 0, 200, 32))

	openMenu2.OnClick = func() {
		fmt.Println("CLICKED!!")
		// menu.Visible = true
		hor.Visible = false
		openMenu.Visible = true
	}
	hor.Add(&openMenu2)
	hor.Add(&openMenu2)
	hor.Add(&openMenu2)
	hor.Add(&openMenu2)
	hor.Add(&openMenu2)
	ui.Add(&hor)

	openMenu.OnClick = func() {
		openMenu.Visible = false
		hor.Visible = true
	}

	return ui
}

func (ui *UI) Clicked(v gfx.Vec) bool {
	for _, c := range ui.clickable {
		if c.IsClicked(v) {
			c.Click(v)
			return true
		}
	}
	return false
}

func (ui *UI) Hovered(v gfx.Vec) bool {
	fmt.Println("hover")
	for _, h := range ui.hoverable {
		fmt.Println("hover2")
		if h.IsHovered(v) {
			h.Hover(v)
			return true
		}
	}
	return false
}

func (ui *UI) Render(screen *ebiten.Image) {
	offset := gfx.V(0, 0)
	for _, r := range ui.renderable {
		r.Render(screen, offset)
	}
}

func (ui *UI) Add(val interface{}) {

	if v, ok := val.(ebitenui.Clickable); ok {
		ui.clickable = append(ui.clickable, v)
		log.Printf("clickable", v)
	}
	if v, ok := val.(ebitenui.Hoverable); ok {
		ui.hoverable = append(ui.hoverable, v)
	}
	if v, ok := val.(ebitenui.Renderable); ok {
		ui.renderable = append(ui.renderable, v)
		log.Printf("renderable", v)
	}
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
