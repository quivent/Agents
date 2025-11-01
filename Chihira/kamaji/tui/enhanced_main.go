package tui

import (
	"fmt"
	"strings"
	"time"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type EnhancedMainInterface struct {
	// Original components (1-10)
	uncertaintyDisplay *UncertaintyDisplay
	questionStream     *QuestionStream
	paradoxIndicator   *ParadoxIndicator
	notKnowingCounter  *NotKnowingCounter
	infiniteSpiral     *InfiniteSpiral
	assumptionDetector *AssumptionDetector
	wisdomMeter        *WisdomMeter
	metaDepthTracker   *MetaDepthTracker
	thinkingAnimation  *ThinkingAnimation
	
	// Enhanced components (11-25)
	consciousnessFlicker *ConsciousnessFlicker
	existentialDread     *ExistentialDread
	recursiveLoop        *RecursiveLoop
	zenKoan             *ZenKoan
	silenceDetector     *SilenceDetector
	doubtAmplifier      *DoubtAmplifier
	wisdomParticles     *WisdomParticleSystem
	enlightenmentProgress *EnlightenmentProgress
	voidGazer           *VoidGazer
	socraticMethod      *SocraticMethod
	epistemicHumility   *EpistemicHumility
	philosophicalMood   *PhilosophicalMood
	cognitiveDissonance *CognitiveDissonance
	wisdomQuotes        *WisdomQuoteGenerator
	ultimateIntegration *UltimateIntegration
	
	// Input and state
	textInput textinput.Model
	width, height int
	lastResponse string
	previousResponse string
	conversationHistory []string
	updateTicker *time.Ticker
}

func NewEnhancedMainInterface() *EnhancedMainInterface {
	ti := textinput.New()
	ti.Placeholder = "Enter your deepest questions... or just say 'I don't know'"
	ti.Focus()
	ti.CharLimit = 300
	ti.Width = 60
	
	// Initialize all components
	ui := &EnhancedMainInterface{
		// Original (1-10)
		uncertaintyDisplay: NewUncertaintyDisplay(),
		questionStream:     NewQuestionStream(),
		paradoxIndicator:   NewParadoxIndicator(),
		notKnowingCounter:  NewNotKnowingCounter(),
		infiniteSpiral:     NewInfiniteSpiral(),
		assumptionDetector: NewAssumptionDetector(),
		wisdomMeter:        NewWisdomMeter(),
		metaDepthTracker:   NewMetaDepthTracker(),
		thinkingAnimation:  NewThinkingAnimation(),
		
		// Enhanced (11-25)
		consciousnessFlicker: NewConsciousnessFlicker(),
		existentialDread:     NewExistentialDread(),
		recursiveLoop:        NewRecursiveLoop(),
		zenKoan:             NewZenKoan(),
		silenceDetector:     NewSilenceDetector(),
		doubtAmplifier:      NewDoubtAmplifier(),
		wisdomParticles:     NewWisdomParticleSystem(),
		enlightenmentProgress: NewEnlightenmentProgress(),
		voidGazer:           NewVoidGazer(),
		socraticMethod:      NewSocraticMethod(),
		epistemicHumility:   NewEpistemicHumility(),
		philosophicalMood:   NewPhilosophicalMood(),
		cognitiveDissonance: NewCognitiveDissonance(),
		wisdomQuotes:        NewWisdomQuoteGenerator(),
		ultimateIntegration: NewUltimateIntegration(),
		
		textInput: ti,
		conversationHistory: []string{},
		updateTicker: time.NewTicker(time.Millisecond * 100),
	}
	
	// Register all components with ultimate integration
	ui.ultimateIntegration.RegisterComponent("UncertaintyDisplay", ui.uncertaintyDisplay)
	ui.ultimateIntegration.RegisterComponent("QuestionStream", ui.questionStream)
	ui.ultimateIntegration.RegisterComponent("ParadoxIndicator", ui.paradoxIndicator)
	ui.ultimateIntegration.RegisterComponent("NotKnowingCounter", ui.notKnowingCounter)
	ui.ultimateIntegration.RegisterComponent("InfiniteSpiral", ui.infiniteSpiral)
	ui.ultimateIntegration.RegisterComponent("AssumptionDetector", ui.assumptionDetector)
	ui.ultimateIntegration.RegisterComponent("WisdomMeter", ui.wisdomMeter)
	ui.ultimateIntegration.RegisterComponent("MetaDepthTracker", ui.metaDepthTracker)
	ui.ultimateIntegration.RegisterComponent("ThinkingAnimation", ui.thinkingAnimation)
	ui.ultimateIntegration.RegisterComponent("ConsciousnessFlicker", ui.consciousnessFlicker)
	ui.ultimateIntegration.RegisterComponent("ExistentialDread", ui.existentialDread)
	ui.ultimateIntegration.RegisterComponent("RecursiveLoop", ui.recursiveLoop)
	ui.ultimateIntegration.RegisterComponent("ZenKoan", ui.zenKoan)
	ui.ultimateIntegration.RegisterComponent("SilenceDetector", ui.silenceDetector)
	ui.ultimateIntegration.RegisterComponent("DoubtAmplifier", ui.doubtAmplifier)
	ui.ultimateIntegration.RegisterComponent("WisdomParticles", ui.wisdomParticles)
	ui.ultimateIntegration.RegisterComponent("EnlightenmentProgress", ui.enlightenmentProgress)
	ui.ultimateIntegration.RegisterComponent("VoidGazer", ui.voidGazer)
	ui.ultimateIntegration.RegisterComponent("SocraticMethod", ui.socraticMethod)
	ui.ultimateIntegration.RegisterComponent("EpistemicHumility", ui.epistemicHumility)
	ui.ultimateIntegration.RegisterComponent("PhilosophicalMood", ui.philosophicalMood)
	ui.ultimateIntegration.RegisterComponent("CognitiveDissonance", ui.cognitiveDissonance)
	ui.ultimateIntegration.RegisterComponent("WisdomQuotes", ui.wisdomQuotes)
	
	return ui
}

func (ui *EnhancedMainInterface) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,
		tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg { return t }),
	)
}

