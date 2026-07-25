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
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	baseWidth    int     = 1280
	baseHeight   int     = 720
	spriteWidth  float64 = 64
	spriteHeight float64 = 64
	velocity     float64 = 10
	frameCount   int     = 2
	mainMenu     int     = iota
	gameWorld
	gameOver
)

var (
	scene          int
	biome          int
	points         int
	multiplier     float64
	windowWidth    int
	windowHeight   int
	animationIndex int
	icon           *ebiten.Image
	player         *ebiten.Image
	enemies        *ebiten.Image
	p              *Player
	e              *Enemy
	g              *Game
	op             *ebiten.DrawImageOptions
)

func init() {
	var err error
	// define window size
	windowWidth, windowHeight = 1280, 720
	windowWidth = windowWidth * baseHeight / windowHeight
	// load images
	icon, _, err = ebitenutil.NewImageFromFile("assets/base.png")
	player, _, err = ebitenutil.NewImageFromFile("assets/sprites/player.png")
	enemies, _, err = ebitenutil.NewImageFromFile("assets/sprites/enemies.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	tick        int
	enemies     Enemies
	initialised bool
}

func collide(a image.Rectangle, b image.Rectangle) bool {
	return a.Min.X < b.Max.X &&
		a.Max.X > b.Min.X &&
		a.Min.Y < b.Max.Y &&
		a.Max.Y > b.Min.Y
}

func (g *Game) Init() {
	defer func() {
		g.initialised = true
	}()

	op = &ebiten.DrawImageOptions{}
	scene = gameWorld

	// initialise enemies
	e = &Enemy{}
	g.enemies.enemies = make([]*Enemy, 2048)
	g.enemies.count = 12 * windowWidth / baseWidth
	for i := range g.enemies.enemies {
		x := float64(rand.Intn(1280))
		y := -float64(windowHeight) - float64(rand.Intn(windowHeight))
		state := rand.Intn(5)
		width := 64.0
		height := 64.0
		var vel float64
		switch state {
		case 1:
			vel = 11
		case 3:
			vel = 16
		case 4:
			vel = 12
		default:
			vel = 10
		}
		g.enemies.enemies[i] = &Enemy{
			x:      x,
			y:      y,
			width:  width,
			height: height,
			state:  state,
			vel:    vel,
		}
	}

	// initialise player
	p = &Player{}
	p.x = float64(windowWidth) / 2
	p.y = float64(windowHeight)
	p.width = 64
	p.height = 64
	p.vel = 10
}

func (g *Game) Update() error {
	if !g.initialised {
		g.Init()
	}
	g.tick++
	animationIndex = (g.tick / 5) % frameCount
	windowWidth, windowHeight = ebiten.WindowSize()
	windowWidth = windowWidth * baseHeight / windowHeight
	g.enemies.count = 12 * windowWidth / baseWidth
	g.enemies.Update()
	p.Update()
	return nil
}

func (g *Game) Reset() {
	// reset player
	points = 0
	p.x = float64(windowWidth) / 2
	p.y = float64(baseHeight)
	p.width = 64
	p.height = 64
	p.vel = 10
	// reset enemies
	e.x = float64(rand.Intn(windowWidth))
	e.y = 0 - float64(windowHeight) - float64(rand.Intn(windowHeight))
	e.state = rand.Intn(5)
	switch e.state {
	case 1:
		e.vel = 11
	case 3:
		e.vel = 16
	case 4:
		e.vel = 12
	default:
		e.vel = 10
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch scene {
	case mainMenu:
		screen.Fill(color.RGBA{50, 50, 50, 255})
	case gameWorld:
		screen.Fill(color.RGBA{0, 50, 0, 255})
		// to-do: draw world
		// to-do: draw clouds
		// draw enemies
		for i := range g.enemies.count {
			e := g.enemies.enemies[i]
			op.GeoM.Reset()
			op.GeoM.Translate(-e.width/2, -e.height)
			op.GeoM.Scale(1, 1)
			op.GeoM.Translate(e.x, e.y)
			screen.DrawImage(enemies.SubImage(image.Rect(e.sheetX, e.sheetY, e.sheetX+int(spriteWidth), e.sheetY+int(spriteHeight))).(*ebiten.Image), op)
		}

		// draw player
		op.GeoM.Reset()
		op.GeoM.Translate(-p.width/2, -p.height)
		op.GeoM.Scale(1, 1)
		op.GeoM.Translate(p.x, p.y)
		screen.DrawImage(player.SubImage(image.Rect(p.sheetX, p.sheetY, p.sheetX+int(spriteWidth), p.sheetY+int(spriteHeight))).(*ebiten.Image), op)
		ebitenutil.DebugPrint(screen, "32-bit Squadron v1.1.0-alpha.1 - In Game - "+fmt.Sprintf("Points: %d", points))
	case gameOver:
		screen.Fill(color.RGBA{0, 0, 0, 255})
		ebitenutil.DebugPrint(screen, "32-bit Squadron v1.1.0-alpha.1 - Game Over (Press Space to Restart) - "+fmt.Sprintf("Points: %d", points))
	default:
		scene = gameWorld
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, baseHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("32-bit Squadron")
	ebiten.SetWindowIcon([]image.Image{icon})
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// main game loop
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
