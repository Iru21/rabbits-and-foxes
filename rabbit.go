package main

import "github.com/google/uuid"

var RabbitInitialEnergy = 10.0
var RabbitMaxEnergy = 20.0
var RabbitEnergyLoss = 0.2
var RabbitReproductionCooldown = 2.0
var RabbitEatGain = 5.0

type Rabbit struct {
	EntityBehavior
	Entity
}

func (r *Rabbit) GetEntity() *Entity {
	if r == nil {
		return nil
	}
	return &r.Entity
}

func (r *Rabbit) Reproduce() EntityBehavior {
	return NewRabbit(r.X, r.Y)
}

func (r *Rabbit) Eat() {
	world := CurrentGame.World
	tile := world.GetTileAt(r.X, r.Y)
	if tile.GrassDensity > 0.5 {
		r.energy += RabbitEatGain
		if r.energy < RabbitMaxEnergy {
			r.energy = RabbitMaxEnergy
		}
		tile.GrassDensity -= RabbitEatGain / 10
		if tile.GrassDensity < 0 {
			tile.GrassDensity = 0
		}
	}
}

func (r *Rabbit) GetEnergyLoss() float64 {
	return RabbitEnergyLoss
}

func (r *Rabbit) GetInitialEnergy() float64 {
	return RabbitInitialEnergy
}

func (r *Rabbit) GetReproductionCooldown() float64 {
	return RabbitReproductionCooldown
}

func NewRabbit(x, y int) *Rabbit {
	rabbit := &Rabbit{
		Entity: Entity{
			uuid:              uuid.New(),
			X:                 x,
			Y:                 y,
			sprite:            RabbitSprite,
			isFlipped:         false,
			species:           RabbitSpecies,
			energy:            RabbitInitialEnergy,
			reproductionClock: 10,
		},
	}
	rabbit.Entity.EntityBehavior = rabbit
	return rabbit
}
