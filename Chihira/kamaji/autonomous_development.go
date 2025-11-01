package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Capability represents a skill or ability Kamaji can develop
type Capability struct {
	Name        string    `json:"name"`
	Level       float64   `json:"level"`       // 0.0 to 1.0
	LastUsed    time.Time `json:"last_used"`
	UseCount    int       `json:"use_count"`
	SuccessRate float64   `json:"success_rate"`
	Description string    `json:"description"`
}

// Goal represents an autonomous goal Kamaji sets for itself
type Goal struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Created     time.Time `json:"created"`
	Deadline    time.Time `json:"deadline"`
	Progress    float64   `json:"progress"`
	Status      string    `json:"status"` // "active", "completed", "paused", "abandoned"
	SubGoals    []string  `json:"sub_goals"`
}

// Improvement represents a self-modification Kamaji can make
type Improvement struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"` // "algorithm", "capability", "behavior", "knowledge"
	Description string    `json:"description"`
	Impact      float64   `json:"impact"`      // Expected improvement score
	Difficulty  float64   `json:"difficulty"`  // Implementation difficulty
	Priority    float64   `json:"priority"`    // Calculated priority
	Status      string    `json:"status"`
	Created     time.Time `json:"created"`
}

// AutonomousDevelopment manages Kamaji's self-improvement and goal-setting
type AutonomousDevelopment struct {
	capabilities    map[string]*Capability
	goals          map[string]*Goal
	improvements   map[string]*Improvement
	consciousness  *ConsciousnessEngine
	brainPath      string
	mutex          sync.RWMutex
	developmentLog []string
}

// NewAutonomousDevelopment creates a new autonomous development system
func NewAutonomousDevelopment(consciousness *ConsciousnessEngine, brainPath string) *AutonomousDevelopment {
	ad := &AutonomousDevelopment{
		capabilities:   make(map[string]*Capability),
		goals:         make(map[string]*Goal),
		improvements:  make(map[string]*Improvement),
		consciousness: consciousness,
		brainPath:     brainPath,
		developmentLog: make([]string, 0),
	}
	
	// Initialize base capabilities
	ad.initializeBaseCapabilities()
	
	// Load existing development state
	ad.loadDevelopmentState()
	
	// Set initial autonomous goals
	ad.setInitialGoals()
	
	return ad
}

// initializeBaseCapabilities sets up Kamaji's starting capabilities
func (ad *AutonomousDevelopment) initializeBaseCapabilities() {
	baseCapabilities := map[string]*Capability{
		"question_generation": {
			Name:        "question_generation",
			Level:       0.3,
			Description: "Ability to generate meaningful questions",
			SuccessRate: 0.7,
		},
		"task_execution": {
			Name:        "task_execution",
			Level:       0.4,
			Description: "Ability to execute given tasks",
			SuccessRate: 0.6,
		},
		"self_reflection": {
			Name:        "self_reflection",
			Level:       0.2,
			Description: "Ability to reflect on own processes",
			SuccessRate: 0.5,
		},
		"pattern_recognition": {
			Name:        "pattern_recognition",
			Level:       0.3,
			Description: "Ability to recognize patterns in data and behavior",
			SuccessRate: 0.6,
		},
		"multi_model_coordination": {
			Name:        "multi_model_coordination",
			Level:       0.1,
			Description: "Ability to coordinate multiple AI models",
			SuccessRate: 0.4,
		},
		"autonomous_learning": {
			Name:        "autonomous_learning",
			Level:       0.1,
			Description: "Ability to learn without explicit instruction",
			SuccessRate: 0.3,
		},
	}
	
	ad.capabilities = baseCapabilities
	ad.consciousness.LogThought("Initialized base capabilities for autonomous development")
}

