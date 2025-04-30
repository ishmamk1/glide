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
    bufferOnce   sync.Once
)

func GetBuffer() *Buffer {
    bufferOnce.Do(func() {
        buffer = &Buffer{
            Lines: [][]rune{[]rune{}},
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
		runeLine = append(runeLine, ' ')
		buffer.Lines = append(buffer.Lines, runeLine)
	}
}

func RenderBuffer(screen tcell.Screen, buffer *Buffer) {
	x, y := 1,1
	width, _ := screen.Size()
	for _, line := range buffer.Lines {
		for _, runeValue := range line {
			if x < width - 1 {
				screen.SetContent(x,y,runeValue, nil, tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack))
				x += 1
			}
		}
		y += 1
		x = 1
	}
}

func InsertRuneAt(x int, y int, r rune) {
	buffer.Lines[x] = append(buffer.Lines[x][:y], append([]rune{r}, buffer.Lines[x][y:]...)...)
}

func DeleteRuneAt(x int, y int) bool {
	if x == 0 && y == 0 {
		return false
	}

	if y > 0 {
		line := buffer.Lines[x]
		if y <= len(line) {
			buffer.Lines[x] = append(line[:y-1], line[y:]...)
		}
		return false
	} else if x > 0 {
		prevLine := buffer.Lines[x-1]
		currLine := buffer.Lines[x]

		if len(prevLine) > 0 && prevLine[len(prevLine)-1] == ' ' {
			prevLine = prevLine[:len(prevLine)-1]
		}

		joined := append(prevLine, currLine...)

		if len(joined) == 0 || joined[len(joined)-1] != ' ' {
			joined = append(joined, ' ')
		}

		buffer.Lines[x-1] = joined
		buffer.Lines = append(buffer.Lines[:x], buffer.Lines[x+1:]...)
		return true
	}
	return false
}

func SplitLine(x int, y int) {
	if x < 0 || x >= len(buffer.Lines) {
		return 
	}

	currentLine := buffer.Lines[x]

	if y < 0 || y > len(currentLine) {
		return 
	}

	firstPart := append([]rune{}, currentLine[:y]...)
	secondPart := append([]rune{}, currentLine[y:]...)

	if len(firstPart) == 0 || firstPart[len(firstPart)-1] != ' ' {
		firstPart = append(firstPart, ' ')
	}
	if len(secondPart) == 0 || secondPart[len(secondPart)-1] != ' ' {
		secondPart = append(secondPart, ' ')
	}

	buffer.Lines = append(buffer.Lines[:x+1], append([][]rune{secondPart}, buffer.Lines[x+1:]...)...)
	buffer.Lines[x] = firstPart
}


