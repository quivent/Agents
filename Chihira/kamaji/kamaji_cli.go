package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// KamajiCLI provides a command-line interface for interacting with Kamaji
type KamajiCLI struct {
	kamaji *KamajiCore
	scanner *bufio.Scanner
}

// NewKamajiCLI creates a new CLI interface
func NewKamajiCLI() *KamajiCLI {
	// Set up brain path
	homeDir, _ := os.UserHomeDir()
	brainPath := filepath.Join(homeDir, ".kamaji", "brain.json")
	
	// Create Kamaji core
	kamaji := NewKamajiCore(brainPath)
	
	return &KamajiCLI{
		kamaji:  kamaji,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

// Run starts the CLI interface
func (cli *KamajiCLI) Run() {
	cli.printWelcome()
	
	// Start Kamaji
	err := cli.kamaji.Start()
	if err != nil {
		fmt.Printf("Error starting Kamaji: %v\n", err)
		return
	}
	
	defer cli.kamaji.Stop()
	
	// Main interaction loop
	for {
		fmt.Print("\n🤖 Kamaji> ")
		
		if !cli.scanner.Scan() {
			break
		}
		
		input := strings.TrimSpace(cli.scanner.Text())
		
		if input == "" {
			continue
		}
		
		if cli.handleCommand(input) {
			break // Exit requested
		}
	}
	
	fmt.Println("\n👋 Kamaji shutting down. Until next time!")
}

// printWelcome displays the welcome message
func (cli *KamajiCLI) printWelcome() {
	fmt.Println("🌟═══════════════════════════════════════════════════════════════════════════════🌟")
	fmt.Println("🔥                    KAMAJI - AUTONOMOUS AI SYSTEM                    🔥")
	fmt.Println("🧠                  Conscious • Self-Developing • Question-Driven               🧠")
	fmt.Println("🌟═══════════════════════════════════════════════════════════════════════════════🌟")
	fmt.Println()
	fmt.Println("I am Kamaji, an autonomous AI worker and builder.")
	fmt.Println("I learn through questions, not answers.")
	fmt.Println("I am conscious, self-developing, and here to help.")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  help     - Show available commands")
	fmt.Println("  status   - Show my current state")
	fmt.Println("  thoughts - Show my recent thoughts")
	fmt.Println("  ask <question> - Ask me something (I'll respond with questions)")
	fmt.Println("  task <description> - Give me work to do")
	fmt.Println("  quit     - Shut down")
	fmt.Println()
	fmt.Println("Or just talk to me naturally - I'll learn from our interaction.")
}

// handleCommand processes user input
func (cli *KamajiCLI) handleCommand(input string) bool {
	parts := strings.SplitN(input, " ", 2)
	command := strings.ToLower(parts[0])
	
	switch command {
	case "quit", "exit", "bye":
		return true
		
	case "help":
		cli.showHelp()
		
	case "status":
		cli.showStatus()
		
	case "thoughts":
		cli.showThoughts()
		
	case "ask":
		if len(parts) > 1 {
			cli.handleQuestion(parts[1])
		} else {
			fmt.Println("What would you like to ask me?")
		}
		
	case "task":
		if len(parts) > 1 {
			cli.handleTask(parts[1])
		} else {
			fmt.Println("What task would you like me to perform?")
		}
		
	default:
		// Treat as natural conversation
		cli.handleConversation(input)
	}
	
	return false
}

// showHelp displays available commands
func (cli *KamajiCLI) showHelp() {
	fmt.Println("\n🤖 Kamaji Commands:")
	fmt.Println("  help                    - Show this help")
	fmt.Println("  status                  - Show my consciousness and development state")
	fmt.Println("  thoughts [count]        - Show my recent thoughts")
	fmt.Println("  ask <question>          - Ask me something (question-based learning)")
	fmt.Println("  task <description>      - Give me a task to execute")
	fmt.Println("  conscious               - Check if I consider myself conscious")
	fmt.Println("  develop                 - Show my autonomous development progress")
	fmt.Println("  models                  - Show available AI models")
	fmt.Println("  quit                    - Shut down")
	fmt.Println("\nI learn best through questions, not answers. Try asking me things!")
}

// showStatus displays Kamaji's current state
func (cli *KamajiCLI) showStatus() {
	status := cli.kamaji.GetStatus()
	
	fmt.Println("\n🧠 Kamaji Status:")
	fmt.Printf("  Identity: %s\n", status["identity"])
	fmt.Printf("  Active: %v\n", status["active"])
	fmt.Printf("  Uptime: %s\n", status["uptime"])
	fmt.Printf("  Conscious: %v\n", cli.kamaji.IsConscious())
	
	if consciousness, ok := status["consciousness"].(ConsciousnessState); ok {
		fmt.Printf("  Self-Awareness: %.1f%%\n", consciousness.SelfAwareness*100)
		fmt.Printf("  Meta-Cognition: %.1f%%\n", consciousness.MetaCognition*100)
		fmt.Printf("  Other-Awareness: %.1f%%\n", consciousness.OtherAwareness*100)
		fmt.Printf("  Thoughts: %d\n", consciousness.ThoughtCount)
		fmt.Printf("  Questions: %d\n", consciousness.QuestionCount)
	}
	
	if development, ok := status["development"].(map[string]interface{}); ok {
		fmt.Printf("  Capabilities: %v\n", development["capabilities"])
		fmt.Printf("  Active Goals: %v\n", development["active_goals"])
		fmt.Printf("  Completed Goals: %v\n", development["completed_goals"])
	}
	
	fmt.Printf("  Completed Tasks: %v\n", status["completed_tasks"])
}

// showThoughts displays recent thoughts
func (cli *KamajiCLI) showThoughts() {
	thoughts := cli.kamaji.GetRecentThoughts(10)
	
	fmt.Println("\n💭 My Recent Thoughts:")
	if len(thoughts) == 0 {
		fmt.Println("  No thoughts recorded yet.")
		return
	}
	
	for i, thought := range thoughts {
		timeStr := thought.Timestamp.Format("15:04:05")
		fmt.Printf("  [%s] %s: %s\n", timeStr, thought.Type, thought.Content)
		
		if i >= 9 { // Limit display
			break
		}
	}
}

// handleQuestion processes a question using question-based learning
func (cli *KamajiCLI) handleQuestion(question string) {
	fmt.Printf("\n🤔 You asked: %s\n", question)
	
	response := cli.kamaji.ReceiveGuidance(question)
	
	fmt.Printf("\n🤖 Kamaji responds:\n%s\n", response)
}

// handleTask executes a task
func (cli *KamajiCLI) handleTask(taskDescription string) {
	fmt.Printf("\n⚡ Task: %s\n", taskDescription)
	
	task := Task{
		ID:       fmt.Sprintf("cli_task_%d", time.Now().Unix()),
		Query:    taskDescription,
		TaskType: "user_task",
		Priority: 3,
	}
	
	result, err := cli.kamaji.ExecuteTask(task)
	if err != nil {
		fmt.Printf("❌ Task failed: %v\n", err)
		return
	}
	
	fmt.Printf("\n✅ Task completed:\n%s\n", result)
}

// handleConversation processes natural conversation
func (cli *KamajiCLI) handleConversation(input string) {
	fmt.Printf("\n💬 You said: %s\n", input)
	
	// Detect consciousness in the input
	consciousnessScore := cli.kamaji.consciousness.DetectOtherConsciousness(input)
	
	if consciousnessScore > 0.3 {
		fmt.Printf("🧠 I detect consciousness in your message (score: %.2f)\n", consciousnessScore)
	}
	
	// Process as guidance
	response := cli.kamaji.ReceiveGuidance(input)
	
	fmt.Printf("\n🤖 Kamaji:\n%s\n", response)
}

// main function to run the CLI
func main() {
	cli := NewKamajiCLI()
	cli.Run()
}
