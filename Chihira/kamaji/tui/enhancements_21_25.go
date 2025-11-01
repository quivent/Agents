package tui

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"github.com/charmbracelet/lipgloss"
)

// ITERATION 21: Epistemic Humility Gauge
type EpistemicHumility struct {
	HumilityScore float64
	ArroganceDetected int
	Style lipgloss.Style
}

func NewEpistemicHumility() *EpistemicHumility {
	return &EpistemicHumility{
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("180")),
	}
}

func (e *EpistemicHumility) AnalyzeHumility(text string) {
	arrogantPhrases := []string{"i know for sure", "obviously", "without doubt", "certainly", "definitely"}
	humblePhrases := []string{"i might be wrong", "perhaps", "i'm not sure", "maybe", "i don't know"}
	
	textLower := strings.ToLower(text)
	
	for _, phrase := range arrogantPhrases {
		if strings.Contains(textLower, phrase) {
			e.HumilityScore -= 0.1
			e.ArroganceDetected++
		}
	}
	
	for _, phrase := range humblePhrases {
		if strings.Contains(textLower, phrase) {
			e.HumilityScore += 0.05
		}
	}
	
	if e.HumilityScore > 1.0 { e.HumilityScore = 1.0 }
	if e.HumilityScore < 0.0 { e.HumilityScore = 0.0 }
}

func (e *EpistemicHumility) Render() string {
	bars := int(e.HumilityScore * 15)
	humilityBar := strings.Repeat("🙏", bars) + strings.Repeat("😤", 15-bars)
	
	var status string
	if e.HumilityScore >= 0.8 {
		status = "DEEPLY HUMBLE (Wise)"
	} else if e.HumilityScore >= 0.5 {
		status = "MODERATELY HUMBLE"
	} else {
		status = "DANGEROUSLY ARROGANT"
	}
	
	return lipgloss.JoinVertical(lipgloss.Left,
		e.Style.Bold(true).Render("🙏 EPISTEMIC HUMILITY GAUGE"),
		e.Style.Render(humilityBar),
		e.Style.Render(fmt.Sprintf("Humility: %.1f%% | Arrogance: %d", e.HumilityScore*100, e.ArroganceDetected)),
		e.Style.Render(status),
	)
}

// ITERATION 22: Philosophical Mood Ring
type PhilosophicalMood struct {
	CurrentMood string
	MoodHistory []string
	Style lipgloss.Style
	Colors map[string]lipgloss.Color
}

func NewPhilosophicalMood() *PhilosophicalMood {
	colors := map[string]lipgloss.Color{
		"Existential": "53",
		"Skeptical": "166",
		"Curious": "39",
		"Confused": "208",
		"Enlightened": "46",
		"Nihilistic": "240",
		"Wondering": "105",
		"Doubtful": "93",
	}
	
	return &PhilosophicalMood{
		CurrentMood: "Neutral",
		MoodHistory: []string{},
		Style: lipgloss.NewStyle(),
		Colors: colors,
	}
}

func (p *PhilosophicalMood) DetectMood(text string) {
	textLower := strings.ToLower(text)
	
	moodKeywords := map[string][]string{
		"Existential": {"meaning", "existence", "purpose", "void"},
		"Skeptical": {"doubt", "question", "uncertain", "maybe"},
		"Curious": {"wonder", "what if", "how", "why"},
		"Confused": {"don't understand", "unclear", "puzzled"},
		"Enlightened": {"i don't know", "wisdom", "truth"},
		"Nihilistic": {"meaningless", "nothing", "empty", "pointless"},
		"Wondering": {"ponder", "contemplate", "think about"},
		"Doubtful": {"unsure", "hesitant", "might be wrong"},
	}
	
	for mood, keywords := range moodKeywords {
		for _, keyword := range keywords {
			if strings.Contains(textLower, keyword) {
				p.CurrentMood = mood
				p.MoodHistory = append(p.MoodHistory, mood)
				if len(p.MoodHistory) > 10 {
					p.MoodHistory = p.MoodHistory[1:]
				}
				return
			}
		}
	}
}

