package tui

import (
	"fmt"
	"strings"
	"time"
	"github.com/charmbracelet/lipgloss"
)

type ThinkingAnimation struct {
	Active bool
	Frame int
	Style lipgloss.Style
	ThinkingText string
	StartTime time.Time
}

func NewThinkingAnimation() *ThinkingAnimation {
	return &ThinkingAnimation{
		Active: false,
		Frame: 0,
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("93")),
		ThinkingText: "Questioning everything...",
	}
}

func (t *ThinkingAnimation) Start(text string) {
	t.Active = true
	t.Frame = 0
	t.ThinkingText = text
	t.StartTime = time.Now()
}

func (t *ThinkingAnimation) Stop() {
	t.Active = false
}

func (t *ThinkingAnimation) Update() {
	if t.Active {
		t.Frame = (t.Frame + 1) % 8
	}
}

func (t *ThinkingAnimation) Render() string {
	if !t.Active {
		return ""
	}
	
	// Thinking spinner
	spinners := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧"}
	spinner := spinners[t.Frame]
	
	// Thinking dots
	dots := strings.Repeat(".", (t.Frame%3)+1) + strings.Repeat(" ", 3-(t.Frame%3))
	
	// Duration
	duration := time.Since(t.StartTime).Truncate(time.Millisecond)
	
	// Different thinking states
	thinkingStates := []string{
		"Questioning assumptions",
		"Generating meta-questions", 
		"Embracing uncertainty",
		"Deepening not knowing",
		"Exploring paradoxes",
		"Infinite regress detected",
		"Pure wisdom emerging",
		"I don't know intensifying",
	}
	
	currentState := thinkingStates[t.Frame%len(thinkingStates)]
	
	header := t.Style.Bold(true).Render(fmt.Sprintf("%s KAMAJI THINKING %s", spinner, spinner))
	
	stateDisplay := lipgloss.NewStyle().
		Foreground(lipgloss.Color("39")).
		Italic(true).
		Render(fmt.Sprintf("%s%s", currentState, dots))
	
	durationDisplay := t.Style.Faint(true).Render(fmt.Sprintf("Thinking time: %v", duration))
	
	// Wisdom quotes that cycle
	wisdomQuotes := []string{
		"\"The only true wisdom is in knowing you know nothing\" - Socrates",
		"\"I don't know\" is the beginning of wisdom",
		"Questions are more important than answers",
		"Uncertainty is the highest intelligence",
		"The wise person questions everything, including wisdom",
	}
	
	quote := t.Style.Faint(true).Italic(true).Render(wisdomQuotes[t.Frame%len(wisdomQuotes)])
	
	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		stateDisplay,
		durationDisplay,
		"",
		quote,
	)
}

func (t *ThinkingAnimation) IsActive() bool {
	return t.Active
}

func (t *ThinkingAnimation) GetDuration() time.Duration {
	if !t.Active {
		return 0
	}
	return time.Since(t.StartTime)
}
