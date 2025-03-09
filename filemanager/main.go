package main

import (
	"fmt"
	"glide/internal/components"
	"os"
	//"github.com/eiannone/keyboard"
	"github.com/rivo/tview"
)

func main() {
	currentDir, _ := os.Getwd()

	app := tview.NewApplication()

	treeView, referencePath := components.FileExplorer(currentDir)
	
	if referencePath != "" {
		fmt.Println(referencePath)
	} else {
		fmt.Println("no path")
	}
	flex := tview.NewFlex().
		AddItem(treeView, 40, 1, true). 
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right Panel"), 0, 2, false)

	app.SetFocus(treeView)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

