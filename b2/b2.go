package b2

import (
	"github.com/ByteArena/box2d"
)

type Vector box2d.B2Vec2
type World struct {
	b2w box2d.B2World
}

type Entity struct {
	body *box2d.B2Body
}

func NewWorld(gravity Vector) World {
	return World{b2w: box2d.MakeB2World(box2d.B2Vec2(gravity))}
}

func (w World) Raw() box2d.B2World {
	return w.b2w
}

func (w World) Step(t float64) {
	w.b2w.Step(t, 8, 3)
}

type StaticEntityConf struct {
	Pos   Vector
	Shape Shape
}

type Shape interface {
	shape() box2d.B2ShapeInterface
}

type BoxShape struct {
	Width, Height float64
}

func (b BoxShape) shape() box2d.B2ShapeInterface {
	// s := box2d.MakeB2EdgeShape()
	// s.Set(box2d.MakeB2Vec2(-20.0, 0.0), box2d.MakeB2Vec2(20.0, 0.0))
	s := box2d.MakeB2PolygonShape()
	s.SetAsBox(b.Width/2, b.Height/2)
	return &s
}

type CircleShape struct {
	Center Vector
	Radius float64
}

func (b CircleShape) shape() box2d.B2ShapeInterface {
	s := box2d.MakeB2CircleShape()
	s.M_p.Set(b.Center.X, b.Center.Y)
	s.M_radius = b.Radius
	return s
}

type CustomShape struct {
	shape box2d.B2ShapeInterface
}

func (w *World) NewStaticEntity(c StaticEntityConf) Entity {
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Position.Set(c.Pos.X, c.Pos.Y)
	body := w.b2w.CreateBody(&bodyDef)
	body.CreateFixture(c.Shape.shape(), 0)
	return Entity{body: body}
}

type DynamicEntityConf struct {
	Pos     Vector
	Shape   Shape
	Density float64
}

func (w *World) NewDynamicEntity(c DynamicEntityConf, config func(def *box2d.B2FixtureDef)) Entity {
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Position.Set(c.Pos.X, c.Pos.Y)
	bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	body := w.b2w.CreateBody(&bodyDef)
	fd := box2d.MakeB2FixtureDef()
	fd.Shape = c.Shape.shape()
	fd.Density = c.Density
	if config != nil {
		config(&fd)
	}
	body.CreateFixtureFromDef(&fd)
	return Entity{body: body}
}
