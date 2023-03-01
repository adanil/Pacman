package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

var OppositeDirection = map[int]int{
	UP:    DOWN,
	DOWN:  UP,
	LEFT:  RIGHT,
	RIGHT: LEFT,
}

type Movable interface {
	Move(direction, widthModulo, heightModulo int)
	CalculateNextPosition(direction, widthModulo, heightModulo int) (int, int)
	ChangeDirection(direction int)
	GetDirection() int
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
