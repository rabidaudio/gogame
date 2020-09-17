package main

import (
	"github.com/hajimehoshi/ebiten"
)

// http://gameprogrammingpatterns.com/dirty-flag.html
type Entity interface {
	Tick(g *Game) error
	Draw(screen *ebiten.Image)
	// Bounds() image.Rectangle
}
