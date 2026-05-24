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
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// TO-DO: Add collisions
	//"https://github.com/tducasse/ebiten-collisions"
)

const (
	spriteWidth  float64 = 64
	spriteHeight float64 = 64
	velocity     float64 = 10
	frameCount   int     = 2
)

var (
	transform      float64
	scaling        float64
	multiplier     float64
	width          int
	height         int
	monitorWidth   int
	monitorHeight  int
	animationIndex int
	icon           *ebiten.Image
	player         *ebiten.Image
	enemies        *ebiten.Image
	p              *Player
	e              *Enemy
)

func init() {
	var err error
	monitorWidth, monitorHeight = ebiten.Monitor().Size()
	width, height = monitorWidth/3*2, monitorHeight/3*2
	// Load Images
	icon, _, err = ebitenutil.NewImageFromFile("assets/base.png")
	player, _, err = ebitenutil.NewImageFromFile("assets/player.png")
	enemies, _, err = ebitenutil.NewImageFromFile("assets/enemies.png")
	p = &Player{}
	p.Init()
	e = &Enemy{}
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	op          *ebiten.DrawImageOptions
	tick        int
	enemies     Enemies
	initialised bool
}

func (g *Game) Init() {
	defer func() {
		g.initialised = true
	}()

	g.op = &ebiten.DrawImageOptions{}

	g.enemies.enemies = make([]*Enemy, 100)
	g.enemies.count = 16

	for i := range g.enemies.enemies {
		x := float64(rand.Intn(1280))
		y := 0 - float64(rand.Intn(height))
		state := rand.Intn(5)
		width := 64.0
		height := 64.0
		vel := 10.0
		g.enemies.enemies[i] = &Enemy{
			x:      x,
			y:      y,
			width:  width,
			height: height,
			state:  state,
			vel:    vel,
		}
	}
}

func (g *Game) Reset() {
	p.Init()
}

func (g *Game) Update() error {
	if !g.initialised {
		g.Init()
	}
	g.tick++
	animationIndex = (g.tick / 5) % frameCount
	width, height = ebiten.WindowSize()
	p.Update()
	g.enemies.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	scaling = float64(height) / 720
	transform = float64(width) / 1280

	screen.Fill(color.RGBA{0, 50, 0, 255})
	ebitenutil.DebugPrint(screen, "v1.1.0-alpha.1")
	p.op = &ebiten.DrawImageOptions{}
	// draw player
	p.op.GeoM.Translate(-p.width/2, -p.height)
	p.op.GeoM.Scale(1*scaling, 1*scaling)
	p.op.GeoM.Translate(p.x*transform, p.y)
	screen.DrawImage(player.SubImage(image.Rect(p.sheetX, p.sheetY, p.sheetX+int(spriteWidth), p.sheetY+int(spriteHeight))).(*ebiten.Image), p.op)
	// draw enemies
	for i := 0; i < g.enemies.count; i++ {
		e := g.enemies.enemies[i]
		g.op.GeoM.Reset()
		g.op.GeoM.Translate(-e.width/2, -e.height)
		g.op.GeoM.Scale(1*scaling, 1*scaling)
		g.op.GeoM.Translate(e.x*transform, e.y)
		screen.DrawImage(enemies.SubImage(image.Rect(e.sheetX, e.sheetY, e.sheetX+int(spriteWidth), e.sheetY+int(spriteHeight))).(*ebiten.Image), g.op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("32-bit Squadron")
	ebiten.SetWindowIcon([]image.Image{icon})
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// main game loop
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
