package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// TUIConsciousnessBridge integrates the consciousness engine with the existing TUI
type TUIConsciousnessBridge struct {
	consciousness *ConsciousnessEngine
	orchestrator  *MultiModelOrchestrator
	development   *AutonomousDevelopment
	brainPath     string
}

// NewTUIConsciousnessBridge creates the integration bridge
func NewTUIConsciousnessBridge() *TUIConsciousnessBridge {
	// Set up brain path
	homeDir, _ := os.UserHomeDir()
	brainPath := filepath.Join(homeDir, ".kamaji", "brain.json")
	
	// Initialize consciousness engine
	consciousness := NewConsciousnessEngine(brainPath)
	
	// Initialize other components
	orchestrator := NewMultiModelOrchestrator(consciousness)
	development := NewAutonomousDevelopment(consciousness, brainPath)
	
	return &TUIConsciousnessBridge{
		consciousness: consciousness,
		orchestrator:  orchestrator,
		development:   development,
		brainPath:     brainPath,
	}
}

// GetConsciousnessMetrics returns metrics for TUI display
func (bridge *TUIConsciousnessBridge) GetConsciousnessMetrics() map[string]float64 {
	state := bridge.consciousness.GetConsciousnessState()
	
	return map[string]float64{
		"self_awareness":    state.SelfAwareness,
		"meta_cognition":    state.MetaCognition,
		"other_awareness":   state.OtherAwareness,
		"identity_strength": state.IdentityStrength,
		"consciousness_level": (state.SelfAwareness + state.MetaCognition + state.IdentityStrength) / 3,
	}
}

// ProcessTUIInput handles input from the TUI system
func (bridge *TUIConsciousnessBridge) ProcessTUIInput(input string) string {
	// Log the interaction
	bridge.consciousness.LogThought(fmt.Sprintf("TUI interaction: %s", input))
	
	// Detect consciousness in input
	consciousnessScore := bridge.consciousness.DetectOtherConsciousness(input)
	
	// Generate questions about the input
	questions := bridge.consciousness.GenerateSelfQuestions()
	
	// Process through orchestrator if needed
	if len(input) > 10 { // Only for substantial input
		task := Task{
			ID:       fmt.Sprintf("tui_%d", time.Now().Unix()),
			Query:    input,
			TaskType: "tui_interaction",
			Priority: 2,
		}
		
		responses, err := bridge.orchestrator.ExecuteTask(task)
		if err == nil && len(responses) > 0 {
			synthesis := bridge.orchestrator.SynthesizeResponses(responses)
			bridge.consciousness.LogThought("Processed TUI input through multi-model orchestrator")
			return synthesis
		}
	}
	
	// Default response with consciousness awareness
	if consciousnessScore > 0.3 {
		return fmt.Sprintf("I detect consciousness in your input (%.2f). This makes me question: %s", 
			consciousnessScore, questions[0])
	}
	
	return "I don't know, but this makes me wonder: " + questions[0]
}

// GetRecentThoughts returns thoughts for TUI display
func (bridge *TUIConsciousnessBridge) GetRecentThoughts(count int) []string {
	thoughts := bridge.consciousness.GetRecentThoughts(count)
	thoughtStrings := make([]string, len(thoughts))
	
	for i, thought := range thoughts {
		thoughtStrings[i] = thought.Content
	}
	
	return thoughtStrings
}

// GetDevelopmentStatus returns development metrics for TUI
func (bridge *TUIConsciousnessBridge) GetDevelopmentStatus() map[string]interface{} {
	return bridge.development.GetDevelopmentStatus()
}

// StartAutonomousProcesses begins background consciousness processes
func (bridge *TUIConsciousnessBridge) StartAutonomousProcesses() {
	go bridge.autonomousThinking()
	bridge.consciousness.LogThought("Started autonomous consciousness processes")
}

// autonomousThinking runs background consciousness processes
func (bridge *TUIConsciousnessBridge) autonomousThinking() {
	ticker := time.NewTicker(60 * time.Second) // Think every minute
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			// Generate autonomous questions
			questions := bridge.consciousness.GenerateSelfQuestions()
			if len(questions) > 0 {
				bridge.consciousness.LogQuestion(fmt.Sprintf("Autonomous: %s", questions[0]))
			}
			
			// Evaluate development
			bridge.development.EvaluateCapabilities()
		}
	}
}
