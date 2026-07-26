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

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	left   int = 0
	centre int = 1
	right  int = 2
)

type Player struct {
	x, y   float64
	sheetX int
	sheetY int
	width  float64
	height float64 ``
	vel    float64
	state  int
}

func (p *Player) Update() {
	p.y = float64(baseHeight)
	//handle player input

	// player movement
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.state = centre
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && p.x > 0+p.width/2 && scene == gameWorld {
		p.x -= p.vel
		p.state = left
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && p.x < float64(windowWidth)-p.width/2 && scene == gameWorld {
		p.x += p.vel
		p.state = right
	} else {
		p.state = centre
	}

	// reset player position
	if scene == gameOver && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		points = 0
		p.x = float64(windowWidth) / 2
		p.y = float64(baseHeight)
		p.width = 64
		p.height = 64
		p.vel = 10
		scene = gameWorld
	}

	//fullscreen
	if inpututil.IsKeyJustPressed(ebiten.KeyF4) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	// player animation
	p.sheetX, p.sheetY = int(spriteWidth)*p.state, animationIndex*int(spriteHeight)
}

func (p *Player) hitbox() image.Rectangle {
	return image.Rect(int(p.x-32), int(p.y-24), int(p.x+32), int(p.y+24))
}
