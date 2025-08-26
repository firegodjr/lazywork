package gui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

// Function name must be capitalized to make it public
func Layout(g *gocui.Gui) error {
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
