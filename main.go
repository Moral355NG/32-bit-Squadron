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
	Scaling        float64
	velocity       float64
	monitorWidth   int
	monitorHeight  int
	icon           *ebiten.Image
	player         *ebiten.Image
	playerPosition float64
	enemies        *ebiten.Image
)

func init() {
	var err error
	monitorWidth, monitorHeight = ebiten.Monitor().Size()
	icon, _, err = ebitenutil.NewImageFromFile("assets/base.png")
	player, _, err = ebitenutil.NewImageFromFile("assets/player.png")
	playerPosition = (1280 / 2) - 32
	enemies, _, err = ebitenutil.NewImageFromFile("assets/enemies.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	playerMovement()
	playerAnim()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var width, height int = ebiten.WindowSize()
	// use when scaling is implemented Scaling := (float64(width) + float64(height)) / (1280 + 720)

	player_op := &ebiten.DrawImageOptions{}
	player_op.GeoM.Translate(playerPosition, float64(height)-64)
	player_op.GeoM.Scale(1, 1)
	enemy_op := &ebiten.DrawImageOptions{}
	enemy_op.GeoM.Translate(200, 200)
	enemy_op.GeoM.Scale(float64(width), 1)

	screen.Fill(color.RGBA{0, 50, 0, 255})
	ebitenutil.DebugPrint(screen, "v1.1.0-Alpha1")
	screen.DrawImage(enemies, enemy_op)
	// formatting for subimage process(x, y, x + width, y + height)
	screen.DrawImage(player.SubImage(image.Rect(playerSheetX, playerSheetY, playerSheetX+64, playerSheetY+64)).(*ebiten.Image), player_op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	// add when scaling is implemented ebiten.SetWindowSize(monitorWidth/2, monitorHeight/2)
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("32-bit Squadron")
	ebiten.SetWindowIcon([]image.Image{icon})
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// main game loop
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
