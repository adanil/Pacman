package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/jpeg"
	_ "image/png"
	"log"
	bs "pacman/internal/base"
	"pacman/internal/base/states"
)

func main() {
	ebiten.SetWindowSize(bs.GameScreenWidth, bs.GameScreenHeight)
	ebiten.SetWindowTitle("Pacman")
	game := &bs.Game{}
	game.SetState(states.NewStartState(game))
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
