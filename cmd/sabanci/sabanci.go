package main

import (
	"io"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/schicho/sabanci/internal/model"

	// initialize the service and web package
	_ "github.com/schicho/sabanci/service"
)

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
	} else {
		log.SetOutput(io.Discard)
	}

	p := tea.NewProgram(model.NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalln(err)
	}
}
