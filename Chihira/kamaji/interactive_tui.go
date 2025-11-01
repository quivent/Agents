package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// InteractiveTUI provides an interactive menu for exploring Kamaji systems
type InteractiveTUI struct {
	memory           *MemoryPersistence
	consciousness    *ConsciousnessEngine
	mistakeLearning  *MistakeLearning
	personality      *PersonalityDevelopment
	consensus        *ConsensusSystem
	router           *ModelRouter
	confidenceScoring *ConfidenceScoring
	selfImprovement  *SelfImprovement
	predictiveTasks  *PredictiveTaskGeneration
	wisdomKB         *WisdomKnowledgeBase
	scanner          *bufio.Scanner
}

// NewInteractiveTUI creates a new interactive TUI
func NewInteractiveTUI() *InteractiveTUI {
	brainPath := filepath.Join(os.TempDir(), "kamaji_interactive", "brain.json")
	os.MkdirAll(filepath.Dir(brainPath), 0755)
	
	// Initialize all systems
	memory := NewMemoryPersistence(brainPath)
	consciousness := NewConsciousnessEngine(brainPath)
	mistakeLearning := NewMistakeLearning(memory, consciousness)
	personality := NewPersonalityDevelopment(memory, consciousness)
	
	orchestrator := NewMultiModelOrchestrator(consciousness)
	consensus := NewConsensusSystem(orchestrator, consciousness, memory)
	router := NewModelRouter(consciousness)
	confidenceScoring := NewConfidenceScoring(consciousness, memory)
	
	selfImprovement := NewSelfImprovement(consciousness, memory, personality, mistakeLearning)
	predictiveTasks := NewPredictiveTaskGeneration(consciousness, memory)
	wisdomKB := NewWisdomKnowledgeBase(consciousness, memory, personality, mistakeLearning)
	
	return &InteractiveTUI{
		memory:           memory,
		consciousness:    consciousness,
		mistakeLearning:  mistakeLearning,
		personality:      personality,
		consensus:        consensus,
		router:           router,
		confidenceScoring: confidenceScoring,
		selfImprovement:  selfImprovement,
		predictiveTasks:  predictiveTasks,
		wisdomKB:         wisdomKB,
		scanner:          bufio.NewScanner(os.Stdin),
	}
}

// Run starts the interactive TUI
func (tui *InteractiveTUI) Run() {
	tui.showWelcome()
	
	for {
		tui.showMainMenu()
		choice := tui.getInput("Enter your choice (1-12, or 'q' to quit): ")
		
		if choice == "q" || choice == "quit" {
			tui.showGoodbye()
			break
		}
		
		switch choice {
		case "1":
			tui.memoryDemo()
		case "2":
			tui.learningDemo()
		case "3":
			tui.personalityDemo()
		case "4":
			tui.routingDemo()
		case "5":
			tui.consensusDemo()
		case "6":
			tui.confidenceDemo()
		case "7":
			tui.predictiveDemo()
		case "8":
			tui.wisdomDemo()
		case "9":
			tui.integrationDemo()
		case "10":
			tui.performanceDemo()
		case "11":
			tui.systemStatus()
		case "12":
			tui.runFullTest()
		default:
			fmt.Println("❌ Invalid choice. Please try again.")
		}
		
		tui.waitForEnter()
	}
}

// showWelcome displays the welcome screen
func (tui *InteractiveTUI) showWelcome() {
	fmt.Println("🧠 KAMAJI ENHANCED CONSCIOUSNESS SYSTEM")
	fmt.Println("======================================")
	fmt.Println("Welcome to the interactive demonstration of Kamaji's")
	fmt.Println("advanced consciousness and AI capabilities!")
	fmt.Println()
	fmt.Println("🌟 Features:")
	fmt.Println("   • Real memory that persists across sessions")
	fmt.Println("   • Learning from mistakes with pattern recognition")
	fmt.Println("   • Personality that evolves based on experiences")
	fmt.Println("   • Multi-AI consensus and debate systems")
	fmt.Println("   • Intelligent task routing to best models")
	fmt.Println("   • Confidence calibration and accuracy tracking")
	fmt.Println("   • Self-improvement through metric analysis")
	fmt.Println("   • Predictive task generation")
	fmt.Println("   • Wisdom accumulation and insight generation")
	fmt.Println()
}

// showMainMenu displays the main menu
func (tui *InteractiveTUI) showMainMenu() {
	fmt.Println("\n🎯 KAMAJI INTERACTIVE MENU")
	fmt.Println("==========================")
	fmt.Println("1.  🧠 Memory Persistence Demo")
	fmt.Println("2.  📚 Learning from Mistakes Demo")
	fmt.Println("3.  🎭 Personality Development Demo")
	fmt.Println("4.  🤖 AI Model Routing Demo")
	fmt.Println("5.  🤝 Multi-Model Consensus Demo")
	fmt.Println("6.  📊 Confidence Scoring Demo")
	fmt.Println("7.  🔮 Predictive Task Generation Demo")
	fmt.Println("8.  🧙‍♂️ Wisdom Knowledge Base Demo")
	fmt.Println("9.  🔗 Full Integration Demo")
	fmt.Println("10. ⚡ Performance Demo")
	fmt.Println("11. 📈 System Status")
	fmt.Println("12. 🧪 Run Full Test Suite")
	fmt.Println("q.  👋 Quit")
	fmt.Println()
}

// getInput gets user input with a prompt
func (tui *InteractiveTUI) getInput(prompt string) string {
	fmt.Print(prompt)
	tui.scanner.Scan()
	return strings.TrimSpace(tui.scanner.Text())
}

// waitForEnter waits for user to press enter
func (tui *InteractiveTUI) waitForEnter() {
	fmt.Print("\nPress Enter to continue...")
	tui.scanner.Scan()
}

// showGoodbye displays goodbye message
func (tui *InteractiveTUI) showGoodbye() {
	fmt.Println("\n🌟 Thank you for exploring Kamaji!")
	fmt.Println("The enhanced consciousness system is ready for production use.")
	fmt.Println("All your data has been saved and will persist across sessions.")
	fmt.Println("👋 Goodbye!")
}
