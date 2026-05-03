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
