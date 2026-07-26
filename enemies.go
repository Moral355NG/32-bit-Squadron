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
)

type Enemy struct {
	x, y   float64
	sheetX int
	sheetY int
	width  float64
	height float64
	vel    float64
	state  int
}

func (e *Enemy) Update() {
	if scene == gameWorld {
		if e.y > float64(baseHeight*2) {
			// enemy spawning when game is being played
			points += 1
			e.x = float64(rand.Intn(windowWidth))
			e.y = 0 - float64(rand.Intn(windowHeight))
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
			// move enemy down the screen
			e.y += e.vel
		}
		// handle enemy animation
		e.sheetX, e.sheetY = int(spriteWidth)*e.state, animationIndex*int(spriteHeight)
		// handle player enemy collision
		if collide(p.hitbox(), e.hitbox()) {
			scene = gameOver
		}
	} else {
		// reset enemy position
		if g.start() {
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
	}
}

func (e *Enemy) hitbox() image.Rectangle {
	return image.Rect(int(e.x-32), int(e.y-24), int(e.x+32), int(e.y+24))
}

type Enemies struct {
	enemies []*Enemy
	count   int
}

func (e *Enemies) Update() {
	for i := range e.count {
		e.enemies[i].Update()
	}
}
