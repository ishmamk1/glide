package components

import (
	//"fmt"
	//"os"
	//"strings"
	"sync"
	//"github.com/gdamore/tcell/v2"
)

type Cursor struct {
	Row int `default:"1"`
	Column int `default:"1"`
}

var (
    cursor *Cursor
	cursorOnce   sync.Once
)

func NewCursor() *Cursor {
	cursorOnce.Do(func() {
		cursor = &Cursor{
			Row:1,
			Column:1,
		}
	})
	return cursor
}

func MoveUp() {
	cursor.Column = max(cursor.Column-1, 1)
}

func MoveDown(maxHeight int) {
	cursor.Column = min(cursor.Column+1, maxHeight)
}

func MoveLeft() {
	cursor.Row = max(1, cursor.Row-1)
}

func MoveRight(maxWidth int) {
	cursor.Row = min(cursor.Row + 1, maxWidth)
}







