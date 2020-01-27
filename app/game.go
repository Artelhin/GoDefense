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
		image, _ := ebiten.NewImage(cellPixelSizeX, cellPixelSizeY, ebiten.FilterNearest)
		image.Fill(color.RGBA{0x00, 0xff, 0x00, 0xff})
		return image
	}()

	game.TowerImage = func() *ebiten.Image {
		image, _ := ebiten.NewImage(cellPixelSizeX, cellPixelSizeY, ebiten.FilterNearest)
		image.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
		return image
	}()
}

func (game *Game) Render(screen *ebiten.Image) {
	GridDrawOptions := func() *ebiten.DrawImageOptions {
		opts := &ebiten.DrawImageOptions{}
		x, y := TransformToPixel(cellSize / 2, cellSize / 2)
		opts.GeoM.Translate(-x, -y)
		return opts
	}
	//todo поправка на то, что 0 0 не в центре картинки
	for _, e := range game.Units {
		opts := GridDrawOptions()
		opts.GeoM.Translate(TransformToPixel(e.X, e.Y))
		screen.DrawImage(game.EnemyImage, opts)
	}
	for _, t := range game.Towers {
		opts := GridDrawOptions()
		opts.GeoM.Translate(TransformToPixel(t.Y, t.X)) // Note: (x,y) from ebiten is (y,x) from maze grid
		screen.DrawImage(game.TowerImage, opts)
	}
	for i := 0; i < len(game.Path.Points)-1; i++ {
		start, end := game.Path.Points[i], game.Path.Points[i+1]
		x1, y1 := TransformToPixel(start.X, start.Y)
		x2, y2 := TransformToPixel(end.X, end.Y)
		ebitenutil.DrawLine(screen,
			x1, y1,
			x2, y2,
			color.RGBA{0xff, 0x00, 0x00, 0xff})
	}
}

func TransformToPixel(width, height float64) (float64, float64) {
	scaleW := float64(Application().GraphOptions.ResolutionW) / float64(MazeSize * cellSize)
	scaleH := float64(Application().GraphOptions.ResolutionH) / float64(MazeSize * cellSize)
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
			10, 180, 0,
		},
	}
	game.InitTextures()
	return game
}
