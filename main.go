package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	tea "github.com/charmbracelet/bubbletea"

	"nexus/widgets"
)

type model struct {
	widgets []widgets.Widget
	config  widgets.Config
}

func initialModel() model {
	// Load config from yaml file
	configData, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var config widgets.Config
	if err := yaml.Unmarshal(configData, &config); err != nil {
		log.Fatal(err)
	}

	// Initialize widgets based on config
	var showWidgets []widgets.Widget
	for name, widgetConfig := range config.Widgets {
		if !widgetConfig.Enabled {
			continue
		}

		switch name {
		case "todo":
			showWidgets = append(showWidgets, widgets.NewTodoWidget(widgetConfig))
		case "timer":
			showWidgets = append(showWidgets, widgets.NewTimerWidget(widgetConfig))
		}
	}

	return model{
		widgets: showWidgets,
		config:  config,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	// Update all widgets
	for i, w := range m.widgets {
		updatedWidget, cmd := w.Update(msg)
		m.widgets[i] = updatedWidget
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var views []string
	for _, w := range m.widgets {
		views = append(views, w.View())
	}
	return strings.Join(views, "\n")
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error has occurred: %v", err)
		os.Exit(1)
	}
}
