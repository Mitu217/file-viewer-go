package gui

import (
	"github.com/jroimartin/gocui"
)

type Binding struct {
	ViewName string
	Handler  func(*gocui.Gui, *gocui.View) error
	Key      interface{} // FIXME: find out how to get `gocui.Key | rune`
}

func (gui *Gui) handleNextLine(g *gocui.Gui, v *gocui.View) error {
	selectedLine := gui.state.selectedLine
	if selectedLine == -1 || selectedLine == len(gui.state.files)-1 {
		return nil
	}
	gui.state.selectedLine += 1
	return gui.renderFiles(g)
}

func (gui *Gui) handlePrevLine(g *gocui.Gui, v *gocui.View) error {
	selectedLine := gui.state.selectedLine
	if selectedLine == -1 || selectedLine == 0 {
		return nil
	}

	gui.state.selectedLine -= 1
	return gui.renderFiles(g)
}

func (gui *Gui) getKeybindings() []*Binding {
	bindings := []*Binding{
		{
			ViewName: "",
			Key:      'q',
			Handler:  gui.quit,
		},
		{
			ViewName: "",
			Key:      gocui.KeyCtrlC,
			Handler:  gui.quit,
		},
		{
			ViewName: "",
			Key:      gocui.KeyEsc,
			Handler:  gui.quit,
		},
		{
			ViewName: "main",
			Key:      gocui.KeyArrowUp,
			Handler:  gui.handlePrevLine,
		},
		{
			ViewName: "main",
			Key:      gocui.KeyArrowDown,
			Handler:  gui.handleNextLine,
		},
	}
	return bindings
}

func (gui *Gui) keybindings(g *gocui.Gui) error {
	bindings := gui.getKeybindings()

	for _, binding := range bindings {
		if err := g.SetKeybinding(binding.ViewName, binding.Key, gocui.ModNone, binding.Handler); err != nil {
			return err
		}
	}
	return nil
}
