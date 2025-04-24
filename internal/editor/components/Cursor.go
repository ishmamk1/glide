package components

import (
	//"fmt"
	//"os"
	//"strings"
	"sync"
	//"github.com/gdamore/tcell/v2"
)

type Cursor struct {
	X int `default:"1"`
	Y int `default:"1"`
}

var (
    cursor *Cursor
	cursorOnce   sync.Once
)

func NewCursor() *Cursor {
	cursorOnce.Do(func() {
		cursor = &Cursor{
			X:1,
			Y:1,
		}
	})
	return cursor
}

func MoveUp(buffer *Buffer) {
	cursor.X = max(cursor.X-1, 1)
	if cursor.X + 1 > 1 {
		cursor.Y = min(cursor.Y, len(buffer.Lines[cursor.X-1]))
	}
}

func MoveDown(buffer *Buffer, maxHeight int) {
	cursor.X = min(cursor.X+1, maxHeight)
	cursor.Y = min(cursor.Y, len(buffer.Lines[cursor.X-1]))
}

func MoveLeft(buffer *Buffer) {
	cursor.Y = max(1, cursor.Y-1)
}

func MoveRight(buffer *Buffer, maxWidth int) {
	cursor.Y = min(cursor.Y + 1, maxWidth)
}







