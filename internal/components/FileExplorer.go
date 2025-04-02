package components

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"strings"
)

func AddFiles(target *tview.TreeNode, path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		node := tview.NewTreeNode(file.Name()).
			SetReference(filepath.Join(path, file.Name())).
			SetSelectable(true) 
		if file.IsDir() {
			node.SetColor(tcell.ColorGreen)
		}
		target.AddChild(node)
	}
}

// TODO: IMPLEMENT FUNCTION TO PRESERVE FILE EXPANDED STATE LOGIC AFTER REFRESH

/*
func TrackExpandedState(root *tview.TreeNode, path string, currentDir string) {

	pathParts := strings.Split(path[len(currentDir):], "/")
	
	for _, file := range pathParts {
		
	}
}
*/

func FileExplorer(app *tview.Application, currentDir string, pathChannel chan string, cliPath chan string, refreshTreeView chan bool) *tview.TreeView { 
	root := tview.NewTreeNode(currentDir).SetColor(tcell.ColorWhite)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	tree.SetBorder(true)
    tree.SetTitle("File Explorer")
    tree.SetTitleColor(tcell.ColorGreen)
    tree.SetBorderColor(tcell.ColorGreen)


	AddFiles(root, currentDir)


	go func() {
		for range refreshTreeView {
			app.QueueUpdateDraw(func() {
				root.ClearChildren()
				AddFiles(root, currentDir)
			})
		}
		// TrackExpandedState(tree, <-cliPath, currentDir)
		
	}()

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return
		}

		referencePath := reference.(string)

		fileInfo, err := os.Stat(referencePath)
		if err != nil {
			fmt.Println("Error reading file information", err)
		}
		
		if fileInfo.IsDir() {
			children := node.GetChildren()
			if len(children) == 0 {
				path := reference.(string)
				AddFiles(node, path)
			} else {
				node.SetExpanded(!node.IsExpanded())
			}
			cliPath <- referencePath
		} else {
			pathChannel <- referencePath
			endSlash := strings.LastIndex(referencePath, "/")
			cliPath <- referencePath[:endSlash]
		}
	})

	return tree
}

