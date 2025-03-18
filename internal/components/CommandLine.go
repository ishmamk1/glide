package components

import (
	//"fmt"
	//"os"
	//"path/filepath"
	//"path/filepath"
	//"os"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CommandLine(app *tview.Application, pathChannel chan string, refreshTreeView chan bool) *tview.InputField {
	commandLine := tview.NewInputField().
		SetFieldTextColor(tcell.ColorWhite).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetLabelColor(tcell.ColorWhite).
		SetLabel("/")
	
	var createFilePath string
	


	go func() {
		for filePath := range pathChannel {
			app.QueueUpdateDraw(func() {
                commandLine.SetLabel(filePath + ": ")
            })
			createFilePath = filePath
		}
	}()

	commandLine.SetDoneFunc(func(key tcell.Key) {
		command := commandLine.GetText()
		commandParts := strings.Split(command, " ")

		switch commandParts[0] {
		case "cd":
		
		case "create":
			os.Create(createFilePath + "/" + commandParts[1])
			refreshTreeView <- true
		default:
				
		}
		commandLine.SetText("")

	})



	return commandLine
}



