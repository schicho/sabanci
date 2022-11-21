package login

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/schicho/sabanci/internal/model/command"
	"github.com/schicho/sabanci/service"
)

type Model struct {
	loginRetry bool
	focusIndex int
	inputs     []textinput.Model
}

func NewModel() Model {
	m := Model{
		inputs: make([]textinput.Model, 2),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 256

		switch i {
		case 0:
			t.Placeholder = "Username"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '*'
		}
		m.inputs[i] = t
	}
	return m
}

func (m Model) Init() tea.Cmd {
	err := service.Login("", "")
	if err != nil {
		log.Println("Session could not be restored. Login required.")
		return textinput.Blink
	}
	log.Println("Login Success: Session reestablished")
	return tea.Batch(command.ExecuteLoginSuccess(), textinput.Blink)
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "down", "shift+tab", "up", "enter":
			s := msg.String()

			// attempt to login, when enter is pressed on the input field.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				username := m.inputs[0].Value()
				password := m.inputs[1].Value()
				err := service.Login(username, password)
				if err != nil {
					log.Printf("Login failed for user %v: %v", username, err)
					m.loginRetry = true
					return m, nil
				}
				log.Printf("Login success for user %v", username)
				m.loginRetry = false
				return m, command.ExecuteLoginSuccess()
			}

			// cycle through the inputs
			if s == "tab" || s == "down" {
				m.focusIndex++
			} else if s == "shift+tab" || s == "up" {
				m.focusIndex--
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}
			return m.updateFocus()
		}
	}
	cmd := m.updateInputs(msg)
	return m, cmd
}

func (m Model) updateFocus() (Model, tea.Cmd) {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		if i == m.focusIndex {
			cmds[i] = m.inputs[i].Focus()
			m.inputs[i].PromptStyle = focusedStyle
			m.inputs[i].TextStyle = focusedStyle
		} else {
			m.inputs[i].Blur()
			m.inputs[i].PromptStyle = blurredStyle
			m.inputs[i].TextStyle = blurredStyle
		}
	}
	return m, tea.Batch(cmds...)
}

func (m Model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return tea.Batch(cmds...)
}

func (m Model) View() string {
	var b strings.Builder
	b.WriteString(titleStyle.Render("Login to mySU"))
	b.WriteRune('\n')
	b.WriteRune('\n')

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	if m.loginRetry {
		fmt.Fprintf(&b, "\n%s\n\n", blurredStyle.Render("Login failed. Please try again."))
	}

	return b.String()
}
