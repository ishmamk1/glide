package components

import (
	"fmt"
	"os"
	"github.com/rivo/tview"
	"bytes"
    "github.com/alecthomas/chroma/v2/formatters"
    "github.com/alecthomas/chroma/v2/lexers"
    "github.com/alecthomas/chroma/v2/styles"
	"github.com/gdamore/tcell/v2"
)

func SyntaxHighlighter(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)

	if err != nil {
        return "", err
    }

	lexer := lexers.Match(filePath)
	if lexer == nil {
		lexer = lexers.Fallback
	}

	formatter := formatters.Get("terminal256")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	style := styles.Get("github-dark")
	if style == nil {
	style = styles.Fallback
	}

	iterator, err := lexer.Tokenise(nil, string(content))


	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = formatter.Format(&buf, style, iterator)

	if err != nil {
		return "", err
	}

	buf_string := buf.String()

	highlightedContent := tview.TranslateANSI(buf_string)

	return highlightedContent, nil
}


func FileViewer(app *tview.Application, pathChannel chan string) *tview.TextView {
	textView := tview.NewTextView().
		SetText("Select a file to view content").
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true).
        SetRegions(true).          
        SetWrap(true)
	
	textView.SetBorder(true)
	textView.SetTitle("File Viewer")
	textView.SetTitleColor(tcell.ColorBlue)
	textView.SetBorderColor(tcell.ColorBlue)
	textView.SetTitleAlign(tview.AlignLeft)

	go func() {
		for filePath := range pathChannel {
			highlightedContent, err := SyntaxHighlighter(filePath)
			if err != nil || highlightedContent == "" {
				highlightedContent = "Error reading file"
			}


			app.QueueUpdateDraw(func() {
				textView.SetText(fmt.Sprintf("[yellow]File:[white] %s\n\n[white]%s", filePath, highlightedContent))
			})
			
		}
	}()

	return textView
}