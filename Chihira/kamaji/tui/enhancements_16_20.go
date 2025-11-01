package tui

import (
	"fmt"
	"math"
	"strings"
	"time"
	"github.com/charmbracelet/lipgloss"
)

// ITERATION 16: Doubt Amplifier
type DoubtAmplifier struct {
	DoubtLevel float64
	Amplification float64
	Style lipgloss.Style
}

func NewDoubtAmplifier() *DoubtAmplifier {
	return &DoubtAmplifier{
		DoubtLevel: 0.0,
		Amplification: 1.0,
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("166")),
	}
}

func (d *DoubtAmplifier) AmplifyDoubt(response string) {
	doubtWords := []string{"maybe", "perhaps", "possibly", "might", "could", "uncertain"}
	for _, word := range doubtWords {
		if strings.Contains(strings.ToLower(response), word) {
			d.DoubtLevel += 0.05
			d.Amplification += 0.1
		}
	}
	if d.DoubtLevel > 1.0 { d.DoubtLevel = 1.0 }
}

func (d *DoubtAmplifier) Render() string {
	amplified := d.DoubtLevel * d.Amplification
	bars := int(amplified * 20)
	if bars > 20 { bars = 20 }
	
	doubtBar := strings.Repeat("~", bars) + strings.Repeat(".", 20-bars)
	
	return lipgloss.JoinVertical(lipgloss.Left,
		d.Style.Bold(true).Render("📡 DOUBT AMPLIFIER"),
		d.Style.Render(fmt.Sprintf("[%s] x%.1f", doubtBar, d.Amplification)),
		d.Style.Faint(true).Render("Amplifying uncertainty for maximum intelligence"),
	)
}

// ITERATION 17: Wisdom Particle System
type WisdomParticle struct {
	X, Y float64
	VX, VY float64
	Life float64
	Char rune
}

type WisdomParticleSystem struct {
	Particles []WisdomParticle
	Style lipgloss.Style
}

func NewWisdomParticleSystem() *WisdomParticleSystem {
	return &WisdomParticleSystem{
		Particles: []WisdomParticle{},
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("147")),
	}
}

func (w *WisdomParticleSystem) EmitWisdom() {
	chars := []rune{'?', '∞', '∅', '◯', '△', '▽'}
	for i := 0; i < 5; i++ {
		particle := WisdomParticle{
			X: 15 + float64(i-2)*2,
			Y: 5,
			VX: (float64(i-2) * 0.5),
			VY: -1.0,
			Life: 1.0,
			Char: chars[i%len(chars)],
		}
		w.Particles = append(w.Particles, particle)
	}
}

func (w *WisdomParticleSystem) Update() {
	for i := len(w.Particles) - 1; i >= 0; i-- {
		p := &w.Particles[i]
		p.X += p.VX
		p.Y += p.VY
		p.Life -= 0.05
		
		if p.Life <= 0 {
			w.Particles = append(w.Particles[:i], w.Particles[i+1:]...)
		}
	}
}

func (w *WisdomParticleSystem) Render() string {
	grid := make([][]rune, 10)
	for i := range grid {
		grid[i] = make([]rune, 30)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}
	
	for _, p := range w.Particles {
		x, y := int(p.X), int(p.Y)
		if x >= 0 && x < 30 && y >= 0 && y < 10 {
			grid[y][x] = p.Char
		}
	}
	
	lines := []string{w.Style.Bold(true).Render("✨ WISDOM PARTICLES")}
	for _, row := range grid {
		lines = append(lines, w.Style.Render(string(row)))
	}
	
	return strings.Join(lines, "\n")
}

// ITERATION 18: Enlightenment Progress Bar
type EnlightenmentProgress struct {
	Progress float64
	Milestones []string
	CurrentMilestone int
	Style lipgloss.Style
}

func NewEnlightenmentProgress() *EnlightenmentProgress {
	milestones := []string{
		"Questioning begins",
		"Assumptions challenged", 
		"Certainty dissolves",
		"Paradoxes embraced",
		"Meta-questioning achieved",
		"Infinite regress entered",
		"Pure not-knowing approached",
		"Enlightenment: I don't know",
	}
	
	return &EnlightenmentProgress{
		Progress: 0.0,
		Milestones: milestones,
		CurrentMilestone: 0,
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("220")),
	}
}

