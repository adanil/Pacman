package base

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"strconv"
)

const gameOverFontSize = 32

type GameOverState struct {
	g     *Game
	score int
	font  font.Face
}

func NewGameOverState(g *Game, score int) GameOverState {
	defaultFont, _ := GetFont(baseFont, gameOverFontSize, defaultDPI)
	return GameOverState{g: g, score: score, font: defaultFont}
}

func (g GameOverState) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.g.SetState(NewPlayState(g.g))
	}
	return nil
}

func (g GameOverState) Draw(screen *ebiten.Image) {
	texts := []string{"SCORE:" + strconv.Itoa(g.score), "GAME OVER"}
	for ind, t := range texts {
		x := (gameScreenWidth - len(t)*gameOverFontSize) / 2
		text.Draw(screen, t, g.font, x, gameScreenHeight/2-30+50*ind, yellowColor)
	}
	restartText := "PRESS SPACE TO RESTART"
	x := (gameScreenWidth - len(restartText)*gameOverFontSize) / 2
	text.Draw(screen, restartText, g.font, x, gameScreenHeight/2+150, yellowColor)
}
