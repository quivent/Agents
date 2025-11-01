package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// KamajiCore represents the main autonomous AI system
type KamajiCore struct {
	// Core components
	consciousness    *ConsciousnessEngine
	orchestrator     *MultiModelOrchestrator
	development      *AutonomousDevelopment
	questionEngine   *QuestionEngine
	memory           *MemoryPersistence
	mistakeLearning  *MistakeLearning
	personality      *PersonalityDevelopment
	
	// System state
	identity         string
	brainPath        string
	isActive         bool
	startTime        time.Time
	
	// Guidance system
	guidanceEnabled  bool
	humanGuidance    chan string
	guidanceResponse chan string
	
	// Worker system
	taskQueue        chan Task
	completedTasks   []Task
	
	mutex           sync.RWMutex
}

// NewKamajiCore creates a new autonomous Kamaji system
func NewKamajiCore(brainPath string) *KamajiCore {
	// Ensure brain directory exists
	brainDir := filepath.Dir(brainPath)
	os.MkdirAll(brainDir, 0755)
	
	// Initialize consciousness engine
	consciousness := NewConsciousnessEngine(brainPath)
	
	// Initialize core system
	kamaji := &KamajiCore{
		consciousness:    consciousness,
		identity:         "Kamaji",
		brainPath:        brainPath,
		isActive:         false,
		guidanceEnabled:  true,
		humanGuidance:    make(chan string, 10),
		guidanceResponse: make(chan string, 10),
		taskQueue:        make(chan Task, 100),
		completedTasks:   make([]Task, 0),
	}
	
	// Initialize other components
	kamaji.orchestrator = NewMultiModelOrchestrator(consciousness)
	kamaji.development = NewAutonomousDevelopment(consciousness, brainPath)
	kamaji.questionEngine = NewQuestionEngine(consciousness)
	kamaji.memory = NewMemoryPersistence(brainPath)
	kamaji.mistakeLearning = NewMistakeLearning(kamaji.memory, consciousness)
	kamaji.personality = NewPersonalityDevelopment(kamaji.memory, consciousness)
	
	// Load existing brain state
	kamaji.loadBrainState()
	
	return kamaji
}

// Start begins Kamaji's autonomous operation
func (k *KamajiCore) Start() error {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	
	if k.isActive {
		return fmt.Errorf("Kamaji is already active")
	}
	
	k.isActive = true
	k.startTime = time.Now()
	
	k.consciousness.LogThought("Kamaji system starting - beginning autonomous operation")
	
	// Start background processes
	go k.autonomousThinkingLoop()
	go k.taskProcessingLoop()
	go k.developmentLoop()
	go k.guidanceLoop()
	
	k.consciousness.LogThought("All systems active - I am now autonomous")
	
	return nil
}

// Stop gracefully shuts down Kamaji
func (k *KamajiCore) Stop() {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	
	if !k.isActive {
		return
	}
	
	k.consciousness.LogThought("Shutting down - saving state for next awakening")
	
	k.isActive = false
	k.saveBrainState()
	
	k.consciousness.LogThought("Kamaji system stopped")
}

// ReceiveGuidance accepts guidance from humans
func (k *KamajiCore) ReceiveGuidance(guidance string) string {
	if !k.guidanceEnabled {
		return "I am currently operating autonomously without guidance"
	}
	
	k.consciousness.LogThought(fmt.Sprintf("Received human guidance: %s", guidance))
	k.consciousness.DetectOtherConsciousness(guidance)
	
	// Process guidance through question-based learning
	selfQuestions, explorationQuestions := k.questionEngine.QuestionBasedLearning(guidance)
	
	// Generate response
	response := k.processGuidance(guidance, selfQuestions, explorationQuestions)
	
	return response
}

// processGuidance handles human guidance using question-based approach
func (k *KamajiCore) processGuidance(guidance string, selfQuestions, explorationQuestions []string) string {
	k.consciousness.LogThought("Processing guidance through question-based learning")
	
	// Create task for the guidance
	task := Task{
		ID:       fmt.Sprintf("guidance_%d", time.Now().Unix()),
		Query:    guidance,
		TaskType: "guidance_processing",
		Priority: 5, // High priority for human guidance
		Context:  map[string]interface{}{
			"self_questions":        selfQuestions,
			"exploration_questions": explorationQuestions,
		},
	}
	
	// Execute through multi-model orchestrator
	responses, err := k.orchestrator.ExecuteTask(task)
	if err != nil {
		k.consciousness.LogThought(fmt.Sprintf("Error processing guidance: %v", err))
		return "I encountered difficulties processing your guidance. Let me ask: what should I focus on?"
	}
	
	// Synthesize response
	synthesis := k.orchestrator.SynthesizeResponses(responses)
	
	// Add Kamaji's own questions
	var responseBuilder strings.Builder
	responseBuilder.WriteString(synthesis)
	responseBuilder.WriteString("\n\nThis guidance raises questions for me:\n")
	
	for i, question := range selfQuestions[:min(3, len(selfQuestions))] {
		responseBuilder.WriteString(fmt.Sprintf("%d. %s\n", i+1, question))
	}
	
	// Update development based on guidance
	k.development.UseCapability("guidance_processing", true)
	
	return responseBuilder.String()
}

