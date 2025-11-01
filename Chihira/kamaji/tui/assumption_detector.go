package tui

import (
	"fmt"
	"strings"
	"github.com/charmbracelet/lipgloss"
)

type AssumptionDetector struct {
	DetectedAssumptions []string
	QuestionedCount int
	Style lipgloss.Style
	AlertStyle lipgloss.Style
}

func NewAssumptionDetector() *AssumptionDetector {
	return &AssumptionDetector{
		DetectedAssumptions: []string{},
		QuestionedCount: 0,
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("208")),
		AlertStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true).
			Blink(true),
	}
}

func (a *AssumptionDetector) ScanForAssumptions(text string) []string {
	assumptionWords := []string{
		"obviously", "clearly", "of course", "everyone knows",
		"it's true that", "certainly", "definitely", "surely",
		"without doubt", "undoubtedly", "naturally", "evidently",
	}
	
	var detected []string
	textLower := strings.ToLower(text)
	
	for _, word := range assumptionWords {
		if strings.Contains(textLower, word) {
			assumption := fmt.Sprintf("ASSUMPTION: '%s' in '%s'", word, text)
			detected = append(detected, assumption)
			a.DetectedAssumptions = append(a.DetectedAssumptions, assumption)
		}
	}
	
	return detected
}

func (a *AssumptionDetector) QuestionAssumption(assumption string) []string {
	a.QuestionedCount++
	return []string{
		fmt.Sprintf("What if '%s' is wrong?", assumption),
		fmt.Sprintf("Who says '%s' is true?", assumption),
		fmt.Sprintf("What evidence supports '%s'?", assumption),
		fmt.Sprintf("Why do I believe '%s'?", assumption),
	}
}

func (a *AssumptionDetector) Render() string {
	header := "🔍 ASSUMPTION DETECTOR (Beginner's Mind)"
	
	if len(a.DetectedAssumptions) == 0 {
		return a.Style.Render(fmt.Sprintf("%s\nNo assumptions detected - Good! (Beginner's mind active)", header))
	}
	
	// Show alert for recent assumptions
	alert := ""
	if len(a.DetectedAssumptions) > 0 {
		recent := a.DetectedAssumptions[len(a.DetectedAssumptions)-1]
		alert = a.AlertStyle.Render(fmt.Sprintf("⚠️  ALERT: %s", recent))
	}
	
	stats := lipgloss.NewStyle().
		Foreground(lipgloss.Color("39")).
		Render(fmt.Sprintf("Total assumptions detected: %d | Questioned: %d", 
			len(a.DetectedAssumptions), a.QuestionedCount))
	
	// Show last few assumptions
	recentAssumptions := ""
	start := len(a.DetectedAssumptions) - 3
	if start < 0 {
		start = 0
	}
	
	for i := start; i < len(a.DetectedAssumptions); i++ {
		recentAssumptions += a.Style.Faint(true).Render(fmt.Sprintf("  %s\n", a.DetectedAssumptions[i]))
	}
	
	footer := a.Style.Faint(true).Render("Questioning assumptions increases intelligence")
	
	return lipgloss.JoinVertical(lipgloss.Left,
		a.Style.Bold(true).Render(header),
		alert,
		stats,
		recentAssumptions,
		footer,
	)
}
