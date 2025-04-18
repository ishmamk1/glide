package components

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"github.com/gdamore/tcell/v2"
)

type Buffer struct {
    Lines [][]rune
}

var (
    buffer *Buffer
    once   sync.Once
)

func GetBuffer() *Buffer {
    once.Do(func() {
        buffer = &Buffer{
            Lines: [][]rune{},
        }
    })
    return buffer
}


func LoadFile(buffer *Buffer, filePath string) {
	buffer.Lines = [][]rune{}
	contentBytes, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file: ", err)
	}

	fileContent := string(contentBytes)
	strings.Split(fileContent, "\n")
	fileContentArray := strings.Split(fileContent, "\n")

	for _,line := range fileContentArray {
		runeLine := []rune(line)
		buffer.Lines = append(buffer.Lines, runeLine)
	}
}

func RenderBuffer(screen tcell.Screen, buffer *Buffer) {
	x, y := 1,1

	for _, line := range buffer.Lines {
		for _, runeValue := range line {
			screen.SetContent(x,y,runeValue, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))
			x += 1
		}
		y += 1
		x = 1
	}
}