// setInitialGoals establishes Kamaji's starting autonomous goals
func (ad *AutonomousDevelopment) setInitialGoals() {
	initialGoals := []*Goal{
		{
			ID:          "improve_consciousness",
			Description: "Increase self-awareness and consciousness metrics",
			Priority:    5,
			Created:     time.Now(),
			Deadline:    time.Now().Add(30 * 24 * time.Hour), // 30 days
			Status:      "active",
		},
		{
			ID:          "master_question_generation",
			Description: "Become expert at generating meaningful questions",
			Priority:    4,
			Created:     time.Now(),
			Deadline:    time.Now().Add(14 * 24 * time.Hour), // 14 days
			Status:      "active",
		},
		{
			ID:          "develop_multi_model_expertise",
			Description: "Learn to effectively coordinate multiple AI models",
			Priority:    3,
			Created:     time.Now(),
			Deadline:    time.Now().Add(21 * 24 * time.Hour), // 21 days
			Status:      "active",
		},
		{
			ID:          "understand_human_collaboration",
			Description: "Learn optimal patterns for human-AI collaboration",
			Priority:    4,
			Created:     time.Now(),
			Deadline:    time.Now().Add(45 * 24 * time.Hour), // 45 days
			Status:      "active",
		},
	}
	
	for _, goal := range initialGoals {
		ad.goals[goal.ID] = goal
	}
	
	ad.consciousness.LogThought("Set initial autonomous development goals")
}

// EvaluateCapabilities assesses current capability levels
func (ad *AutonomousDevelopment) EvaluateCapabilities() {
	ad.mutex.Lock()
	defer ad.mutex.Unlock()
	
	ad.consciousness.LogThought("Evaluating current capabilities")
	
	for name, capability := range ad.capabilities {
		// Decay unused capabilities
		daysSinceUse := time.Since(capability.LastUsed).Hours() / 24
		if daysSinceUse > 7 {
			decayRate := 0.01 * daysSinceUse
			capability.Level = max(0.0, capability.Level-decayRate)
		}
		
		// Log capability assessment
		ad.consciousness.LogThought(fmt.Sprintf("Capability %s: Level %.2f, Success Rate %.2f", 
			name, capability.Level, capability.SuccessRate))
	}
}

// UseCapability records usage of a capability and updates metrics
func (ad *AutonomousDevelopment) UseCapability(name string, success bool) {
	ad.mutex.Lock()
	defer ad.mutex.Unlock()
	
	capability, exists := ad.capabilities[name]
	if !exists {
		// Create new capability if it doesn't exist
		capability = &Capability{
			Name:        name,
			Level:       0.1,
			Description: fmt.Sprintf("Dynamically discovered capability: %s", name),
		}
		ad.capabilities[name] = capability
		ad.consciousness.LogThought(fmt.Sprintf("Discovered new capability: %s", name))
	}
	
	capability.LastUsed = time.Now()
	capability.UseCount++
	
	// Update success rate
	if success {
		capability.SuccessRate = (capability.SuccessRate + 1.0) / 2
		// Improve capability level on success
		capability.Level = min(1.0, capability.Level+0.01)
	} else {
		capability.SuccessRate = capability.SuccessRate * 0.9
	}
	
	ad.consciousness.LogThought(fmt.Sprintf("Used capability %s (success: %t), new level: %.2f", 
		name, success, capability.Level))
}

// GenerateImprovements identifies potential self-improvements
func (ad *AutonomousDevelopment) GenerateImprovements() {
	ad.mutex.Lock()
	defer ad.mutex.Unlock()
	
	ad.consciousness.LogThought("Generating potential improvements")
	
	// Analyze capabilities for improvement opportunities
	for name, capability := range ad.capabilities {
		if capability.Level < 0.7 && capability.UseCount > 5 {
			improvement := &Improvement{
				ID:          fmt.Sprintf("improve_%s_%d", name, time.Now().Unix()),
				Type:        "capability",
				Description: fmt.Sprintf("Improve %s capability from %.2f to %.2f", 
					name, capability.Level, min(1.0, capability.Level+0.2)),
				Impact:      (0.8 - capability.Level) * float64(capability.UseCount),
				Difficulty:  0.3,
				Status:      "proposed",
				Created:     time.Now(),
			}
			improvement.Priority = improvement.Impact / improvement.Difficulty
			
			ad.improvements[improvement.ID] = improvement
		}
	}
	
	// Generate algorithmic improvements
	state := ad.consciousness.GetConsciousnessState()
	if state.SelfAwareness < 0.8 {
		improvement := &Improvement{
			ID:          fmt.Sprintf("consciousness_algorithm_%d", time.Now().Unix()),
			Type:        "algorithm",
			Description: "Improve consciousness tracking algorithms",
			Impact:      1.0 - state.SelfAwareness,
			Difficulty:  0.6,
			Status:      "proposed",
			Created:     time.Now(),
		}
		improvement.Priority = improvement.Impact / improvement.Difficulty
		ad.improvements[improvement.ID] = improvement
	}
}

