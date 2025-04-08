package components

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
    "strings"
)

func TerminalView(app *tview.Application, terminalOutputChannel chan string) *tview.TextView {
    terminalView := tview.NewTextView().
        SetDynamicColors(true).
        SetRegions(true).
        SetScrollable(true).
        SetWrap(false).
        SetTextAlign(tview.AlignLeft)

    terminalView.SetBorder(true).
        SetTitle("Terminal View").
        SetTitleColor(tcell.ColorRed).
        SetBorderColor(tcell.ColorRed).
        SetTitleAlign(tview.AlignLeft).
        SetBackgroundColor(tcell.ColorBlack)

    go func() {
        var messages []string
        for message := range terminalOutputChannel {
            app.QueueUpdateDraw(func() {
                messages = append(messages, message)

                coloredText := strings.Join(messages, "\n")

                terminalView.Clear()
                terminalView.SetText(coloredText)
                terminalView.ScrollToEnd()
            })
        }
    }()

    return terminalView
}
