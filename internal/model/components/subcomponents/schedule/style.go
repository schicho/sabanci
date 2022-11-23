package schedule

import "github.com/charmbracelet/lipgloss"

const (
	blockWidth     = 44
	textWidth      = blockWidth - 2
	blockMaxHeight = 23
)

var (
	white       = lipgloss.Color("#ffffff")
	gray        = lipgloss.Color("#a9a9a9")
	borderStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(white)

	blockStyle = borderStyle.Copy().Width(blockWidth).MaxHeight(blockMaxHeight)

	lineEvenStyle = lipgloss.NewStyle().Foreground(white).Width(textWidth)
	lineOddStyle  = lipgloss.NewStyle().Foreground(gray).Width(textWidth)

	errorColor = lipgloss.Color("1") // terminal color 1 (red)
	errorStyle = borderStyle.Copy().BorderBackground(errorColor).Background(errorColor).Foreground(white)
)
