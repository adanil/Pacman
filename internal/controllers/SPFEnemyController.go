package controllers

import (
	"pacman/internal/command"
	"pacman/internal/entities"
	"pacman/internal/level"
	"pacman/pkg/algorithms"
	"pacman/pkg/queue"
)

type coordinates = level.Coordinates

type SPFEnemyController struct {
	level       *level.Level
	graph       map[coordinates][]coordinates
	routes      map[entities.Enemy][]coordinates
	routesIndex map[entities.Enemy]int
}

func NewSPFEnemyController(level_ *level.Level) SPFEnemyController {
	g := formGraph(level_)
	return SPFEnemyController{level: level_, graph: g, routes: make(map[entities.Enemy][]coordinates), routesIndex: make(map[entities.Enemy]int)}
}

func (e *SPFEnemyController) GetCommands() []command.Command {
	var commands []command.Command
	for _, enemy := range e.level.Enemies() {
		if c := e.GetCommand(enemy); c != nil {
			commands = append(commands, c)
		}
	}
	return commands
}

func (e *SPFEnemyController) GetCommand(enemy *entities.Enemy) command.Command {
	x, y := enemy.GetCoords()
	if x%e.level.TileSize() != 0 || y%e.level.TileSize() != 0 {
		return nil
	}
	x /= e.level.TileSize()
	y /= e.level.TileSize()
	route, exist := e.routes[*enemy]
	if !exist || e.routesIndex[*enemy] >= len(route) || enemy.GetStopped() {
		ex, ey := enemy.GetCoords()
		start := level.NewCoordinate(ex/e.level.TileSize(), ey/e.level.TileSize())
		ex, ey = e.level.Player().GetCoords()
		target := level.NewCoordinate(ex/e.level.TileSize(), ey/e.level.TileSize())
		route = findPath(start, target, e.graph)
		if len(route) <= 1 {
			return nil
		}
		e.routes[*enemy] = route
		e.routesIndex[*enemy] = 1
	}
	nextCoords := route[e.routesIndex[*enemy]]
	var direction int
	if nextCoords.X-x == 1 {
		direction = entities.RIGHT
	} else if nextCoords.X-x == -1 {
		direction = entities.LEFT
	} else if nextCoords.Y-y == 1 {
		direction = entities.DOWN
	} else {
		direction = entities.UP
	}
	if direction == entities.OppositeDirection[enemy.GetDirection()] {
		return nil
	}
	e.routesIndex[*enemy]++
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
	prev[start] = level.NewCoordinate(-1, -1)

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
	stop := level.NewCoordinate(-1, -1)
	for current != stop {
		path = append(path, current)
		current = prev[current]
	}

	algorithms.ReverseSlice(path)
	return path
}

func formGraph(lv *level.Level) map[coordinates][]coordinates {
	g := make(map[coordinates][]coordinates)
	for h := 0; h < lv.Height(); h++ {
		for w := 0; w < lv.Width(); w++ {
			if lv.GetEntityByCoordinates(level.NewCoordinate(w, h)) == level.Wall {
				continue
			}
			currCoordinates := level.NewCoordinate(w, h)

			wRight := (w + 1) % lv.Width()
			wLeft := (w - 1 + lv.Width()) % lv.Width()
			hUp := (h + 1) % lv.Height()
			hDown := (h - 1 + lv.Height()) % lv.Height()

			if lv.GetEntityByCoordinates(level.NewCoordinate(wRight, h)) != level.Wall {
				g[currCoordinates] = append(g[currCoordinates], level.NewCoordinate(wRight, h))
			}
			if lv.GetEntityByCoordinates(level.NewCoordinate(wLeft, h)) != level.Wall {
				g[currCoordinates] = append(g[currCoordinates], level.NewCoordinate(wLeft, h))
			}
			if lv.GetEntityByCoordinates(level.NewCoordinate(w, hUp)) != level.Wall {
				g[currCoordinates] = append(g[currCoordinates], level.NewCoordinate(w, hUp))
			}
			if lv.GetEntityByCoordinates(level.NewCoordinate(w, hDown)) != level.Wall {
				g[currCoordinates] = append(g[currCoordinates], level.NewCoordinate(w, hDown))
			}
		}
	}
	return g
}
