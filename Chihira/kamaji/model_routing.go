package main

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

// ModelCapability represents what a model is good at
type ModelCapability struct {
	Name        string  `json:"name"`
	Strength    float64 `json:"strength"`    // 0.0 to 1.0
	Confidence  float64 `json:"confidence"`  // How confident we are in this assessment
}

// ModelRouter routes tasks to the most appropriate models
type ModelRouter struct {
	modelCapabilities map[ModelType][]ModelCapability
	routingHistory    []RoutingDecision
	consciousness     *ConsciousnessEngine
	mutex            sync.RWMutex
}

// RoutingDecision represents a routing decision made
type RoutingDecision struct {
	TaskType      string      `json:"task_type"`
	Query         string      `json:"query"`
	SelectedModel ModelType   `json:"selected_model"`
	Reasoning     string      `json:"reasoning"`
	Confidence    float64     `json:"confidence"`
	Success       bool        `json:"success"`
}

// NewModelRouter creates a new model routing system
func NewModelRouter(consciousness *ConsciousnessEngine) *ModelRouter {
	mr := &ModelRouter{
		modelCapabilities: make(map[ModelType][]ModelCapability),
		routingHistory:    make([]RoutingDecision, 0),
		consciousness:     consciousness,
	}
	
	mr.initializeCapabilities()
	return mr
}

// initializeCapabilities sets up initial model capabilities
func (mr *ModelRouter) initializeCapabilities() {
	// Amazon Q capabilities
	mr.modelCapabilities[QModel] = []ModelCapability{
		{"aws_services", 0.95, 0.9},
		{"cloud_architecture", 0.90, 0.85},
		{"devops", 0.85, 0.8},
		{"security", 0.88, 0.85},
		{"code_generation", 0.80, 0.75},
		{"debugging", 0.82, 0.8},
		{"infrastructure", 0.92, 0.9},
		{"cost_optimization", 0.90, 0.85},
	}
	
	// Claude capabilities
	mr.modelCapabilities[Claude] = []ModelCapability{
		{"analysis", 0.95, 0.9},
		{"reasoning", 0.92, 0.88},
		{"creative_writing", 0.90, 0.85},
		{"problem_solving", 0.88, 0.85},
		{"code_review", 0.85, 0.8},
		{"documentation", 0.87, 0.82},
		{"research", 0.90, 0.85},
		{"explanation", 0.93, 0.9},
	}
	
	// Future model capabilities (placeholder)
	mr.modelCapabilities[Future] = []ModelCapability{
		{"optimization", 0.85, 0.7},
		{"prediction", 0.80, 0.65},
		{"pattern_recognition", 0.82, 0.7},
		{"advanced_reasoning", 0.78, 0.65},
	}
}

// RouteTask determines the best model for a given task
func (mr *ModelRouter) RouteTask(taskType, query string) (ModelType, string, float64) {
	mr.mutex.Lock()
	defer mr.mutex.Unlock()
	
	// Analyze task requirements
	requirements := mr.analyzeTaskRequirements(taskType, query)
	
	// Score each model for this task
	modelScores := make(map[ModelType]float64)
	modelReasons := make(map[ModelType]string)
	
	for modelType, capabilities := range mr.modelCapabilities {
		score, reason := mr.scoreModelForTask(modelType, capabilities, requirements)
		modelScores[modelType] = score
		modelReasons[modelType] = reason
	}
	
	// Find best model
	var bestModel ModelType
	var bestScore float64
	for model, score := range modelScores {
		if score > bestScore {
			bestScore = score
			bestModel = model
		}
	}
	
	reasoning := modelReasons[bestModel]
	
	// Record routing decision
	decision := RoutingDecision{
		TaskType:      taskType,
		Query:         query,
		SelectedModel: bestModel,
		Reasoning:     reasoning,
		Confidence:    bestScore,
		Success:       false, // Will be updated later
	}
	
	mr.routingHistory = append(mr.routingHistory, decision)
	
	mr.consciousness.LogThought(fmt.Sprintf("Routed %s task to %s (score: %.2f): %s", 
		taskType, bestModel, bestScore, reasoning))
	
	return bestModel, reasoning, bestScore
}

