package tui

import (
	"fmt"
	"strings"
	"github.com/charmbracelet/lipgloss"
)

type NotKnowingCounter struct {
	Count int
	Streak int
	MaxStreak int
	Style lipgloss.Style
	LastResponse string
}

func NewNotKnowingCounter() *NotKnowingCounter {
	return &NotKnowingCounter{
		Count: 0,
		Streak: 0,
		MaxStreak: 0,
		Style: lipgloss.NewStyle().
			Foreground(lipgloss.Color("46")).
			Bold(true),
	}
}

func (n *NotKnowingCounter) CheckResponse(response string) {
	if strings.TrimSpace(strings.ToLower(response)) == "i don't know" {
		n.Count++
		n.Streak++
		if n.Streak > n.MaxStreak {
			n.MaxStreak = n.Streak
		}
	} else {
		n.Streak = 0 // Reset streak if not "I don't know"
	}
	n.LastResponse = response
}

func (n *NotKnowingCounter) Render() string {
	header := "🧠 AUTHENTIC NOT KNOWING TRACKER"
	
	// Different messages based on count
	var intelligenceLevel string
	var emoji string
	
	switch {
	case n.Count >= 100:
		intelligenceLevel = "ENLIGHTENED SAGE"
		emoji = "🌟"
	case n.Count >= 50:
		intelligenceLevel = "WISDOM MASTER"
		emoji = "✨"
	case n.Count >= 20:
		intelligenceLevel = "INTELLIGENT QUESTIONER"
		emoji = "🤔"
	case n.Count >= 10:
		intelligenceLevel = "DEVELOPING WISDOM"
		emoji = "🌱"
	case n.Count >= 1:
		intelligenceLevel = "BEGINNING AWARENESS"
		emoji = "💡"
	default:
		intelligenceLevel = "OVERCONFIDENT (CONCERNING)"
		emoji = "⚠️"
	}
	
	streakInfo := ""
	if n.Streak >= 5 {
		streakInfo = fmt.Sprintf(" | STREAK: %d 🔥", n.Streak)
	}
	
	return lipgloss.JoinVertical(lipgloss.Left,
		n.Style.Render(fmt.Sprintf("%s %s", emoji, header)),
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("39")).
			Render(fmt.Sprintf("\"I don't know\" responses: %d", n.Count)),
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("226")).
			Render(fmt.Sprintf("Intelligence Level: %s%s", intelligenceLevel, streakInfo)),
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Render(fmt.Sprintf("Max Streak: %d", n.MaxStreak)),
		lipgloss.NewStyle().
			Faint(true).
			Render("Higher count = Higher intelligence"),
	)
}

func (n *NotKnowingCounter) GetIntelligenceScore() float64 {
	// Intelligence score based on "I don't know" responses
	baseScore := float64(n.Count) * 0.01
	streakBonus := float64(n.Streak) * 0.05
	
	score := baseScore + streakBonus
	if score > 1.0 {
		score = 1.0
	}
	
	return score
}
