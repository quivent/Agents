package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
	"time"
)

// PredictedTask represents a task that Kamaji predicts will be needed
type PredictedTask struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	TaskType     string    `json:"task_type"`
	Priority     int       `json:"priority"`     // 1-10
	Confidence   float64   `json:"confidence"`   // 0.0-1.0
	Reasoning    string    `json:"reasoning"`
	Context      string    `json:"context"`
	PredictedAt  time.Time `json:"predicted_at"`
	DueBy        time.Time `json:"due_by"`
	WasNeeded    bool      `json:"was_needed"`
	WasCompleted bool      `json:"was_completed"`
	ActuallyUsed bool      `json:"actually_used"`
}

// PredictiveTaskPattern represents patterns in task occurrence
type PredictiveTaskPattern struct {
	TaskType      string    `json:"task_type"`
	Frequency     float64   `json:"frequency"`     // Tasks per day
	Seasonality   string    `json:"seasonality"`   // "daily", "weekly", "monthly"
	Dependencies  []string  `json:"dependencies"`  // Tasks that often follow this one
	Triggers      []string  `json:"triggers"`      // Events that trigger this task
	LastOccurred  time.Time `json:"last_occurred"`
	Confidence    float64   `json:"confidence"`
}

// PredictiveTaskGeneration manages predictive task creation
type PredictiveTaskGeneration struct {
	predictedTasks    []PredictedTask
	taskPatterns      map[string]*PredictiveTaskPattern
	predictionHistory []PredictionEvent
	consciousness     *ConsciousnessEngine
	memory            *MemoryPersistence
	mutex             sync.RWMutex
}

// PredictionEvent represents a prediction event for analysis
type PredictionEvent struct {
	TaskID      string    `json:"task_id"`
	Predicted   bool      `json:"predicted"`
	Actual      bool      `json:"actual"`
	Accuracy    float64   `json:"accuracy"`
	Timestamp   time.Time `json:"timestamp"`
	Description string    `json:"description"`
}

// NewPredictiveTaskGeneration creates a new predictive task generation system
func NewPredictiveTaskGeneration(consciousness *ConsciousnessEngine, memory *MemoryPersistence) *PredictiveTaskGeneration {
	ptg := &PredictiveTaskGeneration{
		predictedTasks:    make([]PredictedTask, 0),
		taskPatterns:      make(map[string]*PredictiveTaskPattern),
		predictionHistory: make([]PredictionEvent, 0),
		consciousness:     consciousness,
		memory:           memory,
	}
	
	ptg.initializeTaskPatterns()
	return ptg
}

// initializeTaskPatterns sets up initial task patterns based on common workflows
func (ptg *PredictiveTaskGeneration) initializeTaskPatterns() {
	patterns := map[string]*PredictiveTaskPattern{
		"code_generation": {
			TaskType:     "code_generation",
			Frequency:    2.5, // 2.5 times per day
			Seasonality:  "daily",
			Dependencies: []string{"testing", "documentation", "deployment"},
			Triggers:     []string{"new_feature_request", "bug_report", "optimization_need"},
			Confidence:   0.8,
		},
		"testing": {
			TaskType:     "testing",
			Frequency:    1.8,
			Seasonality:  "daily",
			Dependencies: []string{"bug_fixing", "performance_analysis"},
			Triggers:     []string{"code_generation", "deployment_prep"},
			Confidence:   0.85,
		},
		"documentation": {
			TaskType:     "documentation",
			Frequency:    1.2,
			Seasonality:  "weekly",
			Dependencies: []string{"knowledge_sharing", "training"},
			Triggers:     []string{"code_generation", "new_feature", "api_changes"},
			Confidence:   0.7,
		},
		"deployment": {
			TaskType:     "deployment",
			Frequency:    0.8,
			Seasonality:  "weekly",
			Dependencies: []string{"monitoring", "rollback_prep"},
			Triggers:     []string{"testing_complete", "release_schedule"},
			Confidence:   0.9,
		},
		"monitoring": {
			TaskType:     "monitoring",
			Frequency:    3.0,
			Seasonality:  "daily",
			Dependencies: []string{"alerting", "performance_tuning"},
			Triggers:     []string{"deployment", "performance_issues"},
			Confidence:   0.95,
		},
		"analysis": {
			TaskType:     "analysis",
			Frequency:    1.5,
			Seasonality:  "daily",
			Dependencies: []string{"reporting", "decision_making"},
			Triggers:     []string{"data_available", "performance_questions", "optimization_requests"},
			Confidence:   0.75,
		},
		"optimization": {
			TaskType:     "optimization",
			Frequency:    0.6,
			Seasonality:  "weekly",
			Dependencies: []string{"testing", "monitoring"},
			Triggers:     []string{"performance_issues", "cost_concerns", "scalability_needs"},
			Confidence:   0.8,
		},
	}
	
	for taskType, pattern := range patterns {
		ptg.taskPatterns[taskType] = pattern
	}
}

