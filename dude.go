package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	dudeImg *ebiten.Image
)

func init() {
	// ctx := gg.NewContext(10, 10)
	// ctx.SetColor(color.RGBA{R: 0xFF})
	// ctx.DrawCircle(5, 5, 5)
	// dudeImg, _ = ebiten.NewImageFromImage(ctx.Image(), ebiten.FilterDefault)
	// dudeImg.Fill(color.RGBA{R: 0xFF})
	dudeImg, _ = ebitenutil.NewImageFromURL("https://raw.githubusercontent.com/hajimehoshi/ebiten/master/examples/resources/images/flappy/gopher.png")
}

type Vector struct{ X, Y float64 }

type Dude struct {
	pos   image.Point
	speed Vector
	size  int
	color color.Color
}

func (d *Dude) Tick(g *Game) error {
	// d.pos.X = float64(d.pos.X) +d.speed.X
	// d.pos.Y += d.speed.Y
	// b := d.Bounds()
	// if b.Left < g.Bounds.Left {
	// 	// collide left wall
	// } else if b.Right > g.Bounds.Right {
	// 	// collide right wall
	// } else if b.Top < g.Bounds.Top {
	// 	// collide top wall
	// } else if b.Bottom > g.Bounds.Bottom {
	// 	// collide bottom wall
	// }
	// if !g.bounds.ContainsRect(d.Bounds()) {
	// 	d.pos.x -= d.speed.x
	// 	d.pos.y -= d.speed.y
	// 	d.speed =
	// }
	return nil
}

// func (d Dude) Bounds() Rect {
// 	return Rect{d.pos.x, d.pos.x + d.size, d.pos.y, d.pos.y + d.size}
// }

func (d *Dude) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := dudeImg.Size()
	op.GeoM.Scale(float64(d.size)/float64(w), float64(d.size)/float64(h))
	op.GeoM.Translate(float64(d.pos.X), float64(d.pos.Y))
	screen.DrawImage(dudeImg, op)
	/*
		w, h := gopherImage.Size()
		op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
		op.GeoM.Rotate(float64(g.vy16) / 96.0 * math.Pi / 6)
		op.GeoM.Translate(float64(w)/2.0, float64(h)/2.0)
		op.GeoM.Translate(float64(g.x16/16.0)-float64(g.cameraX), float64(g.y16/16.0)-float64(g.cameraY))
		op.Filter = ebiten.FilterLinear
		screen.DrawImage(gopherImage, op)
	*/
}
