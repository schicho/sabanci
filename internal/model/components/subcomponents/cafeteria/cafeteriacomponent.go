package cafeteria

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/schicho/sabanci/data"
	"github.com/schicho/sabanci/internal/model/components/login"
	"github.com/schicho/sabanci/service"
)

type Model struct {
	cafeteria *data.Cafeteria
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
	// On login success, retrieve the cafeteria data.
	case login.LoginSuccess:
		m.cafeteria, m.err = service.GetCafeteria()
	}

	return m, nil
}

func (m Model) View() string {
	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	}

	sb := strings.Builder{}
	for i, food := range m.cafeteria.Menu {

		cal, err := strconv.Atoi(food.Calories)
		if err != nil {
			cal = 0
		}
		name := food.Name
		if len(name) > 30 {
			name = strings.ToValidUTF8(name[:30], "") + "..."
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
