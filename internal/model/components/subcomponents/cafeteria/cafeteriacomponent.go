package cafeteria

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/schicho/sabanci/data"
	"github.com/schicho/sabanci/internal/model/command"
	"github.com/schicho/sabanci/service"
)

type Model struct {
	cafeteria *data.Cafeteria
	err       error
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
		m.cafeteria, m.err = service.GetCafeteria()
	}

	return m, nil
}

func (m Model) View() string {
	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	}

	// cafeteria might not be initialized, beffore the first call to View()
	if m.cafeteria == nil {
		return ""
	}

	sb := strings.Builder{}
	for i, food := range m.cafeteria.Menu {

		cal, err := strconv.Atoi(food.Calories)
		if err != nil {
			cal = 0
		}
		name := food.Name
		if len(name) > textWidth {
			name = strings.ToValidUTF8(name[:textWidth - 2], "") + "..."
		}

		switch {
		case cal < 200:
			sb.WriteString(lowCalStyle.Render(name))
		case cal < 300:
			sb.WriteString(midCalStyle.Render(name))
		default:
			sb.WriteString(highCalStyle.Render(name))
		}
		if i < len(m.cafeteria.Menu)-1 {
			sb.WriteRune('\n')
		}
	}
	return borderStyle.Render(sb.String())
}
