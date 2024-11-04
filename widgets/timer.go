package widgets

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TimerWidget struct {
	BaseWidget
	duration time.Duration
	elapsed  time.Duration
	running  bool
}

type tickMsg time.Time

func NewTimerWidget(config WidgetConfig) *TimerWidget {
	return &TimerWidget{
		BaseWidget: NewBaseWidget(config),
		duration:   5 * time.Minute, // Default 5 min timer
		elapsed:    0,
		running:    false,
	}
}

func (w *TimerWidget) Init() tea.Cmd {
	return nil
}

func (w *TimerWidget) tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (w *TimerWidget) Update(msg tea.Msg) (Widget, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return w, tea.Quit
		case " ": // Space to start/stop
			w.running = !w.running
			if w.running {
				return w, w.tick()
			}
			return w, nil
		case "r": // Reset timer
			w.elapsed = 0
			w.running = false
			return w, nil
		}
	case tickMsg:
		if !w.running {
			return w, nil
		}
		w.elapsed += time.Second
		if w.elapsed >= w.duration {
			w.running = false
			w.elapsed = w.duration
			return w, nil
		}
		return w, w.tick()
	}
	return w, nil
}

func (w *TimerWidget) View() string {
	remaining := w.duration - w.elapsed
	minutes := int(remaining.Minutes())
	seconds := int(remaining.Seconds()) % 60

	status := "⏸"
	if w.running {
		status = "▶"
	}

	s := fmt.Sprintf("%s %02d:%02d\n", status, minutes, seconds)
	s += "\nSpace: Start/Stop | r: Reset | q: Quit\n"

	return w.style.Render(s)
}

func (w *TimerWidget) GetConfig() WidgetConfig {
	return w.config
}
