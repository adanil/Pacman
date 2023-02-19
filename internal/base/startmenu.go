package base

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const menuFontSize = 32
const URLFontSize = 16

type StartState struct {
	g *Game
}

func NewStartState(g *Game) StartState {
	return StartState{g: g}
}

func (s StartState) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyW) || inpututil.IsKeyJustPressed(ebiten.KeyA) || inpututil.IsKeyJustPressed(ebiten.KeyS) ||
		inpututil.IsKeyJustPressed(ebiten.KeyD) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.g.SetState(NewPlayState(s.g))
	}
	return nil
}

// TODO Add background
func (s StartState) Draw(screen *ebiten.Image) {
	menuFont, _ := GetFont(baseFont, menuFontSize, defaultDPI)
	titleText := "PACMAN"
	x := (gameScreenWidth - len(titleText)*menuFontSize) / 2
	text.Draw(screen, titleText, menuFont, x, gameScreenHeight/2-120, yellowColor)

	howToStartTexts := []string{"PRESS SPACE BUTTON", "OR W/A/S/D BUTTON", "OR TOUCH SCREEN"}
	for ind, t := range howToStartTexts {
		x = (gameScreenWidth - len(t)*menuFontSize) / 2
		text.Draw(screen, t, menuFont, x, gameScreenHeight/2+50*ind, yellowColor)
	}

	githubFont, _ := GetFont(baseFont, URLFontSize, defaultDPI)
	githubURL := "github.com/adanil/Pacman"
	x = (gameScreenWidth - len(githubURL)*URLFontSize) / 2
	text.Draw(screen, githubURL, githubFont, x, gameScreenHeight-20, yellowColor)
}
