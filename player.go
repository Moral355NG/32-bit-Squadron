
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

var (
	playerSheetX   int
	playerSheetY   int
	playerState    int
	playerWidth    float64
	PlayerHeight   float64
	playerPosition float64
)

func playerInit() {
	playerPosition = 1280 / 2
	playerWidth, PlayerHeight = 64, 64
}

func playerMovement() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && playerPosition > 0+playerWidth/2 {
		playerPosition -= 10
		playerState = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && playerPosition < 1280-playerWidth/2 {
		playerPosition += 10
		playerState = 1
	} else {
		playerState = 0
	}
}

func playerAnimation() {
	switch playerState {
	case -1: // Left
		playerSheetX, playerSheetY = 0, 0
	case 0: // Centre
		playerSheetX, playerSheetY = 64, 0
	case 1: // Right
		playerSheetX, playerSheetY = 128, 0
	default: // Fallback to Centre
		playerSheetX, playerSheetY = 64, 0
	}
}
