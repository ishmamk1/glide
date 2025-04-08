package components

import (
	"os"
	"strings"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"fmt"
)

func CommandLine(app *tview.Application, pathChannel chan string, refreshTreeView chan bool, terminalOutputChannel chan string) *tview.InputField {
	commandLine := tview.NewInputField().
		SetFieldTextColor(tcell.ColorWhite).
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetLabelColor(tcell.ColorWhite)
	
	commandLine.SetBorder(true)
	commandLine.SetTitleColor(tcell.ColorYellow)
	commandLine.SetBorderColor(tcell.ColorYellow)
	commandLine.SetTitleAlign(tview.AlignLeft)

	curr_wd, err := os.Getwd()

	if err != nil {
		return commandLine
	}

	commandLine.SetLabel(curr_wd + ": ")
	
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
			if len(commandParts) > 1 {
				newPath := commandParts[1]
				err := os.Chdir(newPath)
				if err != nil {
					terminalOutputChannel <- fmt.Sprintf("[red]Error: %s[white]", err.Error())
				} else {
					curr_wd, _ = os.Getwd() // Get the new working directory
					terminalOutputChannel <- fmt.Sprintf("[green]Changed directory to:[white] %s", curr_wd)
					pathChannel <- curr_wd  // Update the path in other components
					refreshTreeView <- true // Refresh the tree view
				}
			} else {
				terminalOutputChannel <- "[yellow]Error: No directory specified.[white]"
			}
		case "create":
			if strings.ContainsRune(commandParts[1], '.') {
				os.Create(createFilePath + "/" + commandParts[1])
				refreshTreeView <- true
			} else {
				os.Mkdir(createFilePath + "/" + commandParts[1], 0755)
				refreshTreeView <- true
			}
		case "delete":
			os.Remove(createFilePath + "/" + commandParts[1])
			refreshTreeView <- true
		default:
				
		}
		commandLine.SetText("")

	})

	return commandLine
}



