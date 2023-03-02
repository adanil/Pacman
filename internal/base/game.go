package base

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"os"
	"pacman/internal/utility"
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
	Images      map[string]*ebiten.Image
)

func init() {
	if err := initFont(); err != nil {
		log.Fatal(err)
	}
	if err := initImages(); err != nil {
		log.Fatal(err)
	}
	//nolint:gomnd
	PacmanColor = color.RGBA{R: 230, G: 245, B: 39, A: 255}
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return GameScreenWidth, GameScreenHeight
}

func initFont() error {
	fontReader, err := os.ReadFile("fonts/quinquefive/QuinqueFive.ttf")
	if err != nil {
		log.Fatal(err)
	}
	PacmanFont, err = opentype.Parse(fontReader)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func initImages() error {
	Images = make(map[string]*ebiten.Image)
	imageMap, err := utility.ReadImage("images/map.png")
	if err != nil {
		return err
	}
	Images["map"] = imageMap

	foodImage, err := utility.ReadImage("images/food_10px.png")
	if err != nil {
		return err
	}
	Images["food"] = foodImage

	strawberryImage, _ := utility.ReadImage("images/strawberry.png")
	if err != nil {
		return err
	}
	Images["strawberry"] = strawberryImage

	boosterImage, _ := utility.ReadImage("images/booster.png")
	if err != nil {
		return err
	}
	Images["booster"] = boosterImage

	pacmanImage, _ := utility.ReadImage("images/pacman.png")
	if err != nil {
		return err
	}
	Images["pacman"] = pacmanImage

	blueEnemyImage, _ := utility.ReadImage("images/blueEnemy.png")
	if err != nil {
		return err
	}
	Images["blueEnemy"] = blueEnemyImage

	pinkEnemyImage, _ := utility.ReadImage("images/pinkEnemy.png")
	if err != nil {
		return err
	}
	Images["pinkEnemy"] = pinkEnemyImage

	redEnemyImage, _ := utility.ReadImage("images/redEnemy.png")
	if err != nil {
		return err
	}
	Images["redEnemy"] = redEnemyImage

	yellowEnemyImage, _ := utility.ReadImage("images/yellowEnemy.png")
	if err != nil {
		return err
	}
	Images["yellowEnemy"] = yellowEnemyImage

	ghostImage, _ := utility.ReadImage("images/ghost.png")
	if err != nil {
		return err
	}
	Images["ghost"] = ghostImage

	return nil
}
