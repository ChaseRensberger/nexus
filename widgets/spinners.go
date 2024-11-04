package widgets

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Available spinners
	spinners = []spinner.Spinner{
		spinner.Line,
		spinner.Dot,
		spinner.MiniDot,
		spinner.Jump,
		spinner.Pulse,
		spinner.Points,
		spinner.Globe,
		spinner.Moon,
		spinner.Monkey,
	}

	textStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
)

type SpinnerWidget struct {
	BaseWidget
	index   int
	spinner spinner.Model
}

func NewSpinnerWidget(config WidgetConfig) *SpinnerWidget {
	return &SpinnerWidget{
		BaseWidget: NewBaseWidget(config),
	}
}

func (w *SpinnerWidget) Init() tea.Cmd {
	return w.spinner.Tick
}

func (w *SpinnerWidget) Update(msg tea.Msg) (Widget, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return w, tea.Quit
		case "h", "left":
			w.index--
			if w.index < 0 {
				w.index = len(spinners) - 1
			}
			w.resetSpinner()
			return w, w.spinner.Tick
		case "l", "right":
			w.index++
			if w.index >= len(spinners) {
				w.index = 0
			}
			w.resetSpinner()
			return w, w.spinner.Tick
		default:
			return w, nil
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		w.spinner, cmd = w.spinner.Update(msg)
		return w, cmd
	default:
		return w, nil
	}
}

func (w *SpinnerWidget) resetSpinner() {
	w.spinner = spinner.New()
	w.spinner.Style = spinnerStyle
	w.spinner.Spinner = spinners[w.index]
}

func (w *SpinnerWidget) View() (s string) {
	var gap string
	switch w.index {
	case 1:
		gap = ""
	default:
		gap = " "
	}

	s += fmt.Sprintf("\n %s%s%s\n\n", w.spinner.View(), gap, textStyle("Spinning..."))
	s += helpStyle("h/l, ←/→: change spinner • q: exit\n")
	return
}

func (w *SpinnerWidget) GetConfig() WidgetConfig {
	return w.config
}