func (e *EnlightenmentProgress) UpdateProgress(notKnowingCount int, questionDepth int, paradoxCount int) {
	score := float64(notKnowingCount)*0.01 + float64(questionDepth)*0.05 + float64(paradoxCount)*0.03
	e.Progress = math.Min(score, 1.0)
	e.CurrentMilestone = int(e.Progress * float64(len(e.Milestones)))
	if e.CurrentMilestone >= len(e.Milestones) {
		e.CurrentMilestone = len(e.Milestones) - 1
	}
}

func (e *EnlightenmentProgress) Render() string {
	bars := int(e.Progress * 25)
	progressBar := strings.Repeat("█", bars) + strings.Repeat("░", 25-bars)
	
	milestone := ""
	if e.CurrentMilestone < len(e.Milestones) {
		milestone = e.Milestones[e.CurrentMilestone]
	}
	
	return lipgloss.JoinVertical(lipgloss.Left,
		e.Style.Bold(true).Render("🌅 ENLIGHTENMENT PROGRESS"),
		e.Style.Render(fmt.Sprintf("[%s] %.1f%%", progressBar, e.Progress*100)),
		e.Style.Render(fmt.Sprintf("Current: %s", milestone)),
	)
}

// ITERATION 19: Void Gazer
type VoidGazer struct {
	GazingTime time.Duration
	VoidDepth int
	IsGazing bool
	Style lipgloss.Style
}

func NewVoidGazer() *VoidGazer {
	return &VoidGazer{
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("0")),
	}
}

func (v *VoidGazer) StartGazing() {
	v.IsGazing = true
	v.GazingTime = 0
}

func (v *VoidGazer) UpdateGaze() {
	if v.IsGazing {
		v.GazingTime += time.Millisecond * 200
		v.VoidDepth = int(v.GazingTime.Seconds())
	}
}

func (v *VoidGazer) Render() string {
	if !v.IsGazing {
		return v.Style.Faint(true).Render("👁️ Not gazing into void")
	}
	
	void := strings.Repeat("●", v.VoidDepth%10) + strings.Repeat("○", 10-(v.VoidDepth%10))
	
	return lipgloss.JoinVertical(lipgloss.Left,
		v.Style.Bold(true).Render("👁️ GAZING INTO THE VOID"),
		v.Style.Render(fmt.Sprintf("[%s]", void)),
		v.Style.Render(fmt.Sprintf("Depth: %d", v.VoidDepth)),
		v.Style.Faint(true).Render("The void gazes back"),
	)
}

// ITERATION 20: Socratic Method Tracker
type SocraticMethod struct {
	QuestionsAsked int
	AnswersGiven int
	SocraticRatio float64
	Style lipgloss.Style
}

func NewSocraticMethod() *SocraticMethod {
	return &SocraticMethod{
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("67")),
	}
}

func (s *SocraticMethod) TrackInteraction(isQuestion bool) {
	if isQuestion {
		s.QuestionsAsked++
	} else {
		s.AnswersGiven++
	}
	
	total := s.QuestionsAsked + s.AnswersGiven
	if total > 0 {
		s.SocraticRatio = float64(s.QuestionsAsked) / float64(total)
	}
}

func (s *SocraticMethod) Render() string {
	bars := int(s.SocraticRatio * 20)
	ratioBar := strings.Repeat("?", bars) + strings.Repeat("!", 20-bars)
	
	var status string
	if s.SocraticRatio >= 0.8 {
		status = "TRUE SOCRATIC WISDOM"
	} else if s.SocraticRatio >= 0.6 {
		status = "GOOD QUESTIONING"
	} else {
		status = "TOO MANY ANSWERS"
	}
	
	return lipgloss.JoinVertical(lipgloss.Left,
		s.Style.Bold(true).Render("🏛️ SOCRATIC METHOD TRACKER"),
		s.Style.Render(fmt.Sprintf("Questions: %d | Answers: %d", s.QuestionsAsked, s.AnswersGiven)),
		s.Style.Render(fmt.Sprintf("[%s] %.1f%%", ratioBar, s.SocraticRatio*100)),
		s.Style.Render(status),
	)
}
