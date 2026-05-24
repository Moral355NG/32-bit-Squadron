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

func (p *Player) Init() {
	p.x = 1280 / 2
	p.y = float64(height)
	p.width = 64
	p.height = 64
	p.vel = 10
}

func (p *Player) Update() {
	p.y = float64(height)
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
	// player speed control
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		p.vel = 5
	} else {
		p.vel = 10
	}
	// player animation
	p.sheetX, p.sheetY = int(spriteWidth)*p.state, animationIndex*int(spriteHeight)
}

func (p *Player) hitbox() image.Rectangle {
	return image.Rect(int(p.x)-32, int(p.y)-32, int(p.x)+32, int(p.y)+32)
}
