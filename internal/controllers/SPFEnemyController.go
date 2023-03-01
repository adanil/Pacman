package controllers

import (
	"pacman/internal/command"
	"pacman/internal/entities"
	"pacman/internal/level"
	"pacman/pkg/algorithms"
	"pacman/pkg/queue"
)

type coordinates struct {
	x, y int
}

type SPFEnemyController struct {
	level       *level.Level
	updateFreq  int
	graph       map[coordinates][]coordinates
	routes      map[entities.Playable][]coordinates
	routesIndex map[entities.Playable]int
}

func NewSPFEnemyController(level_ *level.Level) SPFEnemyController {
	g := formGraph(level_)
	return SPFEnemyController{level: level_, graph: g, routes: make(map[entities.Playable][]coordinates), routesIndex: make(map[entities.Playable]int)}
}

func (e *SPFEnemyController) GetCommands() []command.Command {
	var commands []command.Command
	for _, enemy := range e.level.Enemies {
		if c := e.GetCommand(enemy); c != nil {
			commands = append(commands, c)
		}
	}
	return commands
}

func (e *SPFEnemyController) GetCommand(enemy entities.Playable) command.Command {
	x, y := enemy.GetCoords()
	if x%e.level.TileSize != 0 || y%e.level.TileSize != 0 {
		return nil
	}
	x /= e.level.TileSize
	y /= e.level.TileSize
	route, exist := e.routes[enemy]
	if !exist || e.routesIndex[enemy] >= len(route) || enemy.GetStopped() {
		ex, ey := enemy.GetCoords()
		start := coordinates{ex / e.level.TileSize, ey / e.level.TileSize}
		ex, ey = e.level.Player.GetCoords()
		target := coordinates{ex / e.level.TileSize, ey / e.level.TileSize}
		route = findPath(start, target, e.graph)
		if len(route) <= 1 {
			return nil
		}
		e.routes[enemy] = route
		e.routesIndex[enemy] = 1
	}
	nextCoords := route[e.routesIndex[enemy]]
	var direction int
	if nextCoords.x-x == 1 {
		direction = entities.RIGHT
	} else if nextCoords.x-x == -1 {
		direction = entities.LEFT
	} else if nextCoords.y-y == 1 {
		direction = entities.DOWN
	} else {
		direction = entities.UP
	}
	if direction == entities.OppositeDirection[enemy.GetDirection()] {
		return nil
	}
	e.routesIndex[enemy]++
	if enemy.GetDirection() == direction {
		return nil
	}
	cdCommand := command.NewChangeDirectionCommand(direction, enemy, e.level)
	return &cdCommand

}

// Dijkstra algorithm
func findPath(start, target coordinates, g map[coordinates][]coordinates) []coordinates {
	var nodeInProcess queue.Queue[coordinates]
	nodeInProcess.Push(start)
	dist := make(map[coordinates]int)
	dist[start] = 0
	prev := make(map[coordinates]coordinates)
	prev[start] = coordinates{x: -1, y: -1}

	for nodeInProcess.Size() != 0 {
		current := nodeInProcess.Front()
		nodeInProcess.Pop()

		if current == target {
			break
		}
		for _, neighbor := range g[current] {
			d := dist[current] + 1
			_, exist := dist[neighbor]
			if !exist || dist[neighbor] > d {
				dist[neighbor] = d
				nodeInProcess.Push(neighbor)
				prev[neighbor] = current
			}
		}
	}

	var path []coordinates
	current := target
	stop := coordinates{-1, -1}
	for current != stop {
		path = append(path, current)
		current = prev[current]
	}

	algorithms.ReverseSlice(path)
	return path
}

func formGraph(lv *level.Level) map[coordinates][]coordinates {
	g := make(map[coordinates][]coordinates)
	for h := 0; h < lv.Height; h++ {
		for w := 0; w < lv.Width; w++ {
			if lv.LevelTiles[w][h] == level.Wall {
				continue
			}
			currCoord := coordinates{
				x: w,
				y: h,
			}

			wright := (w + 1) % lv.Width
			wleft := (w - 1 + lv.Width) % lv.Width
			hup := (h + 1) % lv.Height
			hdown := (h - 1 + lv.Height) % lv.Height

			if lv.LevelTiles[wright][h] != level.Wall {
				g[currCoord] = append(g[currCoord], coordinates{x: wright, y: h})
			}
			if lv.LevelTiles[wleft][h] != level.Wall {
				g[currCoord] = append(g[currCoord], coordinates{x: wleft, y: h})
			}
			if lv.LevelTiles[w][hup] != level.Wall {
				g[currCoord] = append(g[currCoord], coordinates{x: w, y: hup})
			}
			if lv.LevelTiles[w][hdown] != level.Wall {
				g[currCoord] = append(g[currCoord], coordinates{x: w, y: hdown})
			}
		}
	}
	return g
}
