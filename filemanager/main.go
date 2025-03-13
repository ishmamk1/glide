package main

import (
	"fmt"
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

	app := tview.NewApplication()

	treeView := components.FileExplorer(currentDir, pathChannel)

	textView := tview.NewTextView().
		SetText("Select a file to view content").
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true)
	
	go func() {
		for filePath := range pathChannel {
			// Read file content
			content, err := os.ReadFile(filePath)
			if err != nil {
				content = []byte("Error reading file")
			}

			// Update TextView dynamically
			app.QueueUpdateDraw(func() {
				textView.SetText(fmt.Sprintf("[yellow]File:[white] %s\n\n[white]%s", filePath, string(content)))
			})
			
		}
	}()

	inputField := tview.NewInputField().
		SetLabel("Enter a number: ").
		SetFieldWidth(10).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})
	
	flex := tview.NewFlex().SetDirection(tview.FlexRow). // Vertical layout
	AddItem(
		tview.NewFlex(). // Inner Flex for tree + text view
			AddItem(treeView, 40, 1, true). 
			AddItem(textView, 0, 2, false),
		0, 1, true, // This fills most of the screen
	).
	AddItem(inputField, 3, 1, false) // Fixed height input field at the bottom
	

	app.SetFocus(inputField)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 't':
			app.SetFocus(treeView)
		case 'i':
			app.SetFocus(inputField)
		}
		return event
	})
	

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

