package dashboard

import (
	"log"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/schicho/sabanci/internal/model/command"
	"github.com/schicho/sabanci/internal/model/components/subcomponents/cafeteria"
	"github.com/schicho/sabanci/internal/model/components/subcomponents/schedule"
	"github.com/schicho/sabanci/internal/model/components/subcomponents/wallet"
)

// Model is the model of the dashboard component.
// It contains further components, which compose it.
// Model is similar to the tea.Model of the bubbletea package,
// but it is not the same as it does not implement the tea.Model interface.
// As a component, it returns itself in the Update method.
type Model struct {
	help      help.Model
	wallet    wallet.Model
	schedule  schedule.Model
	cafeteria cafeteria.Model
}

func NewModel() Model {
	return Model{
		wallet:    wallet.NewModel(),
		cafeteria: cafeteria.NewModel(),
		schedule:  schedule.NewModel(),
		help:      help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Update is the update method of the dashboard component.
// It is called by the Update method of the tea.Model of the application.
// Note that the Update method of the dashboard component returns itself, and not the tea.Model.
// Thus it does not implement the tea.Model interface.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd, walletCmd, cafeteriaCmd, scheduleCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch msg.String() {
		// We can finally add Q to quit the app.
		case "q":
			cmd = command.ExecuteShutdown()
		case "f5":
			cmd = command.ExecuteRetrieveData()
			log.Println("refreshing data")
		}

	case command.LoginSuccess:
		cmd = command.ExecuteRetrieveData()
	}

	// run updates of components
	m.wallet, walletCmd = m.wallet.Update(msg)
	m.cafeteria, cafeteriaCmd = m.cafeteria.Update(msg)
	m.schedule, scheduleCmd = m.schedule.Update(msg)

	cmds = append(cmds, cmd, walletCmd, cafeteriaCmd, scheduleCmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	w := m.wallet.View()
	c := m.cafeteria.View()
	s := m.schedule.View()
	h := m.help.View(keys)

	// join the views of the components
	ws := lipgloss.JoinVertical(lipgloss.Top, w, s)
	wsc := lipgloss.JoinHorizontal(lipgloss.Top, ws, c)

	dashboard := lipgloss.JoinVertical(lipgloss.Top, wsc, h)

	return dashboard
}
