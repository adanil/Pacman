package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"pacman/internal/base"
	"time"
)

type Enemy struct {
	health               int
	x, y                 int
	startX, startY       int
	speed                int
	rotation             int
	graphic              *ebiten.Image
	baseGraphic          *ebiten.Image
	ghostGraphic         *ebiten.Image
	stopped              bool
	nightMode            bool
	nightModeExpiredTime time.Time
}

func CreateEnemy(x, y int, baseGraphic *ebiten.Image) Enemy {
	return Enemy{
		health:               1,
		x:                    x,
		y:                    y,
		startX:               x,
		startY:               y,
		speed:                2,
		rotation:             UP,
		graphic:              baseGraphic,
		baseGraphic:          baseGraphic,
		ghostGraphic:         base.Images["ghost"],
		stopped:              false,
		nightMode:            false,
		nightModeExpiredTime: time.Time{},
	}
}

func (e *Enemy) SetCoords(x, y int) {
	e.x = x
	e.y = y
}
func (e *Enemy) GetCoords() (int, int) {
	return e.x, e.y
}

func (e *Enemy) GetStartCoords() (int, int) {
	return e.startX, e.startY
}

func (e *Enemy) GetSpeed() int {
	return e.speed
}

func (e *Enemy) Move(direction, widthModulo, heightModulo int) {
	newX, newY := e.CalculateNextPosition(direction, widthModulo, heightModulo)
	e.SetCoords(newX, newY)
}

func (e *Enemy) CalculateNextPosition(direction, widthModulo, heightModulo int) (int, int) {
	x := e.x
	y := e.y
	switch direction {
	case UP:
		y -= e.speed
	case DOWN:
		y += e.speed
	case LEFT:
		x -= e.speed
	case RIGHT:
		x += e.speed
	}
	x += widthModulo
	y += heightModulo
	x %= widthModulo
	y %= heightModulo
	return x, y
}

func (e *Enemy) ChangeDirection(direction int) {
	e.rotation = direction
}

func (e *Enemy) GetDirection() int {
	return e.rotation
}

func (e *Enemy) SetGraphic(g *ebiten.Image) {
	e.graphic = g
}

func (e *Enemy) GetGraphic() *ebiten.Image {
	return e.graphic
}

func (e *Enemy) SetStopped(stop bool) {
	e.stopped = stop
}

func (e *Enemy) GetStopped() bool {
	return e.stopped
}

func (e *Enemy) NightModeExpiredTime() time.Time {
	return e.nightModeExpiredTime
}

func (e *Enemy) SetNightModeExpiredTime(nightModeExpiredTime time.Time) {
	e.nightModeExpiredTime = nightModeExpiredTime
}

func (e *Enemy) NightMode() bool {
	return e.nightMode
}

func (e *Enemy) SetNightMode(nightMode bool) {
	if nightMode {
		e.SetGraphic(e.ghostGraphic)
		e.SetNightModeExpiredTime(time.Now().Add(5 * time.Second))
	} else {
		e.SetGraphic(e.baseGraphic)
	}
	e.nightMode = nightMode
}
