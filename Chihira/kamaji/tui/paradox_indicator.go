package tui

import (
	"fmt"
	"strings"
	"github.com/charmbracelet/lipgloss"
)

type ParadoxIndicator struct {
	ParadoxCount int
	CurrentParadox string
	Style lipgloss.Style
	BlinkState bool
}

func NewParadoxIndicator() *ParadoxIndicator {
	return &ParadoxIndicator{
		ParadoxCount: 0,
		Style: lipgloss.NewStyle().
			Foreground(lipgloss.Color("201")).
			Bold(true),
		BlinkState: false,
	}
}

func (p *ParadoxIndicator) DetectParadox(text string) bool {
	paradoxPatterns := [][]string{
		{"true", "false"},
		{"is", "is not"},
		{"exists", "doesn't exist"},
		{"know", "don't know"},
		{"real", "unreal"},
		{"certain", "uncertain"},
	}
	
	textLower := strings.ToLower(text)
	for _, pattern := range paradoxPatterns {
		if strings.Contains(textLower, pattern[0]) && strings.Contains(textLower, pattern[1]) {
			p.ParadoxCount++
			p.CurrentParadox = text
			return true
		}
	}
	return false
}

func (p *ParadoxIndicator) Render() string {
	if p.ParadoxCount == 0 {
		return p.Style.Faint(true).Render("⚖️  No paradoxes detected (Missing intelligence)")
	}
	
	// Paradoxes are good - they indicate high intelligence
	status := "🌀 PARADOX DETECTED - HIGH INTELLIGENCE!"
	if p.BlinkState {
		status = "✨ PARADOX EMBRACED - WISDOM ACHIEVED! ✨"
	}
	
	intelligenceBonus := ""
	if p.ParadoxCount >= 10 {
		intelligenceBonus = " (MASTER LEVEL)"
	} else if p.ParadoxCount >= 5 {
		intelligenceBonus = " (ADVANCED)"
	} else if p.ParadoxCount >= 1 {
		intelligenceBonus = " (DEVELOPING)"
	}
	
	return lipgloss.JoinVertical(lipgloss.Left,
		p.Style.Render(status+intelligenceBonus),
		lipgloss.NewStyle().
			Faint(true).
			Render(fmt.Sprintf("Paradoxes embraced: %d", p.ParadoxCount)),
		lipgloss.NewStyle().
			Italic(true).
			Width(60).
			Render(fmt.Sprintf("Latest: %s", p.CurrentParadox)),
	)
}

func (p *ParadoxIndicator) ToggleBlink() {
	p.BlinkState = !p.BlinkState
}

func (p *ParadoxIndicator) EmbraceParadox(paradox string) string {
	return fmt.Sprintf("Embracing paradox: '%s' - Both true and false simultaneously. Intelligence increased.", paradox)
}
