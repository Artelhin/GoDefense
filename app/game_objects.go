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
		dist := math.Sqrt(math.Pow(end.X - start.X,2) + math.Pow(end.Y - start.Y,2))
		if distance - dist <= 0 {
			resultX = start.X + (end.X - start.X) * (distance / dist)
			resultY = start.Y + (end.Y - start.Y) * (distance / dist)
			return resultX, resultY
		}
		distance -= dist
	}
	log.Println("returning 0 0")
	return 0, 0
}

func FormPath(points []MazePoint) *Path {
	//todo
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
