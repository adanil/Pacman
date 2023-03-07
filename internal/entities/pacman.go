package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Pacman struct {
	health         int
	x, y           int
	startX, startY int
	speed          int
	rotation       Direction
	graphic        *ebiten.Image
}

func CreatePacman(x, y int, baseGraphic *ebiten.Image) Pacman {
	return Pacman{
		health:   1,
		x:        x,
		y:        y,
		startX:   x,
		startY:   y,
		speed:    defaultSpeed,
		rotation: RIGHT,
		graphic:  baseGraphic,
	}
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

func (p *Pacman) Move(direction Direction, widthModulo, heightModulo int) {
	newX, newY := p.CalculateNextPosition(direction, widthModulo, heightModulo)
	p.SetCoords(newX, newY)
}

func (p *Pacman) CalculateNextPosition(direction Direction, widthModulo, heightModulo int) (int, int) {
	x := p.x
	y := p.y
	switch direction {
	case UP:
		y -= p.speed
	case DOWN:
		y += p.speed
	case LEFT:
		x -= p.speed
	case RIGHT:
		x += p.speed
	}
	x += widthModulo
	y += heightModulo
	x %= widthModulo
	y %= heightModulo
	return x, y
}

func (p *Pacman) ChangeDirection(direction Direction) {
	p.rotation = direction
}

func (p *Pacman) GetDirection() Direction {
	return p.rotation
}

func (p *Pacman) SetGraphic(g *ebiten.Image) {
	p.graphic = g
}

func (p *Pacman) GetGraphic() *ebiten.Image {
	return p.graphic
}

func (p *Pacman) DecreaseHealth() {
	p.health--
}
