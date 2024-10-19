package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jpxcz/sqlterm/databases"
	"github.com/jpxcz/sqlterm/tui"
)

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

    databases.NewDatabases()

	p := tui.NewTeaProgram()
	if _, err := p.Run(); err != nil {
		log.Fatal(err)

	}
}
