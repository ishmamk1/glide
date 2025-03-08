package main

import (
	"fmt"
	"os"
	"glide/internal/components" 
	"github.com/eiannone/keyboard"
)

func main() {
	currentDir, _ := os.Getwd()
	err := keyboard.Open()
	if err != nil {
		fmt.Println("Failed to open keyboard listener:", err)
		return
	}
	defer keyboard.Close()

	components.HandleFileNavigation(currentDir)
}
