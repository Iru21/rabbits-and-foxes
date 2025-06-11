package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
	"math"
)

const (
	StartingRabbits = 100
	StartingFoxes   = 20
	WorldHeight     = 60
	WorldWidth      = 60
	TileSize        = 16
	Render          = true
	TPS             = 100
)

var CurrentGame *Game

func main() {
	LoadAssets()
	CurrentGame = NewGame()

	windowWidth := WorldWidth * TileSize
	windowHeight := WorldHeight * TileSize
	CurrentGame.ui.AddButton(*NewUIButton("Quit", windowWidth-50, 10, 40, 25, func() {
		CurrentGame.Stop()
	}))
	CurrentGame.ui.AddButton(*NewUIButton("Reset", windowWidth-50, 40, 40, 25, func() {
		CurrentGame.Reset()
	}))
	CurrentGame.ui.AddButton(*NewUIButton("Speed Up", windowWidth-80, 70, 70, 25, func() {
		CurrentGame.simulationSpeed += 0.01
	}))
	CurrentGame.ui.AddButton(*NewUIButton("Slow Down", windowWidth-80, 100, 70, 25, func() {
		CurrentGame.simulationSpeed = math.Max(CurrentGame.simulationSpeed-0.01, 0.001)
	}))
	CurrentGame.ui.AddButton(*NewUIButton("Bigger Speed Up", windowWidth-120, 130, 110, 25, func() {
		CurrentGame.simulationSpeed += 0.1
	}))
	CurrentGame.ui.AddButton(*NewUIButton("Bigger Slow Down", windowWidth-120, 160, 110, 25, func() {
		CurrentGame.simulationSpeed = math.Max(CurrentGame.simulationSpeed-0.1, 0.001)
	}))

	ebiten.SetWindowTitle("Rabbits and Foxes")
	ebiten.SetWindowIcon([]image.Image{FoxSprite48, FoxSprite32, FoxSprite16})
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetTPS(TPS)

	if err := ebiten.RunGame(CurrentGame); err != nil {
		log.Fatal(err)
	}
}
