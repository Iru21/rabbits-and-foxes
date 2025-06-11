package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"math/rand"
)

const (
	GrassSpreadChance = 0.01
	GrassGrowthRate   = 0.01
)

type Tile struct {
	X, Y                      int
	GrassDensity              float64
	AvailableSpreadDirections []int
	RandomTileRotation        int
}

func NewTile(x, y int, density float64) *Tile {
	return &Tile{
		X:                         x,
		Y:                         y,
		GrassDensity:              density,
		AvailableSpreadDirections: []int{1, 2, 3, 4, 5, 6, 7, 8},
		RandomTileRotation:        rand.Intn(4) * 90,
	}
}

func (t *Tile) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	rand.NewSource(int64(t.X * t.Y))
	s := GrassSprite.Bounds().Size()
	op.GeoM.Translate(-float64(s.X)/2, -float64(s.Y)/2)
	op.GeoM.Rotate(float64(t.RandomTileRotation) * (math.Pi / 180))
	op.GeoM.Translate(float64(t.X*TileSize+s.X/2), float64(t.Y*TileSize+s.Y/2))

	if t.GrassDensity > 0 {
		alpha := float32(0.5 + (t.GrassDensity / 2))
		op.ColorScale.Scale(alpha, alpha, alpha, alpha)
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
		t.GrassDensity = Clamp(t.GrassDensity+GrassGrowthRate, 0.0, 1.0)
	}

	if rand.Float32() < GrassSpreadChance && len(t.AvailableSpreadDirections) > 0 {
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
				world.Tiles[newY][newX].GrassDensity = 0.2 + rand.Float64()*0.3
				spreadTo = dir
				break
			}
		}
		if spreadTo != 0 {
			t.AvailableSpreadDirections = RemoveFromSlice(t.AvailableSpreadDirections, spreadTo)
		}
	}
}
