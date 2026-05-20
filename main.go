package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	PlayerImage *ebiten.Image
	X, Y        float64
}

func (g *Game) Update() error {

	var velocity float64 = 2

	if ebiten.IsKeyPressed(ebiten.KeyW) && g.Y > 0 {
		g.Y = g.Y - velocity
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) && g.Y < 240-16 {
		g.Y = g.Y + velocity
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) && g.X < 320-16 {
		g.X = g.X + velocity
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) && g.X > 0 {
		g.X = g.X - velocity
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{120, 180, 255, 255})

	playerOptions := ebiten.DrawImageOptions{}
	playerOptions.GeoM.Translate(g.X, g.Y)

	screen.DrawImage(
		g.PlayerImage.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&playerOptions,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Sample RPG")

	playerImage, _, err := ebitenutil.NewImageFromFile("assets/images/villager_sprite_sheet.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{PlayerImage: playerImage, X: 100, Y: 100}); err != nil {
		log.Fatal(err)
	}
}
