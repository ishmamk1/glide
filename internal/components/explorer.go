package components

import (
	"fmt"
	"os"
	"strings"
	"github.com/eiannone/keyboard"
)

func HandleFileNavigation(currentDir string) {
	for {
		files, err := os.ReadDir(currentDir) 
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}

		selectedFileIndex := 0 

		for {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Current Directory:", currentDir)

			for index, file := range files { 
				if index == selectedFileIndex {
					fmt.Println("→", file.Name()) 
				} else {
					fmt.Println("  ", file.Name()) 
				}
			}
			fmt.Println("\nPress 'q' to quit, '←' to go back")

			char, key, err := keyboard.GetKey()
			if err != nil {
				fmt.Println("Keyboard error:", err)
				return
			}

			if key == keyboard.KeyArrowUp && selectedFileIndex > 0 {
				selectedFileIndex--
			} else if key == keyboard.KeyArrowDown && selectedFileIndex < len(files)-1 {
				selectedFileIndex++
			} else if key == keyboard.KeyEnter {
				selectedFile := files[selectedFileIndex]
				if selectedFile.IsDir() {
					currentDir = currentDir + "/" + selectedFile.Name()
					break
				}
			} else if key == keyboard.KeyArrowLeft {
				if currentDir != "/" {
					parentDir := currentDir[:len(currentDir)-len(currentDir[strings.LastIndex(currentDir, "/"):])]
					if parentDir == "" {
						parentDir = "/"
					}
					currentDir = parentDir
					break
				}
			} else if char == 'q' { 
				return
			}
		}
	}
}
