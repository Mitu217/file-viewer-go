package app

import (
	"io"
	"log"

	"github.com/jroimartin/gocui"

	"github.com/Mitu217/file-viewer/pkg/gui"
)

type App struct {
	closers []io.Closer
	gui     *gui.Gui
}

func NewApp() (*App, error) {
	gui, err := gui.NewGui()
	if err != nil {
		return nil, err
	}

	return &App{
		closers: []io.Closer{},
		gui:     gui,
	}, nil
}

func (app *App) Run() {
	for {
		if err := app.gui.Run(); err != nil {
			if err == gocui.ErrQuit {
				break
			} else {
				log.Panicln(err)
			}
		}
	}
}
