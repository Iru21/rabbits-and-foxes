# Rabbits and Foxes

![fox.png](assets/fox.png)![fox.png](assets/fox.png)![fox.png](assets/fox.png)![fox.png](assets/fox.png)![fox.png](assets/fox.png)![fox.png](assets/fox.png)![fox.png](assets/fox.png)![fox.png](assets/fox.png)![fox.png](assets/fox.png)

A simple simulation of a predator-prey model written in go using ebitengine

![rabbit.png](assets/rabbit.png)![rabbit.png](assets/rabbit.png)![rabbit.png](assets/rabbit.png)![rabbit.png](assets/rabbit.png)![rabbit.png](assets/rabbit.png)![rabbit.png](assets/rabbit.png)![rabbit.png](assets/rabbit.png)![rabbit.png](assets/rabbit.png)![rabbit.png](assets/rabbit.png)

## Description

### Features

* Grass grows on tiles over time and spreads to neighboring empty spaces.
* Rabbits run away from foxes, find partner rabbits to reproduce, and die of starvation when they can't find any grass.
* Foxes seek out rabbits if hungry, find partner foxes to reproduce, and also die of starvation when they can't find any rabbits to eat.
* After the simulation finishes, a chart is generated showing the population of rabbits and foxes over time.
* 16x16 pixel art sprites
* Manage the simulation using keyboard hotkeys or on-screen buttons

### Detailed mechanics

#### 1. Energy

   * Rabbits and foxes have an energy level that decreases every turn.
   * Rabbits gain energy by eating grass, while foxes gain energy by eating rabbits.
   * When an animal's energy reaches zero, it dies.
   * Animals can only reproduce if they have more than a certain energy threshold.

#### 2. Reproduction

   * Rabbits and foxes can reproduce when they have enough energy, find a partner of the same species, and don't have any other priority actions (like fleeing or hunting).
   * Animals can reproduce only if both are the only animals on their tile.

#### 3. Movement

   * Animals can move to adjacent tiles if they are empty or occupied by grass.
   * Rabbits flee from foxes, find partners to reproduce, and search for grass to eat.
   * Foxes will move towards the nearest rabbit if they are hungry, and will also search for partners to reproduce.
   * Animals can move to tiles that are occupied by other animals.

#### 4. Grass

   * Grass grows on empty tiles over time.
   * Grass can spread to neighboring empty tiles.
   * Rabbits can eat grass to gain energy.

### Observations

* The simulation can exhibit oscillations in populations, where the number of rabbits and foxes rise and fall over time.
* The populations can stabilize at certain levels, depending on the initial conditions and parameters.
* Usual end case: foxes eat all rabbits, then die out due to starvation, leaving only grass behind.
* Rabbits run away in groups from foxes, which creates interesting patterns of movement at higher simulation speeds.
* Further tweaking of behavior parameters can lead to more complex interactions and behaviors which in turn can lead to a longer lasting simulation. 

## Keyboard hotkeys

| Key        | Action |
|------------| --- |
| `Up arrow` | Speed up simulation by 0.01x |
| `Down arrow` | Slow down simulation by 0.01x |
| `Left arrow` | Decrease simulation speed by 1x |
| `Right arrow` | Increase simulation speed by 1x |
| `R` | Reset simulation |
| `Space` | Pause simulation |


## Installation

Download the latest release from the [releases page](https://github.com/Iru21/rabbits-and-foxes/releases)

## Building from source

### Requirements

- go v1.24.3

### Instructions

```bash
$ git clone https://github.com/Iru21/rabbits-and-foxes && cd rabbits-and-foxes
```

```bash
$ go mod tidy
```

```bash
$ go build
```

### Customization

You can customize the simulation by modifying the constants at the top of [main.go](main.go)

```go
const (
	StartingRabbits = 100   // Initial number of rabbits and foxes in the simulation
	StartingFoxes   = 20    //
	WorldHeight     = 60    // Height of the world in tiles
	WorldWidth      = 60    // Width of the world in tiles
	TileSize        = 16    // Size of each tile in pixels
	Render          = true  // Whether to render the simulation
	TPS             = 100   // Maxium ticks per second
)
```

## Screenshots

![simulation.png](screenshots/simulation.png)
![chart.png](screenshots/chart.png)