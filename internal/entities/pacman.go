package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Pacman struct {
	Health   int
	x, y     int
	speed    int
	rotation int
	graphic  *ebiten.Image
	stopped  bool
}

func CreatePlayer(x, y, tileSize int, graphic *ebiten.Image) Pacman {
	p := Pacman{}
	p.x = x * tileSize
	p.y = y * tileSize
	p.rotation = RIGHT
	p.Health = 1
	p.speed = 2
	p.graphic = graphic
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
