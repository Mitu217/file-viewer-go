package gui

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jroimartin/gocui"
)

type State struct {
	files        []os.FileInfo
	selectedLine int
}

type Gui struct {
	g     *gocui.Gui
	state *State
}

func NewGui() (*Gui, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return &Gui{
		state: &State{
			files:        files,
			selectedLine: 0,
		},
	}, nil
}

func (gui *Gui) renderFiles(g *gocui.Gui) error {
	g.Update(func(*gocui.Gui) error {
		v, err := g.View("main")
		if err != nil {
			return nil
		}

		v.Clear()
		if err := v.SetOrigin(0, 0); err != nil {
			return err
		}

		for _, file := range gui.state.files {
			fmt.Fprintln(v, " "+file.Name())
		}

		v.SetCursor(0, gui.state.selectedLine)

		return nil
	})
	return nil
}

func (gui *Gui) layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("main", -1, -1, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		gui.renderFiles(g)

		if _, err := g.SetCurrentView("main"); err != nil {
			return err
		}
	}
	return nil
}

func (gui *Gui) Run() error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.InputEsc = true // Change InputMode.InputEsc

	g.SetManagerFunc(gui.layout)

	if err = gui.keybindings(g); err != nil {
		return err
	}

	return g.MainLoop()
}

func (gui *Gui) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
