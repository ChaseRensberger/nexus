package widgets

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TableWidget struct {
	BaseWidget
	table table.Model
}

func NewTableWidget(config WidgetConfig) *TableWidget {
	columns := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "City", Width: 10},
		{Title: "Country", Width: 10},
		{Title: "Population", Width: 10},
	}

	rows := []table.Row{
		{"1", "Tokyo", "Japan", "37,274,000"},
		{"2", "Delhi", "India", "32,065,760"},
		{"3", "Shanghai", "China", "28,516,904"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(config.Position.Height),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return &TableWidget{
		BaseWidget: NewBaseWidget(config),
		table:      t,
	}
}

func (w *TableWidget) Init() tea.Cmd {
	return nil
}

func (w *TableWidget) Update(msg tea.Msg) (Widget, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if w.table.Focused() {
				w.table.Blur()
			} else {
				w.table.Focus()
			}
		case "enter":
			return w, nil
		}
	}
	w.table, cmd = w.table.Update(msg)
	return w, cmd
}

func (w *TableWidget) View() string {
	return w.style.Render(w.table.View() + "\n" + w.table.HelpView())
}

func (w *TableWidget) GetConfig() WidgetConfig {
	return w.config
}
