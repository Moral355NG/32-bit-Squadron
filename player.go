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

import "github.com/hajimehoshi/ebiten/v2"

const (
	left   int = 0
	centre int = 1
	right  int = 2
)

var (
	playerSheetX   int
	playerSheetY   int
	playerState    int
	playerPosition float64
)

type Player struct {
	x, y   float64
	sheetX int
	sheetY int
	width  float64
	height float64
	vel    float64
	state  int
}

func (p *Player) init() {
	p.x = 1280 / 2
	p.vel = 10
}

func (p *Player) update() {
	// player movement and state
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && p.x > 0+spriteWidth/2 {
		p.x -= p.vel
		p.state = left
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && p.x < 1280-spriteWidth/2 {
		p.x += p.vel
		p.state = right
	} else {
		p.state = centre
	}
	// player animation
	p.sheetX, p.sheetY = int(spriteWidth)*p.state, animationIndex*int(spriteHeight)
}