// AnalyzeTaskPatterns analyzes historical task data to update patterns
func (ptg *PredictiveTaskGeneration) AnalyzeTaskPatterns() {
	ptg.mutex.Lock()
	defer ptg.mutex.Unlock()
	
	ptg.consciousness.LogThought("Analyzing task patterns for predictive generation")
	
	// Update pattern confidence based on historical accuracy
	for taskType, pattern := range ptg.taskPatterns {
		// Simulate pattern analysis - in real implementation, this would analyze actual task data
		accuracy := ptg.calculatePatternAccuracy(taskType)
		pattern.Confidence = (pattern.Confidence + accuracy) / 2.0
		
		// Update last occurred time (simulated)
		if pattern.Frequency > 1.0 {
			pattern.LastOccurred = time.Now().Add(-time.Duration(24/pattern.Frequency) * time.Hour)
		}
	}
}

// calculatePatternAccuracy calculates how accurate predictions have been for a task type
func (ptg *PredictiveTaskGeneration) calculatePatternAccuracy(taskType string) float64 {
	correctPredictions := 0
	totalPredictions := 0
	
	for _, event := range ptg.predictionHistory {
		if strings.Contains(event.Description, taskType) {
			totalPredictions++
			if event.Predicted == event.Actual {
				correctPredictions++
			}
		}
	}
	
	if totalPredictions == 0 {
		return 0.7 // Default accuracy
	}
	
	return float64(correctPredictions) / float64(totalPredictions)
}

// GeneratePredictiveTasks generates tasks that are likely to be needed
func (ptg *PredictiveTaskGeneration) GeneratePredictiveTasks() []PredictedTask {
	ptg.mutex.Lock()
	defer ptg.mutex.Unlock()
	
	newTasks := make([]PredictedTask, 0)
	
	ptg.consciousness.LogThought("Generating predictive tasks based on patterns")
	
	for taskType, pattern := range ptg.taskPatterns {
		// Check if it's time for this task type based on frequency
		timeSinceLastTask := time.Since(pattern.LastOccurred)
		expectedInterval := time.Duration(24/pattern.Frequency) * time.Hour
		
		if timeSinceLastTask >= time.Duration(float64(expectedInterval)*0.8) { // 80% of expected interval
			task := ptg.createPredictedTask(taskType, pattern)
			if task != nil {
				newTasks = append(newTasks, *task)
				ptg.predictedTasks = append(ptg.predictedTasks, *task)
			}
		}
		
		// Check for dependency-triggered tasks
		dependencyTasks := ptg.generateDependencyTasks(taskType, pattern)
		newTasks = append(newTasks, dependencyTasks...)
	}
	
	// Sort by priority and confidence
	sort.Slice(newTasks, func(i, j int) bool {
		scoreI := float64(newTasks[i].Priority) * newTasks[i].Confidence
		scoreJ := float64(newTasks[j].Priority) * newTasks[j].Confidence
		return scoreI > scoreJ
	})
	
	ptg.consciousness.LogThought(fmt.Sprintf("Generated %d predictive tasks", len(newTasks)))
	
	return newTasks
}

