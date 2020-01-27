package app

import (
	"log"
	"math"
	"time"
)

//todo redo
type Tower struct {
	X, Y        float64 // position
	AttackPower int
	Name        string
}

func (*Tower) Transparent() bool {
	return false
}

type PathPoint struct {
	X, Y float64
}

type Path struct {
	Points []PathPoint
	Length float64
}

func (p *Path) Follow(distance float64) (float64, float64) {
	n := len(p.Points)
	if n < 2 {
		panic("bad path: less than 2 points")
	}
	if distance >= p.Length {
		lastPoint := p.Points[len(p.Points)-1]
		return lastPoint.X, lastPoint.Y
	}
	var (
		resultX, resultY float64
	)
	for i := 0; i < n-1; i++ {
		start := p.Points[i]
		end := p.Points[i+1]
		dist := math.Sqrt(math.Pow(end.X-start.X, 2) + math.Pow(end.Y-start.Y, 2))
		if distance-dist <= 0 {
			resultX = start.X + (end.X-start.X)*(distance/dist)
			resultY = start.Y + (end.Y-start.Y)*(distance/dist)
			return resultX, resultY
		}
		distance -= dist
	}
	log.Println("returning 0 0")
	return 0, 0
}

const cellSize = 100

func FormPath(points []MazePoint) *Path {
	// optimize points, left only corners and start/end
	optimized := []MazePoint{points[0]}
	if len(points) > 2 {
		for i := 1; i < len(points)-1; i++ {
			if !(points[i-1].X-points[i].X == points[i].X-points[i+1].X ||
				points[i-1].Y-points[i].Y == points[i].Y-points[i+1].Y) { // so this is a corner and can't be removed => add it to the path
				optimized = append(optimized, points[i])
			}
		}
	}
	optimized = append(optimized, points[len(points)-1])

	// transform from grid to game object coordinates
	path := new(Path)
	path.Points = make([]PathPoint, 0)
	for _, p := range optimized {
		point := PathPoint{
			X: float64(cellSize/2 + p.Y*cellSize), // Note: (x,y) from ebiten is (y,x) from maze grid
			Y: float64(cellSize/2 + p.X*cellSize),
		}
		path.Points = append(path.Points, point)
	}
	for i := 0; i < len(path.Points)-1; i++ {
		path.Length += math.Sqrt(math.Pow(path.Points[i].X-path.Points[i+1].X, 2) + math.Pow(path.Points[i].Y-path.Points[i+1].Y, 2)) // calculate distance between two points
	}

	return path
}

type Enemy struct {
	Path *Path

	X, Y     float64
	HP       int
	Speed    float64
	Distance float64
}

func (e *Enemy) Move(delta time.Duration) {
	e.Distance += delta.Seconds() * e.Speed
	e.X, e.Y = e.Path.Follow(e.Distance)
}

type Bullet struct {
	X, Y   float64 // position
	Vx, Vy float64 // speed
	Speed  float64
	Damage int
}
