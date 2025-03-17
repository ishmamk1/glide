package components

import (
	"fmt"
	"os"
	//"path/filepath"
	//"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func FileViewer(app *tview.Application, pathChannel chan string) *tview.TextView {
	textView := tview.NewTextView().
		SetText("Select a file to view content").
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true)


	go func() {
		for filePath := range pathChannel {
			content, err := os.ReadFile(filePath)
			if err != nil {
				content = []byte("Error reading file")
			}

			app.QueueUpdateDraw(func() {
                textView.SetText(fmt.Sprintf("[yellow]File:[white] %s\n\n[white]%s", filePath, string(content)))
            })
			
		}
	}()

	return textView
}