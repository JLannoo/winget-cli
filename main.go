package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model := CLIModel{
		choices:    []string{},
		cursor:     0,
		selected:   make(map[int]bool),
		wingetList: NewWingetList(),
	}

	model.ClearScreen()

	if err := model.wingetList.FetchFromGist(); err != nil {
		fmt.Println("Error fetching from gist:", err)
		return
	}

	model.choices = model.wingetList.IDs

	p := tea.NewProgram(model)

	m, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		return
	}

	model.ClearScreen()

	err = model.wingetList.RunInstall(m.(CLIModel))
	if err != nil {
		fmt.Println("Error running install:", err)
		return
	}

	fmt.Println("Successfully installed all packages!")
}
