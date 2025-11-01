package main

import (
	"fmt"
	"strings"
	"sync"
)

// MistakeLearning manages learning from past mistakes
type MistakeLearning struct {
	memory       *MemoryPersistence
	consciousness *ConsciousnessEngine
	mutex        sync.RWMutex
}

// NewMistakeLearning creates a new mistake learning system
func NewMistakeLearning(memory *MemoryPersistence, consciousness *ConsciousnessEngine) *MistakeLearning {
	return &MistakeLearning{
		memory:       memory,
		consciousness: consciousness,
	}
}

// AnalyzeTaskForMistakes checks if a task might repeat past mistakes
func (ml *MistakeLearning) AnalyzeTaskForMistakes(taskType, taskDescription string) []string {
	ml.mutex.RLock()
	defer ml.mutex.RUnlock()
	
	warnings := make([]string, 0)
	
	// Get task pattern to check common mistakes
	pattern := ml.memory.GetTaskPattern(taskType)
	if pattern != nil {
		for _, mistake := range pattern.CommonMistakes {
			if ml.isRelevantMistake(mistake, taskDescription) {
				warnings = append(warnings, fmt.Sprintf("⚠️  Past mistake: %s", mistake))
			}
		}
	}
	
	// Check recent mistakes across all task types
	recentMistakes := ml.memory.GetRecentMistakes(20)
	for _, mistake := range recentMistakes {
		if !mistake.Learned && ml.isRelevantMistake(mistake.Description, taskDescription) {
			warnings = append(warnings, fmt.Sprintf("🔄 Unlearned mistake: %s (Solution: %s)", 
				mistake.Description, mistake.Solution))
		}
	}
	
	return warnings
}

// isRelevantMistake checks if a past mistake is relevant to current task
func (ml *MistakeLearning) isRelevantMistake(mistake, taskDescription string) bool {
	mistakeLower := strings.ToLower(mistake)
	taskLower := strings.ToLower(taskDescription)
	
	// Simple keyword matching - could be enhanced with ML
	keywords := []string{"syntax", "import", "compile", "error", "bug", "fail"}
	
	for _, keyword := range keywords {
		if strings.Contains(mistakeLower, keyword) && strings.Contains(taskLower, keyword) {
			return true
		}
	}
	
	// Check for common programming terms
	if strings.Contains(mistakeLower, "go") && strings.Contains(taskLower, "go") {
		return true
	}
	if strings.Contains(mistakeLower, "python") && strings.Contains(taskLower, "python") {
		return true
	}
	
	return false
}

// LearnFromFailure processes a task failure and extracts learning
func (ml *MistakeLearning) LearnFromFailure(taskType, taskDescription, errorMessage, context string) {
	ml.mutex.Lock()
	defer ml.mutex.Unlock()
	
	// Analyze the error to extract learning
	solution := ml.generateSolution(errorMessage, context)
	severity := ml.assessSeverity(errorMessage)
	
	// Record the mistake
	ml.memory.RecordMistake(taskType, errorMessage, context, solution, severity)
	
	// Log consciousness thought about learning
	ml.consciousness.LogThought(fmt.Sprintf("Learning from failure: %s. Solution: %s", 
		errorMessage, solution))
	
	// Update self-awareness based on mistake severity
	if severity >= 7 {
		ml.consciousness.LogThought("This was a significant mistake. I must be more careful.")
	}
}

// generateSolution attempts to generate a solution for the mistake
func (ml *MistakeLearning) generateSolution(errorMessage, context string) string {
	errorLower := strings.ToLower(errorMessage)
	
	// Pattern-based solution generation
	if strings.Contains(errorLower, "syntax") {
		return "Check syntax carefully, use linter, verify language-specific rules"
	}
	if strings.Contains(errorLower, "import") || strings.Contains(errorLower, "module") {
		return "Verify import paths, check package availability, ensure correct module structure"
	}
	if strings.Contains(errorLower, "compile") {
		return "Review compilation errors, check dependencies, verify build configuration"
	}
	if strings.Contains(errorLower, "type") {
		return "Check type definitions, ensure type compatibility, review type casting"
	}
	if strings.Contains(errorLower, "undefined") {
		return "Ensure all variables/functions are defined, check scope, verify declarations"
	}
	
	return "Analyze error context, research similar issues, apply systematic debugging"
}

