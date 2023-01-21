package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Pacman struct {
	Score    int
	Name     string
	Health   int
	X, Y     int
	Speed    int
	Rotation int
	Graphic  *ebiten.Image
	stopped  bool
}

func CreatePlayer(x, y, tileSize int) Pacman {
	p := Pacman{}
	p.X = x * tileSize
	p.Y = y * tileSize
	p.Rotation = RIGHT
	p.Health = 1
	p.Speed = 2
	return p
}

func (p *Pacman) SetStopped(stop bool) {
	p.stopped = stop
}

func (p *Pacman) GetStopped() bool {
	return p.stopped
}

func (p *Pacman) SetCoords(x, y int) {
	p.X = x
	p.Y = y
}
func (p *Pacman) GetCoords() (int, int) {
	return p.X, p.Y
}

func (p *Pacman) GetSpeed() int {
	return p.Speed
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
	p.X += widthModulo
	p.Y += heightModulo
	p.X %= widthModulo
	p.Y %= heightModulo
}

func (p *Pacman) ChangeDirection(direction int) {
	p.Rotation = direction
}

func (p *Pacman) GetDirection() int {
	return p.Rotation
}

func (p *Pacman) GetGraphic() *ebiten.Image {
	return p.Graphic
}

func (p *Pacman) MoveDown() {
	p.Y += p.Speed
}

func (p *Pacman) MoveUp() {
	p.Y -= p.Speed
}

func (p *Pacman) MoveRight() {
	p.X += p.Speed
}

func (p *Pacman) MoveLeft() {
	p.X -= p.Speed
}

//pacmanImage.SubImage(image.Rect(165/9*6, 0, 165/9*7, tileSize)).(*ebiten.Image)
