package game

type Coordinate struct {
	X int
	Y int
}

func Offset(coord Coordinate, xOffset int, yOffset int) Coordinate {
	coord.X += xOffset
	coord.Y += yOffset
	return coord
}
func Adjacent(coord Coordinate) []Coordinate {
	neighbors := []Coordinate{}
	for yOffset := -1; yOffset <= 1; yOffset++ {
		for xOffset := -1; xOffset <= 1; xOffset++ {
			xAdjusted, yAdjusted := coord.X+xOffset, coord.Y+yOffset
			neighbors = append(neighbors, Coordinate{xAdjusted, yAdjusted})
		}
	}
	return neighbors
}
