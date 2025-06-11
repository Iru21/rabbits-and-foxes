package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"math"
	"os"
)

var isHoldingLeftMouseButton = false

type Game struct {
	ui                      *UI
	World                   *World
	ticks                   int
	simulationSpeed         float64
	rabbitPopulationHistory []int
	foxPopulationHistory    []int
}

func NewGame() *Game {
	return &Game{
		ui:                      NewUI(),
		World:                   NewWorld(),
		ticks:                   0,
		simulationSpeed:         0.2,
		rabbitPopulationHistory: make([]int, 0),
		foxPopulationHistory:    make([]int, 0),
	}
}

func (g *Game) Reset() {
	def := NewGame()
	g.World = def.World
	g.ticks = def.ticks
	g.simulationSpeed = def.simulationSpeed
	g.rabbitPopulationHistory = make([]int, 0)
	g.foxPopulationHistory = make([]int, 0)
}

func (g *Game) CheckHotkeys() {
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
			g.Reset()
		case ebiten.KeyQ:
			g.Stop()
		}
	}
}

func (g *Game) CheckClick() {
	isPressed := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || isHoldingLeftMouseButton
	if isPressed {
		isHoldingLeftMouseButton = true
		x, y := ebiten.CursorPosition()
		g.ui.HandleClick(x, y)
		if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			isHoldingLeftMouseButton = false
		}
	}
}

func (g *Game) Simulate() {
	modulo := math.Max(math.Pow(g.simulationSpeed, -1), 1.0)
	if g.ticks%int(modulo) == 0 {
		g.World.update()
	}
	g.ticks++
}

func (g *Game) Update() error {
	g.CheckHotkeys()
	g.CheckClick()
	g.Simulate()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if Render {
		for y := 0; y < g.World.Height; y++ {
			for x := 0; x < g.World.Width; x++ {
				tile := g.World.Tiles[y][x]
				tile.Draw(screen)
			}
		}

		for _, entity := range g.World.Entities {
			entity.GetEntity().Draw(screen)
		}
	}

	g.ui.Draw(screen)
	g.Debug(screen)
}

func (g *Game) Debug(screen *ebiten.Image) {
	rabbits := g.World.Count(RabbitSpecies)
	foxes := g.World.Count(FoxSpecies)
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"FPS: %.2f\n"+
			"TPS: %.2f\n"+
			"Ticks: %d\n"+
			"Simulation Speed: %.2f\n"+
			"Rabbit Count: %d\n"+
			"Foxes Count: %d",
		ebiten.ActualFPS(),
		ebiten.ActualTPS(),
		g.ticks,
		g.simulationSpeed,
		rabbits,
		foxes,
	))

}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return g.World.Width * TileSize, g.World.Height * TileSize
}

func (g *Game) Stop() {
	DrawChart(g.rabbitPopulationHistory, g.foxPopulationHistory)
	os.Exit(0)
}
