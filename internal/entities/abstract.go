package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
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
	ChangeDirection(direction int)
	SetCoords(x, y int)
	GetSpeed() int
	GetCoords() (int, int)
	GetStartCoords() (int, int)
	GetDirection() int
	SetStopped(stop bool)
	GetStopped() bool
	NightMode() bool
	SetNightMode(nightMode bool)
	NightModeExpiredTime() time.Time
	SetNightModeExpiredTime(nightModeExpiredTime time.Time)
}

type Graphical interface {
	GetGraphic() *ebiten.Image
	SetGraphic(*ebiten.Image)
}

type Playable interface {
	Movable
	Graphical
}
