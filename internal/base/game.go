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
	PacmanColor = color.RGBA{R: 230, G: 245, B: 39, A: 255}
}

func initFont() error {
	fontReader, err := os.ReadFile("fonts/quinquefive/QuinqueFive.ttf")
	PacmanFont, err = opentype.Parse(fontReader)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func initImages() error {
	Images = make(map[string]*ebiten.Image)
	imageMap, err := utility.ReadImage("images/map_v3.png")
	if err != nil {
		return err
	}
	Images["map"] = imageMap

	foodImage, err := utility.ReadImage("images/pacman-pack_v2/food_10px.png")
	if err != nil {
		return err
	}
	Images["food"] = foodImage

	strawberryImage, _ := utility.ReadImage("images/pacman-pack_v2/Strawberry_25px.png")
	if err != nil {
		return err
	}
	Images["strawberry"] = strawberryImage

	boosterImage, _ := utility.ReadImage("images/pacman-pack_v2/NightModeBooster_20px.png")
	if err != nil {
		return err
	}
	Images["booster"] = boosterImage

	pacmanImage, _ := utility.ReadImage("images/pacman-pack_v2/Pacmanx2.png")
	if err != nil {
		return err
	}
	Images["pacman"] = pacmanImage

	blueEnemyImage, _ := utility.ReadImage("images/pacman-pack_v2/BlueEnemyx2.png")
	if err != nil {
		return err
	}
	Images["blueEnemy"] = blueEnemyImage

	pinkEnemyImage, _ := utility.ReadImage("images/pacman-pack_v2/PinkEnemyx2.png")
	if err != nil {
		return err
	}
	Images["pinkEnemy"] = pinkEnemyImage

	redEnemyImage, _ := utility.ReadImage("images/pacman-pack_v2/RedEnemyx2.png")
	if err != nil {
		return err
	}
	Images["redEnemy"] = redEnemyImage

	yellowEnemyImage, _ := utility.ReadImage("images/pacman-pack_v2/YellowEnemyv3.png")
	if err != nil {
		return err
	}
	Images["yellowEnemy"] = yellowEnemyImage

	ghostImage, _ := utility.ReadImage("images/pacman-pack_v2/Ghostv1.png")
	if err != nil {
		return err
	}
	Images["ghost"] = ghostImage

	return nil
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
