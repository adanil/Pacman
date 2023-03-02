package states

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"pacman/internal/base"
	"pacman/internal/utility"
)

const menuFontSize = 32
const URLFontSize = 16

type StartState struct {
	g *base.Game
}

func NewStartState(g *base.Game) StartState {
	return StartState{g: g}
}

func (s StartState) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyS) ||
		inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		playState := NewPlayState(s.g)
		s.g.SetState(&playState)
	}
	return nil
}

// TODO Add background
//
//nolint:gomnd
func (s StartState) Draw(screen *ebiten.Image) {
	menuFont, _ := utility.GetFont(base.PacmanFont, menuFontSize, base.DefaultDPI)
	titleText := "PACMAN"
	x := (base.GameScreenWidth - len(titleText)*menuFontSize) / 2
	text.Draw(screen, titleText, menuFont, x, base.GameScreenHeight/2-120, base.PacmanColor)

	howToStartTexts := []string{"PRESS SPACE BUTTON", "OR W/A/S/D BUTTON", "OR TOUCH SCREEN"}
	for ind, t := range howToStartTexts {
		x = (base.GameScreenWidth - len(t)*menuFontSize) / 2
		text.Draw(screen, t, menuFont, x, base.GameScreenHeight/2+50*ind, base.PacmanColor)
	}

	githubFont, _ := utility.GetFont(base.PacmanFont, URLFontSize, base.DefaultDPI)
	githubURL := "github.com/adanil/Pacman"
	x = (base.GameScreenWidth - len(githubURL)*URLFontSize) / 2
	text.Draw(screen, githubURL, githubFont, x, base.GameScreenHeight-20, base.PacmanColor)
}
