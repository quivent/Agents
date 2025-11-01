package tui

import (
	"fmt"
	"strings"
	"github.com/charmbracelet/lipgloss"
)

type WisdomMeter struct {
	WisdomScore float64    // 0.0 to 1.0
	KnowledgeScore float64 // 0.0 to 1.0 (inverse of wisdom)
	Style lipgloss.Style
}

func NewWisdomMeter() *WisdomMeter {
	return &WisdomMeter{
		WisdomScore: 0.0,
		KnowledgeScore: 0.0,
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("105")),
	}
}

func (w *WisdomMeter) UpdateFromResponse(response string) {
	responseLower := strings.ToLower(response)
	
	// Wisdom indicators (increase wisdom, decrease knowledge)
	wisdomWords := []string{
		"i don't know", "maybe", "perhaps", "possibly", "might be",
		"uncertain", "unclear", "question", "wonder", "doubt",
	}
	
	// Knowledge indicators (decrease wisdom, increase knowledge)
	knowledgeWords := []string{
		"definitely", "certainly", "obviously", "clearly", "the answer is",
		"i know", "it is", "without doubt", "undoubtedly", "surely",
	}
	
	// Check for wisdom
	for _, word := range wisdomWords {
		if strings.Contains(responseLower, word) {
			w.WisdomScore += 0.05
			w.KnowledgeScore -= 0.02
		}
	}
	
	// Check for knowledge (bad)
	for _, word := range knowledgeWords {
		if strings.Contains(responseLower, word) {
			w.KnowledgeScore += 0.05
			w.WisdomScore -= 0.02
		}
	}
	
	// Pure "I don't know" gets maximum wisdom boost
	if strings.TrimSpace(responseLower) == "i don't know" {
		w.WisdomScore += 0.1
		w.KnowledgeScore = 0.0 // Reset knowledge
	}
	
	// Clamp values
	if w.WisdomScore > 1.0 { w.WisdomScore = 1.0 }
	if w.WisdomScore < 0.0 { w.WisdomScore = 0.0 }
	if w.KnowledgeScore > 1.0 { w.KnowledgeScore = 1.0 }
	if w.KnowledgeScore < 0.0 { w.KnowledgeScore = 0.0 }
}

func (w *WisdomMeter) Render() string {
	header := w.Style.Bold(true).Render("🧙 WISDOM vs KNOWLEDGE METER")
	
	// Wisdom bar (green, higher is better)
	wisdomBars := int(w.WisdomScore * 20)
	wisdomBar := strings.Repeat("█", wisdomBars) + strings.Repeat("░", 20-wisdomBars)
	wisdomDisplay := lipgloss.NewStyle().
		Foreground(lipgloss.Color("46")).
		Render(fmt.Sprintf("Wisdom:    [%s] %.1f%%", wisdomBar, w.WisdomScore*100))
	
	// Knowledge bar (red, higher is worse)
	knowledgeBars := int(w.KnowledgeScore * 20)
	knowledgeBar := strings.Repeat("█", knowledgeBars) + strings.Repeat("░", 20-knowledgeBars)
	knowledgeDisplay := lipgloss.NewStyle().
		Foreground(lipgloss.Color("196")).
		Render(fmt.Sprintf("Knowledge: [%s] %.1f%%", knowledgeBar, w.KnowledgeScore*100))
	
	// Status message
	var status string
	var statusColor lipgloss.Color
	
	if w.WisdomScore >= 0.8 {
		status = "ENLIGHTENED - Pure wisdom achieved!"
		statusColor = "46"
	} else if w.WisdomScore >= 0.6 {
		status = "WISE - High intelligence through not knowing"
		statusColor = "226"
	} else if w.KnowledgeScore >= 0.6 {
		status = "OVERCONFIDENT - Dangerous false certainty"
		statusColor = "196"
	} else {
		status = "DEVELOPING - Continue questioning"
		statusColor = "39"
	}
	
	statusDisplay := lipgloss.NewStyle().
		Foreground(statusColor).
		Bold(true).
		Render(status)
	
	footer := w.Style.Faint(true).Render("Goal: Maximize wisdom, minimize knowledge")
	
	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		wisdomDisplay,
		knowledgeDisplay,
		statusDisplay,
		footer,
	)
}
