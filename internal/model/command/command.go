package command

import tea "github.com/charmbracelet/bubbletea"

type LoginSuccess struct{}

func ExecuteLoginSuccess() tea.Cmd {
	return func() tea.Msg {
		return LoginSuccess{}
	}
}

type Shutdown struct{}

// ExecuteShutdown defines a custom command for bubbletea,
// which allows us to define a shutdown hook.
// We call ExecuteShutdown on commands like "q", "^C" or "ESC"
// and handle it centralized in the main model.
func ExecuteShutdown() tea.Cmd {
	return func() tea.Msg {
		return Shutdown{}
	}
}
