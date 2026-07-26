package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (p *Player) move() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		p.state = centre
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && p.x > 0+p.width/2 && scene == gameWorld {
		p.x -= p.vel
		p.state = left
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && p.x < float64(windowWidth)-p.width/2 && scene == gameWorld {
		p.x += p.vel
		p.state = right
	} else {
		p.state = centre
	}
}

func (g *Game) start() bool {
	if scene != gameWorld && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true
	}
	return false
}
