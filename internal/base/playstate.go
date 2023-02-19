package base

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image"
	"pacman/internal/level"
	"strconv"
)

const fontSize = 16

type PlayState struct {
	g    *Game
	font font.Face
}

func NewPlayState(g *Game) PlayState {
	defaultFont, _ := GetFont(baseFont, fontSize, defaultDPI)
	return PlayState{g: g, font: defaultFont}
}

func (p PlayState) Update() error {
	if gameLevel.Player.Score == foodCount {
		p.g.SetState(NewWinState(p.g, gameLevel.Player.Score))
		return nil
	}
	commands := keyboardController.GetKeyboardCommands()
	commands = append(commands, enemyController.GetCommands()...)
	for _, com := range commands {
		com.Execute()
	}
	ok := gameLevel.UpdateAll()
	if ok == false {
		p.g.SetState(NewGameOverState(p.g, gameLevel.Player.Score))
	}
	return nil
}

func (p PlayState) Draw(screen *ebiten.Image) {
	frameNumber = (frameNumber + 1) % frameModulo
	//TODO refactor duplicate code
	for x := 0; x < widthTiles; x++ {
		for y := 0; y < heightTiles; y++ {
			if gameLevel.LevelTiles[x][y] == level.Wall {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
				screen.DrawImage(wallImage, op)
			} else if gameLevel.LevelTiles[x][y] == level.Food {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*tileSize)+tileSize/2, float64(y*tileSize)+tileSize/2)
				screen.DrawImage(foodImage, op)
			}
		}
	}

	p.drawPacman(screen)
	p.drawEnemies(screen)
	p.drawTitle(screen)
	p.drawScore(screen)
}

func (p PlayState) drawTitle(screen *ebiten.Image) {
	x := (gameScreenWidth - len(title)*fontSize) / 2
	text.Draw(screen, title, p.font, x, 20, yellowColor)
}

func (p PlayState) drawScore(screen *ebiten.Image) {
	text.Draw(screen, "Score: "+strconv.Itoa(gameLevel.Player.Score), p.font, 35, gameScreenHeight-12, yellowColor)
}

func (p PlayState) drawMap(screen *ebiten.Image) {
	for x := 0; x < widthTiles; x++ {
		for y := 0; y < heightTiles; y++ {
			if gameLevel.LevelTiles[x][y] == level.Wall {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*tileSize), float64(y*tileSize))
				screen.DrawImage(wallImage, op)
			}
		}
	}
}

func (p PlayState) drawEnemies(screen *ebiten.Image) {
	state := frameNumber / (frameModulo / 2)
	for _, enemy := range gameLevel.Enemies {
		op := &ebiten.DrawImageOptions{}
		px, py := enemy.GetCoords()
		op.GeoM.Translate(float64(px), float64(py))
		enemyImage := enemy.GetGraphic().SubImage(image.Rect(6+40*(enemy.GetDirection()*2+state), 0, 6+40*(enemy.GetDirection()*2+state)+30, 38)).(*ebiten.Image)
		screen.DrawImage(enemyImage, op)
	}
}

func (p PlayState) drawPacman(screen *ebiten.Image) {
	state := frameNumber / (frameModulo / 3)
	pacman := gameLevel.Player
	op := &ebiten.DrawImageOptions{}
	px, py := pacman.GetCoords()
	op.GeoM.Translate(float64(px), float64(py))
	pacmanImage := (*ebiten.Image)(nil)
	if state == 2 {
		pacmanImage = pacman.GetGraphic().SubImage(image.Rect(296, 0, 324, 30)).(*ebiten.Image)
	} else {
		state = (state + 1) % 2
		pacmanImage = pacman.GetGraphic().SubImage(image.Rect(38*(pacman.GetDirection()*2+state), 0, 38*(pacman.GetDirection()*2+state)+30, 30)).(*ebiten.Image)
	}
	screen.DrawImage(pacmanImage, op)
}
