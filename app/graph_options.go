package app

import (
	"github.com/hajimehoshi/ebiten"
)

type GraphOptions struct {
	Fullscreen bool
	Borderless bool
	ResolutionW, ResolutionH int
	ScaleFactor float64
	CustomCursor bool
	MaxFPS float64
	VSync bool
}

func DefaultGraphOptions() *GraphOptions {
	graph := new(GraphOptions)
	graph.Fullscreen = true
	graph.Borderless = !graph.Fullscreen
	graph.ResolutionW, graph.ResolutionH = ebiten.ScreenSizeInFullscreen()
	graph.ScaleFactor = ebiten.DeviceScaleFactor()
	graph.CustomCursor = false
	graph.MaxFPS = 60
	graph.VSync = false

	return graph
}