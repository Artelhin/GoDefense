package app

import (
	"github.com/hajimehoshi/ebiten"
	"image"
)

func (game *Game) GridImage(screen *ebiten.Image) *ebiten.Image {
	x, y := Application().GraphOptions.ResolutionW, Application().GraphOptions.ResolutionH
	r := image.Rectangle{
		Min:image.Point{
			X: x / 2 - y / 2,
			Y: 0,
		},
		Max: image.Point{
			X: x / 2 + y / 2,
			Y: y,
		},
	}
	return screen.SubImage(r).(*ebiten.Image)
}

func TransformToPixelWithImage(img *ebiten.Image, width, height float64) (float64, float64) {
	x, y := img.Size()
	scaleW := float64(x) / float64(MazeSize * cellSize)
	scaleH := float64(y) / float64(MazeSize * cellSize)
	return width * scaleW, height * scaleH
}