package base

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"strconv"
)

const winFontSize = 32

type WinState struct {
	g     *Game
	score int
	font  font.Face
}

func NewWinState(g *Game, score int) WinState {
	defaultFont, _ := GetFont(baseFont, winFontSize, defaultDPI)
	return WinState{g: g, score: score, font: defaultFont}
}

func (w WinState) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		w.g.SetState(NewPlayState(w.g))
	}
	return nil
}

func (w WinState) Draw(screen *ebiten.Image) {
	texts := []string{"SCORE:" + strconv.Itoa(w.score), "YOU WIN!"}
	for ind, t := range texts {
		x := (gameScreenWidth - len(t)*winFontSize) / 2
		text.Draw(screen, t, w.font, x, gameScreenHeight/2-30+50*ind, yellowColor)
	}

	restartText := "PRESS SPACE TO RESTART"
	x := (gameScreenWidth - len(restartText)*winFontSize) / 2
	text.Draw(screen, restartText, w.font, x, gameScreenHeight/2+150, yellowColor)
}