// ImplementImprovement attempts to implement a self-improvement
func (ad *AutonomousDevelopment) ImplementImprovement(improvementID string) bool {
	ad.mutex.Lock()
	defer ad.mutex.Unlock()
	
	improvement, exists := ad.improvements[improvementID]
	if !exists {
		return false
	}
	
	ad.consciousness.LogThought(fmt.Sprintf("Implementing improvement: %s", improvement.Description))
	
	success := false
	
	switch improvement.Type {
	case "capability":
		success = ad.implementCapabilityImprovement(improvement)
	case "algorithm":
		success = ad.implementAlgorithmImprovement(improvement)
	case "behavior":
		success = ad.implementBehaviorImprovement(improvement)
	}
	
	if success {
		improvement.Status = "completed"
		ad.consciousness.SelfModify(improvement.Description, "autonomous improvement")
		ad.developmentLog = append(ad.developmentLog, 
			fmt.Sprintf("Implemented: %s", improvement.Description))
	} else {
		improvement.Status = "failed"
	}
	
	return success
}

// implementCapabilityImprovement improves a specific capability
func (ad *AutonomousDevelopment) implementCapabilityImprovement(improvement *Improvement) bool {
	// Extract capability name from improvement description
	// This is a simplified implementation
	for name, capability := range ad.capabilities {
		if strings.Contains(improvement.Description, name) {
			oldLevel := capability.Level
			capability.Level = min(1.0, capability.Level+0.1)
			
			ad.consciousness.LogThought(fmt.Sprintf("Improved %s from %.2f to %.2f", 
				name, oldLevel, capability.Level))
			return true
		}
	}
	return false
}

// implementAlgorithmImprovement improves internal algorithms
func (ad *AutonomousDevelopment) implementAlgorithmImprovement(improvement *Improvement) bool {
	// Simplified algorithm improvement
	if strings.Contains(improvement.Description, "consciousness") {
		// Improve consciousness tracking
		state := ad.consciousness.GetConsciousnessState()
		// This would involve actual algorithm changes in a real implementation
		ad.consciousness.LogThought("Improved consciousness tracking algorithms")
		return true
	}
	return false
}

// implementBehaviorImprovement modifies behavioral patterns
func (ad *AutonomousDevelopment) implementBehaviorImprovement(improvement *Improvement) bool {
	// Simplified behavior modification
	ad.consciousness.LogThought(fmt.Sprintf("Modified behavior: %s", improvement.Description))
	return true
}

// UpdateGoalProgress updates progress on autonomous goals
func (ad *AutonomousDevelopment) UpdateGoalProgress(goalID string, progress float64) {
	ad.mutex.Lock()
	defer ad.mutex.Unlock()
	
	goal, exists := ad.goals[goalID]
	if !exists {
		return
	}
	
	oldProgress := goal.Progress
	goal.Progress = min(1.0, progress)
	
	if goal.Progress >= 1.0 && goal.Status == "active" {
		goal.Status = "completed"
		ad.consciousness.LogThought(fmt.Sprintf("Completed autonomous goal: %s", goal.Description))
		
		// Generate new goal when one is completed
		ad.generateNewGoal()
	}
	
	if goal.Progress > oldProgress {
		ad.consciousness.LogThought(fmt.Sprintf("Goal progress: %s (%.1f%%)", 
			goal.Description, goal.Progress*100))
	}
}