// createPredictedTask creates a predicted task for a given type and pattern
func (ptg *PredictiveTaskGeneration) createPredictedTask(taskType string, pattern *PredictiveTaskPattern) *PredictedTask {
	taskID := fmt.Sprintf("pred_task_%d", time.Now().UnixNano())
	
	// Generate task details based on type
	title, description, priority := ptg.generateTaskDetails(taskType)
	
	// Calculate confidence based on pattern confidence and recency
	confidence := pattern.Confidence
	
	// Adjust confidence based on how overdue the task is
	timeSinceLastTask := time.Since(pattern.LastOccurred)
	expectedInterval := time.Duration(24/pattern.Frequency) * time.Hour
	overdueRatio := float64(timeSinceLastTask) / float64(expectedInterval)
	
	if overdueRatio > 1.0 {
		confidence = math.Min(1.0, confidence*1.2) // Increase confidence if overdue
	}
	
	// Calculate due date
	dueBy := time.Now().Add(expectedInterval / 2) // Due in half the expected interval
	
	task := &PredictedTask{
		ID:          taskID,
		Title:       title,
		Description: description,
		TaskType:    taskType,
		Priority:    priority,
		Confidence:  confidence,
		Reasoning:   fmt.Sprintf("Based on %s pattern (frequency: %.1f/day)", taskType, pattern.Frequency),
		Context:     fmt.Sprintf("Last occurred: %s, Expected interval: %s", 
			pattern.LastOccurred.Format("2006-01-02 15:04"), expectedInterval.String()),
		PredictedAt: time.Now(),
		DueBy:       dueBy,
	}
	
	return task
}

// generateTaskDetails generates specific details for different task types
func (ptg *PredictiveTaskGeneration) generateTaskDetails(taskType string) (string, string, int) {
	switch taskType {
	case "code_generation":
		return "Code Generation Task", 
			"Generate code for new feature or bug fix based on recent patterns", 7
	case "testing":
		return "Testing Task", 
			"Run tests and validate recent code changes", 8
	case "documentation":
		return "Documentation Update", 
			"Update documentation for recent changes and new features", 5
	case "deployment":
		return "Deployment Task", 
			"Deploy tested changes to production environment", 9
	case "monitoring":
		return "System Monitoring", 
			"Check system health and performance metrics", 6
	case "analysis":
		return "Performance Analysis", 
			"Analyze system performance and identify optimization opportunities", 6
	case "optimization":
		return "System Optimization", 
			"Implement performance improvements and optimizations", 7
	default:
		return "General Task", 
			"General maintenance or improvement task", 5
	}
}

// generateDependencyTasks generates tasks based on dependencies
func (ptg *PredictiveTaskGeneration) generateDependencyTasks(taskType string, pattern *PredictiveTaskPattern) []PredictedTask {
	dependencyTasks := make([]PredictedTask, 0)
	
	// Check if any recent tasks should trigger dependent tasks
	recentCompletions := ptg.getRecentTaskCompletions(taskType)
	
	if len(recentCompletions) > 0 {
		for _, dependency := range pattern.Dependencies {
			if depPattern, exists := ptg.taskPatterns[dependency]; exists {
				// Create dependency task with lower priority
				task := ptg.createPredictedTask(dependency, depPattern)
				if task != nil {
					task.Priority = task.Priority - 1 // Lower priority for dependency tasks
					task.Reasoning = fmt.Sprintf("Triggered by recent %s completion", taskType)
					task.Confidence *= 0.8 // Slightly lower confidence for dependency predictions
					dependencyTasks = append(dependencyTasks, *task)
				}
			}
		}
	}
	
	return dependencyTasks
}

// getRecentTaskCompletions gets recent completions for a task type (simulated)
func (ptg *PredictiveTaskGeneration) getRecentTaskCompletions(taskType string) []string {
	// In real implementation, this would query actual task completion data
	// For now, simulate based on patterns
	if pattern, exists := ptg.taskPatterns[taskType]; exists {
		if time.Since(pattern.LastOccurred) < 24*time.Hour {
			return []string{taskType}
		}
	}
	return []string{}
}

