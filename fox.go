package main

var FoxInitialEnergy = 15.0
var FoxMaxEnergy = 25.0
var FoxEnergyLoss = 2.0
var FoxReproductionCooldown = 100.0
var FoxEatGain = 3.0

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
	rabbit := world.GetEntityOfSpeciesAt(f.X, f.Y, RabbitSpecies)
	if rabbit != nil {
		f.energy += FoxEatGain
		if f.energy < FoxMaxEnergy {
			f.energy = FoxMaxEnergy
		}
		rabbit.energy = 0
	}
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
