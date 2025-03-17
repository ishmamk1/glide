package components

import (
	//"fmt"
	//"os"
	//"path/filepath"
	//"path/filepath"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CommandLine(app *tview.Application, pathChannel chan string) *tview.InputField {
	inputField := tview.NewInputField().
		SetFieldWidth(10).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		}).
		SetFieldTextColor(tcell.ColorWhite).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetLabelColor(tcell.ColorWhite).
		SetLabel("/")


	go func() {
		for filePath := range pathChannel {
			app.QueueUpdateDraw(func() {
                inputField.SetLabel(filePath + ": ")
            })
		}
	}()

	

	return inputField
}



