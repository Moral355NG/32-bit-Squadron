package main

import "github.com/hajimehoshi/ebiten/v2"

var (
	playerSheetX    int
	playerSheetY    int
	playerAnimation int
)

func playerMovement() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && playerPosition > 0 {
		playerPosition -= 10
		playerAnimation = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && playerPosition < 1280-64 {
		playerPosition += 10
		playerAnimation = 1
	} else {
		playerAnimation = 0
	}
}

func playerAnim() {
	switch playerAnimation {
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