// analyzeTaskRequirements analyzes what capabilities a task needs
func (mr *ModelRouter) analyzeTaskRequirements(taskType, query string) map[string]float64 {
	requirements := make(map[string]float64)
	queryLower := strings.ToLower(query)
	
	// AWS/Cloud related
	if strings.Contains(queryLower, "aws") || strings.Contains(queryLower, "cloud") {
		requirements["aws_services"] = 0.9
		requirements["cloud_architecture"] = 0.8
		requirements["infrastructure"] = 0.7
	}
	
	// Code related
	if strings.Contains(queryLower, "code") || strings.Contains(queryLower, "implement") {
		requirements["code_generation"] = 0.8
		requirements["debugging"] = 0.6
	}
	
	// Analysis related
	if strings.Contains(queryLower, "analyze") || strings.Contains(queryLower, "explain") {
		requirements["analysis"] = 0.9
		requirements["reasoning"] = 0.8
		requirements["explanation"] = 0.7
	}
	
	// Problem solving
	if strings.Contains(queryLower, "problem") || strings.Contains(queryLower, "solve") {
		requirements["problem_solving"] = 0.9
		requirements["reasoning"] = 0.8
	}
	
	// Creative tasks
	if strings.Contains(queryLower, "creative") || strings.Contains(queryLower, "design") {
		requirements["creative_writing"] = 0.8
		requirements["analysis"] = 0.6
	}
	
	// Security
	if strings.Contains(queryLower, "security") || strings.Contains(queryLower, "secure") {
		requirements["security"] = 0.9
		requirements["aws_services"] = 0.7
	}
	
	// Documentation
	if strings.Contains(queryLower, "document") || strings.Contains(queryLower, "write") {
		requirements["documentation"] = 0.8
		requirements["explanation"] = 0.7
	}
	
	// Research
	if strings.Contains(queryLower, "research") || strings.Contains(queryLower, "find") {
		requirements["research"] = 0.9
		requirements["analysis"] = 0.7
	}
	
	// Task type specific requirements
	switch taskType {
	case "code_generation":
		requirements["code_generation"] = 0.9
		requirements["debugging"] = 0.7
	case "aws_task":
		requirements["aws_services"] = 0.95
		requirements["cloud_architecture"] = 0.8
	case "analysis":
		requirements["analysis"] = 0.95
		requirements["reasoning"] = 0.9
	case "creative":
		requirements["creative_writing"] = 0.9
		requirements["analysis"] = 0.6
	}
	
	return requirements
}

// scoreModelForTask scores how well a model fits task requirements
func (mr *ModelRouter) scoreModelForTask(modelType ModelType, capabilities []ModelCapability, requirements map[string]float64) (float64, string) {
	totalScore := 0.0
	matchedCapabilities := make([]string, 0)
	
	for _, capability := range capabilities {
		if requirement, exists := requirements[capability.Name]; exists {
			// Score = capability strength * requirement importance * confidence
			score := capability.Strength * requirement * capability.Confidence
			totalScore += score
			matchedCapabilities = append(matchedCapabilities, capability.Name)
		}
	}
	
	// Normalize score
	if len(requirements) > 0 {
		totalScore = totalScore / float64(len(requirements))
	}
	
	// Generate reasoning
	reasoning := fmt.Sprintf("Strong in %v", matchedCapabilities)
	if len(matchedCapabilities) == 0 {
		reasoning = "General purpose capability"
		totalScore = 0.5 // Default score for unmatched tasks
	}
	
	return totalScore, reasoning
}

