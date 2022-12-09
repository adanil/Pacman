package player

import (
	"log"
	"math/rand"
	"pacman/internal/level"
)

const (
	LEFT = iota
	RIGHT
	UP
	DOWN
)

type Player struct {
	X, Y     int
	health   int
	speed    int
	Rotation int
}

func CreatePlayer(lv level.Level, tileSize int) Player {
	p := Player{}
	for {
		x := rand.Intn(len(lv.LevelTiles))
		y := rand.Intn(len(lv.LevelTiles[0]))
		log.Println("CreatePlayer X: ", x, " Y: ", y)
		if lv.LevelTiles[x][y] == level.Free {
			p.X = x * tileSize
			p.Y = y * tileSize
			break
		}
	}
	p.Rotation = RIGHT
	p.health = 10
	p.speed = 1
	return p
}

func (p *Player) ChangeDirection(direction int) {
	p.Rotation = direction
}

func (p *Player) SetCoords(x, y int) {
	p.X = x
	p.Y = y
}

func (p *Player) Move(rotation int) {
	switch rotation {
	case UP:
		p.MoveUp()
	case DOWN:
		p.MoveDown()
	case LEFT:
		p.MoveLeft()
	case RIGHT:
		p.MoveRight()
	}
}

func (p *Player) MoveDown() {
	p.Y += p.speed
}

func (p *Player) MoveUp() {
	p.Y -= p.speed
}

func (p *Player) MoveRight() {
	p.X += p.speed
}

func (p *Player) MoveLeft() {
	p.X -= p.speed
}