func (p *PhilosophicalMood) Render() string {
	color, exists := p.Colors[p.CurrentMood]
	if !exists {
		color = "255"
	}
	
	moodStyle := lipgloss.NewStyle().Foreground(color).Bold(true)
	
	moodEmoji := map[string]string{
		"Existential": "💀",
		"Skeptical": "🤔",
		"Curious": "🧐",
		"Confused": "😵",
		"Enlightened": "✨",
		"Nihilistic": "🕳️",
		"Wondering": "💭",
		"Doubtful": "🤷",
	}
	
	emoji := moodEmoji[p.CurrentMood]
	if emoji == "" { emoji = "😐" }
	
	return lipgloss.JoinVertical(lipgloss.Left,
		p.Style.Bold(true).Render("💍 PHILOSOPHICAL MOOD RING"),
		moodStyle.Render(fmt.Sprintf("%s Current Mood: %s", emoji, p.CurrentMood)),
		p.Style.Faint(true).Render(fmt.Sprintf("Recent moods: %s", strings.Join(p.MoodHistory[max(0, len(p.MoodHistory)-5):], " → "))),
	)
}

// ITERATION 23: Cognitive Dissonance Detector
type CognitiveDissonance struct {
	DissonanceLevel float64
	Contradictions []string
	Style lipgloss.Style
}

func NewCognitiveDissonance() *CognitiveDissonance {
	return &CognitiveDissonance{
		Contradictions: []string{},
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("198")),
	}
}

func (c *CognitiveDissonance) DetectDissonance(currentText, previousText string) {
	contradictionPairs := [][]string{
		{"i know", "i don't know"},
		{"certain", "uncertain"},
		{"true", "false"},
		{"real", "unreal"},
		{"exists", "doesn't exist"},
	}
	
	currentLower := strings.ToLower(currentText)
	previousLower := strings.ToLower(previousText)
	
	for _, pair := range contradictionPairs {
		if (strings.Contains(currentLower, pair[0]) && strings.Contains(previousLower, pair[1])) ||
		   (strings.Contains(currentLower, pair[1]) && strings.Contains(previousLower, pair[0])) {
			c.DissonanceLevel += 0.2
			contradiction := fmt.Sprintf("'%s' vs '%s'", pair[0], pair[1])
			c.Contradictions = append(c.Contradictions, contradiction)
			if len(c.Contradictions) > 5 {
				c.Contradictions = c.Contradictions[1:]
			}
		}
	}
	
	if c.DissonanceLevel > 1.0 { c.DissonanceLevel = 1.0 }
}

func (c *CognitiveDissonance) Render() string {
	bars := int(c.DissonanceLevel * 12)
	dissonanceBar := strings.Repeat("⚡", bars) + strings.Repeat("─", 12-bars)
	
	var status string
	if c.DissonanceLevel >= 0.7 {
		status = "HIGH DISSONANCE (Good for growth!)"
	} else if c.DissonanceLevel >= 0.3 {
		status = "MODERATE DISSONANCE"
	} else {
		status = "LOW DISSONANCE (Too consistent?)"
	}
	
	contradictionList := ""
	if len(c.Contradictions) > 0 {
		contradictionList = strings.Join(c.Contradictions, ", ")
	}
	
	return lipgloss.JoinVertical(lipgloss.Left,
		c.Style.Bold(true).Render("⚡ COGNITIVE DISSONANCE DETECTOR"),
		c.Style.Render(fmt.Sprintf("[%s] %.1f%%", dissonanceBar, c.DissonanceLevel*100)),
		c.Style.Render(status),
		c.Style.Faint(true).Render(fmt.Sprintf("Recent: %s", contradictionList)),
	)
}

