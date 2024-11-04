package widgets

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type TodoWidget struct {
	BaseWidget
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func NewTodoWidget(config WidgetConfig) *TodoWidget {
	return &TodoWidget{
		BaseWidget: NewBaseWidget(config),
		choices:    []string{"Financial Report", "Buy groceries", "Make root canal appt"},
		selected:   make(map[int]struct{}),
	}
}

func (w *TodoWidget) Init() tea.Cmd {
	return nil
}

func (w *TodoWidget) Update(msg tea.Msg) (Widget, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return w, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if w.cursor > 0 {
				w.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if w.cursor < len(w.choices)-1 {
				w.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := w.selected[w.cursor]
			if ok {
				delete(w.selected, w.cursor)
			} else {
				w.selected[w.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return w, nil
}

func (w *TodoWidget) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range w.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if w.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := w.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func (w *TodoWidget) GetConfig() WidgetConfig {
	return w.config
}
