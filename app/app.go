package app

import (
	"fmt"
	"github.com/artelhin/GoDefense/utils"
	"github.com/hajimehoshi/ebiten"
	"log"
)

type App struct {
	Version string
	GraphOptions *GraphOptions
	State State

	ShouldQuit bool
}

const (
	NormalQuit = "quited"
)

var (
	app *App
)

func (app *App) Update(screen *ebiten.Image) error {
	if app.ShouldQuit {
		return utils.ErrNormalQuit
	}
	if err := app.State.Tick(); err != nil {
		return err
	}
	app.State.Render(screen)
	return nil
}

func Application() *App {
	return app
}

func (app *App) Run() {
	ebiten.SetFullscreen(app.GraphOptions.Fullscreen)
	ebiten.SetCursorVisible(app.GraphOptions.CustomCursor)
	err := ebiten.Run(app.Update,
		int(float64(app.GraphOptions.ResolutionW)*app.GraphOptions.ScaleFactor),
		int(float64(app.GraphOptions.ResolutionH)*app.GraphOptions.ScaleFactor), 1/app.GraphOptions.ScaleFactor,
		fmt.Sprintf("GoDefense %s", app.Version))
	if err != nil {
		if err != utils.ErrNormalQuit {
			log.Println("Error occurred: ", err)
		}
	}
}

func NewApp() *App {
	app = new(App)
	app.GraphOptions = DefaultGraphOptions()
	app.ShouldQuit = false
	app.State = NewMenuState()
	return app
}