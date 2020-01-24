package app

import (
	"github.com/artelhin/GoDefense/input"
	"github.com/artelhin/GoDefense/utils"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"math"
	"time"
)

type Game struct {
	Input input.InputSolver

	Path   *Path
	Units  []*Enemy
	Towers []*Tower

	Maze *Maze

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

var (
	cellsize = int(math.Min(float64(Application().GraphOptions.ResolutionW), float64(Application().GraphOptions.ResolutionH)) / 32)

	EnemyImage = func() *ebiten.Image {
		image, _ := ebiten.NewImage(30, 30, ebiten.FilterNearest)
		image.Fill(color.RGBA{0x00, 0xff, 0x00, 0xff})
		return image
	}()
	TowerImage = func() *ebiten.Image {
		image, _ := ebiten.NewImage(30, 30, ebiten.FilterNearest)
		image.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
		return image
	}()
)

func (game *Game) Render(screen *ebiten.Image) {
	GridDrawOptions := func() *ebiten.DrawImageOptions {
		opts := &ebiten.DrawImageOptions{}
		//opts.GeoM.Translate(float64(cellsize/2), -float64(cellsize/2))
		return opts
	}
	//todo поправка на то, что 0 0 не в центре картинки
	for _, e := range game.Units {
		opts := GridDrawOptions()
		opts.GeoM.Translate(e.X, e.Y)
		screen.DrawImage(EnemyImage, opts)
	}
	for _, t := range game.Towers {
		opts := GridDrawOptions()
		opts.GeoM.Translate(t.Y, t.X) // Note: (x,y) from ebiten is (y,x) from maze grid
		screen.DrawImage(TowerImage, opts)
	}
	for i := 0; i < len(game.Path.Points)-1; i++ {
		start, end := game.Path.Points[i], game.Path.Points[i+1]
		ebitenutil.DrawLine(screen,
			start.X,
			start.Y,
			end.X,
			end.Y,
			color.RGBA{0x00, 0xff, 0x00, 0xff})
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
	game.Towers = make([]*Tower, 0)
	game.LastTick = time.Now()
	game.Maze = NewMaze(32, 32)
	tower := &Tower{X: float64(cellsize/2 + 1*cellsize), Y: float64(cellsize / 2)}
	game.Maze.Cell[1][0] = tower
	game.Towers = append(game.Towers, tower)
	solution, _ := game.Maze.SolveMaze()
	game.Path = FormPath(solution)
	game.Units = []*Enemy{
		{
			game.Path,
			game.Path.Points[0].X,
			game.Path.Points[0].Y,
			10, 180, 0,
		},
	}
	return game
}