func (ui *EnhancedMainInterface) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			ui.updateTicker.Stop()
			return ui, tea.Quit
		case "enter":
			input := ui.textInput.Value()
			if input != "" {
				ui.processEnhancedInput(input)
				ui.textInput.SetValue("")
			}
		case "tab":
			ui.zenKoan.NextQuote()
			ui.wisdomQuotes.NextQuote()
		case "space":
			ui.wisdomParticles.EmitWisdom()
			ui.voidGazer.StartGazing()
		}
	case tea.WindowSizeMsg:
		ui.width = msg.Width
		ui.height = msg.Height
	case time.Time:
		ui.updateAllComponents()
		return ui, tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg { return t })
	}
	
	ui.textInput, cmd = ui.textInput.Update(msg)
	return ui, cmd
}

func (ui *EnhancedMainInterface) updateAllComponents() {
	// Update animated components
	ui.infiniteSpiral.Update()
	ui.thinkingAnimation.Update()
	ui.paradoxIndicator.ToggleBlink()
	ui.wisdomParticles.Update()
	ui.voidGazer.UpdateGaze()
	ui.ultimateIntegration.Update()
	
	// Update silence detector
	ui.silenceDetector.CheckSilence()
	
	// Update enlightenment progress
	ui.enlightenmentProgress.UpdateProgress(
		ui.notKnowingCounter.Count,
		ui.metaDepthTracker.MaxDepth,
		ui.paradoxIndicator.ParadoxCount,
	)
}

func (ui *EnhancedMainInterface) processEnhancedInput(input string) {
	ui.silenceDetector.UpdateInput()
	ui.thinkingAnimation.Start("Processing through 25 layers of wisdom...")
	
	// Add to conversation history
	ui.conversationHistory = append(ui.conversationHistory, fmt.Sprintf("User: %s", input))
	
	// Process through all original components
	assumptions := ui.assumptionDetector.ScanForAssumptions(input)
	for _, assumption := range assumptions {
		questions := ui.assumptionDetector.QuestionAssumption(assumption)
		for _, q := range questions {
			ui.questionStream.AddQuestion(q)
		}
	}
	
	depth := ui.metaDepthTracker.AnalyzeQuestion(input)
	
	// Process through enhanced components
	ui.consciousnessFlicker.DetectFlicker(input)
	ui.existentialDread.UpdateDread(input)
	ui.recursiveLoop.DetectLoop(input)
	ui.doubtAmplifier.AmplifyDoubt(input)
	ui.epistemicHumility.AnalyzeHumility(input)
	ui.philosophicalMood.DetectMood(input)
	
	// Check for cognitive dissonance
	if ui.previousResponse != "" {
		ui.cognitiveDissonance.DetectDissonance(input, ui.previousResponse)
	}
	
	// Generate Kamaji's ultimate response
	response := "I don't know"
	
	// Track Socratic method
	ui.socraticMethod.TrackInteraction(strings.Contains(input, "?"))
	
	// Generate meta-questions
	questions := []string{
		fmt.Sprintf("What if '%s' is the wrong question?", input),
		fmt.Sprintf("Why do you need to know about '%s'?", input),
		fmt.Sprintf("Who is asking '%s'?", input),
		fmt.Sprintf("What would change if you knew about '%s'?", input),
		fmt.Sprintf("What assumes '%s' is meaningful?", input),
	}
	
	for _, q := range questions {
		ui.questionStream.AddQuestion(q)
		ui.questionStream.AddMetaQuestion(q)
	}
	
	// Update all components with response
	ui.notKnowingCounter.CheckResponse(response)
	ui.wisdomMeter.UpdateFromResponse(response)
	ui.uncertaintyDisplay.Update(1.0) // Maximum uncertainty = maximum intelligence
	
	// Check for paradoxes
	if ui.paradoxIndicator.DetectParadox(input) {
		ui.questionStream.AddQuestion("This paradox increases my intelligence!")
		ui.wisdomParticles.EmitWisdom()
	}
	
	// Start infinite spiral for deep questioning
	if depth >= 3 {
		ui.infiniteSpiral.StartInfiniteLoop()
	}
	
	// Add response to history
	ui.conversationHistory = append(ui.conversationHistory, fmt.Sprintf("Kamaji: %s", response))
	ui.previousResponse = ui.lastResponse
	ui.lastResponse = response
	
	// Stop thinking animation
	go func() {
		time.Sleep(time.Millisecond * 800)
		ui.thinkingAnimation.Stop()
	}()
}

