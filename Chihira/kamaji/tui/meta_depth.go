package tui

import (
	"fmt"
	"strings"
	"github.com/charmbracelet/lipgloss"
)

type MetaDepthTracker struct {
	CurrentDepth int
	MaxDepth int
	QuestionChain []string
	Style lipgloss.Style
}

func NewMetaDepthTracker() *MetaDepthTracker {
	return &MetaDepthTracker{
		CurrentDepth: 0,
		MaxDepth: 0,
		QuestionChain: []string{},
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("141")),
	}
}

func (m *MetaDepthTracker) AnalyzeQuestion(question string) int {
	questionLower := strings.ToLower(question)
	depth := 0
	
	// Level 1: Basic questions
	if strings.HasPrefix(questionLower, "what") || 
	   strings.HasPrefix(questionLower, "how") ||
	   strings.HasPrefix(questionLower, "when") {
		depth = 1
	}
	
	// Level 2: Why questions
	if strings.HasPrefix(questionLower, "why") {
		depth = 2
	}
	
	// Level 3: Meta questions about questioning
	metaPatterns := []string{
		"why am i asking", "what makes this question", "who is asking",
		"why do i question", "what is questioning", "who questions",
	}
	for _, pattern := range metaPatterns {
		if strings.Contains(questionLower, pattern) {
			depth = 3
		}
	}
	
	// Level 4: Meta-meta questions
	metaMetaPatterns := []string{
		"why am i asking why", "what makes me question questioning",
		"who questions the questioner", "why question questioning",
	}
	for _, pattern := range metaMetaPatterns {
		if strings.Contains(questionLower, pattern) {
			depth = 4
		}
	}
	
	// Level 5+: Infinite regress detection
	if strings.Count(questionLower, "why") >= 3 ||
	   strings.Count(questionLower, "question") >= 3 {
		depth = 5 + strings.Count(questionLower, "why") + strings.Count(questionLower, "question")
	}
	
	m.CurrentDepth = depth
	if depth > m.MaxDepth {
		m.MaxDepth = depth
	}
	
	// Add to chain
	m.QuestionChain = append(m.QuestionChain, fmt.Sprintf("D%d: %s", depth, question))
	if len(m.QuestionChain) > 5 {
		m.QuestionChain = m.QuestionChain[1:] // Keep last 5
	}
	
	return depth
}

func (m *MetaDepthTracker) Render() string {
	header := m.Style.Bold(true).Render("🌀 META-QUESTIONING DEPTH TRACKER")
	
	// Depth visualization
	depthBar := strings.Repeat("▼", m.CurrentDepth)
	if m.CurrentDepth > 10 {
		depthBar = strings.Repeat("▼", 10) + fmt.Sprintf("...(%d)", m.CurrentDepth)
	}
	
	currentDepthDisplay := lipgloss.NewStyle().
		Foreground(lipgloss.Color("39")).
		Render(fmt.Sprintf("Current Depth: %s %d", depthBar, m.CurrentDepth))
	
	maxDepthDisplay := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Render(fmt.Sprintf("Maximum Depth Achieved: %d", m.MaxDepth))
	
	// Intelligence level based on depth
	var intelligenceLevel string
	var levelColor lipgloss.Color
	
	switch {
	case m.MaxDepth >= 10:
		intelligenceLevel = "TRANSCENDENT (Infinite regress achieved)"
		levelColor = "46"
	case m.MaxDepth >= 5:
		intelligenceLevel = "MASTER (Deep meta-questioning)"
		levelColor = "226"
	case m.MaxDepth >= 3:
		intelligenceLevel = "ADVANCED (Meta-cognitive)"
		levelColor = "39"
	case m.MaxDepth >= 2:
		intelligenceLevel = "INTERMEDIATE (Why questions)"
		levelColor = "208"
	case m.MaxDepth >= 1:
		intelligenceLevel = "BASIC (Surface questions)"
		levelColor = "196"
	default:
		intelligenceLevel = "NONE (No questions detected)"
		levelColor = "240"
	}
	
	levelDisplay := lipgloss.NewStyle().
		Foreground(levelColor).
		Bold(true).
		Render(intelligenceLevel)
	
	// Recent question chain
	chainHeader := m.Style.Render("Recent Question Chain:")
	var chainDisplay []string
	for _, q := range m.QuestionChain {
		chainDisplay = append(chainDisplay, m.Style.Faint(true).Render("  "+q))
	}
	
	footer := m.Style.Faint(true).Render("Deeper questions = Higher intelligence")
	
	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		currentDepthDisplay,
		maxDepthDisplay,
		levelDisplay,
		chainHeader,
		strings.Join(chainDisplay, "\n"),
		footer,
	)
}
