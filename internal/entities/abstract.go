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

type Movable interface {
	Move(direction, widthModulo, heightModulo int)
	ChangeDirection(direction int)
	SetCoords(x, y int)
	GetSpeed() int
	GetCoords() (int, int)
	GetDirection() int
}

type Graphical interface {
	GetGraphic() *ebiten.Image
}

type Playable interface {
	Movable
	Graphical
}
