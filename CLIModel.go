package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type CLIModel struct {
	choices    []string     // items on the to-do list
	cursor     int          // which to-do list item our cursor is pointing at
	selected   map[int]bool // which to-do items are selected
	wingetList *WingetList
}

func (m CLIModel) Init() tea.Cmd {
	return nil
}

func (m CLIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, tea.Quit

		case "down":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}

		case "up":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}

		case " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = true
			}

		case "left":
			m.selected = make(map[int]bool)

		case "right":
			for i := 0; i < len(m.choices); i++ {
				m.selected[i] = true
			}

		case "ctrl+c":
			os.Exit(0)

		default:
			return m, nil
		}

	default:
		return m, nil
	}

	return m, nil
}

func (m CLIModel) View() string {
	message := "Press enter to finish. Use the arrow keys to navigate.\nPress space to select/unselect an item.\nPress left to unselect all items. Press right to select all items.\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		message += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	return message
}

func (m CLIModel) GetSelected() []string {
	ret := []string{}
	for i := 0; i < len(m.choices); i++ {
		if _, ok := m.selected[i]; ok {
			ret = append(ret, m.choices[i])
		}
	}
	return ret
}

func (m CLIModel) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}