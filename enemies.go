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

const ()

// WIP ADD COLLISIONS AND MULTIPLE ENEMY SPAWNING LATER
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

func (e *Enemy) init() {
	e.x = float64(rand.Intn(1280))
	e.y = float64(rand.Intn(32))
	e.width = 64
	e.height = 64
	e.vel = 10
}

func (e *Enemy) update() {
	// enemy movement
	if e.y > float64(height*2) {
		e.x = float64(rand.Intn(1280))
		e.y = float64(rand.Intn(32))
		e.state = rand.Intn(5)
	} else {
		e.y += e.vel * scaling
	}
	// enemy animation
	e.sheetX, e.sheetY = int(spriteWidth)*e.state, animationIndex*int(spriteHeight)
}

func (e *Enemy) draw(screen *ebiten.Image) {
	e.op = &ebiten.DrawImageOptions{}
	// define player position and size on screen
	e.op.GeoM.Translate(-e.width/2, -e.height)
	e.op.GeoM.Scale(1*scaling, 1*scaling)
	e.op.GeoM.Translate(e.x*transform, e.y)
	// formatting for subimage process(x, y, x + width, y + height)
	screen.DrawImage(enemies.SubImage(image.Rect(e.sheetX, e.sheetY, e.sheetX+int(spriteWidth), e.sheetY+int(spriteHeight))).(*ebiten.Image), e.op)
}
