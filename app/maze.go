package app

import (
	"errors"
	"github.com/artelhin/GoDefense/utils"
)

type MazeObject interface {
	Transparent() bool
}

type MazePoint struct {
	X, Y int
}

type Maze struct {
	Rows, Columns int // size
	Cell          [][]MazeObject
	MazePoints    []MazePoint
}

var (
	ErrInvalidMazePoints = errors.New("no path between points")
)

//SolveMaze solves the maze in receiver and returns a sequence of points, representing the solution.
//Returns error if maze has no solution
func (maze *Maze) SolveMaze() ([]MazePoint, error) {
	path := make([]MazePoint, 0)
	for i := 0; i < len(maze.MazePoints)-1; i++ {
		var (
			visited  = make([][]bool, maze.Rows)
			ancestor = make([][]MazePoint, maze.Rows)
		)
		for j := range visited {
			visited[j] = make([]bool, maze.Columns)
			ancestor[j] = make([]MazePoint, maze.Columns)
		}
		section := maze.Solve(maze.MazePoints[i], maze.MazePoints[i+1])
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
func (maze *Maze) Solve(from, to MazePoint) []MazePoint {

	var (
		visited  [][]bool
		ancestor [][]MazePoint
	)

	visited = make([][]bool, maze.Rows)
	ancestor = make([][]MazePoint, maze.Rows)
	for j := range visited {
		visited[j] = make([]bool, maze.Columns)
		ancestor[j] = make([]MazePoint, maze.Columns)
	}

	q := utils.NewQueue()
	visited[from.X][from.Y] = true
	ancestor[from.X][from.Y] = from
	q.Push(from)
	for q.Len() != 0 {
		p := q.Pop().(MazePoint)
		if p != to {
			nearbyPoints := []MazePoint{
				{p.X + 1, p.Y},
				{p.X - 1, p.Y},
				{p.X, p.Y + 1},
				{p.X, p.Y - 1},
			}
			for _, np := range nearbyPoints {
				if np.X >= 0 && np.X < maze.Rows &&
					np.Y >= 0 && np.Y < maze.Columns &&
					(maze.Cell[np.X][np.Y] == nil || maze.Cell[np.X][np.Y].Transparent()) &&
					!visited[np.X][np.Y] {
					visited[np.X][np.Y] = true
					ancestor[np.X][np.Y] = p
					q.Push(np)
				}
			}
		} else {
			section := []MazePoint{p}
			for ancestor[p.X][p.Y] != p {
				p = ancestor[p.X][p.Y]
				section = append([]MazePoint{p}, section...)
			}
			return section
		}
	}
	return nil
}

//NewMaze returns default empty maze with start at left-top corner and exit at bottom-right
func NewMaze(rows, columns int) *Maze {
	maze := new(Maze)
	maze.Rows, maze.Columns = rows, columns
	maze.Cell = make([][]MazeObject, rows)
	for i := range maze.Cell {
		maze.Cell[i] = make([]MazeObject, columns)
	}
	maze.MazePoints = []MazePoint{
		{0, 0},
		{rows - 1, columns - 1},
	}
	return maze
}
