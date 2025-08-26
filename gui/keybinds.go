package gui

import "github.com/awesome-gocui/gocui"

func RegisterKeybinds(g *gocui.Gui) error {
	// Register global keybinds
	// Ctrl+C Quit
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	// hjkl navigation
	// TODO: View focus system in the same vein as lazygit

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
