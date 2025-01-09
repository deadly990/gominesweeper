package game

type Coordinate struct {
	X int
	Y int
}

// Returns a new Coordinate with X and Y values augmented by corresponding arguments.
func (coord Coordinate) Offset(xOffset int, yOffset int) Coordinate {
	return Coordinate{coord.X + xOffset, coord.Y + yOffset}
}

// Returns a slice containing adjacent Coordinates to input param.
func (coord Coordinate) Adjacent() []Coordinate {
	neighbors := []Coordinate{}
	for yOffset := -1; yOffset <= 1; yOffset++ {
		for xOffset := -1; xOffset <= 1; xOffset++ {
			neighbors = append(neighbors, coord.Offset(xOffset, yOffset))
		}
	}
	return neighbors
}
