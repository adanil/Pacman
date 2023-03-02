package states

import (
	"context"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image"
	"pacman/internal/base"
	"pacman/internal/controllers"
	"pacman/internal/level"
	"pacman/internal/utility"
	"strconv"
	"time"
)

const fontSize = 16
const enemyRandomControlPercentage = 40
const textOnMapFontSize = 8

var (
	frameNumber        int
	gameLevel          level.Level
	keyboardController controllers.KeyboardHandler
	enemyController    controllers.MixedEnemyController
)

type PlayState struct {
	g      *base.Game
	font   font.Face
	ctx    context.Context
	cancel context.CancelFunc
}

func NewPlayState(g *base.Game) PlayState {
	defaultFont, _ := utility.GetFont(base.PacmanFont, fontSize, base.DefaultDPI)
	p := PlayState{g: g, font: defaultFont}
	p.initGame()
	return p
}

func (p *PlayState) Update() error {
	if gameLevel.IsAllFoodEaten() {
		p.cancel()
		p.g.SetState(NewWinState(p.g, gameLevel.Score()))
		return nil
	}
	commands := keyboardController.GetKeyboardCommands()
	commands = append(commands, enemyController.GetCommands()...)
	for _, com := range commands {
		com.Execute()
	}
	ok := gameLevel.UpdateAll()
	if !ok {
		p.cancel()
		p.g.SetState(NewGameOverState(p.g, gameLevel.Score()))
	}
	return nil
}

func (p *PlayState) Draw(screen *ebiten.Image) {
	frameNumber = (frameNumber + 1) % base.FrameModulo
	p.drawMap(screen)
	for x := 0; x < base.WidthTiles; x++ {
		for y := 0; y < base.HeightTiles; y++ {
			entity := gameLevel.GetEntityByCoordinates(level.NewCoordinate(x, y))
			switch entity {
			case level.Food:
				p.drawEntityAtCenter(screen, base.Images["food"], x, y)
			case level.Strawberry:
				p.drawEntityAtCenter(screen, base.Images["strawberry"], x, y)
			case level.NightModeBooster:
				p.drawEntityAtCenter(screen, base.Images["booster"], x, y)
			}
		}
	}
	p.drawMapText(screen)
	p.drawPacman(screen)
	p.drawEnemies(screen)
	p.drawTitle(screen)
	p.drawScore(screen)
}

func (p *PlayState) drawEntityAtCenter(screen, entityImage *ebiten.Image, x int, y int) {
	op := &ebiten.DrawImageOptions{}
	foodWidth, foodHeight := entityImage.Size()
	op.GeoM.Translate(float64(x*base.TileSize-foodWidth/2)+base.TileSize/2, float64(y*base.TileSize-foodHeight/2)+base.TileSize/2)
	screen.DrawImage(entityImage, op)
}

//nolint:gomnd
func (p *PlayState) drawTitle(screen *ebiten.Image) {
	x := (base.GameScreenWidth - len(base.Title)*fontSize) / 2
	text.Draw(screen, base.Title, p.font, x, 25, base.PacmanColor)
}

//nolint:gomnd
func (p *PlayState) drawScore(screen *ebiten.Image) {
	text.Draw(screen, "score: "+strconv.Itoa(gameLevel.Score()), p.font, 35, base.GameScreenHeight-14, base.PacmanColor)
}

func (p *PlayState) drawMap(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(base.Images["map"], op)
}

//nolint:gomnd
func (p *PlayState) drawEnemies(screen *ebiten.Image) {
	const ghostStatesNumber = 2
	state := frameNumber / (base.FrameModulo / 2)
	for _, enemy := range gameLevel.Enemies() {
		op := &ebiten.DrawImageOptions{}
		px, py := enemy.GetCoords()
		op.GeoM.Translate(float64(px)+5, float64(py)+5)
		var enemyImage *ebiten.Image
		if enemy.NightMode() {
			state %= ghostStatesNumber
			enemyImage = enemy.GetGraphic().SubImage(image.Rect(6+40*(state), 0, 6+40*(state)+30, 38)).(*ebiten.Image)
		} else {
			enemyImage = enemy.GetGraphic().SubImage(image.Rect(6+40*(enemy.GetDirection()*2+state), 0, 6+40*(enemy.GetDirection()*2+state)+30, 38)).(*ebiten.Image)
		}
		screen.DrawImage(enemyImage, op)
	}
}

//nolint:gomnd
func (p *PlayState) drawPacman(screen *ebiten.Image) {
	state := frameNumber / (base.FrameModulo / 3)
	pacman := gameLevel.Player()
	op := &ebiten.DrawImageOptions{}
	px, py := pacman.GetCoords()
	op.GeoM.Translate(float64(px+5), float64(py+5))
	pacmanImage := (*ebiten.Image)(nil)
	if state == 2 {
		pacmanImage = pacman.GetGraphic().SubImage(image.Rect(296, 0, 324, 30)).(*ebiten.Image)
	} else {
		state = (state + 1) % 2
		pacmanImage = pacman.GetGraphic().SubImage(image.Rect(38*(pacman.GetDirection()*2+state), 0, 38*(pacman.GetDirection()*2+state)+30, 30)).(*ebiten.Image)
	}
	screen.DrawImage(pacmanImage, op)
}

//nolint:gomnd
func (p *PlayState) drawMapText(screen *ebiten.Image) {
	textFont, _ := utility.GetFont(base.PacmanFont, textOnMapFontSize, base.DefaultDPI)
	for _, screenText := range gameLevel.Texts() {
		if screenText.ExpiredTime().Before(time.Now()) {
			continue
		}
		text.Draw(screen, screenText.Text(), textFont, screenText.X()*base.TileSize+5, screenText.Y()*(base.TileSize)+base.TileSize/2+5, base.PacmanColor)
	}
}

func (p *PlayState) initGame() {
	levelCreator := &level.ReadLevel{Filepath: "maps/base"}
	initLevel(levelCreator)

	p.ctx, p.cancel = context.WithCancel(context.Background())
	keyboardController = controllers.NewKeyboardHandler(&gameLevel)
	go keyboardController.HandlePressedButtons(p.ctx)
	enemyController = controllers.NewMixedEnemyController(&gameLevel, enemyRandomControlPercentage)
}

func initLevel(creator level.Creator) {
	gameLevel = creator.CreateLevel(base.WidthTiles, base.HeightTiles, base.TileSize)
	gameLevel.CreateEntities()
}
