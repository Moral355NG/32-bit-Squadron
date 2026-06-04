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
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var g *Game
var gameOver bool = false

type Enemy struct {
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

func (e *Enemy) Update() {
	// enemy movement
	if gameOver == true {
		e.x = float64(rand.Intn(1280))
		e.y = 0 - float64(rand.Intn(height))
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
	} else if e.y > float64(height*2) {
		e.x = float64(rand.Intn(1280))
		e.y = 0 - float64(height) - float64(rand.Intn(height))
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
	} else {
		e.y += e.vel * scaling
	}
	// handle enemy animation
	e.sheetX, e.sheetY = int(spriteWidth)*e.state, animationIndex*int(spriteHeight)
	// handle player enemy collision
	if collide(p.hitbox(), e.hitbox()) {
		gameOver = true
		g.Reset()
	}
}

func (e *Enemy) hitbox() image.Rectangle {
	return image.Rect(int(e.x)-32, int(e.y)-32, int(e.x)+32, int(e.y)+32)
}

type Enemies struct {
	enemies []*Enemy
	count   int
}

func (e *Enemies) Update() {
	for i := 0; i < e.count; i++ {
		e.enemies[i].Update()
	}
}
