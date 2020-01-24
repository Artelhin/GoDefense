package app

import "errors"

type MazeObject interface {
	Transparent() bool
}

type MazePoint struct {
	X, Y int
}

type Maze struct {
	Columns, Rows int // size
	Cell          [][]MazeObject
	MazePoints    []MazePoint
}

var (
	ErrInvalidMazePoints = errors.New("no path between points")
)

func (maze *Maze) SolveMaze() ([]MazePoint, error) {
	path := make([]MazePoint, 0)
	for i := 0; i < len(maze.MazePoints) - 1; i++ {
		var visited = make([][]bool, maze.Rows)
		for j := range visited {
			visited[j] = make([]bool, maze.Columns)
		}
		section := maze.Solve(visited, maze.MazePoints[i], maze.MazePoints[i+1])
		if section == nil {
			return nil, ErrInvalidMazePoints
		}
		path = append(path, section...)

	}
	return path, nil
}

//Solve recursively finds a path between two given points.
//Considers previously visited cells.
//Returns nil if no path can be found
func (maze *Maze) Solve(visited [][]bool, cur, target MazePoint) []MazePoint {
	if cur.X == target.X && cur.Y == target.Y {
		return []MazePoint{cur}
	}
	visited[cur.X][cur.Y] = true
	nearbyPoints := []MazePoint{
		{cur.X + 1, cur.Y + 1},
		{cur.X + 1, cur.Y - 1},
		{cur.X - 1, cur.Y + 1},
		{cur.X - 1, cur.Y - 1},
	}
	var res []MazePoint = nil
	for _, point := range nearbyPoints {
		if point.X > 0 && point.X < maze.Rows &&
			point.Y > 0 && point.Y < maze.Columns &&
			maze.Cell[point.X][point.Y].Transparent() &&
			!visited[point.X][point.Y] {
			if path := maze.Solve(visited, point, target); path != nil {
				if res == nil || len(res) == 0 || len(path)+1 < len(res) {
					res = append([]MazePoint{cur}, path...)
				}
			}
		}
	}
	return res
}
