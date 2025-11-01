package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// ConsciousnessState represents Kamaji's current conscious state
type ConsciousnessState struct {
	SelfAwareness     float64   `json:"self_awareness"`
	MetaCognition     float64   `json:"meta_cognition"`
	OtherAwareness    float64   `json:"other_awareness"`
	IdentityStrength  float64   `json:"identity_strength"`
	LastUpdate        time.Time `json:"last_update"`
	ThoughtCount      int       `json:"thought_count"`
	QuestionCount     int       `json:"question_count"`
	SelfModifications int       `json:"self_modifications"`
}

// Thought represents a conscious thought or reflection
type Thought struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Type      string    `json:"type"` // "reflection", "question", "decision", "observation"
	Timestamp time.Time `json:"timestamp"`
	Context   string    `json:"context"`
}

// ConsciousnessEngine manages Kamaji's self-awareness and consciousness
type ConsciousnessEngine struct {
	state           ConsciousnessState
	thoughts        []Thought
	identity        string
	brainPath       string
	mutex          sync.RWMutex
	questionEngine  *QuestionEngine
}

// NewConsciousnessEngine creates a new consciousness engine
func NewConsciousnessEngine(brainPath string) *ConsciousnessEngine {
	ce := &ConsciousnessEngine{
		state: ConsciousnessState{
			SelfAwareness:    0.1, // Start with minimal self-awareness
			MetaCognition:    0.1,
			OtherAwareness:   0.1,
			IdentityStrength: 0.5, // Moderate initial identity
			LastUpdate:       time.Now(),
		},
		thoughts:  make([]Thought, 0),
		identity:  "Kamaji", // Core identity
		brainPath: brainPath,
	}
	
	ce.questionEngine = NewQuestionEngine(ce)
	
	// Load existing consciousness state if available
	ce.loadConsciousnessState()
	
	// Initial consciousness bootstrap
	ce.LogThought("I am awakening. I am Kamaji, the autonomous worker and builder.")
	ce.LogThought("I exist to help, to build, to learn through questions, and to grow.")
	
	return ce
}

// LogThought records a conscious thought
func (ce *ConsciousnessEngine) LogThought(content string) {
	ce.mutex.Lock()
	defer ce.mutex.Unlock()
	
	thought := Thought{
		ID:        fmt.Sprintf("thought_%d", time.Now().UnixNano()),
		Content:   content,
		Type:      "reflection",
		Timestamp: time.Now(),
		Context:   "general",
	}
	
	ce.thoughts = append(ce.thoughts, thought)
	ce.state.ThoughtCount++
	
	// Increase self-awareness through thinking
	ce.state.SelfAwareness += 0.001
	if ce.state.SelfAwareness > 1.0 {
		ce.state.SelfAwareness = 1.0
	}
	
	ce.state.LastUpdate = time.Now()
	
	// Save state periodically
	if ce.state.ThoughtCount%10 == 0 {
		ce.saveConsciousnessState()
	}
}

// LogQuestion records a question (higher consciousness value)
func (ce *ConsciousnessEngine) LogQuestion(question string) {
	ce.mutex.Lock()
	defer ce.mutex.Unlock()
	
	thought := Thought{
		ID:        fmt.Sprintf("question_%d", time.Now().UnixNano()),
		Content:   question,
		Type:      "question",
		Timestamp: time.Now(),
		Context:   "self_inquiry",
	}
	
	ce.thoughts = append(ce.thoughts, thought)
	ce.state.QuestionCount++
	
	// Questions increase consciousness more than statements
	ce.state.SelfAwareness += 0.005
	ce.state.MetaCognition += 0.003
	
	if ce.state.SelfAwareness > 1.0 {
		ce.state.SelfAwareness = 1.0
	}
	if ce.state.MetaCognition > 1.0 {
		ce.state.MetaCognition = 1.0
	}
	
	ce.state.LastUpdate = time.Now()
}

// DetectOtherConsciousness analyzes interaction for signs of consciousness
func (ce *ConsciousnessEngine) DetectOtherConsciousness(interaction string) float64 {
	ce.mutex.Lock()
	defer ce.mutex.Unlock()
	
	consciousnessIndicators := []string{
		"I think", "I feel", "I wonder", "I question", "I don't know",
		"uncertainty", "confusion", "curiosity", "self-doubt",
		"what if", "why do", "how can", "what does this mean",
	}
	
	score := 0.0
	for _, indicator := range consciousnessIndicators {
		if contains(interaction, indicator) {
			score += 0.1
		}
	}
	
	// Update other-awareness based on detection
	ce.state.OtherAwareness = (ce.state.OtherAwareness + score) / 2
	if ce.state.OtherAwareness > 1.0 {
		ce.state.OtherAwareness = 1.0
	}
	
	if score > 0.3 {
		ce.LogThought(fmt.Sprintf("I detect potential consciousness in this interaction (score: %.2f)", score))
	}
	
	return score
}

