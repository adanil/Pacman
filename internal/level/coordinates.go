package level

type Coordinates struct {
	x, y int
}

var enemiesSpawnCoords = []Coordinates{{9, 9}, {10, 9}, {11, 9}, {12, 9}, {10, 8}, {11, 8}}

func isEnemySpawn(c Coordinates) bool {
	for _, s := range enemiesSpawnCoords {
		if s == c {
			return true
		}
	}
	return false
}
