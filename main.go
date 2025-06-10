package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
)

const (
	WorldHeight     = 50
	WorldWidth      = 50
	TileSize        = 16
	SimulationSpeed = 0.2
)

type Game struct {
	World *World
}

var ticks int

func (g *Game) Update() error {
	modulo := int(math.Pow(SimulationSpeed, -1))
	if ticks%modulo == 0 {
		g.World.update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for y := 0; y < g.World.Height; y++ {
		for x := 0; x < g.World.Width; x++ {
			tile := g.World.Tiles[y][x]
			tile.Draw(screen, x*TileSize, y*TileSize)
		}
	}

	for _, entity := range g.World.Entities {
		entity.Draw(screen)
	}

	g.Debug(screen)
}

func (g *Game) Debug(screen *ebiten.Image) {
	rabbits := g.World.count(Rabbit)
	foxes := g.World.count(Fox)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.2f\nSimulation Speed: %.2f\nRabbit count: %d\nFoxes count: %d", ebiten.ActualFPS(), SimulationSpeed, rabbits, foxes))

}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return g.World.Width * TileSize, g.World.Height * TileSize
}

func main() {
	LoadAssets()
	world := NewWorld()
	g := Game{World: world}
	ebiten.SetWindowSize(WorldWidth*TileSize, WorldHeight*TileSize)
	ebiten.SetWindowTitle("Rabbits and Foxes")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
