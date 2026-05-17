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
	height float64
	vel    float64
	state  int
	img    *ebiten.Image
	op     *ebiten.DrawImageOptions
}

func (p *Player) init() {
	p.x = 1280 / 2
	p.width = 64
	p.height = 64
	p.vel = 10
}

func (p *Player) update() {
	// player movement and state
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && p.x > 0+p.width/2 {
		p.x -= p.vel
		p.state = left
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && p.x < 1280-p.width/2 {
		p.x += p.vel
		p.state = right
	} else {
		p.state = centre
	}
	// player animation
	p.sheetX, p.sheetY = int(spriteWidth)*p.state, animationIndex*int(spriteHeight)
}

func (p *Player) draw(screen *ebiten.Image) {
	p.op = &ebiten.DrawImageOptions{}
	// define player position and size on screen
	p.op.GeoM.Translate(-p.width/2, -p.height)
	p.op.GeoM.Scale(1*scaling, 1*scaling)
	p.op.GeoM.Translate(p.x*transform, float64(height))
	// formatting for subimage process(x, y, x + width, y + height)
	screen.DrawImage(player.SubImage(image.Rect(p.sheetX, p.sheetY, p.sheetX+int(spriteWidth), p.sheetY+int(spriteHeight))).(*ebiten.Image), p.op)
}
