package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

const (
	StartingRabbits = 100
	StartingFoxes   = 20
	WorldHeight     = 30
	WorldWidth      = 30
	TileSize        = 32
)

var CurrentGame *Game

func main() {
	LoadAssets()
	CurrentGame = NewGame()
	ebiten.SetWindowTitle("Rabbits and Foxes")
	ebiten.SetWindowIcon([]image.Image{FoxSprite48, FoxSprite32, FoxSprite16})
	ebiten.SetWindowSize(WorldWidth*TileSize, WorldHeight*TileSize)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetTPS(999)
	if err := ebiten.RunGame(CurrentGame); err != nil {
		log.Fatal(err)
	}
}