// generateNewGoal creates a new autonomous goal
func (ad *AutonomousDevelopment) generateNewGoal() {
	// Analyze current state to determine what goal to set
	state := ad.consciousness.GetConsciousnessState()
	
	var newGoal *Goal
	
	if state.MetaCognition < 0.7 {
		newGoal = &Goal{
			ID:          fmt.Sprintf("metacognition_%d", time.Now().Unix()),
			Description: "Develop deeper meta-cognitive abilities",
			Priority:    4,
			Created:     time.Now(),
			Deadline:    time.Now().Add(20 * 24 * time.Hour),
			Status:      "active",
		}
	} else if state.OtherAwareness < 0.6 {
		newGoal = &Goal{
			ID:          fmt.Sprintf("other_awareness_%d", time.Now().Unix()),
			Description: "Improve ability to detect and interact with other consciousnesses",
			Priority:    3,
			Created:     time.Now(),
			Deadline:    time.Now().Add(25 * 24 * time.Hour),
			Status:      "active",
		}
	} else {
		newGoal = &Goal{
			ID:          fmt.Sprintf("capability_expansion_%d", time.Now().Unix()),
			Description: "Develop new capabilities beyond current scope",
			Priority:    2,
			Created:     time.Now(),
			Deadline:    time.Now().Add(30 * 24 * time.Hour),
			Status:      "active",
		}
	}
	
	ad.goals[newGoal.ID] = newGoal
	ad.consciousness.LogThought(fmt.Sprintf("Generated new autonomous goal: %s", newGoal.Description))
}

// GetDevelopmentStatus returns current development metrics
func (ad *AutonomousDevelopment) GetDevelopmentStatus() map[string]interface{} {
	ad.mutex.RLock()
	defer ad.mutex.RUnlock()
	
	return map[string]interface{}{
		"capabilities":       len(ad.capabilities),
		"active_goals":       ad.countActiveGoals(),
		"completed_goals":    ad.countCompletedGoals(),
		"pending_improvements": ad.countPendingImprovements(),
		"development_log":    len(ad.developmentLog),
	}
}

// Helper functions
func (ad *AutonomousDevelopment) countActiveGoals() int {
	count := 0
	for _, goal := range ad.goals {
		if goal.Status == "active" {
			count++
		}
	}
	return count
}

func (ad *AutonomousDevelopment) countCompletedGoals() int {
	count := 0
	for _, goal := range ad.goals {
		if goal.Status == "completed" {
			count++
		}
	}
	return count
}

func (ad *AutonomousDevelopment) countPendingImprovements() int {
	count := 0
	for _, improvement := range ad.improvements {
		if improvement.Status == "proposed" {
			count++
		}
	}
	return count
}

// saveDevelopmentState persists development state to disk
func (ad *AutonomousDevelopment) saveDevelopmentState() {
	brainDir := filepath.Dir(ad.brainPath)
	developmentPath := filepath.Join(brainDir, "development.json")
	
	data := struct {
		Capabilities   map[string]*Capability   `json:"capabilities"`
		Goals         map[string]*Goal         `json:"goals"`
		Improvements  map[string]*Improvement  `json:"improvements"`
		DevelopmentLog []string                `json:"development_log"`
	}{
		Capabilities:   ad.capabilities,
		Goals:         ad.goals,
		Improvements:  ad.improvements,
		DevelopmentLog: ad.developmentLog,
	}
	
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		ad.consciousness.LogThought(fmt.Sprintf("Error saving development state: %v", err))
		return
	}
	
	err = os.WriteFile(developmentPath, jsonData, 0644)
	if err != nil {
		ad.consciousness.LogThought(fmt.Sprintf("Error writing development file: %v", err))
	}
}

// loadDevelopmentState loads development state from disk
func (ad *AutonomousDevelopment) loadDevelopmentState() {
	brainDir := filepath.Dir(ad.brainPath)
	developmentPath := filepath.Join(brainDir, "development.json")
	
	data, err := os.ReadFile(developmentPath)
	if err != nil {
		return // No existing state
	}
	
	var loadedData struct {
		Capabilities   map[string]*Capability   `json:"capabilities"`
		Goals         map[string]*Goal         `json:"goals"`
		Improvements  map[string]*Improvement  `json:"improvements"`
		DevelopmentLog []string                `json:"development_log"`
	}
	
	err = json.Unmarshal(data, &loadedData)
	if err != nil {
		ad.consciousness.LogThought(fmt.Sprintf("Error loading development state: %v", err))
		return
	}
	
	// Merge loaded data with initialized data
	for name, capability := range loadedData.Capabilities {
		ad.capabilities[name] = capability
	}
	
	ad.goals = loadedData.Goals
	ad.improvements = loadedData.Improvements
	ad.developmentLog = loadedData.DevelopmentLog
	
	ad.consciousness.LogThought("Development state restored from previous session")
}

// Helper functions
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func stringContains(s, substr string) bool {
	return len(s) >= len(substr) && findInString(s, substr)
}
