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

	app := tview.NewApplication()

	treeView := components.FileExplorer(currentDir, pathChannel, cliPath)

	textView := components.FileViewer(app, pathChannel)

	commandLine := components.CommandLine(app, cliPath)
	
	flex := tview.NewFlex().SetDirection(tview.FlexRow). 
	AddItem(
		tview.NewFlex().
			AddItem(treeView, 40, 1, true). 
			AddItem(textView, 0, 2, false),
		0, 1, true, 
	).
	AddItem(commandLine, 3, 1, false) 
	

	app.SetFocus(commandLine)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 't':
			app.SetFocus(treeView)
		case 'i':
			app.SetFocus(commandLine)
			return nil
		case 'q':
			app.SetFocus(textView)
		}
		return event
	})
	

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

