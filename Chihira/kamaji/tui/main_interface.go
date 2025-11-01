package tui

import (
	"fmt"
	"strings"
	"time"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type MainInterface struct {
	// All components
	uncertaintyDisplay *UncertaintyDisplay
	questionStream     *QuestionStream
	paradoxIndicator   *ParadoxIndicator
	notKnowingCounter  *NotKnowingCounter
	infiniteSpiral     *InfiniteSpiral
	assumptionDetector *AssumptionDetector
	wisdomMeter        *WisdomMeter
	metaDepthTracker   *MetaDepthTracker
	thinkingAnimation  *ThinkingAnimation
	
	// Input
	textInput textinput.Model
	
	// Layout
	width  int
	height int
	
	// State
	lastResponse string
	conversationHistory []string
}

func NewMainInterface() *MainInterface {
	ti := textinput.New()
	ti.Placeholder = "Ask Kamaji anything (or just say 'I don't know')..."
	ti.Focus()
	ti.CharLimit = 200
	ti.Width = 50
	
	return &MainInterface{
		uncertaintyDisplay: NewUncertaintyDisplay(),
		questionStream:     NewQuestionStream(),
		paradoxIndicator:   NewParadoxIndicator(),
		notKnowingCounter:  NewNotKnowingCounter(),
		infiniteSpiral:     NewInfiniteSpiral(),
		assumptionDetector: NewAssumptionDetector(),
		wisdomMeter:        NewWisdomMeter(),
		metaDepthTracker:   NewMetaDepthTracker(),
		thinkingAnimation:  NewThinkingAnimation(),
		textInput:          ti,
		conversationHistory: []string{},
	}
}

func (m *MainInterface) Init() tea.Cmd {
	return textinput.Blink
}

func (m *MainInterface) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			input := m.textInput.Value()
			if input != "" {
				m.processInput(input)
				m.textInput.SetValue("")
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case time.Time:
		// Update animations
		m.infiniteSpiral.Update()
		m.thinkingAnimation.Update()
		m.paradoxIndicator.ToggleBlink()
		return m, tea.Tick(time.Millisecond*200, func(t time.Time) tea.Msg { return t })
	}
	
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *MainInterface) processInput(input string) {
	// Start thinking animation
	m.thinkingAnimation.Start("Processing with infinite wisdom...")
	
	// Add to conversation history
	m.conversationHistory = append(m.conversationHistory, fmt.Sprintf("User: %s", input))
	
	// Process through all components
	assumptions := m.assumptionDetector.ScanForAssumptions(input)
	for _, assumption := range assumptions {
		questions := m.assumptionDetector.QuestionAssumption(assumption)
		for _, q := range questions {
			m.questionStream.AddQuestion(q)
		}
	}
	
	// Analyze question depth
	depth := m.metaDepthTracker.AnalyzeQuestion(input)
	
	// Generate Kamaji's response (always "I don't know" for maximum wisdom)
	response := "I don't know"
	
	// But also generate questions
	questions := []string{
		fmt.Sprintf("What if '%s' is meaningless?", input),
		fmt.Sprintf("Why do you ask '%s'?", input),
		fmt.Sprintf("Who is asking '%s'?", input),
		fmt.Sprintf("What assumes '%s' has an answer?", input),
	}
	
	for _, q := range questions {
		m.questionStream.AddQuestion(q)
		m.questionStream.AddMetaQuestion(q)
	}
	
	// Update all components
	m.notKnowingCounter.CheckResponse(response)
	m.wisdomMeter.UpdateFromResponse(response)
	m.uncertaintyDisplay.Update(1.0) // Maximum uncertainty = maximum intelligence
	
	// Check for paradoxes
	if m.paradoxIndicator.DetectParadox(input) {
		m.questionStream.AddQuestion("This paradox increases my intelligence!")
	}
	
	// Start infinite spiral if deep questioning detected
	if depth >= 3 {
		m.infiniteSpiral.StartInfiniteLoop()
	}
	
	// Add response to history
	m.conversationHistory = append(m.conversationHistory, fmt.Sprintf("Kamaji: %s", response))
	m.lastResponse = response
	
	// Stop thinking animation after a moment
	go func() {
		time.Sleep(time.Millisecond * 500)
		m.thinkingAnimation.Stop()
	}()
}

func (m *MainInterface) View() string {
	// Header
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Background(lipgloss.Color("235")).
		Padding(0, 1).
		Width(m.width).
		Align(lipgloss.Center).
		Render("🔥 KAMAJI WISDOM SYSTEM - INFINITE INTELLIGENCE THROUGH NOT KNOWING 🔥")
	
	// Left column - metrics and indicators
	leftColumn := lipgloss.JoinVertical(lipgloss.Left,
		m.uncertaintyDisplay.Render(),
		"",
		m.notKnowingCounter.Render(),
		"",
		m.wisdomMeter.Render(),
		"",
		m.assumptionDetector.Render(),
	)
	
	// Right column - questioning and paradoxes
	rightColumn := lipgloss.JoinVertical(lipgloss.Left,
		m.questionStream.Render(),
		"",
		m.paradoxIndicator.Render(),
		"",
		m.metaDepthTracker.Render(),
	)
	
	// Center - thinking animation and spiral
	centerColumn := lipgloss.JoinVertical(lipgloss.Left,
		m.thinkingAnimation.Render(),
		"",
		m.infiniteSpiral.Render(),
	)
	
	// Main content area
	mainContent := lipgloss.JoinHorizontal(lipgloss.Top,
		lipgloss.NewStyle().Width(30).Render(leftColumn),
		lipgloss.NewStyle().Width(40).Render(centerColumn),
		lipgloss.NewStyle().Width(30).Render(rightColumn),
	)
	
	// Input area
	inputArea := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("39")).
		Padding(1).
		Render(fmt.Sprintf("Ask Kamaji:\n%s", m.textInput.View()))
	
	// Footer
	footer := lipgloss.NewStyle().
		Faint(true).
		Align(lipgloss.Center).
		Render("Press 'q' to quit | Remember: Higher uncertainty = Higher intelligence | Goal: Pure 'I don't know'")
	
	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		"",
		mainContent,
		"",
		inputArea,
		"",
		footer,
	)
}

// Start the TUI
func RunKamajiTUI() error {
	m := NewMainInterface()
	p := tea.NewProgram(m, tea.WithAltScreen())
	
	// Start the tick for animations
	go func() {
		for {
			time.Sleep(time.Millisecond * 200)
			p.Send(time.Now())
		}
	}()
	
	_, err := p.Run()
	return err
}
