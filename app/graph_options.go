package app

import (
	"github.com/hajimehoshi/ebiten"
)

type GraphOptions struct {
	Fullscreen bool
	ResolutionW, ResolutionH int
	ScaleFactor float64
	CustomCursor bool
}

func DefaultGraphOptions() *GraphOptions {
	graph := new(GraphOptions)
	graph.Fullscreen = true
	graph.ResolutionW, graph.ResolutionH = ebiten.ScreenSizeInFullscreen()
	graph.ScaleFactor = ebiten.DeviceScaleFactor()
	graph.CustomCursor = false
	return graph
}