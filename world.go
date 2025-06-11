package main

import (
	"math/rand"
)

type World struct {
	Width    int
	Height   int
	Tiles    [][]*Tile
	Entities []EntityBehavior
}

func NewWorld() *World {
	tiles := make([][]*Tile, WorldHeight)
	for y := 0; y < WorldHeight; y++ {
		tiles[y] = make([]*Tile, WorldWidth)
		for x := 0; x < WorldWidth; x++ {
			var density = 0.0
			if rand.Float32() < 0.05 {
				density = 0.2 + rand.Float64()*0.3
			}
			tiles[y][x] = NewTile(x, y, density)
		}
	}
	entities := make([]EntityBehavior, 0, StartingRabbits+StartingFoxes)
	for i := 0; i < StartingRabbits; i++ {
		entities = append(entities, NewRabbit(rand.Intn(WorldWidth), rand.Intn(WorldHeight)))
	}
	for i := 0; i < StartingFoxes; i++ {
		entities = append(entities, NewFox(rand.Intn(WorldWidth), rand.Intn(WorldHeight)))
	}
	return &World{
		Width:    WorldWidth,
		Height:   WorldHeight,
		Tiles:    tiles,
		Entities: entities,
	}
}

func (w *World) Count(s Species) int {
	count := 0
	for _, e := range w.Entities {
		if e.GetEntity().species == s {
			count++
		}
	}
	return count
}

func (w *World) CountEntitiesAt(x, y int) int {
	count := 0
	for _, e := range w.Entities {
		entity := e.GetEntity()
		if entity != nil && entity.X == x && entity.Y == y {
			count++
		}
	}
	return count
}

func (w *World) GetTileAt(x, y int) *Tile {
	if x < 0 || x >= w.Width || y < 0 || y >= w.Height {
		return nil
	}
	return w.Tiles[y][x]
}

func (w *World) AddEntity(entity EntityBehavior) {
	if entity == nil {
		return
	}
	w.Entities = append(w.Entities, entity)
}

func (w *World) RemoveEntity(entity EntityBehavior) {
	if entity == nil {
		return
	}
	for i, e := range w.Entities {
		if e.GetEntity().uuid == entity.GetEntity().uuid {
			if i == len(w.Entities) {
				w.Entities = w.Entities[:i]
				return
			} else {
				w.Entities = append(w.Entities[:i], w.Entities[i+1:]...)
			}
			return
		}
	}
}

func (w *World) GetEntityOfSpeciesAt(x, y int, species Species) *Entity {
	for _, e := range w.Entities {
		entity := e.GetEntity()
		if entity.X == x && entity.Y == y && entity.species == species {
			return entity
		}
	}
	return nil
}

func (w *World) update() {
	if CurrentGame.ticks%10 == 0 {
		rabbitCount := w.Count(RabbitSpecies)
		foxCount := w.Count(FoxSpecies)
		CurrentGame.rabbitPopulationHistory = append(CurrentGame.rabbitPopulationHistory, rabbitCount)
		CurrentGame.foxPopulationHistory = append(CurrentGame.foxPopulationHistory, foxCount)
		if rabbitCount <= 1 || foxCount <= 1 {
			println("Rabbits population:", rabbitCount, "Foxes Population:", foxCount)
			CurrentGame.Stop()
		}
	}

	for _, tileRow := range w.Tiles {
		for _, tile := range tileRow {
			tile.Update(w)
		}
	}

	for _, entity := range w.Entities {
		go entity.GetEntity().Update()
	}
}
