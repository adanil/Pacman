package base

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"os"
)

const (
	FrameModulo      = 150
	TileSize         = 40
	GameScreenWidth  = 880
	GameScreenHeight = 840
	WidthTiles       = GameScreenWidth / TileSize
	HeightTiles      = GameScreenHeight / TileSize
	DefaultDPI       = 64
	Title            = "Pacman"
)

var (
	PacmanFont  *opentype.Font
	PacmanColor color.RGBA
)

func init() {
	fontReader, err := os.ReadFile("fonts/quinquefive/QuinqueFive.ttf")
	PacmanFont, err = opentype.Parse(fontReader)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	PacmanColor = color.RGBA{
		R: 230,
		G: 245,
		B: 39,
		A: 255,
	}
}

type Game struct {
	currentState State
}

func (g *Game) SetState(state State) {
	g.currentState = state
}

func (g *Game) Update() error {
	return g.currentState.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentState.Draw(screen)
}

// TODO CHECK params
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GameScreenWidth, GameScreenHeight
}
