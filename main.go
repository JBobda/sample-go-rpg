package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Image *ebiten.Image
	X, Y  float64
}

type Player struct {
	*Sprite
	Health uint
}

type Enemy struct {
	*Sprite
	FollowPlayer bool
}

type Game struct {
	player       *Player
	enemies      []*Enemy
	tilemapJSON  *TilemapJSON
	tilemapImage *ebiten.Image
}

func (g *Game) Update() error {

	var velocity float64 = 2

	if ebiten.IsKeyPressed(ebiten.KeyW) && g.player.Y > 0 {
		g.player.Y -= velocity
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) && g.player.Y < 240-16 {
		g.player.Y += velocity
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) && g.player.X < 320-16 {
		g.player.X += velocity
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) && g.player.X > 0 {
		g.player.X -= velocity
	}

	for _, sprite := range g.enemies {
		if sprite.X < g.player.X {
			sprite.X += 1
		} else if sprite.X > g.player.X {
			sprite.X -= 1
		}

		if sprite.Y < g.player.Y {
			sprite.Y += 1
		} else if sprite.Y > g.player.Y {
			sprite.Y -= 1
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{120, 180, 255, 255})

	options := ebiten.DrawImageOptions{}

	for _, layer := range g.tilemapJSON.Layers {
		for index, id := range layer.Data {
			x := index % layer.Width
			y := index / layer.Width

			x *= 16
			y *= 16

			srcX := (id - 1) % 22
			srcY := (id - 1) / 22

			srcX *= 16
			srcY *= 16

			options.GeoM.Translate(float64(x), float64(y))

			screen.DrawImage(
				g.tilemapImage.SubImage(
					image.Rect(srcX, srcY, srcX+16, srcY+16),
				).(*ebiten.Image),
				&options,
			)

			options.GeoM.Reset()
		}
	}

	options.GeoM.Translate(g.player.X, g.player.Y)

	screen.DrawImage(
		g.player.Image.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&options,
	)

	options.GeoM.Reset()

	for _, sprite := range g.enemies {
		options.GeoM.Translate(sprite.X, sprite.Y)

		screen.DrawImage(
			sprite.Image.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image),
			&options,
		)

		options.GeoM.Reset()
	}
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

	knightImage, _, err := ebitenutil.NewImageFromFile("assets/images/knight_sprite_sheet.png")
	if err != nil {
		log.Fatal(err)
	}

	tilemapImage, _, err := ebitenutil.NewImageFromFile("assets/images/TilesetFloor.png")
	if err != nil {
		log.Fatal(err)
	}

	tilemapJSON, err := NewTilemapJSON("assets/maps/spawn.json")
	if err != nil {
		log.Fatal(err)
	}

	game := Game{
		player: &Player{
			&Sprite{
				Image: playerImage,
				X:     100,
				Y:     100,
			},
			100,
		},
		enemies: []*Enemy{
			{
				&Sprite{
					Image: knightImage,
					X:     50,
					Y:     50,
				},
				true,
			},
			{
				&Sprite{
					Image: knightImage,
					X:     75,
					Y:     75,
				},
				true,
			},
		},
		tilemapJSON:  tilemapJSON,
		tilemapImage: tilemapImage,
	}

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
