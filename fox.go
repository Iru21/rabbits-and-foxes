package main

import (
	"github.com/google/uuid"
	"math/rand"
)

var FoxInitialEnergy = 15.0
var FoxMaxEnergy = 25.0
var FoxEnergyLoss = 0.3
var FoxReproductionCooldown = 7.0
var FoxEatGain = 10.0

type Fox struct {
	EntityBehavior
	Entity
}

func (f *Fox) GetEntity() *Entity {
	return &f.Entity
}

func (f *Fox) Reproduce() EntityBehavior {
	return NewFox(f.X, f.Y)
}

func (f *Fox) Eat() {
	world := CurrentGame.World
	rabbit := world.GetEntityOfSpeciesAt(f.X, f.Y, RabbitSpecies, nil)
	if rabbit != nil && rabbit.energy > 0 {
		f.energy += FoxEatGain
		if f.energy < FoxMaxEnergy {
			f.energy = FoxMaxEnergy
		}
		rabbit.energy = 0
	}
}

func (f *Fox) FindRabbitIfHungry() *Entity {
	if f.energy >= FoxMaxEnergy/2 {
		return nil
	}
	world := CurrentGame.World
	return world.FindNearestEntityOfSpeciesWithLimitedDistance(f.X, f.Y, RabbitSpecies, nil, 10)
}

func (f *Fox) Move() {
	world := CurrentGame.World

	dx, dy := 0, 0
	if rabbit := f.FindRabbitIfHungry(); rabbit != nil {
		dx = Clamp(rabbit.X-f.X, -1, 1)
		dy = Clamp(rabbit.Y-f.Y, -1, 1)
	} else if partner := f.FindPartnerIfCanRepoduce(); partner != nil && partner.CanReproduce() {
		dx = Clamp(partner.X-f.X, -1, 1)
		dy = Clamp(partner.Y-f.Y, -1, 1)
	} else {
		dx = rand.Intn(3) - 1
		dy = rand.Intn(3) - 1
	}

	f.X = Clamp(f.X+dx, 0, world.Width-1)
	f.Y = Clamp(f.Y+dy, 0, world.Height-1)
	f.isFlipped = dx > 0
}

func (f *Fox) GetEnergyLoss() float64 {
	return FoxEnergyLoss
}

func (f *Fox) GetInitialEnergy() float64 {
	return FoxInitialEnergy
}

func (f *Fox) GetReproductionCooldown() float64 {
	return FoxReproductionCooldown
}

func NewFox(x, y int) *Fox {
	fox := &Fox{
		Entity: Entity{
			uuid:              uuid.New(),
			X:                 x,
			Y:                 y,
			sprite:            FoxSprite,
			isFlipped:         false,
			species:           FoxSpecies,
			energy:            FoxInitialEnergy,
			reproductionClock: FoxReproductionCooldown,
		},
	}
	fox.Entity.EntityBehavior = fox
	return fox
}
