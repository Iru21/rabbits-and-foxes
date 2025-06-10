package main

import "github.com/hajimehoshi/ebiten/v2"

type Species int

const (
	Rabbit Species = iota
	Fox
)

type Entity struct {
	X, Y      int
	sprite    *ebiten.Image
	isFlipped bool
	species   Species
}

func (e *Entity) Draw(screen *ebiten.Image) {
	if e.sprite == nil {
		return
	}

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(e.X*TileSize), float64(e.Y*TileSize))

	if e.isFlipped {
		opts.GeoM.Scale(-1, 1)
		opts.GeoM.Translate(float64(e.sprite.Bounds().Dx()), 0)
		screen.DrawImage(e.sprite, opts)
	} else {
		screen.DrawImage(e.sprite, opts)
	}
}

func (e *Entity) Move(dx, dy int, world *World) {
	newX := Clamp(e.X+dx, 0, world.Width-1)
	newY := Clamp(e.Y+dy, 0, world.Height-1)

	if world.Tiles[newY][newX] == nil {
		e.X = newX
		e.Y = newY
	}

	if dx < 0 {
		e.isFlipped = false
	} else if dx > 0 {
		e.isFlipped = true
	}
}
