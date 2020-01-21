package app

import (
	"github.com/artelhin/GoDefense/input"
	"github.com/artelhin/GoDefense/utils"
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type Game struct {
	Input input.InputSolver

	Path *Path
	Units []Enemy
	Towers []Tower

	Red bool
}

func (game *Game) Tick() error {
	keys := game.Input.GetInput()
	for _, key := range keys {
		if key.State != input.Pressed && key.State != input.Holded {
			continue
		}
		switch key.Key {
		case ebiten.KeyEscape:
			Application().ShouldQuit = true
			return utils.ErrNormalQuit
		case ebiten.KeySpace:
			game.Red = true
		}
	}
	return nil
}

func (game *Game) Render(screen *ebiten.Image) {
	if game.Red {
		screen.Fill(color.RGBA{0,0,0xff,0xff})
		game.Red = false
	} else {
		screen.Fill(color.RGBA{0, 0xff, 0, 0xff})
	}
}

func NewGameState() *Game {
	return &Game{
		Input:input.NewInputSolver([]ebiten.Key{
			ebiten.KeyEscape,
			ebiten.KeySpace,
		}, nil),
	}
}