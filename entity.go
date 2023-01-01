package main

import (
	"github.com/hajimehoshi/ebiten"
)

// http://gameprogrammingpatterns.com/dirty-flag.html
type Entity interface {
	Img() *ebiten.Image
	Transform(op *ebiten.DrawImageOptions)
}

func DrawEntity(screen *ebiten.Image, ent Entity) {
	op := &ebiten.DrawImageOptions{}
	ent.Transform(op)
	screen.DrawImage(ent.Img(), op)
}

type Ticker interface {
	Tick(g *Game)
}
