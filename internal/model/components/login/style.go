package login

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	highlightColor = lipgloss.Color("004")

	centeredStyle = lipgloss.NewStyle()

	focusedStyle = centeredStyle.Copy().Foreground(highlightColor)
	blurredStyle = centeredStyle.Copy().Foreground(lipgloss.Color("240"))
	cursorStyle  = focusedStyle.Copy()

	titleStyle = focusedStyle.Copy().Bold(true).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(highlightColor)

	focusedButton = focusedStyle.Copy().Render("[ Login ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Login"))
)
