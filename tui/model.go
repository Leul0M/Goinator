package tui

import (
	"Goinator/loader"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Styles
	titleStyle     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#6afff3ff"))
	questionStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")) // gold
	answerStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF87")).Bold(true)
	optionStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF")).Bold(true)
	highlightStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4500")).Bold(true)

	// Multi-colored Gofer ASCII
	goferArt = `
    
 _.---~~~~~------,_---~~~~~----._         
  _,,_,*^____      _____*g*\"*, 
 / __/ /'     ^.  /       ^@q   f 
[  @f | @))    |  | @))   l  0 _/  
 \/    /~____ / __ _____/      
  |           _l__l_           I   
  }          [______]           I  ` + lipgloss.NewStyle().Foreground(lipgloss.Color("#00ccffff")).Render("Goinator!") + `
  ]            | | |            |  
  ]             ~ ~             |  
  |                            |   
   |                           |   
 
`
)

type model struct {
	node     *loader.Node
	done     bool
	selected string // "y" or "n"
}

func New(root *loader.Node) tea.Model {
	return &model{node: root}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "y", "Y":
			m.selected = "y"
			if m.node.Yes != nil {
				m.node = m.node.Yes
			}
		case "n", "N":
			m.selected = "n"
			if m.node.No != nil {
				m.node = m.node.No
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}
		if m.node.Entity != nil {
			m.done = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	title := titleStyle.Render("🎯 Goinator - Guessing Game 🎯")
	var body strings.Builder

	// ASCII Gofer art
	body.WriteString(goferArt + "\n")
	body.WriteString(title + "\n")
	body.WriteString(strings.Repeat("=", 30) + "\n\n")

	if m.done {
		body.WriteString(fmt.Sprintf("🟢 I guess: %s\n", answerStyle.Render(m.node.Entity.Name)))
		return body.String()
	}

	// Style the question
	coloredPrefix := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Bold(true).Render("❓ Question:")
	coloredQuestion := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true).Render(m.node.Question)
	body.WriteString(fmt.Sprintf("%s %s\n\n", coloredPrefix, coloredQuestion))

	// Options with dynamic highlighting
	yes := optionStyle.Render("[Y] Yes")
	no := optionStyle.Render("[N] No")
	if m.selected == "y" {
		yes = highlightStyle.Render("[Y] Yes")
	} else if m.selected == "n" {
		no = highlightStyle.Render("[N] No")
	}
	body.WriteString(fmt.Sprintf("Options: %s  %s  (q to quit)\n", yes, no))

	return body.String()
}
