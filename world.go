package main

import (
	"math/rand"
)

type World struct {
	Width    int
	Height   int
	Tiles    [][]*Tile
	Entities []*Entity
}

func NewWorld() *World {
	tiles := make([][]*Tile, WorldHeight)
	for y := 0; y < WorldHeight; y++ {
		tiles[y] = make([]*Tile, WorldWidth)
		for x := 0; x < WorldWidth; x++ {
			var density float32 = 0.0
			if rand.Float32() < 0.05 {
				density = 0.2 + rand.Float32()*0.3
			}
			tiles[y][x] = NewTile(x, y, density)
		}
	}
	return &World{
		Width:    WorldWidth,
		Height:   WorldHeight,
		Tiles:    tiles,
		Entities: []*Entity{},
	}
}

func (w *World) count(s Species) int {
	count := 0
	for _, e := range w.Entities {
		if e.species == s {
			count++
		}
	}
	return count
}

func (w *World) update() {
	for _, tileRow := range w.Tiles {
		for _, tile := range tileRow {
			tile.Update(w)
		}
	}
}