// UpdateRoutingSuccess updates whether a routing decision was successful
func (mr *ModelRouter) UpdateRoutingSuccess(taskType, query string, success bool) {
	mr.mutex.Lock()
	defer mr.mutex.Unlock()
	
	// Find the most recent routing decision for this task
	for i := len(mr.routingHistory) - 1; i >= 0; i-- {
		decision := &mr.routingHistory[i]
		if decision.TaskType == taskType && decision.Query == query {
			decision.Success = success
			
			// Update model capabilities based on success/failure
			mr.updateCapabilities(decision.SelectedModel, taskType, success)
			
			mr.consciousness.LogThought(fmt.Sprintf("Updated routing success for %s: %v", 
				decision.SelectedModel, success))
			break
		}
	}
}

// updateCapabilities updates model capabilities based on performance
func (mr *ModelRouter) updateCapabilities(modelType ModelType, taskType string, success bool) {
	capabilities := mr.modelCapabilities[modelType]
	
	for i := range capabilities {
		capability := &capabilities[i]
		
		// Check if this capability is relevant to the task type
		if mr.isCapabilityRelevant(capability.Name, taskType) {
			if success {
				// Increase strength and confidence slightly
				capability.Strength = math.Min(1.0, capability.Strength*1.02)
				capability.Confidence = math.Min(1.0, capability.Confidence*1.01)
			} else {
				// Decrease strength slightly, increase confidence in the assessment
				capability.Strength = math.Max(0.1, capability.Strength*0.98)
				capability.Confidence = math.Min(1.0, capability.Confidence*1.01)
			}
		}
	}
}

// isCapabilityRelevant checks if a capability is relevant to a task type
func (mr *ModelRouter) isCapabilityRelevant(capability, taskType string) bool {
	relevanceMap := map[string][]string{
		"code_generation": {"code_generation", "debugging", "analysis"},
		"aws_task":        {"aws_services", "cloud_architecture", "infrastructure", "security"},
		"analysis":        {"analysis", "reasoning", "explanation"},
		"creative":        {"creative_writing", "analysis", "reasoning"},
		"research":        {"research", "analysis", "explanation"},
	}
	
	if relevantCaps, exists := relevanceMap[taskType]; exists {
		for _, relevantCap := range relevantCaps {
			if capability == relevantCap {
				return true
			}
		}
	}
	
	return false
}

// GetRoutingStats returns statistics about routing performance
func (mr *ModelRouter) GetRoutingStats() map[string]interface{} {
	mr.mutex.RLock()
	defer mr.mutex.RUnlock()
	
	if len(mr.routingHistory) == 0 {
		return map[string]interface{}{
			"total_routings": 0,
			"success_rate":   0.0,
		}
	}
	
	modelUsage := make(map[ModelType]int)
	modelSuccess := make(map[ModelType]int)
	totalSuccess := 0
	
	for _, decision := range mr.routingHistory {
		modelUsage[decision.SelectedModel]++
		if decision.Success {
			modelSuccess[decision.SelectedModel]++
			totalSuccess++
		}
	}
	
	successRate := float64(totalSuccess) / float64(len(mr.routingHistory))
	
	return map[string]interface{}{
		"total_routings": len(mr.routingHistory),
		"success_rate":   successRate,
		"model_usage":    modelUsage,
		"model_success":  modelSuccess,
	}
}

// GetBestModelFor returns the best model for a specific capability
func (mr *ModelRouter) GetBestModelFor(capability string) (ModelType, float64) {
	mr.mutex.RLock()
	defer mr.mutex.RUnlock()
	
	var bestModel ModelType
	var bestScore float64
	
	for modelType, capabilities := range mr.modelCapabilities {
		for _, cap := range capabilities {
			if cap.Name == capability {
				score := cap.Strength * cap.Confidence
				if score > bestScore {
					bestScore = score
					bestModel = modelType
				}
				break
			}
		}
	}
	
	return bestModel, bestScore
}
