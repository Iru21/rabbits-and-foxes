package main

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

//go:embed assets/*.png
var tileFiles embed.FS

var (
	GrassSprite *ebiten.Image
	DirtSprite  *ebiten.Image
	// RabbitSprite *ebiten.Image
	// FoxSprite    *ebiten.Image
)

func LoadAssets() {
	load := func(path string) *ebiten.Image {
		file, err := tileFiles.Open(path)
		if err != nil {
			log.Fatalf("Cannot open asset: %v", err)
		}
		img, _, err := image.Decode(file)
		if err != nil {
			log.Fatalf("Cannot decode asset: %v", err)
		}
		return ebiten.NewImageFromImage(img)
	}

	GrassSprite = load("assets/grass.png")
	DirtSprite = load("assets/dirt.png")
	// RabbitSprite = load("assets/rabbit.png")
	// FoxSprite = load("assets/fox.png")
}
