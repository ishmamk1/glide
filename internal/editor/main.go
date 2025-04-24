package main

import (
	"log"
	"github.com/gdamore/tcell/v2"
	"glide/internal/editor/components"
)

func CreateDynamicBox(screen tcell.Screen, text string, x, y, width, height int) {
	for i := 0; i < width; i++ {
		screen.SetContent(x+i, y, tcell.RuneHLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))
		screen.SetContent(x+i, y+height-1, tcell.RuneHLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))
	}
	for i := 0; i < height; i++ {
		screen.SetContent(x, y+i, tcell.RuneVLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))
		screen.SetContent(x+width-1, y+i, tcell.RuneVLine, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))
	}

	screen.SetContent(x, y, tcell.RuneULCorner, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))
	screen.SetContent(x+width-1, y, tcell.RuneURCorner, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))
	screen.SetContent(x, y+height-1, tcell.RuneLLCorner, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))
	screen.SetContent(x+width-1, y+height-1, tcell.RuneLRCorner, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))

	textX := x + 1
	textY := y + 1
	for i, r := range text {
		if textX+i < x+width-1 {
			screen.SetContent(textX+i, textY, r, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))
		}
	}
}

func main() {
    // Initialize screen
    screen, err := tcell.NewScreen()
    if err != nil {
        log.Fatal(err)
    }
    if err := screen.Init(); err != nil {
        log.Fatal(err)
    }
    defer screen.Fini()

    buffer := components.GetBuffer()
	cursor := components.NewCursor()
    
    components.LoadFile(buffer, "/Users/ishmam/glide/internal/editor/components/sample.txt")

    screen.SetStyle(tcell.StyleDefault.
        Background(tcell.ColorBlack).
        Foreground(tcell.ColorWhite))

    for {
        screen.Clear()
        width, height := screen.Size()
        
        CreateDynamicBox(screen, "", 0, 0, width, height)
        
        components.RenderBuffer(screen, buffer)
        
        screen.Show()
        
        switch ev := screen.PollEvent().(type) {
        case *tcell.EventKey:
            if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
                return
            }
			switch ev.Key() {
			case tcell.KeyDown:
				components.MoveDown(buffer, len(buffer.Lines))
			case tcell.KeyUp:
				components.MoveUp(buffer)
			case tcell.KeyLeft:
				components.MoveLeft(buffer,)
			case tcell.KeyRight:
				components.MoveRight(buffer, len(buffer.Lines[cursor.X - 1]))
			case tcell.KeyDelete, tcell.KeyBackspace, tcell.KeyBackspace2:
				components.DeleteRuneAt(cursor.X-1, cursor.Y-1)
				components.MoveLeft(buffer)
			case tcell.KeyRune:
				r := ev.Rune()
				if r >= 32 && r < 127 { 
					components.InsertRuneAt(cursor.X-1, cursor.Y-1, r)
					components.MoveRight(buffer, width)
				}
			}
        case *tcell.EventResize:
            screen.Sync()
        }
		screen.ShowCursor(cursor.Y, cursor.X)
    }
}