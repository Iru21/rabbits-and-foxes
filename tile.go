package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Tile struct {
	X, Y                      int
	GrassDensity              float32
	AvailableSpreadDirections []int
}

func NewTile(x, y int, density float32) *Tile {
	return &Tile{
		X:                         x,
		Y:                         y,
		GrassDensity:              density,
		AvailableSpreadDirections: []int{1, 2, 3, 4, 5, 6, 7, 8},
	}
}

func (t *Tile) Draw(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))

	if t.GrassDensity > 0 {
		alpha := 0.5 + (t.GrassDensity)/2
		op.ColorScale.Scale(1, 1, 1, alpha)
		screen.DrawImage(GrassSprite, op)
	} else {
		screen.DrawImage(DirtSprite, op)
	}
}

func (t *Tile) Update(world *World) {
	if t.GrassDensity == 0 {
		return
	}

	if t.GrassDensity < 1.0 {
		t.GrassDensity = Clamp(t.GrassDensity+0.01, 0.0, 1.0)
	}

	if rand.Float32() < 0.01 {
		direction := []int{1, 2, 3, 4, 5, 6, 7, 8}
		rand.Shuffle(len(direction), func(i, j int) {
			direction[i], direction[j] = direction[j], direction[i]
		})
		spreadTo := 0
		for _, dir := range direction {
			dx, dy := 0, 0
			switch dir {
			case 1: // Up
				dy = -1
			case 2: // Down
				dy = 1
			case 3: // Left
				dx = -1
			case 4: // Right
				dx = 1
			case 5: // Up-Left
				dx, dy = -1, -1
			case 6: // Up-Right
				dx, dy = 1, -1
			case 7: // Down-Left
				dx, dy = -1, 1
			case 8: // Down-Right
				dx, dy = 1, 1
			}

			newX := Clamp(t.X+dx, 0, world.Width-1)
			newY := Clamp(t.Y+dy, 0, world.Height-1)

			if world.Tiles[newY][newX].GrassDensity == 0 {
				world.Tiles[newY][newX].GrassDensity = 0.2 + rand.Float32()*0.3
				spreadTo = dir
				break
			}
		}
		if spreadTo != 0 {
			t.AvailableSpreadDirections = RemoveFromSlice(t.AvailableSpreadDirections, spreadTo)
		}
	}
}
