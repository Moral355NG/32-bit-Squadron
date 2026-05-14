
// Copyright (C) 2026 Moral355NG
// GPL-3.0-or-later

// 32-bit-Squadron is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// 32-bit-Squadron is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// See <https://www.gnu.org/licenses/> for more details.

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
	scaling       float64
	velocity      float64
	width         int
	height        int
	monitorWidth  int
	monitorHeight int
	icon          *ebiten.Image
	player        *ebiten.Image
	enemies       *ebiten.Image
)

func init() {
	var err error
	width, height = 1280, 720
	monitorWidth, monitorHeight = ebiten.Monitor().Size()
	icon, _, err = ebitenutil.NewImageFromFile("assets/base.png")
	player, _, err = ebitenutil.NewImageFromFile("assets/player.png")
	enemies, _, err = ebitenutil.NewImageFromFile("assets/enemies.png")
	playerInit()
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	tick int
}

func (g *Game) Update() error {
	g.tick++
	width, height = ebiten.WindowSize()
	playerMovement()
	playerAnimation()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	frameIndex := 
	scaling = (float64(width) + float64(height)) / (1280 + 720)
	transform := float64(width) / 1280

	player_op := &ebiten.DrawImageOptions{}
	player_op.GeoM.Translate(-playerWidth/2, -PlayerHeight)
	player_op.GeoM.Scale(1*scaling, 1*scaling)
	player_op.GeoM.Translate(playerPosition*transform, float64(height))
	enemy_op := &ebiten.DrawImageOptions{}
	enemy_op.GeoM.Translate(200, 200)
	enemy_op.GeoM.Scale(float64(width), 1)

	screen.Fill(color.RGBA{0, 50, 0, 255})
	ebitenutil.DebugPrint(screen, "v1.1.0-alpha.1")
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
