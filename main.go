package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	Tick int
	// Entities   []Entity
	leftWall   Wall
	rightWall  Wall
	topWall    Wall
	bottomWall Wall
	dude       Dude
}

func NewGame() Game {
	g := Game{}
	// g.Entities = make([]Entity, 10)
	g.bottomWall = Wall{bounds: image.Rect(0, 0, 100, 10), color: HexColor(0x25ae4e)}
	g.leftWall = Wall{bounds: image.Rect(0, 0, 10, 100), color: HexColor(0x25ae4e)}
	g.rightWall = Wall{bounds: image.Rect(90, 0, 100, 100), color: HexColor(0x25ae4e)}
	g.topWall = Wall{bounds: image.Rect(0, 90, 100, 100), color: HexColor(0x25ae4e)}
	// g.Entities = append(g.Entities, &g.bottomWall, &g.leftWall, &g.rightWall, &g.bottomWall)
	g.dude = Dude{pos: image.Point{25, 25}, speed: Vector{0.1, 0.1}, size: 20, color: color.White}
	// g.Entities = append(g.Entities, &g.dude)
	return g
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.Tick += 1
	g.dude.Tick(g)
	// for _, e := range g.Entities {
	// 	if err := e.Tick(g); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, e := range []Entity{&g.bottomWall, &g.topWall, &g.leftWall, &g.rightWall, &g.dude} {
		e.Draw(screen)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	g := NewGame()

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