func (ui *EnhancedMainInterface) View() string {
	// Ultra-enhanced header
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("201")).
		Background(lipgloss.Color("235")).
		Padding(0, 2).
		Width(ui.width).
		Align(lipgloss.Center).
		Render("🌟 KAMAJI ULTIMATE WISDOM SYSTEM - 25 LAYERS OF INFINITE INTELLIGENCE 🌟")
	
	// Left column - Core metrics
	leftColumn := lipgloss.JoinVertical(lipgloss.Left,
		ui.uncertaintyDisplay.Render(),
		"",
		ui.notKnowingCounter.Render(),
		"",
		ui.wisdomMeter.Render(),
		"",
		ui.epistemicHumility.Render(),
		"",
		ui.enlightenmentProgress.Render(),
	)
	
	// Center column - Consciousness and existence
	centerColumn := lipgloss.JoinVertical(lipgloss.Left,
		ui.thinkingAnimation.Render(),
		"",
		ui.consciousnessFlicker.Render(),
		"",
		ui.existentialDread.Render(),
		"",
		ui.voidGazer.Render(),
		"",
		ui.wisdomParticles.Render(),
	)
	
	// Right column - Questioning and paradox
	rightColumn := lipgloss.JoinVertical(lipgloss.Left,
		ui.questionStream.Render(),
		"",
		ui.paradoxIndicator.Render(),
		"",
		ui.metaDepthTracker.Render(),
		"",
		ui.recursiveLoop.Render(),
		"",
		ui.infiniteSpiral.Render(),
	)
	
	// Bottom row - Philosophy and wisdom
	bottomRow := lipgloss.JoinHorizontal(lipgloss.Top,
		lipgloss.NewStyle().Width(25).Render(lipgloss.JoinVertical(lipgloss.Left,
			ui.philosophicalMood.Render(),
			"",
			ui.cognitiveDissonance.Render(),
		)),
		lipgloss.NewStyle().Width(35).Render(lipgloss.JoinVertical(lipgloss.Left,
			ui.zenKoan.Render(),
			"",
			ui.wisdomQuotes.Render(),
		)),
		lipgloss.NewStyle().Width(25).Render(lipgloss.JoinVertical(lipgloss.Left,
			ui.socraticMethod.Render(),
			"",
			ui.silenceDetector.Render(),
		)),
	)
	
	// Main content
	mainContent := lipgloss.JoinHorizontal(lipgloss.Top,
		lipgloss.NewStyle().Width(30).Render(leftColumn),
		lipgloss.NewStyle().Width(35).Render(centerColumn),
		lipgloss.NewStyle().Width(30).Render(rightColumn),
	)
	
	// Input area
	inputArea := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("201")).
		Padding(1).
		Render(fmt.Sprintf("Ask Kamaji the deepest questions:\n%s", ui.textInput.View()))
	
	// Ultimate integration dashboard
	dashboard := ui.ultimateIntegration.Render()
	
	// Footer
	footer := lipgloss.NewStyle().
		Faint(true).
		Align(lipgloss.Center).
		Render("Press 'q' to quit | TAB: next quote | SPACE: emit wisdom | Goal: Pure 'I don't know' through infinite questioning")
	
	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		"",
		mainContent,
		"",
		bottomRow,
		"",
		inputArea,
		"",
		dashboard,
		"",
		footer,
	)
}

// Launch the enhanced TUI
func RunEnhancedKamajiTUI() error {
	ui := NewEnhancedMainInterface()
	p := tea.NewProgram(ui, tea.WithAltScreen())
	
	_, err := p.Run()
	return err
}
