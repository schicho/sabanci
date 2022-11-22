package schedule

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/schicho/sabanci/data"
	"github.com/schicho/sabanci/internal/model/command"
	"github.com/schicho/sabanci/service"
)

type Model struct {
	schedule *data.Schedule
	currentDay	  int
	err    error
}

func getDaySchedule() int {
	d := int(time.Now().Weekday())

	// Convert to the weekday number used in the schedule
	// 0 = Monday, 6 = Sunday
	d--
	if d < 0 {
		d = 6
	}
	return d
}

func NewModel() Model {
	return Model{currentDay: getDaySchedule()}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg.(type) {

	case command.RetrieveData, command.LoginSuccess:
		m.currentDay = getDaySchedule()
		m.schedule, m.err = service.GetSchedule()
	}

	return m, nil
}

func (m Model) View() string {
	if m.err != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	}

	// wallet might not be initialized, beffore the first call to View()
	if m.schedule == nil {
		return ""
	}

	var sb strings.Builder

	for _, class := range m.schedule.Classes[m.currentDay] {
		sb.WriteString(class.ClassCode)
		sb.WriteRune(' ')
		sb.WriteString(class.Building)
		sb.WriteRune(' ')
		sb.WriteString(class.TimeStart)
		sb.WriteString(" - ")
		sb.WriteString(class.TimeEnd)
		sb.WriteRune('\n')
		sb.WriteString(class.Name)
		sb.WriteRune('\n')
	}
	return blockStyle.Render(sb.String())
}
