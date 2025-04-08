package main

import (
	"glide/internal/components"
	"os"
	"github.com/gdamore/tcell/v2"
	//"github.com/eiannone/keyboard"
	"github.com/rivo/tview"
)

func main() {
	currentDir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	pathChannel := make(chan string)
	cliPath := make(chan string)
	refreshTreeView := make(chan bool)
	terminalOutputChannel := make(chan string)

	app := tview.NewApplication()

	treeView := components.FileExplorer(app, currentDir, pathChannel, cliPath, refreshTreeView)

	textView := components.FileViewer(app, pathChannel)

	commandLine := components.CommandLine(app, cliPath, refreshTreeView, terminalOutputChannel)

	terminalView := components.TerminalView(app, terminalOutputChannel)
	
	flex := tview.NewFlex().SetDirection(tview.FlexRow). 
	AddItem(
		tview.NewFlex().
			AddItem(treeView, 40, 1, true). 
			AddItem(textView, 0, 2, false),
		0, 1, true, 
	).
	AddItem(terminalView, 5, 0, false). 
    AddItem(commandLine, 3, 1, false)
	

	app.SetFocus(commandLine)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case '!':
			app.SetFocus(treeView)
		case '@':
			app.SetFocus(commandLine)
			return nil
		case '#':
			app.SetFocus(textView)
		case '$':
			app.SetFocus(terminalView)
		}
		return event
	})
	

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

