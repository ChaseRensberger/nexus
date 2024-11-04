package widgets

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Position struct {
	Top    int `yaml:"top"`
	Left   int `yaml:"left"`
	Height int `yaml:"height"`
	Width  int `yaml:"width"`
}

type Colors struct {
	Label string `yaml:"label"`
	Text  string `yaml:"text"`
}

type WidgetConfig struct {
	Enabled         bool     `yaml:"enabled"`
	Position        Position `yaml:"position"`
	Colors          Colors   `yaml:"colors"`
	RefreshInterval int      `yaml:"refresh_interval"`
}

type Config struct {
	Widgets map[string]WidgetConfig `yaml:"widgets"`
}

type Widget interface {
	Init() tea.Cmd
	Update(tea.Msg) (Widget, tea.Cmd)
	View() string
	GetConfig() WidgetConfig
}

type BaseWidget struct {
	config WidgetConfig
	style  lipgloss.Style
}

func NewBaseWidget(config WidgetConfig) BaseWidget {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color(config.Colors.Text)).
		MarginLeft(config.Position.Left).
		MarginTop(config.Position.Top).
		Width(config.Position.Width).
		Height(config.Position.Height)

	return BaseWidget{
		config: config,
		style:  style,
	}
}
