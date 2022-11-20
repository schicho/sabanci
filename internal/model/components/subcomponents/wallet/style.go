package wallet

import "github.com/charmbracelet/lipgloss"

var (
	white       = lipgloss.Color("#ffffff")
	borderStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(white)

	// terminal color 6 (light blue)
	shuttleColor = lipgloss.Color("6")
	shuttleStyle = borderStyle.Copy().BorderBackground(shuttleColor).Background(shuttleColor).Foreground(white)
	// terminal color 2 (green)
	cafeteriaColor = lipgloss.Color("2")
	cafeteriaStyle = borderStyle.Copy().BorderBackground(cafeteriaColor).Background(cafeteriaColor).Foreground(white)
	// terminal color 9 (light red)
	printColor = lipgloss.Color("9")
	printStyle = borderStyle.Copy().BorderBackground(printColor).Background(printColor).Foreground(white)

	// terminal color 1 (red)
	errorColor = lipgloss.Color("1")
	errorStyle = borderStyle.Copy().BorderBackground(errorColor).Background(errorColor).Foreground(white)
)
