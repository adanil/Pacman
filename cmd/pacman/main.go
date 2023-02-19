package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/jpeg"
	_ "image/png"
	"log"
	bs "pacman/internal/base"
)

/* TODO
*	Add boosters
*	Add pathfind algorithm for enemies
* 	Refactor code
*	Check code by linter
*	Optional: create AI generator for map
*	Optional: create AI controller for enemies
 */

func main() {
	ebiten.SetWindowSize(928, 704) //TODO Refactor constants
	ebiten.SetWindowTitle("Pacman")
	game := &bs.Game{}
	game.SetState(bs.NewStartState(game))
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
