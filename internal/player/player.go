package player

type Player struct {
	x, y   int
	health int
	speed  float64
}

func (p *Player) CreatePlayer(x, y int) {
	p.x = x
	p.y = y
	p.health = 10
	p.speed = 1.0
}