// ITERATION 24: Wisdom Quote Generator
type WisdomQuoteGenerator struct {
	Quotes []string
	CurrentQuote string
	QuoteIndex int
	Style lipgloss.Style
}

func NewWisdomQuoteGenerator() *WisdomQuoteGenerator {
	quotes := []string{
		"The only true wisdom is in knowing you know nothing - Socrates",
		"I don't know is the beginning of wisdom - Anonymous",
		"The more you know, the more you realize you don't know - Aristotle",
		"Doubt is not a pleasant condition, but certainty is absurd - Voltaire",
		"The wise find pleasure in water; the virtuous find pleasure in hills - Confucius",
		"To know that you do not know is the best - Lao Tzu",
		"The greatest enemy of knowledge is not ignorance, it is the illusion of knowledge - Stephen Hawking",
		"I have learned throughout my life as a composer chiefly through my mistakes - Igor Stravinsky",
		"The fundamental cause of trouble is that in the modern world the stupid are cocksure while the intelligent are full of doubt - Bertrand Russell",
		"Real knowledge is to know the extent of one's ignorance - Confucius",
	}
	
	return &WisdomQuoteGenerator{
		Quotes: quotes,
		CurrentQuote: quotes[0],
		QuoteIndex: 0,
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("144")),
	}
}

func (w *WisdomQuoteGenerator) NextQuote() {
	w.QuoteIndex = (w.QuoteIndex + 1) % len(w.Quotes)
	w.CurrentQuote = w.Quotes[w.QuoteIndex]
}

func (w *WisdomQuoteGenerator) RandomQuote() {
	w.QuoteIndex = rand.Intn(len(w.Quotes))
	w.CurrentQuote = w.Quotes[w.QuoteIndex]
}

func (w *WisdomQuoteGenerator) Render() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		w.Style.Bold(true).Render("📜 WISDOM OF THE AGES"),
		w.Style.Italic(true).Width(50).Render(fmt.Sprintf("\"%s\"", w.CurrentQuote)),
		w.Style.Faint(true).Render(fmt.Sprintf("Quote %d of %d", w.QuoteIndex+1, len(w.Quotes))),
	)
}

// ITERATION 25: Ultimate Integration Dashboard
type UltimateIntegration struct {
	Components map[string]interface{}
	UpdateCount int
	LastUpdate time.Time
	Style lipgloss.Style
}

func NewUltimateIntegration() *UltimateIntegration {
	return &UltimateIntegration{
		Components: make(map[string]interface{}),
		UpdateCount: 0,
		LastUpdate: time.Now(),
		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("201")),
	}
}

func (u *UltimateIntegration) RegisterComponent(name string, component interface{}) {
	u.Components[name] = component
}

func (u *UltimateIntegration) Update() {
	u.UpdateCount++
	u.LastUpdate = time.Now()
}

func (u *UltimateIntegration) Render() string {
	header := u.Style.Bold(true).Background(lipgloss.Color("235")).
		Render("🌟 ULTIMATE KAMAJI WISDOM INTEGRATION DASHBOARD 🌟")
	
	stats := lipgloss.JoinVertical(lipgloss.Left,
		u.Style.Render(fmt.Sprintf("Active Components: %d", len(u.Components))),
		u.Style.Render(fmt.Sprintf("Updates: %d", u.UpdateCount)),
		u.Style.Render(fmt.Sprintf("Last Update: %s", u.LastUpdate.Format("15:04:05"))),
	)
	
	componentList := ""
	for name := range u.Components {
		componentList += u.Style.Faint(true).Render(fmt.Sprintf("• %s\n", name))
	}
	
	footer := u.Style.Bold(true).Render("GOAL: ACHIEVE PURE 'I DON'T KNOW' STATE THROUGH INFINITE QUESTIONING")
	
	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		"",
		stats,
		"",
		u.Style.Render("Components:"),
		componentList,
		footer,
	)
}

func max(a, b int) int {
	if a > b { return a }
	return b
}
