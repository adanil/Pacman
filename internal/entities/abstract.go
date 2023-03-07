package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)
const defaultSpeed = 2
const ghostModeDuration = 5

var OppositeDirection = map[Direction]Direction{
	UP:    DOWN,
	DOWN:  UP,
	LEFT:  RIGHT,
	RIGHT: LEFT,
}

type Movable interface {
	Move(direction Direction, widthModulo, heightModulo int)
	CalculateNextPosition(direction Direction, widthModulo, heightModulo int) (int, int)
	ChangeDirection(direction Direction)
	GetDirection() Direction
	SetCoords(x, y int)
	GetCoords() (int, int)
	GetSpeed() int
}

type Graphical interface {
	GetGraphic() *ebiten.Image
	SetGraphic(*ebiten.Image)
}

type Entity interface {
	Movable
	Graphical
}
