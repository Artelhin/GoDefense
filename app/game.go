package app

import (
	"github.com/artelhin/GoDefense/input"
	"github.com/artelhin/GoDefense/utils"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"time"
)

type Game struct {
	Input input.InputSolver

	Path   *Path
	Units  []*Enemy
	Towers []*Tower

	LastTick time.Time
	Red      bool
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

	delta := time.Since(game.LastTick)
	game.LastTick = time.Now()
	for _, e := range game.Units {
		e.Move(delta)
	}
	return nil
}

var EnemyImage = func() *ebiten.Image {
	image, _ := ebiten.NewImage(10, 10, ebiten.FilterNearest)
	image.Fill(color.RGBA{0x00,0xff,0x00,0xff})
	return image
}()

func (game *Game) Render(screen *ebiten.Image) {
	for _, e := range game.Units {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(e.X,e.Y)
		screen.DrawImage(EnemyImage, opts)
	}
}

func NewGameState() *Game {
	game := &Game{
		Input: input.NewInputSolver([]ebiten.Key{
			ebiten.KeyEscape,
			ebiten.KeySpace,
		}, nil),
		Path: &Path{
			Points: []PathPoint{
				{1, 1},
				{140, 1},
				{140, 140},
			},
			Length: 260,
		},
	}
	game.Units = []*Enemy{
		{
			game.Path,
			0, 0, 10, 30, 0,
		},
	}
	game.Towers = make([]*Tower, 0)
	game.LastTick = time.Now()
	return game
}
