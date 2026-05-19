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

const ()

// WIP
type Enemy struct {
	x, y   float64
	sheetX int
	sheetY int
	width  float64
	height float64
	vel    float64
	plane  int
	img    *ebiten.Image
	op     *ebiten.DrawImageOptions
}

func (e *Enemy) draw(screen *ebiten.Image) {
	e.op = &ebiten.DrawImageOptions{}
	// define player position and size on screen
	e.op.GeoM.Translate(-e.width/2, -e.height)
	e.op.GeoM.Scale(1*scaling, 1*scaling)
	e.op.GeoM.Translate(e.x*transform, float64(height))
	// formatting for subimage process(x, y, x + width, y + height)
	screen.DrawImage(enemies.SubImage(image.Rect(e.sheetX, e.sheetY, e.sheetX+int(spriteWidth), e.sheetY+int(spriteHeight))).(*ebiten.Image), e.op)
}
