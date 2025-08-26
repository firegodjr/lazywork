package main

import (
	"fmt"
	"log"

	"firegodjr/lazywork/gui"

	"github.com/awesome-gocui/gocui"
)

func main() {
	// Creates a new gocui instance and handle error
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	// "defer" causes g.Close() to run on program end
	defer g.Close()

	g.SetManagerFunc(gui.Layout)

	if err := gui.RegisterKeybinds(g); err != nil {
		log.Panicln(err)
	}

	// Sets a keybinding. Empty string for viewname makes it global.
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", 0, 0, maxX/2, maxY/2, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.FrameRunes = []rune{'─', '│', '╭', '╮', '╰', '╯'}
		v.Editable = true
		v.Title = "Hello World Containment Field"
		fmt.Fprintln(v, "Hello world!")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
