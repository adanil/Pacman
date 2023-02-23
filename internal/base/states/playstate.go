package states

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/nfnt/resize"
	"golang.org/x/image/font"
	"image"
	"math/rand"
	"os"
	"pacman/internal/base"
	"pacman/internal/controllers"
	"pacman/internal/entities"
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
	enemyController    controllers.SPFEnemyController
	wallImage          *ebiten.Image //TODO delete this variable later
	foodImage          *ebiten.Image //TODO probably this too
	foodCount          int
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

func (p PlayState) Update() error {
	if gameLevel.Score == foodCount {
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
	//TODO refactor duplicate code
	for x := 0; x < base.WidthTiles; x++ {
		for y := 0; y < base.HeightTiles; y++ {
			if gameLevel.LevelTiles[x][y] == level.Wall {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*base.TileSize), float64(y*base.TileSize))
				screen.DrawImage(wallImage, op)
			} else if gameLevel.LevelTiles[x][y] == level.Food {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*base.TileSize)+base.TileSize/2, float64(y*base.TileSize)+base.TileSize/2)
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
	x := (base.GameScreenWidth - len(base.Title)*fontSize) / 2
	text.Draw(screen, base.Title, p.font, x, 20, base.PacmanColor)
}

func (p PlayState) drawScore(screen *ebiten.Image) {
	text.Draw(screen, "Score: "+strconv.Itoa(gameLevel.Score), p.font, 35, base.GameScreenHeight-12, base.PacmanColor)
}

func (p PlayState) drawMap(screen *ebiten.Image) {
	for x := 0; x < base.WidthTiles; x++ {
		for y := 0; y < base.HeightTiles; y++ {
			if gameLevel.LevelTiles[x][y] == level.Wall {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*base.TileSize), float64(y*base.TileSize))
				screen.DrawImage(wallImage, op)
			}
		}
	}
}

func (p PlayState) drawEnemies(screen *ebiten.Image) {
	state := frameNumber / (base.FrameModulo / 2)
	for _, enemy := range gameLevel.Enemies {
		op := &ebiten.DrawImageOptions{}
		px, py := enemy.GetCoords()
		op.GeoM.Translate(float64(px), float64(py))
		enemyImage := enemy.GetGraphic().SubImage(image.Rect(6+40*(enemy.GetDirection()*2+state), 0, 6+40*(enemy.GetDirection()*2+state)+30, 38)).(*ebiten.Image)
		screen.DrawImage(enemyImage, op)
	}
}

func (p PlayState) drawPacman(screen *ebiten.Image) {
	state := frameNumber / (base.FrameModulo / 3)
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

func initGame() {
	lvGenerator := level.Generator{Creator: &level.ReadLevel{Filepath: "maps/base"}}
	gameLevel = lvGenerator.CreateLevel(base.WidthTiles, base.HeightTiles, base.TileSize)

	readerWall, _ := os.Open("images/wall2.jpg")
	imgWall, _, _ := image.Decode(readerWall)
	resizedWallImage := resize.Resize(base.TileSize, base.TileSize, imgWall, resize.NearestNeighbor)
	wallImage = ebiten.NewImageFromImage(resizedWallImage)

	readerFood, _ := os.Open("images/dot.png")
	imgFood, _, _ := image.Decode(readerFood)
	resizedFoodImage := resize.Resize(base.TileSize/5, base.TileSize/5, imgFood, resize.NearestNeighbor)
	foodImage = ebiten.NewImageFromImage(resizedFoodImage)

	pacmanImage, _ := utility.ReadImage("images/pacman-pack_v2/Pacmanx2.png")
	blueEnemyImage, _ := utility.ReadImage("images/pacman-pack_v2/BlueEnemyx2.png")
	pinkEnemyImage, _ := utility.ReadImage("images/pacman-pack_v2/PinkEnemyx2.png")
	redEnemyImage, _ := utility.ReadImage("images/pacman-pack_v2/RedEnemyx2.png")

	pacman := CreateRandomPlayer(gameLevel, pacmanImage)
	blueEnemy := CreateRandomPlayer(gameLevel, blueEnemyImage)
	pinkEnemy := CreateRandomPlayer(gameLevel, pinkEnemyImage)
	redEnemy := CreateRandomPlayer(gameLevel, redEnemyImage)

	gameLevel.Player = pacman
	gameLevel.Enemies = append(gameLevel.Enemies, &blueEnemy, &pinkEnemy, &redEnemy)
	foodCount = gameLevel.CreateFood()

	keyboardController = controllers.NewKeyboardHandler(&gameLevel)
	go keyboardController.HandlePressedButtons() //TODO Exit
	enemyController = controllers.NewSPFEnemyController(&gameLevel)
}

func CreateRandomPlayer(lv level.Level, playerImage *ebiten.Image) entities.Pacman {
	for {
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(lv.Width)
		y := rand.Intn(lv.Height)
		if lv.LevelTiles[x][y] == level.Free {
			p := entities.CreatePlayer(x, y, base.TileSize, playerImage)
			return p
		}
	}
}
