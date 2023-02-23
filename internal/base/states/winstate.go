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

const winFontSize = 32

type WinState struct {
	g     *base.Game
	score int
	font  font.Face
}

func NewWinState(g *base.Game, score int) WinState {
	defaultFont, _ := utility.GetFont(base.PacmanFont, winFontSize, base.DefaultDPI)
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
		x := (base.GameScreenWidth - len(t)*winFontSize) / 2
		text.Draw(screen, t, w.font, x, base.GameScreenHeight/2-30+50*ind, base.PacmanColor)
	}

	restartText := "PRESS SPACE TO RESTART"
	x := (base.GameScreenWidth - len(restartText)*winFontSize) / 2
	text.Draw(screen, restartText, w.font, x, base.GameScreenHeight/2+150, base.PacmanColor)
}
