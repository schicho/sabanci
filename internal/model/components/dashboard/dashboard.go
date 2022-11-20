package dashboard

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/schicho/sabanci/internal/model/components/subcomponents/wallet"
	"github.com/schicho/sabanci/internal/model/shutdown"
)

// Model is the model of the dashboard component.
// It contains further components, which compose it.
// Model is similar to the tea.Model of the bubbletea package,
// but it is not the same as it does not implement the tea.Model interface.
// As a component, it returns itself in the Update method.
type Model struct {
	help   help.Model
	wallet wallet.Model
}

func NewModel() Model {
	return Model{
		wallet: wallet.NewModel(),
		help:   help.New(),
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
	var cmd, walletCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch msg.String() {
		// We can finally add Q to quit the app.
		case "q":
			cmd = shutdown.ExecuteShutdown()
		}
	}

	// run updates of components
	m.wallet, walletCmd = m.wallet.Update(msg)

	cmds = append(cmds, cmd, walletCmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := m.wallet.View()

	helpView := m.help.View(keys)

	return s + strings.Repeat("\n", 8) + helpView
}
