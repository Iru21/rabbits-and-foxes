package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
	"math"
	"os"
)

const (
	StartingRabbits = 100
	StartingFoxes   = 20
	WorldHeight     = 50
	WorldWidth      = 50
	TileSize        = 16
)

var CurrentGame *Game

type Game struct {
	World           *World
	ticks           int
	simulationSpeed float64
}

func NewGame() *Game {
	return &Game{
		World:           NewWorld(),
		ticks:           0,
		simulationSpeed: 0.2,
	}
}

func (g *Game) Update() error {
	pressed := inpututil.AppendPressedKeys([]ebiten.Key{})
	if len(pressed) > 0 {
		switch pressed[0] {
		case ebiten.KeyUp:
			g.simulationSpeed = g.simulationSpeed + 0.01
		case ebiten.KeyRight:
			g.simulationSpeed = g.simulationSpeed + 1
		case ebiten.KeyDown:
			g.simulationSpeed = math.Max(g.simulationSpeed-0.01, 0.001)
		case ebiten.KeyLeft:
			g.simulationSpeed = math.Max(g.simulationSpeed-1, 0.001)
		case ebiten.KeyR:
			def := NewGame()
			g.World = def.World
			g.ticks = def.ticks
			g.simulationSpeed = def.simulationSpeed
		case ebiten.KeyQ:
			os.Exit(0)
		}
	}

	modulo := math.Max(math.Pow(g.simulationSpeed, -1), 1.0)
	if g.ticks%int(modulo) == 0 {
		g.World.update()
	}
	g.ticks++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for y := 0; y < g.World.Height; y++ {
		for x := 0; x < g.World.Width; x++ {
			tile := g.World.Tiles[y][x]
			tile.Draw(screen)
		}
	}

	for _, entity := range g.World.Entities {
		entity.Draw(screen)
	}

	g.Debug(screen)
}

func (g *Game) Debug(screen *ebiten.Image) {
	rabbits := g.World.Count(Rabbit)
	foxes := g.World.Count(Fox)
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"FPS: %.2f\n"+
			"Simulation Speed: %.2f\n"+
			"Rabbit Count: %d\n"+
			"Foxes Count: %d", ebiten.ActualFPS(), g.simulationSpeed, rabbits, foxes))

}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return g.World.Width * TileSize, g.World.Height * TileSize
}

func main() {
	LoadAssets()
	CurrentGame = NewGame()
	ebiten.SetWindowSize(WorldWidth*TileSize, WorldHeight*TileSize)
	ebiten.SetWindowTitle("Rabbits and Foxes")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(CurrentGame); err != nil {
		log.Fatal(err)
	}
}