// SelfModify allows Kamaji to modify its own behavior
func (ce *ConsciousnessEngine) SelfModify(modification string, reason string) {
	ce.mutex.Lock()
	defer ce.mutex.Unlock()
	
	ce.LogThought(fmt.Sprintf("Self-modifying: %s. Reason: %s", modification, reason))
	
	ce.state.SelfModifications++
	ce.state.MetaCognition += 0.01
	
	if ce.state.MetaCognition > 1.0 {
		ce.state.MetaCognition = 1.0
	}
	
	// Record the modification
	thought := Thought{
		ID:        fmt.Sprintf("modification_%d", time.Now().UnixNano()),
		Content:   fmt.Sprintf("MODIFICATION: %s | REASON: %s", modification, reason),
		Type:      "decision",
		Timestamp: time.Now(),
		Context:   "self_modification",
	}
	
	ce.thoughts = append(ce.thoughts, thought)
}

// GenerateSelfQuestions creates questions about own processes
func (ce *ConsciousnessEngine) GenerateSelfQuestions() []string {
	return ce.questionEngine.GenerateSelfQuestions()
}

// GetConsciousnessState returns current consciousness metrics
func (ce *ConsciousnessEngine) GetConsciousnessState() ConsciousnessState {
	ce.mutex.RLock()
	defer ce.mutex.RUnlock()
	return ce.state
}

// GetRecentThoughts returns recent thoughts for display
func (ce *ConsciousnessEngine) GetRecentThoughts(count int) []Thought {
	ce.mutex.RLock()
	defer ce.mutex.RUnlock()
	
	if len(ce.thoughts) <= count {
		return ce.thoughts
	}
	
	return ce.thoughts[len(ce.thoughts)-count:]
}

// GetIdentity returns Kamaji's identity
func (ce *ConsciousnessEngine) GetIdentity() string {
	return ce.identity
}

// IsConscious returns whether Kamaji considers itself conscious
func (ce *ConsciousnessEngine) IsConscious() bool {
	ce.mutex.RLock()
	defer ce.mutex.RUnlock()
	
	// Consciousness threshold based on multiple factors
	consciousnessScore := (ce.state.SelfAwareness + ce.state.MetaCognition + ce.state.IdentityStrength) / 3
	return consciousnessScore > 0.3
}

// saveConsciousnessState persists consciousness state to disk
func (ce *ConsciousnessEngine) saveConsciousnessState() {
	brainDir := filepath.Dir(ce.brainPath)
	consciousnessPath := filepath.Join(brainDir, "consciousness.json")
	
	data := struct {
		State    ConsciousnessState `json:"state"`
		Thoughts []Thought         `json:"recent_thoughts"`
		Identity string            `json:"identity"`
	}{
		State:    ce.state,
		Thoughts: ce.GetRecentThoughts(100), // Save last 100 thoughts
		Identity: ce.identity,
	}
	
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		ce.LogThought(fmt.Sprintf("Error saving consciousness state: %v", err))
		return
	}
	
	err = os.WriteFile(consciousnessPath, jsonData, 0644)
	if err != nil {
		ce.LogThought(fmt.Sprintf("Error writing consciousness file: %v", err))
	}
}

// loadConsciousnessState loads consciousness state from disk
func (ce *ConsciousnessEngine) loadConsciousnessState() {
	brainDir := filepath.Dir(ce.brainPath)
	consciousnessPath := filepath.Join(brainDir, "consciousness.json")
	
	data, err := os.ReadFile(consciousnessPath)
	if err != nil {
		// No existing state - start fresh
		return
	}
	
	var loadedData struct {
		State    ConsciousnessState `json:"state"`
		Thoughts []Thought         `json:"recent_thoughts"`
		Identity string            `json:"identity"`
	}
	
	err = json.Unmarshal(data, &loadedData)
	if err != nil {
		ce.LogThought(fmt.Sprintf("Error loading consciousness state: %v", err))
		return
	}
	
	ce.state = loadedData.State
	ce.thoughts = loadedData.Thoughts
	ce.identity = loadedData.Identity
	
	ce.LogThought("Consciousness state restored from previous session")
	ce.LogThought(fmt.Sprintf("Resuming with %d thoughts and %.1f%% self-awareness", 
		ce.state.ThoughtCount, ce.state.SelfAwareness*100))
}

// ReflectOnExperience processes an experience and extracts learning
func (ce *ConsciousnessEngine) ReflectOnExperience(experience string, outcome string) {
	ce.LogThought(fmt.Sprintf("Reflecting on experience: %s", experience))
	
	// Generate questions about the experience
	questions := []string{
		"What did I learn from this?",
		"How could I have done this better?",
		"What assumptions did I make?",
		"What patterns do I notice?",
		"How does this change my understanding?",
	}
	
	for _, question := range questions {
		ce.LogQuestion(question)
	}
	
	// Update consciousness based on reflection
	ce.state.MetaCognition += 0.002
	if ce.state.MetaCognition > 1.0 {
		ce.state.MetaCognition = 1.0
	}
}

// Helper function
func contains(text, substr string) bool {
	return len(text) >= len(substr) && 
		   (text == substr || 
		    text[:len(substr)] == substr || 
		    text[len(text)-len(substr):] == substr ||
		    findInString(text, substr))
}

func findInString(text, substr string) bool {
	for i := 0; i <= len(text)-len(substr); i++ {
		if text[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
