package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/rabidaudio/gogame/b2"
)

var GAME Game

type Game struct {
	world  b2.World
	ground b2.Entity
	dude   b2.Entity
	// Entities   []Entity
	// leftWall   Wall
	// rightWall  Wall
	// topWall    Wall
	// bottomWall Wall
	// dude       Dude
}

// func createGround(g *Game) {
// 	gnd := g.world.CreateBody(Ground)
// 	shp := box2d.MakeB2PolygonShape()
// 	shp.SetAsBox(50, 10)
// 	gnd.CreateFixture(&shp, 0)
// }

func init() {
	GAME = Game{}
	// 0, -10
	GAME.world = b2.NewWorld(b2.Vector{0, -10})
	GAME.ground = GAME.world.NewStaticEntity(b2.StaticEntityConf{
		Pos:   b2.Vector{0, 0},
		Shape: b2.BoxShape{Width: 50, Height: 10},
	})
	GAME.dude = GAME.world.NewDynamicEntity(b2.DynamicEntityConf{
		Pos:     b2.Vector{0, 4},
		Shape:   b2.BoxShape{Width: 1, Height: 1},
		Density: 1,
	}, nil)
	// createGround(&g)

	// g.Entities = make([]Entity, 10)
	// g.bottomWall = Wall{bounds: image.Rect(0, 0, 100, 10), color: HexColor(0x25ae4e)}
	// g.leftWall = Wall{bounds: image.Rect(0, 0, 10, 100), color: HexColor(0x25ae4e)}
	// g.rightWall = Wall{bounds: image.Rect(90, 0, 100, 100), color: HexColor(0x25ae4e)}
	// g.topWall = Wall{bounds: image.Rect(0, 90, 100, 100), color: HexColor(0x25ae4e)}
	// // g.Entities = append(g.Entities, &g.bottomWall, &g.leftWall, &g.rightWall, &g.bottomWall)
	// g.dude = NewDude(&g)
	// // g.Entities = append(g.Entities, &g.dude)
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.world.Step(1.0 / float64(ebiten.MaxTPS()))
	// for _, e := range g.Entities {
	// 	if v, ok := e.(Ticker); ok {
	// 		v.Tick(g)
	// 	}
	// }
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// for _, e := range []Entity{&g.bottomWall, &g.topWall, &g.leftWall, &g.rightWall, &g.dude} {
	// 	DrawEntity(screen, e)
	// }
	b := g.world.Raw().M_bodyList.M_fixtureCount
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%v", b))
	// x := g.dude.body.GetPosition().X
	// y := g.dude.body.GetPosition().Y
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("%0.2f %0.2f", x, y))
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("%0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&GAME); err != nil {
		log.Fatal(err)
	}
}
