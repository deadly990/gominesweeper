package game

type Coordinate struct {
	x int
	y int
}

func Offset(coord Coordinate, xOffset int, yOffset int) Coordinate {
	coord.x += xOffset
	coord.y += yOffset
	return coord
}
func Adjacent(coord Coordinate) []Coordinate {
	neighbors := []Coordinate{}
	for yOffset := -1; yOffset <= 1; yOffset++ {
		for xOffset := -1; xOffset <= 1; xOffset++ {
			xAdjusted, yAdjusted := coord.x+xOffset, coord.y+yOffset
			neighbors = append(neighbors, Coordinate{xAdjusted, yAdjusted})
		}
	}
	return neighbors
}
