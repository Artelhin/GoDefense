package app

import (
	"github.com/artelhin/GoDefense/input"
	"github.com/artelhin/GoDefense/utils"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"time"
)

type Game struct {
	Input input.InputSolver

	Path   *Path
	Units  []*Enemy
	Towers []*Tower

	EnemyImage *ebiten.Image
	TowerImage *ebiten.Image

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
		case ebiten.KeyQ:
			Application().State = NewMenuState()
			return nil
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
	cellPixelSizeX = Application().GraphOptions.ResolutionW / MazeSize
	cellPixelSizeY = Application().GraphOptions.ResolutionH / MazeSize
)

func (game *Game) InitTextures() {
	game.EnemyImage = func() *ebiten.Image {
		image, _ := ebiten.NewImage(cellPixelSizeY, cellPixelSizeY, ebiten.FilterNearest)
		image.Fill(color.RGBA{0x00, 0xff, 0x00, 0xff})
		return image
	}()

	game.TowerImage = func() *ebiten.Image {
		image, _ := ebiten.NewImage(cellPixelSizeY, cellPixelSizeY, ebiten.FilterNearest)
		image.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
		return image
	}()
}

func (game *Game) Render(screen *ebiten.Image) {
	//find grid upper left point
	rx, ry := Application().GraphOptions.ResolutionW, Application().GraphOptions.ResolutionH
	rx = rx / 2 - ry / 2
	ry = 0
	gridDrawOptions := func() *ebiten.DrawImageOptions {
		opts := &ebiten.DrawImageOptions{}
		x, y := TransformToPixel(cellSize / 2, cellSize / 2) // move cell drawing point to the center of the cell
		opts.GeoM.Translate(-x, -y)
		opts.GeoM.Translate(float64(rx), float64(ry)) // move drawing point to the upper left corner of the grid
		return opts
	}
	grid := game.GridImage(screen)

	// draw enemy units
	for _, e := range game.Units {
		opts := gridDrawOptions()
		opts.GeoM.Translate(TransformToPixelWithImage(grid, e.X, e.Y))
		screen.DrawImage(game.EnemyImage, opts)
	}

	// draw towers
	for _, t := range game.Towers {
		opts := gridDrawOptions()
		opts.GeoM.Translate(TransformToPixelWithImage(grid, t.Y, t.X)) // Note: (x,y) from ebiten is (y,x) from maze grid
		screen.DrawImage(game.TowerImage, opts)
	}

	// draw path
	for i := 0; i < len(game.Path.Points)-1; i++ {
		start, end := game.Path.Points[i], game.Path.Points[i+1]
		x1, y1 := TransformToPixelWithImage(grid, start.X, start.Y)
		x2, y2 := TransformToPixelWithImage(grid, end.X, end.Y)
		ebitenutil.DrawLine(screen,
			x1 + float64(rx), y1 + float64(ry),
			x2 + float64(rx), y2 + float64(ry),
			color.RGBA{0xff, 0x00, 0x00, 0xff})
	}
}

func TransformToPixel(width, height float64) (float64, float64) {
	scaleW := float64(Application().GraphOptions.ResolutionW) / float64(MazeSize*cellSize)
	scaleH := float64(Application().GraphOptions.ResolutionH) / float64(MazeSize*cellSize)
	return width * scaleW, height * scaleH
}

func NewGameState() *Game {
	game := &Game{
		Input: input.NewInputSolver([]ebiten.Key{
			ebiten.KeyEscape,
			ebiten.KeySpace,
			ebiten.KeyQ,
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
	tower := &Tower{X: float64(cellSize/2 + 1*cellSize), Y: float64(cellSize / 2)}
	game.Maze.Cell[1][0] = tower
	game.Towers = append(game.Towers, tower)
	solution, _ := game.Maze.SolveMaze()
	game.Path = FormPath(solution)
	game.Units = []*Enemy{
		{
			game.Path,
			game.Path.Points[0].X,
			game.Path.Points[0].Y,
			10, 500, 0,
		},
	}
	game.InitTextures()
	return game
}
