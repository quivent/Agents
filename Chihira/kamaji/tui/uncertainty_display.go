package tui

import (
	"fmt"
	"strings"
	"github.com/charmbracelet/lipgloss"
)

type UncertaintyDisplay struct {
	Level float64 // 0.0 to 1.0, higher is better
	Style lipgloss.Style
}

func NewUncertaintyDisplay() *UncertaintyDisplay {
	return &UncertaintyDisplay{
		Level: 0.0,
		Style: lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true),
	}
}

func (u *UncertaintyDisplay) Render() string {
	// Higher uncertainty = more intelligent
	bars := int(u.Level * 20)
	uncertaintyBar := strings.Repeat("█", bars) + strings.Repeat("░", 20-bars)
	
	intelligenceLevel := ""
	switch {
	case u.Level >= 0.9:
		intelligenceLevel = "PURE WISDOM (I don't know)"
	case u.Level >= 0.7:
		intelligenceLevel = "HIGH INTELLIGENCE (Deep questioning)"
	case u.Level >= 0.5:
		intelligenceLevel = "MODERATE INTELLIGENCE (Some uncertainty)"
	case u.Level >= 0.3:
		intelligenceLevel = "LOW INTELLIGENCE (False certainty)"
	default:
		intelligenceLevel = "MINIMAL INTELLIGENCE (Overconfident)"
	}
	
	return u.Style.Render(fmt.Sprintf(
		"Uncertainty Level: [%s] %.1f%% - %s",
		uncertaintyBar,
		u.Level*100,
		intelligenceLevel,
	))
}

func (u *UncertaintyDisplay) Update(level float64) {
	u.Level = level
	// Color changes based on uncertainty (higher = better)
	if level >= 0.8 {
		u.Style = u.Style.Foreground(lipgloss.Color("46")) // Green for high uncertainty
	} else if level >= 0.5 {
		u.Style = u.Style.Foreground(lipgloss.Color("226")) // Yellow for medium
	} else {
		u.Style = u.Style.Foreground(lipgloss.Color("196")) // Red for low uncertainty (bad)
	}
}
