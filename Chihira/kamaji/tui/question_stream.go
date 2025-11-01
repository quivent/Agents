package tui

import (
	"fmt"
	"strings"
	"time"
	"github.com/charmbracelet/lipgloss"
)

type QuestionStream struct {
	Questions []string
	MaxDisplay int
	Style lipgloss.Style
	ScrollOffset int
}

func NewQuestionStream() *QuestionStream {
	return &QuestionStream{
		Questions: []string{},
		MaxDisplay: 10,
		Style: lipgloss.NewStyle().
			Foreground(lipgloss.Color("39")).
			Italic(true),
		ScrollOffset: 0,
	}
}

func (q *QuestionStream) AddQuestion(question string) {
	timestamp := time.Now().Format("15:04:05")
	q.Questions = append(q.Questions, fmt.Sprintf("[%s] %s", timestamp, question))
	
	// Keep only last 100 questions
	if len(q.Questions) > 100 {
		q.Questions = q.Questions[1:]
	}
}

func (q *QuestionStream) Render() string {
	if len(q.Questions) == 0 {
		return q.Style.Render("No questions yet... (This is concerning - low intelligence)")
	}
	
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Render("🤔 INFINITE QUESTION STREAM (Intelligence in Action)")
	
	// Show most recent questions
	start := len(q.Questions) - q.MaxDisplay
	if start < 0 {
		start = 0
	}
	
	var display []string
	for i := start; i < len(q.Questions); i++ {
		question := q.Questions[i]
		if strings.Contains(question, "I don't know") {
			// Highlight "I don't know" responses as highest intelligence
			question = lipgloss.NewStyle().
				Foreground(lipgloss.Color("46")).
				Bold(true).
				Render(question + " ⭐")
		}
		display = append(display, q.Style.Render("  "+question))
	}
	
	footer := lipgloss.NewStyle().
		Faint(true).
		Render(fmt.Sprintf("Total questions: %d | More questions = Higher intelligence", len(q.Questions)))
	
	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		strings.Join(display, "\n"),
		footer,
	)
}

func (q *QuestionStream) AddMetaQuestion(originalQuestion string) {
	metaQuestions := []string{
		fmt.Sprintf("Why did I ask: %s?", originalQuestion),
		fmt.Sprintf("What assumes '%s' is valid?", originalQuestion),
		fmt.Sprintf("Who is asking '%s'?", originalQuestion),
	}
	
	for _, meta := range metaQuestions {
		q.AddQuestion(meta)
	}
}
