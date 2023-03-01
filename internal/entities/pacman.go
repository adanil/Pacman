package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Pacman struct {
	Health               int
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

func (p *Pacman) NightModeExpiredTime() time.Time {
	return p.nightModeExpiredTime
}

func (p *Pacman) SetNightModeExpiredTime(nightModeExpiredTime time.Time) {
	p.nightModeExpiredTime = nightModeExpiredTime
}

func (p *Pacman) NightMode() bool {
	return p.nightMode
}

func (p *Pacman) SetNightMode(nightMode bool) {
	if nightMode {
		p.SetGraphic(p.ghostGraphic)
		p.SetNightModeExpiredTime(time.Now().Add(5 * time.Second))
	} else {
		p.SetGraphic(p.baseGraphic)
	}
	p.nightMode = nightMode
}

func CreatePlayer(x, y, tileSize int, baseGraphic, ghostGraphic *ebiten.Image) Pacman {
	p := Pacman{}
	p.x = x * tileSize
	p.y = y * tileSize
	p.startX = p.x
	p.startY = p.y
	p.rotation = RIGHT
	p.Health = 1
	p.speed = 2
	p.graphic = baseGraphic
	p.baseGraphic = baseGraphic
	p.ghostGraphic = ghostGraphic
	return p
}

func (p *Pacman) SetStopped(stop bool) {
	p.stopped = stop
}

func (p *Pacman) GetStopped() bool {
	return p.stopped
}

func (p *Pacman) SetCoords(x, y int) {
	p.x = x
	p.y = y
}
func (p *Pacman) GetCoords() (int, int) {
	return p.x, p.y
}

func (p *Pacman) GetStartCoords() (int, int) {
	return p.startX, p.startY
}

func (p *Pacman) GetSpeed() int {
	return p.speed
}

func (p *Pacman) Move(direction, widthModulo, heightModulo int) {
	switch direction {
	case UP:
		p.MoveUp()
	case DOWN:
		p.MoveDown()
	case LEFT:
		p.MoveLeft()
	case RIGHT:
		p.MoveRight()
	}
	p.x += widthModulo
	p.y += heightModulo
	p.x %= widthModulo
	p.y %= heightModulo
}

func (p *Pacman) ChangeDirection(direction int) {
	p.rotation = direction
}

func (p *Pacman) GetDirection() int {
	return p.rotation
}

func (p *Pacman) SetGraphic(g *ebiten.Image) {
	p.graphic = g
}

func (p *Pacman) GetGraphic() *ebiten.Image {
	return p.graphic
}

func (p *Pacman) MoveDown() {
	p.y += p.speed
}

func (p *Pacman) MoveUp() {
	p.y -= p.speed
}

func (p *Pacman) MoveRight() {
	p.x += p.speed
}

func (p *Pacman) MoveLeft() {
	p.x -= p.speed
}
