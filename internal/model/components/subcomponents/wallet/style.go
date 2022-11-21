package wallet

import "github.com/charmbracelet/lipgloss"

var (
	borderStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder())

	// terminal color 6 (light blue)
	shuttleColor = lipgloss.Color("6")
	shuttleStyle = borderStyle.Copy().Foreground(shuttleColor).BorderForeground(shuttleColor)
	// terminal color 2 (green)
	cafeteriaColor = lipgloss.Color("2")
	cafeteriaStyle = borderStyle.Copy().Foreground(cafeteriaColor).BorderForeground(cafeteriaColor)
	// terminal color 9 (light red)
	printColor = lipgloss.Color("9")
	printStyle = borderStyle.Copy().Foreground(printColor).BorderForeground(printColor)

	// terminal color 1 (red)
	errorColor = lipgloss.Color("1")
	white       = lipgloss.Color("#ffffff")
	errorStyle = borderStyle.Copy().BorderBackground(errorColor).Background(errorColor).Foreground(white)
)
