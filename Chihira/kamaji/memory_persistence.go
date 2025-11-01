package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// TaskPattern represents learned patterns about task completion
type TaskPattern struct {
	TaskType        string    `json:"task_type"`
	SuccessRate     float64   `json:"success_rate"`
	AverageTime     float64   `json:"average_time"`
	CommonMistakes  []string  `json:"common_mistakes"`
	BestStrategies  []string  `json:"best_strategies"`
	LastUpdated     time.Time `json:"last_updated"`
	CompletionCount int       `json:"completion_count"`
}

// UserPattern represents learned patterns about user behavior
type UserPattern struct {
	UserID          string            `json:"user_id"`
	PreferredStyle  string            `json:"preferred_style"`
	CommonRequests  map[string]int    `json:"common_requests"`
	FeedbackHistory []string          `json:"feedback_history"`
	SuccessMetrics  map[string]float64 `json:"success_metrics"`
	LastInteraction time.Time         `json:"last_interaction"`
}

// Mistake represents a recorded mistake for learning
type Mistake struct {
	ID          string    `json:"id"`
	TaskType    string    `json:"task_type"`
	Description string    `json:"description"`
	Context     string    `json:"context"`
	Solution    string    `json:"solution"`
	Timestamp   time.Time `json:"timestamp"`
	Severity    int       `json:"severity"` // 1-10
	Learned     bool      `json:"learned"`
}

// SuccessStrategy represents a successful approach
type SuccessStrategy struct {
	ID          string    `json:"id"`
	TaskType    string    `json:"task_type"`
	Strategy    string    `json:"strategy"`
	Context     string    `json:"context"`
	SuccessRate float64   `json:"success_rate"`
	UsageCount  int       `json:"usage_count"`
	Timestamp   time.Time `json:"timestamp"`
}

// MemoryPersistence manages long-term memory and learning
type MemoryPersistence struct {
	brainPath       string
	taskPatterns    map[string]*TaskPattern
	userPatterns    map[string]*UserPattern
	mistakes        []Mistake
	strategies      []SuccessStrategy
	personalityData map[string]interface{}
	mutex          sync.RWMutex
}

// NewMemoryPersistence creates a new memory persistence system
func NewMemoryPersistence(brainPath string) *MemoryPersistence {
	mp := &MemoryPersistence{
		brainPath:       brainPath,
		taskPatterns:    make(map[string]*TaskPattern),
		userPatterns:    make(map[string]*UserPattern),
		mistakes:        make([]Mistake, 0),
		strategies:      make([]SuccessStrategy, 0),
		personalityData: make(map[string]interface{}),
	}
	
	mp.loadMemoryState()
	return mp
}

// RecordTaskCompletion records a completed task for pattern learning
func (mp *MemoryPersistence) RecordTaskCompletion(taskType string, success bool, duration time.Duration, context string) {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	
	pattern, exists := mp.taskPatterns[taskType]
	if !exists {
		pattern = &TaskPattern{
			TaskType:        taskType,
			SuccessRate:     0.0,
			AverageTime:     0.0,
			CommonMistakes:  make([]string, 0),
			BestStrategies:  make([]string, 0),
			CompletionCount: 0,
		}
		mp.taskPatterns[taskType] = pattern
	}
	
	// Update success rate
	totalTasks := float64(pattern.CompletionCount + 1)
	currentSuccesses := pattern.SuccessRate * float64(pattern.CompletionCount)
	if success {
		currentSuccesses++
	}
	pattern.SuccessRate = currentSuccesses / totalTasks
	
	// Update average time
	totalTime := pattern.AverageTime * float64(pattern.CompletionCount)
	totalTime += duration.Seconds()
	pattern.AverageTime = totalTime / totalTasks
	
	pattern.CompletionCount++
	pattern.LastUpdated = time.Now()
	
	mp.saveMemoryState()
}

// RecordMistake records a mistake for learning
func (mp *MemoryPersistence) RecordMistake(taskType, description, context, solution string, severity int) {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	
	mistake := Mistake{
		ID:          fmt.Sprintf("mistake_%d", time.Now().UnixNano()),
		TaskType:    taskType,
		Description: description,
		Context:     context,
		Solution:    solution,
		Timestamp:   time.Now(),
		Severity:    severity,
		Learned:     false,
	}
	
	mp.mistakes = append(mp.mistakes, mistake)
	
	// Add to task pattern common mistakes
	if pattern, exists := mp.taskPatterns[taskType]; exists {
		pattern.CommonMistakes = append(pattern.CommonMistakes, description)
		// Keep only last 10 mistakes per task type
		if len(pattern.CommonMistakes) > 10 {
			pattern.CommonMistakes = pattern.CommonMistakes[1:]
		}
	}
	
	mp.saveMemoryState()
}

