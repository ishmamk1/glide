package components

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
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

func FileExplorer(currentDir string, pathChannel chan string, cliPath chan string) *tview.TreeView { 
	root := tview.NewTreeNode(currentDir).SetColor(tcell.ColorWhite)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	AddFiles(root, currentDir)

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
		}
	})

	return tree
}

