package app

import (
	"github.com/artelhin/GoDefense/input"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
)

type Menu struct {
	Input input.InputSolver

	White bool
}

func (menu *Menu) Tick() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		Application().ShouldQuit = true
		return ErrNormalQuit
	}
	keys := menu.Input.GetInput()
	for _, key := range keys {
		if !key.IsKey {
			continue
		}
		if key.State == input.Pressed || key.State == input.Holded {
			switch key.Key {
			case ebiten.KeyEscape:
				Application().ShouldQuit = true
				return ErrNormalQuit
			case ebiten.KeyEnter:
			Application().State = NewGameState()
			case ebiten.KeySpace:
				menu.White = true
			}
		}
	}
	return nil
}

func (menu *Menu) Render(screen *ebiten.Image) {
	if menu.White {
		screen.Fill(color.White)
		menu.White = false
	} else {
		screen.Fill(color.Black)
	}
	ebitenutil.DebugPrint(screen, `Press ENTER to proceed to the game`)
}

func NewMenuState() *Menu {
	return &Menu{
		Input:input.NewInputSolver([]ebiten.Key{
			ebiten.KeyEscape,
			ebiten.KeyEnter,
			ebiten.KeySpace,
		}, nil),
	}
}
