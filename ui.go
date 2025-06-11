package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	color "image/color"
)

type UIButton struct {
	UIElement
	Text         string
	ClickHandler func()
}

func NewUIButton(text string, x, y, width, height int, clickHandler func()) *UIButton {
	return &UIButton{
		Text:         text,
		ClickHandler: clickHandler,
		UIElement: UIElement{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
		},
	}
}

func (b *UIButton) GetElement() *UIElement {
	return &b.UIElement
}

func (b *UIButton) HandleClick() {
	if b.ClickHandler != nil {
		b.ClickHandler()
	}
}

func (b *UIButton) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(b.X), float32(b.Y), float32(b.Width), float32(b.Height), color.RGBA{R: 128, G: 128, B: 128, A: 255}, false)
	ebitenutil.DebugPrintAt(screen, b.Text, b.X+5, b.Y+5)
}

type UIElement struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (e *UIElement) InBounds(x, y int) bool {
	return x >= e.X && x <= e.X+e.Width && y >= e.Y && y <= e.Y+e.Height
}

type UI struct {
	Buttons []UIButton
}

func NewUI() *UI {
	return &UI{
		Buttons: make([]UIButton, 0),
	}
}

func (ui *UI) AddButton(button UIButton) {
	ui.Buttons = append(ui.Buttons, button)
}

func (ui *UI) Draw(screen *ebiten.Image) {
	for _, element := range ui.Buttons {
		element.Draw(screen)
	}
}

func (ui *UI) HandleClick(x, y int) {
	for _, button := range ui.Buttons {
		if button.InBounds(x, y) {
			button.HandleClick()
		}
	}
}
