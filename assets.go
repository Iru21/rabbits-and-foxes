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
	GrassSprite  *ebiten.Image
	DirtSprite   *ebiten.Image
	RabbitSprite *ebiten.Image
	FoxSprite    *ebiten.Image
	FoxSprite16  *ebiten.Image
	FoxSprite32  *ebiten.Image
	FoxSprite48  *ebiten.Image
)

func LoadAssets() {
	load := func(path string, resize bool) *ebiten.Image {
		file, err := tileFiles.Open(path)
		if err != nil {
			log.Fatalf("Cannot open asset: %v", err)
		}
		img, _, err := image.Decode(file)
		if err != nil {
			log.Fatalf("Cannot decode asset: %v", err)
		}
		if resize && (img.Bounds().Dx() != TileSize || img.Bounds().Dy() != TileSize) {
			newImg := image.NewRGBA(image.Rect(0, 0, TileSize, TileSize))
			for y := 0; y < TileSize; y++ {
				for x := 0; x < TileSize; x++ {
					srcX := x * img.Bounds().Dx() / TileSize
					srcY := y * img.Bounds().Dy() / TileSize
					newImg.Set(x, y, img.At(srcX, srcY))
				}
			}
			img = newImg
		}
		return ebiten.NewImageFromImage(img)
	}

	GrassSprite = load("assets/grass.png", true)
	DirtSprite = load("assets/dirt.png", true)
	RabbitSprite = load("assets/rabbit.png", true)
	FoxSprite = load("assets/fox.png", true)
	FoxSprite16 = load("assets/fox.png", false)
	FoxSprite32 = load("assets/fox-32.png", false)
	FoxSprite48 = load("assets/fox-48.png", false)
}
