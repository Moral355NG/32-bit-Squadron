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

type Object struct {
	x, y   float64
	sheetX int
	sheetY int
	width  float64
	height float64
	vel    float64
	state  int
}

func (obj *Object) Update() {
	if scene == gameWorld {
		if obj.y > float64(baseHeight*2) || collide(obj.hitbox(), obj.hitbox()) {
			// object spawning when game is being played
			points += 1
			obj.x = float64(rand.Intn(windowWidth))
			obj.y = 0 - float64(rand.Intn(windowHeight))
			obj.state = rand.Intn(3)
			obj.vel = 7
		} else {
			// move object down the screen
			obj.y += e.vel
		}
		// handle object animation
		obj.sheetX, obj.sheetY = int(spriteWidth)*obj.state, animationIndex*int(spriteHeight)
	} else {
		// reset object position
		if g.start() {
			obj.x = float64(rand.Intn(windowWidth))
			obj.y = 0 - float64(windowHeight) - float64(rand.Intn(windowHeight))
			obj.state = rand.Intn(3)
			obj.vel = 7
		}
	}
}

func (obj *Object) hitbox() image.Rectangle {
	return image.Rect(int(obj.x-32), int(obj.y-32), int(obj.x+32), int(obj.y+32))
}

type Objects struct {
	objects []*Object
	count   int
}

func (obj *Objects) Update() {
	for i := range obj.count {
		obj.objects[i].Update()
	}
}
