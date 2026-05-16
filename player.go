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
	playerWidth  float64 = 64
	playerHeight float64 = 64
	left         int     = 0
	centre       int     = 1
	right        int     = 2
)

var (
	playerSheetX   int
	playerSheetY   int
	playerState    int
	playerPosition float64
)

func playerInit() {
	playerPosition = 1280 / 2
}

func updatePlayer() {
	// player movement and state
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && playerPosition > 0+playerWidth/2 {
		playerPosition -= velocity
		playerState = left
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && playerPosition < 1280-playerWidth/2 {
		playerPosition += velocity
		playerState = right
	} else {
		playerState = centre
	}
	// player animation
	playerSheetX, playerSheetY = int(playerWidth)*playerState, animationIndex*int(playerHeight)
}
