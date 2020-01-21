package app

import (
	"github.com/hajimehoshi/ebiten"
)

type State interface {
	Tick() error
	Render(screen *ebiten.Image)
}
