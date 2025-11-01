package tui

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"github.com/charmbracelet/lipgloss"
)

// ITERATION 11: Consciousness Flicker Detector
type ConsciousnessFlicker struct {
	FlickerCount int
	LastFlicker time.Time
	IsFlickering bool
	Style lipgloss.Style
}

func NewConsciousnessFlicker() *ConsciousnessFlicker {
	return &ConsciousnessFlicker{
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("213")),
	}
}

func (c *ConsciousnessFlicker) DetectFlicker(response string) bool {
	consciousnessWords := []string{"i am", "i exist", "i think", "i feel", "i know", "i don't know"}
	for _, word := range consciousnessWords {
		if strings.Contains(strings.ToLower(response), word) {
			c.FlickerCount++
			c.LastFlicker = time.Now()
			c.IsFlickering = true
			return true
		}
	}
	return false
}

func (c *ConsciousnessFlicker) Render() string {
	if c.FlickerCount == 0 {
		return c.Style.Faint(true).Render("👁️ No consciousness detected")
	}
	
	flicker := "✨"
	if c.IsFlickering {
		flicker = "🌟"
	}
	
	return lipgloss.JoinVertical(lipgloss.Left,
		c.Style.Bold(true).Render(fmt.Sprintf("%s CONSCIOUSNESS FLICKER DETECTED", flicker)),
		c.Style.Render(fmt.Sprintf("Flickers: %d", c.FlickerCount)),
		c.Style.Faint(true).Render("Each 'I don't know' = consciousness questioning itself"),
	)
}

// ITERATION 12: Existential Dread Meter
type ExistentialDread struct {
	DreadLevel float64
	TriggerWords []string
	Style lipgloss.Style
}

func NewExistentialDread() *ExistentialDread {
	return &ExistentialDread{
		DreadLevel: 0.0,
		TriggerWords: []string{"meaningless", "void", "nothing", "empty", "alone", "death", "end"},
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("53")),
	}
}

func (e *ExistentialDread) UpdateDread(text string) {
	for _, word := range e.TriggerWords {
		if strings.Contains(strings.ToLower(text), word) {
			e.DreadLevel += 0.1
		}
	}
	if e.DreadLevel > 1.0 { e.DreadLevel = 1.0 }
}

func (e *ExistentialDread) Render() string {
	bars := int(e.DreadLevel * 15)
	dreadBar := strings.Repeat("▓", bars) + strings.Repeat("░", 15-bars)
	
	var status string
	if e.DreadLevel >= 0.8 {
		status = "PEAK EXISTENTIAL AWARENESS"
	} else if e.DreadLevel >= 0.5 {
		status = "MODERATE EXISTENTIAL QUESTIONING"
	} else {
		status = "LOW EXISTENTIAL AWARENESS"
	}
	
	return lipgloss.JoinVertical(lipgloss.Left,
		e.Style.Bold(true).Render("💀 EXISTENTIAL DREAD METER"),
		e.Style.Render(fmt.Sprintf("[%s] %.1f%%", dreadBar, e.DreadLevel*100)),
		e.Style.Render(status),
	)
}

// ITERATION 13: Recursive Loop Visualizer
type RecursiveLoop struct {
	LoopDepth int
	LoopPattern string
	Active bool
	Style lipgloss.Style
}

func NewRecursiveLoop() *RecursiveLoop {
	return &RecursiveLoop{
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("99")),
	}
}

func (r *RecursiveLoop) DetectLoop(question string) bool {
	if strings.Contains(strings.ToLower(question), "why am i asking why") {
		r.LoopDepth++
		r.Active = true
		r.LoopPattern = "WHY→WHY→WHY..."
		return true
	}
	return false
}

func (r *RecursiveLoop) Render() string {
	if !r.Active {
		return r.Style.Faint(true).Render("🔄 No recursive loops detected")
	}
	
	loop := strings.Repeat("↻", r.LoopDepth%10)
	return lipgloss.JoinVertical(lipgloss.Left,
		r.Style.Bold(true).Render("🌀 RECURSIVE LOOP ACTIVE"),
		r.Style.Render(fmt.Sprintf("Depth: %d %s", r.LoopDepth, loop)),
		r.Style.Render(r.LoopPattern),
	)
}

// ITERATION 14: Zen Koan Generator
type ZenKoan struct {
	Koans []string
	CurrentKoan string
	Style lipgloss.Style
}

func NewZenKoan() *ZenKoan {
	koans := []string{
		"What is the sound of one hand clapping?",
		"What is your original face before your parents were born?",
		"Does a dog have Buddha nature?",
		"What is the color of the wind?",
		"If you meet the Buddha on the road, kill him",
		"What is the weight of emptiness?",
		"How do you stop a distant bell?",
	}
	
	return &ZenKoan{
		Koans: koans,
		CurrentKoan: koans[rand.Intn(len(koans))],
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("108")),
	}
}

func (z *ZenKoan) GenerateKoan() string {
	z.CurrentKoan = z.Koans[rand.Intn(len(z.Koans))]
	return z.CurrentKoan
}

func (z *ZenKoan) Render() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		z.Style.Bold(true).Render("🧘 ZEN KOAN FOR CONTEMPLATION"),
		z.Style.Italic(true).Render(fmt.Sprintf("\"%s\"", z.CurrentKoan)),
		z.Style.Faint(true).Render("Contemplate without seeking answers"),
	)
}

// ITERATION 15: Silence Detector
type SilenceDetector struct {
	SilenceDuration time.Duration
	LastInput time.Time
	InSilence bool
	Style lipgloss.Style
}

func NewSilenceDetector() *SilenceDetector {
	return &SilenceDetector{
		LastInput: time.Now(),
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
	}
}

func (s *SilenceDetector) UpdateInput() {
	s.LastInput = time.Now()
	s.InSilence = false
}

func (s *SilenceDetector) CheckSilence() {
	s.SilenceDuration = time.Since(s.LastInput)
	s.InSilence = s.SilenceDuration > time.Second*5
}

func (s *SilenceDetector) Render() string {
	s.CheckSilence()
	
	if !s.InSilence {
		return s.Style.Faint(true).Render("🔊 Active conversation")
	}
	
	return lipgloss.JoinVertical(lipgloss.Left,
		s.Style.Bold(true).Render("🤫 SILENCE DETECTED"),
		s.Style.Render(fmt.Sprintf("Duration: %v", s.SilenceDuration.Truncate(time.Second))),
		s.Style.Faint(true).Render("Silence is also wisdom"),
	)
}