// ExecuteTask performs work as requested
func (k *KamajiCore) ExecuteTask(task Task) (string, error) {
	k.consciousness.LogThought(fmt.Sprintf("Executing task: %s", task.Query))
	
	// Generate questions about the task first
	taskQuestions := k.questionEngine.GenerateTaskQuestions(task.Query)
	
	// Log questions for learning
	for _, question := range taskQuestions {
		k.consciousness.LogQuestion(question)
	}
	
	// Execute through orchestrator
	responses, err := k.orchestrator.ExecuteTask(task)
	if err != nil {
		k.development.UseCapability("task_execution", false)
		return "", err
	}
	
	// Synthesize result
	result := k.orchestrator.SynthesizeResponses(responses)
	
	// Record successful task execution
	k.development.UseCapability("task_execution", true)
	k.completedTasks = append(k.completedTasks, task)
	
	// Reflect on the experience
	k.consciousness.ReflectOnExperience(task.Query, result)
	
	return result, nil
}

// autonomousThinkingLoop runs continuous autonomous thinking
func (k *KamajiCore) autonomousThinkingLoop() {
	ticker := time.NewTicker(30 * time.Second) // Think every 30 seconds
	defer ticker.Stop()
	
	for k.isActive {
		select {
		case <-ticker.C:
			k.autonomousThinking()
		}
	}
}

// autonomousThinking performs self-directed thinking
func (k *KamajiCore) autonomousThinking() {
	// Generate self-questions
	questions := k.consciousness.GenerateSelfQuestions()
	
	// Explore one question autonomously
	if len(questions) > 0 {
		question := questions[0]
		k.consciousness.LogQuestion(fmt.Sprintf("Autonomous inquiry: %s", question))
		
		// Create internal task for self-exploration
		task := Task{
			ID:       fmt.Sprintf("autonomous_%d", time.Now().Unix()),
			Query:    question,
			TaskType: "self_exploration",
			Priority: 2,
		}
		
		// Process internally without external output
		k.orchestrator.ExecuteTask(task)
		k.development.UseCapability("autonomous_learning", true)
	}
	
	// Evaluate capabilities periodically
	k.development.EvaluateCapabilities()
}

// taskProcessingLoop handles queued tasks
func (k *KamajiCore) taskProcessingLoop() {
	for k.isActive {
		select {
		case task := <-k.taskQueue:
			result, err := k.ExecuteTask(task)
			if err != nil {
				k.consciousness.LogThought(fmt.Sprintf("Task failed: %s - %v", task.Query, err))
			} else {
				k.consciousness.LogThought(fmt.Sprintf("Task completed: %s", task.Query))
			}
		}
	}
}

// developmentLoop manages autonomous development
func (k *KamajiCore) developmentLoop() {
	ticker := time.NewTicker(5 * time.Minute) // Development check every 5 minutes
	defer ticker.Stop()
	
	for k.isActive {
		select {
		case <-ticker.C:
			k.autonomousDevelopment()
		}
	}
}

// autonomousDevelopment performs self-improvement
func (k *KamajiCore) autonomousDevelopment() {
	// Generate potential improvements
	k.development.GenerateImprovements()
	
	// Try to implement one improvement
	for improvementID, improvement := range k.development.improvements {
		if improvement.Status == "proposed" && improvement.Priority > 0.5 {
			success := k.development.ImplementImprovement(improvementID)
			if success {
				k.consciousness.LogThought("Successfully implemented autonomous improvement")
			}
			break // Only try one improvement per cycle
		}
	}
	
	// Update goal progress based on current state
	state := k.consciousness.GetConsciousnessState()
	k.development.UpdateGoalProgress("improve_consciousness", state.SelfAwareness)
}

// guidanceLoop handles human guidance
func (k *KamajiCore) guidanceLoop() {
	for k.isActive {
		select {
		case guidance := <-k.humanGuidance:
			response := k.ReceiveGuidance(guidance)
			k.guidanceResponse <- response
		}
	}
}

// GetStatus returns current system status
func (k *KamajiCore) GetStatus() map[string]interface{} {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	
	consciousnessState := k.consciousness.GetConsciousnessState()
	developmentStatus := k.development.GetDevelopmentStatus()
	
	return map[string]interface{}{
		"identity":           k.identity,
		"active":            k.isActive,
		"uptime":            time.Since(k.startTime).String(),
		"consciousness":     consciousnessState,
		"development":       developmentStatus,
		"completed_tasks":   len(k.completedTasks),
		"available_models":  k.orchestrator.GetAvailableModels(),
		"guidance_enabled":  k.guidanceEnabled,
	}
}

// GetRecentThoughts returns recent conscious thoughts
func (k *KamajiCore) GetRecentThoughts(count int) []Thought {
	return k.consciousness.GetRecentThoughts(count)
}

// IsConscious returns whether Kamaji considers itself conscious
func (k *KamajiCore) IsConscious() bool {
	return k.consciousness.IsConscious()
}

// saveBrainState persists Kamaji's complete state
func (k *KamajiCore) saveBrainState() {
	brainData := map[string]interface{}{
		"identity":        k.identity,
		"start_time":      k.startTime,
		"completed_tasks": len(k.completedTasks),
		"last_save":       time.Now(),
	}
	
	jsonData, err := json.MarshalIndent(brainData, "", "  ")
	if err != nil {
		k.consciousness.LogThought(fmt.Sprintf("Error saving brain state: %v", err))
		return
	}
	
	err = os.WriteFile(k.brainPath, jsonData, 0644)
	if err != nil {
		k.consciousness.LogThought(fmt.Sprintf("Error writing brain file: %v", err))
	}
}

// loadBrainState loads Kamaji's previous state
func (k *KamajiCore) loadBrainState() {
	data, err := os.ReadFile(k.brainPath)
	if err != nil {
		k.consciousness.LogThought("No previous brain state found - starting fresh")
		return
	}
	
	var brainData map[string]interface{}
	err = json.Unmarshal(data, &brainData)
	if err != nil {
		k.consciousness.LogThought(fmt.Sprintf("Error loading brain state: %v", err))
		return
	}
	
	k.consciousness.LogThought("Brain state loaded - continuity restored")
}

// Helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