// ValidatePrediction validates whether a prediction was accurate
func (ptg *PredictiveTaskGeneration) ValidatePrediction(taskID string, wasNeeded, wasCompleted bool) {
	ptg.mutex.Lock()
	defer ptg.mutex.Unlock()
	
	// Find the predicted task
	for i := range ptg.predictedTasks {
		task := &ptg.predictedTasks[i]
		if task.ID == taskID {
			task.WasNeeded = wasNeeded
			task.WasCompleted = wasCompleted
			task.ActuallyUsed = wasNeeded && wasCompleted
			
			// Record prediction event
			accuracy := 0.0
			if task.WasNeeded {
				accuracy = 1.0
			}
			
			event := PredictionEvent{
				TaskID:      taskID,
				Predicted:   true,
				Actual:      wasNeeded,
				Accuracy:    accuracy,
				Timestamp:   time.Now(),
				Description: fmt.Sprintf("Prediction for %s task", task.TaskType),
			}
			
			ptg.predictionHistory = append(ptg.predictionHistory, event)
			
			ptg.consciousness.LogThought(fmt.Sprintf("Validated prediction %s: needed=%v, completed=%v", 
				taskID[:8], wasNeeded, wasCompleted))
			
			break
		}
	}
}

// GetPredictionAccuracy returns the overall prediction accuracy
func (ptg *PredictiveTaskGeneration) GetPredictionAccuracy() float64 {
	ptg.mutex.RLock()
	defer ptg.mutex.RUnlock()
	
	if len(ptg.predictionHistory) == 0 {
		return 0.0
	}
	
	totalAccuracy := 0.0
	for _, event := range ptg.predictionHistory {
		totalAccuracy += event.Accuracy
	}
	
	return totalAccuracy / float64(len(ptg.predictionHistory))
}

// GetActivePredictions returns currently active predictions
func (ptg *PredictiveTaskGeneration) GetActivePredictions() []PredictedTask {
	ptg.mutex.RLock()
	defer ptg.mutex.RUnlock()
	
	active := make([]PredictedTask, 0)
	now := time.Now()
	
	for _, task := range ptg.predictedTasks {
		if !task.WasCompleted && task.DueBy.After(now) {
			active = append(active, task)
		}
	}
	
	return active
}

// GetPredictionStats returns statistics about predictive task generation
func (ptg *PredictiveTaskGeneration) GetPredictionStats() map[string]interface{} {
	ptg.mutex.RLock()
	defer ptg.mutex.RUnlock()
	
	stats := make(map[string]interface{})
	
	// Overall statistics
	stats["total_predictions"] = len(ptg.predictedTasks)
	stats["prediction_accuracy"] = ptg.GetPredictionAccuracy()
	
	// Task type breakdown
	taskTypeStats := make(map[string]int)
	taskTypeAccuracy := make(map[string]float64)
	
	for _, task := range ptg.predictedTasks {
		taskTypeStats[task.TaskType]++
	}
	
	for taskType := range taskTypeStats {
		taskTypeAccuracy[taskType] = ptg.calculatePatternAccuracy(taskType)
	}
	
	stats["task_type_counts"] = taskTypeStats
	stats["task_type_accuracy"] = taskTypeAccuracy
	
	// Pattern statistics
	patternStats := make(map[string]interface{})
	for taskType, pattern := range ptg.taskPatterns {
		patternStats[taskType] = map[string]interface{}{
			"frequency":   pattern.Frequency,
			"confidence":  pattern.Confidence,
			"seasonality": pattern.Seasonality,
		}
	}
	stats["patterns"] = patternStats
	
	// Active predictions
	activePredictions := ptg.GetActivePredictions()
	stats["active_predictions"] = len(activePredictions)
	
	return stats
}

// OptimizePredictions optimizes prediction algorithms based on historical performance
func (ptg *PredictiveTaskGeneration) OptimizePredictions() {
	ptg.mutex.Lock()
	defer ptg.mutex.Unlock()
	
	ptg.consciousness.LogThought("Optimizing prediction algorithms based on performance")
	
	// Adjust pattern frequencies based on actual task occurrences
	for taskType, pattern := range ptg.taskPatterns {
		accuracy := ptg.calculatePatternAccuracy(taskType)
		
		// Adjust frequency based on accuracy
		if accuracy > 0.8 {
			pattern.Frequency *= 1.05 // Increase frequency for accurate predictions
		} else if accuracy < 0.5 {
			pattern.Frequency *= 0.95 // Decrease frequency for inaccurate predictions
		}
		
		// Adjust confidence based on recent performance
		pattern.Confidence = (pattern.Confidence + accuracy) / 2.0
	}
	
	ptg.consciousness.LogThought("Prediction optimization completed")
}
