package wallet

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/schicho/sabanci/data"
	"github.com/schicho/sabanci/internal/model/command"
	"github.com/schicho/sabanci/service"
)

type Model struct {
	wallet *data.Wallet
	err    error
}

func NewModel() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg.(type) {

	case command.RetrieveData, command.LoginSuccess:
		m.wallet, m.err = service.GetWallet()
	}

	return m, nil
}

func (m Model) View() string {
	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	}

	// wallet might not be initialized, beffore the first call to View()
	if m.wallet == nil {
		return ""
	}

	shuttle := shuttleStyle.Render(fmt.Sprintf("Shuttle: %v", m.wallet.Shuttle))
	cafeteria := cafeteriaStyle.Render(fmt.Sprintf("Cafeteria: %v", m.wallet.Meal))
	print := printStyle.Render(fmt.Sprintf("Print: %v", m.wallet.Print))

	return lipgloss.JoinHorizontal(lipgloss.Center, shuttle, cafeteria, print)
}
