//nolint:dupl
package states

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"pacman/internal/base"
	"pacman/internal/utility"
	"strconv"
)

const gameOverFontSize = 32

type GameOverState struct {
	g     *base.Game
	score int
	font  font.Face
}

func NewGameOverState(g *base.Game, score int) GameOverState {
	defaultFont, _ := utility.GetFont(base.PacmanFont, gameOverFontSize, base.DefaultDPI)
	return GameOverState{g: g, score: score, font: defaultFont}
}

func (g GameOverState) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		playState := NewPlayState(g.g)
		g.g.SetState(&playState)
	}
	return nil
}

//nolint:gomnd
func (g GameOverState) Draw(screen *ebiten.Image) {
	texts := []string{"SCORE:" + strconv.Itoa(g.score), "GAME OVER"}
	for ind, t := range texts {
		x := (base.GameScreenWidth - len(t)*gameOverFontSize) / 2
		text.Draw(screen, t, g.font, x, base.GameScreenHeight/2-30+50*ind, base.PacmanColor)
	}
	restartText := "PRESS SPACE TO RESTART"
	x := (base.GameScreenWidth - len(restartText)*gameOverFontSize) / 2
	text.Draw(screen, restartText, g.font, x, base.GameScreenHeight/2+150, base.PacmanColor)
}