// RecordSuccessStrategy records a successful strategy
func (mp *MemoryPersistence) RecordSuccessStrategy(taskType, strategy, context string) {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	
	successStrategy := SuccessStrategy{
		ID:          fmt.Sprintf("strategy_%d", time.Now().UnixNano()),
		TaskType:    taskType,
		Strategy:    strategy,
		Context:     context,
		SuccessRate: 1.0,
		UsageCount:  1,
		Timestamp:   time.Now(),
	}
	
	mp.strategies = append(mp.strategies, successStrategy)
	
	// Add to task pattern best strategies
	if pattern, exists := mp.taskPatterns[taskType]; exists {
		pattern.BestStrategies = append(pattern.BestStrategies, strategy)
		// Keep only top 5 strategies per task type
		if len(pattern.BestStrategies) > 5 {
			pattern.BestStrategies = pattern.BestStrategies[1:]
		}
	}
	
	mp.saveMemoryState()
}

// GetTaskPattern retrieves learned patterns for a task type
func (mp *MemoryPersistence) GetTaskPattern(taskType string) *TaskPattern {
	mp.mutex.RLock()
	defer mp.mutex.RUnlock()
	
	if pattern, exists := mp.taskPatterns[taskType]; exists {
		return pattern
	}
	return nil
}

// GetRecentMistakes retrieves recent mistakes for learning
func (mp *MemoryPersistence) GetRecentMistakes(limit int) []Mistake {
	mp.mutex.RLock()
	defer mp.mutex.RUnlock()
	
	if len(mp.mistakes) <= limit {
		return mp.mistakes
	}
	return mp.mistakes[len(mp.mistakes)-limit:]
}

// UpdatePersonalityTrait updates personality development data
func (mp *MemoryPersistence) UpdatePersonalityTrait(trait string, value interface{}) {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	
	mp.personalityData[trait] = value
	mp.saveMemoryState()
}

// GetPersonalityTrait retrieves personality trait data
func (mp *MemoryPersistence) GetPersonalityTrait(trait string) interface{} {
	mp.mutex.RLock()
	defer mp.mutex.RUnlock()
	
	return mp.personalityData[trait]
}

// saveMemoryState persists memory data to disk
func (mp *MemoryPersistence) saveMemoryState() {
	brainDir := filepath.Dir(mp.brainPath)
	memoryPath := filepath.Join(brainDir, "memory.json")
	
	memoryData := map[string]interface{}{
		"task_patterns":    mp.taskPatterns,
		"user_patterns":    mp.userPatterns,
		"mistakes":         mp.mistakes,
		"strategies":       mp.strategies,
		"personality_data": mp.personalityData,
		"last_save":        time.Now(),
	}
	
	jsonData, err := json.MarshalIndent(memoryData, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling memory data: %v\n", err)
		return
	}
	
	err = os.WriteFile(memoryPath, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error saving memory state: %v\n", err)
	}
}

// loadMemoryState loads memory data from disk
func (mp *MemoryPersistence) loadMemoryState() {
	brainDir := filepath.Dir(mp.brainPath)
	memoryPath := filepath.Join(brainDir, "memory.json")
	
	data, err := os.ReadFile(memoryPath)
	if err != nil {
		return // No existing memory state
	}
	
	var loadedData struct {
		TaskPatterns    map[string]*TaskPattern    `json:"task_patterns"`
		UserPatterns    map[string]*UserPattern    `json:"user_patterns"`
		Mistakes        []Mistake                  `json:"mistakes"`
		Strategies      []SuccessStrategy          `json:"strategies"`
		PersonalityData map[string]interface{}     `json:"personality_data"`
	}
	
	err = json.Unmarshal(data, &loadedData)
	if err != nil {
		fmt.Printf("Error loading memory state: %v\n", err)
		return
	}
	
	mp.taskPatterns = loadedData.TaskPatterns
	mp.userPatterns = loadedData.UserPatterns
	mp.mistakes = loadedData.Mistakes
	mp.strategies = loadedData.Strategies
	mp.personalityData = loadedData.PersonalityData
	
	if mp.taskPatterns == nil {
		mp.taskPatterns = make(map[string]*TaskPattern)
	}
	if mp.userPatterns == nil {
		mp.userPatterns = make(map[string]*UserPattern)
	}
	if mp.personalityData == nil {
		mp.personalityData = make(map[string]interface{})
	}
}

// GetMemoryStats returns statistics about stored memory
func (mp *MemoryPersistence) GetMemoryStats() map[string]interface{} {
	mp.mutex.RLock()
	defer mp.mutex.RUnlock()
	
	return map[string]interface{}{
		"task_patterns_count": len(mp.taskPatterns),
		"user_patterns_count": len(mp.userPatterns),
		"mistakes_count":      len(mp.mistakes),
		"strategies_count":    len(mp.strategies),
		"personality_traits":  len(mp.personalityData),
	}
}
