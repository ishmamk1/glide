package components

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
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
				newPath := filepath.Join(createFilePath, commandParts[1])
				err := os.Chdir(newPath)
				if err != nil {
					terminalOutputChannel <- fmt.Sprintf("[red]Error: %s[white]", err.Error())
				} else {
					curr_wd, _ = os.Getwd() 
					terminalOutputChannel <- fmt.Sprintf("[green]Changed directory to:[green] %s", curr_wd)
					pathChannel <- curr_wd 
					refreshTreeView <- true
				}
			} else {
				terminalOutputChannel <- "[yellow]Error: No directory specified.[white]"
			}
		case "create":
			if len(commandParts) > 1 {
                newPath := filepath.Join(createFilePath, commandParts[1])
                if strings.ContainsRune(commandParts[1], '.') {
                    if _, err := os.Create(newPath); err != nil {
                        terminalOutputChannel <- fmt.Sprintf("[red]Error creating file: %s[white]", err)
                    } else {
                        terminalOutputChannel <- fmt.Sprintf("[green]Created file:[white] %s", commandParts[1])
                        refreshTreeView <- true
                    }
                } else {
                    if err := os.Mkdir(newPath, 0755); err != nil {
                        terminalOutputChannel <- fmt.Sprintf("[red]Error creating directory: %s[white]", err)
                    } else {
                        terminalOutputChannel <- fmt.Sprintf("[green]Created directory:[white] %s", commandParts[1])
                        refreshTreeView <- true
                    }
                }
            }
		case "delete":
			if len(commandParts) > 1 {
                targetPath := filepath.Join(createFilePath, commandParts[1])
                if err := os.Remove(targetPath); err != nil {
                    terminalOutputChannel <- fmt.Sprintf("[red]Error deleting: %s[white]", err)
                } else {
                    terminalOutputChannel <- fmt.Sprintf("[red]Deleted:[white] %s", commandParts[1])
                    refreshTreeView <- true
                }
            }
		case "ls":
			current_path, err := os.Getwd()
			if err != nil {
				terminalOutputChannel <- fmt.Sprintf("[red]Error obtaining current directory: %s[white]", err)
				break
			}
			
			files, err := os.ReadDir(current_path)
			if err != nil {
				terminalOutputChannel <- fmt.Sprintf("[red]Error listing content: %s[white]", err)
				break
			}
			
			var files_list []string
			for _, file := range files {
				if file.IsDir() {
					
					files_list = append(files_list, fmt.Sprintf("[green]%s[white]", file.Name()))
				} else {
					files_list = append(files_list, file.Name())
				}
    		}
    		terminalOutputChannel <- fmt.Sprintf("[green]Content:[white] %s", strings.Join(files_list, " "))
		default:
			terminalOutputChannel <- fmt.Sprintf("[yellow]Unknown command:[white] %s", commandParts[0])		
		}
		commandLine.SetText("")

	})

	return commandLine
}



