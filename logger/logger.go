package logger

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func startupLogger() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}
}
