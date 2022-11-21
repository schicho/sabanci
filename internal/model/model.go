package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/schicho/sabanci/internal/model/components/dashboard"
	"github.com/schicho/sabanci/internal/model/components/login"
	"github.com/schicho/sabanci/internal/model/shutdown"
	"github.com/schicho/sabanci/service"
)

// Model is the bubbletea tea.Model for the Elm architecture
// of the application.
type Model struct {
	isLoggedIn bool
	login      login.Model
	dashboard  dashboard.Model
}

func NewModel() Model {
	m := Model{
		login:     login.NewModel(),
		dashboard: dashboard.NewModel(),
	}
	return m
}

func (m Model) Init() tea.Cmd {
	return m.login.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd, loginCmd, dashCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			cmd = shutdown.ExecuteShutdown()
		}

	case login.LoginSuccess:
		m.isLoggedIn = true

	case shutdown.Shutdown:
		service.SaveCookies()
		return m, tea.Quit
	}

	// run updates of components
	if !m.isLoggedIn {
		m.login, loginCmd = m.login.Update(msg)
	} else {
		m.dashboard, dashCmd = m.dashboard.Update(msg)
	}

	cmds = append(cmds, cmd, loginCmd, dashCmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if !m.isLoggedIn {
		return m.login.View()
	}
	return m.dashboard.View()
}
