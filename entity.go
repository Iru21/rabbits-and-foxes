package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

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

func NewRabbit(x, y int) *Entity {
	return &Entity{
		X:         x,
		Y:         y,
		sprite:    RabbitSprite,
		isFlipped: false,
		species:   Rabbit,
	}
}

func NewFox(x, y int) *Entity {
	return &Entity{
		X:         x,
		Y:         y,
		sprite:    FoxSprite,
		isFlipped: false,
		species:   Fox,
	}
}

func (e *Entity) Update() {
	e.Move()
}

func (e *Entity) Draw(screen *ebiten.Image) {
	if e.sprite == nil {
		return
	}

	opts := &ebiten.DrawImageOptions{}
	if e.isFlipped {
		opts.GeoM.Scale(-1, 1)
		opts.GeoM.Translate(float64((e.X+1)*TileSize), float64(e.Y*TileSize))
	} else {
		opts.GeoM.Translate(float64(e.X*TileSize), float64(e.Y*TileSize))

	}
	screen.DrawImage(e.sprite, opts)

}

func (e *Entity) Move() {
	if rand.Float32() < 0.5 {
		dx := rand.Intn(3) - 1
		dy := rand.Intn(3) - 1
		world := CurrentGame.World
		newX := Clamp(e.X+dx, 0, world.Width-1)
		newY := Clamp(e.Y+dy, 0, world.Height-1)

		if !world.IsEntityAt(newX, newY) {
			e.X = newX
			e.Y = newY
		}

		if dx < 0 {
			e.isFlipped = false
		} else if dx > 0 {
			e.isFlipped = true
		}
	}
}
