package cafeteria

import "github.com/charmbracelet/lipgloss"

const (
	blockWidth     = 32
	textWidth      = blockWidth - 2
	blockMaxHeight = 23
)

var (
	white       = lipgloss.Color("#ffffff")
	borderStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(white).Width(blockWidth).MaxHeight(blockMaxHeight)

	lowCal  = lipgloss.Color("2") // terminal color 2 (green)
	midCal  = lipgloss.Color("3") // terminal color 3 (yellow)
	highCal = lipgloss.Color("9") // terminal color 9 (light red)

	lowCalStyle  = lipgloss.NewStyle().Foreground(lowCal)
	midCalStyle  = lipgloss.NewStyle().Foreground(midCal)
	highCalStyle = lipgloss.NewStyle().Foreground(highCal)

	errorColor = lipgloss.Color("1") // terminal color 1 (red)
	errorStyle = borderStyle.Copy().BorderBackground(errorColor).Background(errorColor).Foreground(white)
)
