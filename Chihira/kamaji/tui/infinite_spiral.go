package tui

import (
	"fmt"
	"math"
	"strings"
	"time"
	"github.com/charmbracelet/lipgloss"
)

type InfiniteSpiral struct {
	Depth int
	Rotation float64
	Style lipgloss.Style
	Active bool
}

func NewInfiniteSpiral() *InfiniteSpiral {
	return &InfiniteSpiral{
		Depth: 0,
		Rotation: 0,
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("93")),
		Active: false,
	}
}

func (s *InfiniteSpiral) Update() {
	if s.Active {
		s.Rotation += 0.2
		s.Depth++
	}
}

func (s *InfiniteSpiral) Render() string {
	if !s.Active {
		return s.Style.Faint(true).Render("∞ Infinite questioning: INACTIVE")
	}
	
	// Create spiral pattern
	spiral := s.generateSpiral()
	
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("201")).
		Render("∞ INFINITE QUESTIONING SPIRAL ∞")
	
	depthInfo := lipgloss.NewStyle().
		Foreground(lipgloss.Color("39")).
		Render(fmt.Sprintf("Depth: %d | Rotation: %.1f°", s.Depth, s.Rotation))
	
	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		spiral,
		depthInfo,
		s.Style.Faint(true).Render("Questions generating questions infinitely..."),
	)
}

func (s *InfiniteSpiral) generateSpiral() string {
	lines := []string{}
	centerX, centerY := 15, 5
	
	for i := 0; i < 10; i++ {
		line := strings.Repeat(" ", 30)
		
		for j := 0; j < 5; j++ {
			angle := s.Rotation + float64(i)*0.5 + float64(j)*1.2
			radius := 2 + float64(i)*0.3
			
			x := int(float64(centerX) + radius*math.Cos(angle))
			y := int(float64(centerY) + radius*math.Sin(angle)*0.5)
			
			if x >= 0 && x < 30 && y == i {
				chars := []rune(line)
				if x < len(chars) {
					chars[x] = '?'
				}
				line = string(chars)
			}
		}
		lines = append(lines, s.Style.Render(line))
	}
	
	return strings.Join(lines, "\n")
}

func (s *InfiniteSpiral) StartInfiniteLoop() {
	s.Active = true
}

func (s *InfiniteSpiral) StopInfiniteLoop() {
	s.Active = false
}
