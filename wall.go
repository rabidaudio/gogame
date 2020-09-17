package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type Wall struct {
	bounds image.Rectangle
	color  color.Color
}

func (w *Wall) Tick(g *Game) error {
	return nil // walls don't do anything
}

func (w *Wall) Draw(screen *ebiten.Image) {
	i, _ := ebiten.NewImage(w.bounds.Dx(), w.bounds.Dy(), ebiten.FilterDefault)
	if w.color == nil {
		panic("Wtf")
	}
	i.Fill(w.color)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(w.bounds.Min.X), float64(w.bounds.Min.Y))
	screen.DrawImage(i, op)
}

// func (w *Wall) Bounds() image.Rectangle {
// 	return w.bounds
// }
