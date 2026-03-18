package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	player  *ebiten.Image
	enemies *ebiten.Image
	scaling float64 = 1
)

func init() {
	var err error
	player, _, err = ebitenutil.NewImageFromFile("assets/player.png")
	enemies, _, err = ebitenutil.NewImageFromFile("assets/enemies.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var width, height = ebiten.WindowSize()
	scaling := (float64(width) + float64(height)) / (1280 + 720)

	player_op := &ebiten.DrawImageOptions{}
	player_op.GeoM.Translate(float64(width)/2-96, float64(height)-128)
	player_op.GeoM.Scale(1, 1)
	enemy_op := &ebiten.DrawImageOptions{}
	enemy_op.GeoM.Translate(200, 200)
	enemy_op.GeoM.Scale(1*scaling, 1*scaling)

	screen.Fill(color.RGBA{0, 50, 0, 255})
	ebitenutil.DebugPrint(screen, "v1.1.0-Alpha1")
	screen.DrawImage(enemies, enemy_op)
	screen.DrawImage(player, player_op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("32-bit Squadron")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// main game loop
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
