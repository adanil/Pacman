package base

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	Update() error
	Draw(screen *ebiten.Image)
}
