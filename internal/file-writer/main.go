package main

import (
    "github.com/gdamore/tcell/v2"
    "os"
    "log"
)

func drawBorder(s tcell.Screen) {
    width, height := s.Size()
    borderStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)

    // Draw top and bottom borders
    for x := 0; x < width; x++ {
        s.SetContent(x, 0, '─', nil, borderStyle)         // top
        s.SetContent(x, height-1, '─', nil, borderStyle)  // bottom
    }

    // Draw left and right borders
    for y := 0; y < height; y++ {
        s.SetContent(0, y, '│', nil, borderStyle)         // left
        s.SetContent(width-1, y, '│', nil, borderStyle)   // right
    }

    // Draw corners
    s.SetContent(0, 0, '┌', nil, borderStyle)            // top-left
    s.SetContent(width-1, 0, '┐', nil, borderStyle)      // top-right
    s.SetContent(0, height-1, '└', nil, borderStyle)     // bottom-left
    s.SetContent(width-1, height-1, '┘', nil, borderStyle) // bottom-right
}

func main() {
    // Create a new screen
    screen, err := tcell.NewScreen()
    if err != nil {
        log.Fatalf("Failed to create screen: %v", err)
        os.Exit(1)
    }
    if err := screen.Init(); err != nil {
        log.Fatalf("Failed to initialize screen: %v", err)
        os.Exit(1)
    }
    defer screen.Fini()

    // Set default style
    defStyle := tcell.StyleDefault.
        Background(tcell.ColorBlack).
        Foreground(tcell.ColorWhite)
    screen.SetStyle(defStyle)

    // Clear the screen
    screen.Clear()

    // Event loop
    quit := make(chan struct{})
    go func() {
        for {
            switch ev := screen.PollEvent().(type) {
            case *tcell.EventKey:
                if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
                    close(quit)
                    return
                }
            case *tcell.EventResize:
                screen.Sync()
            }
        }
    }()

    // Main loop
    for {
        select {
        case <-quit:
            return
        default:
			drawBorder(screen)
            screen.Show()
        }
    }
}
