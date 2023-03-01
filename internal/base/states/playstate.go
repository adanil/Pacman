package states

import (
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

var (
	frameNumber        int
	gameLevel          level.Level
	keyboardController controllers.KeyboardHandler
	enemyController    controllers.MixedEnemyController
)

type PlayState struct {
	g    *base.Game
	font font.Face
}

func NewPlayState(g *base.Game) PlayState {
	defaultFont, _ := utility.GetFont(base.PacmanFont, fontSize, base.DefaultDPI)
	initGame() //TODO Maybe fix this
	return PlayState{g: g, font: defaultFont}
}

func initGame() {
	lvGenerator := level.Generator{Creator: &level.ReadLevel{Filepath: "maps/base"}}
	initLevel(lvGenerator)

	keyboardController = controllers.NewKeyboardHandler(&gameLevel)
	go keyboardController.HandlePressedButtons() //TODO Exit
	enemyController = controllers.NewMixedEnemyController(&gameLevel, 20)
}

func initLevel(lvGenerator level.Generator) {
	gameLevel = lvGenerator.CreateLevel(base.WidthTiles, base.HeightTiles, base.TileSize)
	gameLevel.CreateEntities()
}

func (p PlayState) Update() error {
	if gameLevel.IsAllFoodEaten() {
		p.g.SetState(NewWinState(p.g, gameLevel.Score))
		return nil
	}
	commands := keyboardController.GetKeyboardCommands()
	commands = append(commands, enemyController.GetCommands()...)
	for _, com := range commands {
		com.Execute()
	}
	ok := gameLevel.UpdateAll()
	if ok == false {
		p.g.SetState(NewGameOverState(p.g, gameLevel.Score))
	}
	return nil
}

func (p PlayState) Draw(screen *ebiten.Image) {
	frameNumber = (frameNumber + 1) % base.FrameModulo
	p.drawMap(screen)
	//TODO refactor duplicate code
	for x := 0; x < base.WidthTiles; x++ {
		for y := 0; y < base.HeightTiles; y++ {
			if gameLevel.LevelTiles[x][y] == level.Food {
				op := &ebiten.DrawImageOptions{}
				foodWidth, foodHeight := base.Images["food"].Size()
				op.GeoM.Translate(float64(x*base.TileSize-foodWidth/2)+base.TileSize/2, float64(y*base.TileSize-foodHeight/2)+base.TileSize/2)
				screen.DrawImage(base.Images["food"], op)
			} else if gameLevel.LevelTiles[x][y] == level.Strawberry {
				op := &ebiten.DrawImageOptions{}
				strawberryWidth, strawberryHeight := base.Images["strawberry"].Size()
				op.GeoM.Translate(float64(x*base.TileSize-strawberryWidth/2)+base.TileSize/2, float64(y*base.TileSize-strawberryHeight/2)+base.TileSize/2)
				screen.DrawImage(base.Images["strawberry"], op)
			} else if gameLevel.LevelTiles[x][y] == level.NightModeBooster {
				op := &ebiten.DrawImageOptions{}
				boosterWidth, boosterHeight := base.Images["booster"].Size()
				op.GeoM.Translate(float64(x*base.TileSize-boosterWidth/2)+base.TileSize/2, float64(y*base.TileSize-boosterHeight/2)+base.TileSize/2)
				screen.DrawImage(base.Images["booster"], op)
			}
		}
	}
	p.drawMapText(screen)
	p.drawPacman(screen)
	p.drawEnemies(screen)
	p.drawTitle(screen)
	p.drawScore(screen)
}

func (p PlayState) drawTitle(screen *ebiten.Image) {
	x := (base.GameScreenWidth - len(base.Title)*fontSize) / 2
	text.Draw(screen, base.Title, p.font, x, 25, base.PacmanColor)
}

func (p PlayState) drawScore(screen *ebiten.Image) {
	text.Draw(screen, "Score: "+strconv.Itoa(gameLevel.Score), p.font, 35, base.GameScreenHeight-14, base.PacmanColor)
}

func (p PlayState) drawMap(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(base.Images["map"], op)
}

func (p PlayState) drawEnemies(screen *ebiten.Image) {
	state := frameNumber / (base.FrameModulo / 2)
	for _, enemy := range gameLevel.Enemies {
		op := &ebiten.DrawImageOptions{}
		px, py := enemy.GetCoords()
		op.GeoM.Translate(float64(px)+5, float64(py)+5)
		var enemyImage *ebiten.Image
		if enemy.NightMode() {
			state %= 2
			enemyImage = enemy.GetGraphic().SubImage(image.Rect(6+40*(state), 0, 6+40*(state)+30, 38)).(*ebiten.Image)
		} else {
			enemyImage = enemy.GetGraphic().SubImage(image.Rect(6+40*(enemy.GetDirection()*2+state), 0, 6+40*(enemy.GetDirection()*2+state)+30, 38)).(*ebiten.Image)
		}
		screen.DrawImage(enemyImage, op)
	}
}

func (p PlayState) drawPacman(screen *ebiten.Image) {
	state := frameNumber / (base.FrameModulo / 3)
	pacman := gameLevel.Player
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

func (p PlayState) drawMapText(screen *ebiten.Image) {
	textFont, _ := utility.GetFont(base.PacmanFont, 8, base.DefaultDPI)
	for _, screenText := range gameLevel.Texts {
		if screenText.ExpiredTime.Before(time.Now()) {
			continue
		}
		text.Draw(screen, screenText.Text, textFont, screenText.X*base.TileSize+5, screenText.Y*(base.TileSize)+base.TileSize/2+5, base.PacmanColor)
	}
}