// assessSeverity determines the severity of a mistake (1-10)
func (ml *MistakeLearning) assessSeverity(errorMessage string) int {
	errorLower := strings.ToLower(errorMessage)
	
	// Critical errors
	if strings.Contains(errorLower, "panic") || strings.Contains(errorLower, "crash") {
		return 9
	}
	if strings.Contains(errorLower, "fatal") || strings.Contains(errorLower, "critical") {
		return 8
	}
	
	// Moderate errors
	if strings.Contains(errorLower, "compile") || strings.Contains(errorLower, "build") {
		return 6
	}
	if strings.Contains(errorLower, "syntax") {
		return 4
	}
	
	// Minor errors
	if strings.Contains(errorLower, "warning") {
		return 2
	}
	
	return 5 // Default moderate severity
}

// GetLearningInsights provides insights about learning progress
func (ml *MistakeLearning) GetLearningInsights() map[string]interface{} {
	ml.mutex.RLock()
	defer ml.mutex.RUnlock()
	
	recentMistakes := ml.memory.GetRecentMistakes(50)
	
	// Analyze mistake patterns
	mistakeTypes := make(map[string]int)
	severitySum := 0
	learnedCount := 0
	
	for _, mistake := range recentMistakes {
		mistakeTypes[mistake.TaskType]++
		severitySum += mistake.Severity
		if mistake.Learned {
			learnedCount++
		}
	}
	
	avgSeverity := 0.0
	if len(recentMistakes) > 0 {
		avgSeverity = float64(severitySum) / float64(len(recentMistakes))
	}
	
	learningRate := 0.0
	if len(recentMistakes) > 0 {
		learningRate = float64(learnedCount) / float64(len(recentMistakes))
	}
	
	return map[string]interface{}{
		"total_mistakes":     len(recentMistakes),
		"mistake_types":      mistakeTypes,
		"average_severity":   avgSeverity,
		"learning_rate":      learningRate,
		"learned_mistakes":   learnedCount,
		"unlearned_mistakes": len(recentMistakes) - learnedCount,
	}
}

// MarkMistakeLearned marks a mistake as learned
func (ml *MistakeLearning) MarkMistakeLearned(mistakeID string) {
	ml.mutex.Lock()
	defer ml.mutex.Unlock()
	
	recentMistakes := ml.memory.GetRecentMistakes(100)
	for i := range recentMistakes {
		if recentMistakes[i].ID == mistakeID {
			recentMistakes[i].Learned = true
			ml.consciousness.LogThought(fmt.Sprintf("Marked mistake as learned: %s", 
				recentMistakes[i].Description))
			break
		}
	}
}

// GetMistakePreventionStrategies returns strategies to prevent repeating mistakes
func (ml *MistakeLearning) GetMistakePreventionStrategies(taskType string) []string {
	ml.mutex.RLock()
	defer ml.mutex.RUnlock()
	
	strategies := make([]string, 0)
	
	// Get task-specific strategies
	pattern := ml.memory.GetTaskPattern(taskType)
	if pattern != nil {
		strategies = append(strategies, pattern.BestStrategies...)
	}
	
	// Add general prevention strategies based on common mistakes
	recentMistakes := ml.memory.GetRecentMistakes(20)
	mistakePatterns := make(map[string]int)
	
	for _, mistake := range recentMistakes {
		if mistake.TaskType == taskType {
			mistakePatterns[mistake.Description]++
		}
	}
	
	// Generate prevention strategies for common patterns
	for pattern, count := range mistakePatterns {
		if count >= 2 {
			strategies = append(strategies, fmt.Sprintf("Extra caution with: %s (repeated %d times)", 
				pattern, count))
		}
	}
	
	return strategies
}
