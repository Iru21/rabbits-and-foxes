package main

import (
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type EntityBehavior interface {
	Reproduce() EntityBehavior
	Eat()
	GetEntity() *Entity
	GetEnergyLoss() float64
	GetInitialEnergy() float64
	GetReproductionCooldown() float64
}

type Species int

const (
	FoxSpecies Species = iota
	RabbitSpecies
)

type Entity struct {
	EntityBehavior
	uuid              uuid.UUID
	X, Y              int
	sprite            *ebiten.Image
	isFlipped         bool
	species           Species
	energy            float64
	reproductionClock float64
}

func (e *Entity) CanReproduce() bool {
	return e.reproductionClock <= 0 && e.energy > e.GetInitialEnergy()
}

func (e *Entity) Update() {
	world := CurrentGame.World
	e.energy -= e.GetEnergyLoss()
	if e.energy <= 0 {
		world.RemoveEntity(e)
		return
	}
	e.Move()
	e.Eat()

	if other := world.GetEntityOfSpeciesAt(e.X, e.Y, e.species); other != nil {
		if world.CountEntitiesAt(e.X, e.Y) < 3 && e.CanReproduce() && other.CanReproduce() && e != other {
			child := e.Reproduce()
			CurrentGame.World.AddEntity(child)
			e.reproductionClock = e.GetReproductionCooldown()
			other.reproductionClock = other.GetReproductionCooldown()
		}
	}

	if e.reproductionClock > 0 {
		e.reproductionClock--
	}

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
	world := CurrentGame.World
	dx := rand.Intn(3) - 1
	dy := rand.Intn(3) - 1
	newX := Clamp(e.X+dx, 0, world.Width-1)
	newY := Clamp(e.Y+dy, 0, world.Height-1)

	if dx < 0 {
		e.isFlipped = false
	} else if dx > 0 {
		e.isFlipped = true
	}
	e.X = newX
	e.Y = newY

}
