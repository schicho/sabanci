package schedule

import "github.com/charmbracelet/lipgloss"

var (
	white       = lipgloss.Color("#ffffff")
	borderStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(white)

	blockStyle = borderStyle.Copy().Width(44)

	errorColor = lipgloss.Color("1") // terminal color 1 (red)
	errorStyle = borderStyle.Copy().BorderBackground(errorColor).Background(errorColor).Foreground(white)
)
