package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	Scaling       float64
	monitorWidth  int
	monitorHeight int
	player        *ebiten.Image
	enemies       *ebiten.Image
)

func init() {
	var err error
	monitorWidth, monitorHeight = ebiten.Monitor().Size()
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
	var width, height int = ebiten.WindowSize()
	// use when scaling is implemented Scaling := (float64(width) + float64(height)) / (1280 + 720)

	player_op := &ebiten.DrawImageOptions{}
	player_op.GeoM.Translate(float64(width)/2-32, float64(height)-64)
	player_op.GeoM.Scale(1, 1)
	enemy_op := &ebiten.DrawImageOptions{}
	enemy_op.GeoM.Translate(200, 200)
	enemy_op.GeoM.Scale(1, 1)

	screen.Fill(color.RGBA{0, 50, 0, 255})
	ebitenutil.DebugPrint(screen, "v1.1.0-Alpha1")
	screen.DrawImage(enemies, enemy_op)
	// formatting for image.Rect(x, y, x+width, y+height) x, and y are positions on the sprite sheet
	screen.DrawImage(player.SubImage(image.Rect(64, 0, 64+64, 0+64)).(*ebiten.Image), player_op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	// add when scaling is implemented ebiten.SetWindowSize(monitorWidth/2, monitorHeight/2)
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("32-bit Squadron")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// main game loop
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
