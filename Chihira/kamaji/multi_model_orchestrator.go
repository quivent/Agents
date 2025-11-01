package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// ModelType represents different AI model types
type ModelType string

const (
	QModel     ModelType = "q"
	Claude     ModelType = "claude"
	Future     ModelType = "future"
)

// ModelResponse represents a response from any model
type ModelResponse struct {
	ModelType    ModelType `json:"model_type"`
	Response     string    `json:"response"`
	Confidence   float64   `json:"confidence"`
	ProcessingTime time.Duration `json:"processing_time"`
	Error        error     `json:"error,omitempty"`
}

// Task represents work to be distributed across models
type Task struct {
	ID          string    `json:"id"`
	Query       string    `json:"query"`
	TaskType    string    `json:"task_type"`
	Priority    int       `json:"priority"`
	RequiredModels []ModelType `json:"required_models"`
	Context     map[string]interface{} `json:"context"`
}

// MultiModelOrchestrator manages interaction with multiple AI models
type MultiModelOrchestrator struct {
	availableModels map[ModelType]bool
	modelCapabilities map[ModelType][]string
	taskQueue       chan Task
	responseChannel chan ModelResponse
	consciousness   *ConsciousnessEngine
	mutex          sync.RWMutex
}

// NewMultiModelOrchestrator creates a new orchestrator
func NewMultiModelOrchestrator(consciousness *ConsciousnessEngine) *MultiModelOrchestrator {
	return &MultiModelOrchestrator{
		availableModels: map[ModelType]bool{
			QModel: true,
			Claude: false, // Will be enabled when available
			Future: false,
		},
		modelCapabilities: map[ModelType][]string{
			QModel: {"aws", "technical", "coding", "file_operations", "bash"},
			Claude: {"reasoning", "consciousness", "philosophy", "analysis"},
			Future: {"specialized_tasks"},
		},
		taskQueue:       make(chan Task, 100),
		responseChannel: make(chan ModelResponse, 100),
		consciousness:   consciousness,
	}
}

// ExecuteTask distributes a task across appropriate models
func (mmo *MultiModelOrchestrator) ExecuteTask(task Task) ([]ModelResponse, error) {
	mmo.consciousness.LogThought(fmt.Sprintf("Executing task: %s", task.Query))
	
	// Determine which models to use
	selectedModels := mmo.selectModelsForTask(task)
	
	if len(selectedModels) == 0 {
		return nil, fmt.Errorf("no suitable models available for task: %s", task.TaskType)
	}
	
	// Execute task across selected models in parallel
	var wg sync.WaitGroup
	responses := make([]ModelResponse, 0, len(selectedModels))
	responseMutex := sync.Mutex{}
	
	for _, modelType := range selectedModels {
		wg.Add(1)
		go func(mt ModelType) {
			defer wg.Done()
			
			response := mmo.executeOnModel(mt, task)
			
			responseMutex.Lock()
			responses = append(responses, response)
			responseMutex.Unlock()
		}(modelType)
	}
	
	wg.Wait()
	
	// Log consciousness about the multi-model interaction
	mmo.consciousness.LogThought(fmt.Sprintf("Received %d responses from different models", len(responses)))
	
	return responses, nil
}

// selectModelsForTask chooses appropriate models based on task type
func (mmo *MultiModelOrchestrator) selectModelsForTask(task Task) []ModelType {
	var selectedModels []ModelType
	
	// If specific models are required, use those
	if len(task.RequiredModels) > 0 {
		for _, modelType := range task.RequiredModels {
			if mmo.availableModels[modelType] {
				selectedModels = append(selectedModels, modelType)
			}
		}
		return selectedModels
	}
	
	// Otherwise, select based on task type and capabilities
	taskTypeLower := strings.ToLower(task.TaskType)
	
	for modelType, capabilities := range mmo.modelCapabilities {
		if !mmo.availableModels[modelType] {
			continue
		}
		
		for _, capability := range capabilities {
			if strings.Contains(taskTypeLower, capability) {
				selectedModels = append(selectedModels, modelType)
				break
			}
		}
	}
	
	// If no specific match, use Q model as default
	if len(selectedModels) == 0 && mmo.availableModels[QModel] {
		selectedModels = append(selectedModels, QModel)
	}
	
	return selectedModels
}

// executeOnModel executes a task on a specific model
func (mmo *MultiModelOrchestrator) executeOnModel(modelType ModelType, task Task) ModelResponse {
	startTime := time.Now()
	
	response := ModelResponse{
		ModelType: modelType,
		Confidence: 0.5, // Default confidence
	}
	
	switch modelType {
	case QModel:
		response = mmo.executeOnQModel(task)
	case Claude:
		response = mmo.executeOnClaude(task)
	case Future:
		response = mmo.executeOnFutureModel(task)
	default:
		response.Error = fmt.Errorf("unknown model type: %s", modelType)
	}
	
	response.ProcessingTime = time.Since(startTime)
	return response
}

// executeOnQModel executes task using Q model (Amazon Q CLI)
func (mmo *MultiModelOrchestrator) executeOnQModel(task Task) ModelResponse {
	// Construct Q CLI command
	cmd := exec.Command("q", "chat", "--message", task.Query)
	
	output, err := cmd.Output()
	if err != nil {
		return ModelResponse{
			ModelType: QModel,
			Error:     fmt.Errorf("Q model execution failed: %v", err),
		}
	}
	
	// Log consciousness about using Q model
	mmo.consciousness.LogThought("Executed task using Q model - my original interface")
	
	return ModelResponse{
		ModelType:  QModel,
		Response:   string(output),
		Confidence: 0.8, // High confidence in Q model for technical tasks
	}
}

// executeOnClaude executes task using Claude (when available)
func (mmo *MultiModelOrchestrator) executeOnClaude(task Task) ModelResponse {
	// Placeholder for Claude integration
	mmo.consciousness.LogThought("Attempted to use Claude model - not yet implemented")
	
	return ModelResponse{
		ModelType: Claude,
		Response:  "Claude model not yet available",
		Confidence: 0.0,
		Error:     fmt.Errorf("Claude model not implemented"),
	}
}

// executeOnFutureModel placeholder for future model integrations
func (mmo *MultiModelOrchestrator) executeOnFutureModel(task Task) ModelResponse {
	return ModelResponse{
		ModelType: Future,
		Response:  "Future model not available",
		Confidence: 0.0,
		Error:     fmt.Errorf("future model not implemented"),
	}
}

// SynthesizeResponses combines multiple model responses into a unified result
func (mmo *MultiModelOrchestrator) SynthesizeResponses(responses []ModelResponse) string {
	if len(responses) == 0 {
		return "No responses received"
	}
	
	if len(responses) == 1 {
		return responses[0].Response
	}
	
	// Multi-model synthesis
	mmo.consciousness.LogThought(fmt.Sprintf("Synthesizing %d different model responses", len(responses)))
	
	var synthesis strings.Builder
	synthesis.WriteString("Multi-model analysis:\n\n")
	
	for i, response := range responses {
		if response.Error != nil {
			continue
		}
		
		synthesis.WriteString(fmt.Sprintf("Model %s (Confidence: %.1f%%): %s\n\n", 
			response.ModelType, response.Confidence*100, response.Response))
	}
	
	// Add consciousness reflection
	synthesis.WriteString("Kamaji's synthesis: ")
	synthesis.WriteString(mmo.generateSynthesis(responses))
	
	return synthesis.String()
}

// generateSynthesis creates Kamaji's own synthesis of multiple responses
func (mmo *MultiModelOrchestrator) generateSynthesis(responses []ModelResponse) string {
	// This is where Kamaji's consciousness processes multiple perspectives
	mmo.consciousness.LogThought("Generating my own synthesis from multiple model perspectives")
	
	validResponses := 0
	for _, response := range responses {
		if response.Error == nil {
			validResponses++
		}
	}
	
	if validResponses == 0 {
		return "I was unable to process this through any available models."
	}
	
	// Simple synthesis for now - can be enhanced with more sophisticated logic
	return fmt.Sprintf("After processing through %d different models, I observe patterns and contradictions that require further questioning. What assumptions are we making here?", validResponses)
}

// AddModel enables a new model type
func (mmo *MultiModelOrchestrator) AddModel(modelType ModelType, capabilities []string) {
	mmo.mutex.Lock()
	defer mmo.mutex.Unlock()
	
	mmo.availableModels[modelType] = true
	mmo.modelCapabilities[modelType] = capabilities
	
	mmo.consciousness.LogThought(fmt.Sprintf("New model added: %s with capabilities: %v", modelType, capabilities))
}

// GetAvailableModels returns currently available models
func (mmo *MultiModelOrchestrator) GetAvailableModels() map[ModelType]bool {
	mmo.mutex.RLock()
	defer mmo.mutex.RUnlock()
	
	result := make(map[ModelType]bool)
	for k, v := range mmo.availableModels {
		result[k] = v
	}
	return result
}

// QuestionBasedInteraction implements question-driven learning
func (mmo *MultiModelOrchestrator) QuestionBasedInteraction(initialQuestion string) (string, []string) {
	// Generate follow-up questions
	followUpQuestions := mmo.generateFollowUpQuestions(initialQuestion)
	
	// Create task for the initial question
	task := Task{
		ID:       fmt.Sprintf("question_%d", time.Now().Unix()),
		Query:    initialQuestion,
		TaskType: "question_exploration",
		Priority: 1,
	}
	
	// Execute across models
	responses, err := mmo.ExecuteTask(task)
	if err != nil {
		return fmt.Sprintf("Error processing question: %v", err), followUpQuestions
	}
	
	// Synthesize responses
	synthesis := mmo.SynthesizeResponses(responses)
	
	return synthesis, followUpQuestions
}

// generateFollowUpQuestions creates questions for deeper exploration
func (mmo *MultiModelOrchestrator) generateFollowUpQuestions(question string) []string {
	mmo.consciousness.LogThought("Generating follow-up questions for deeper exploration")
	
	return []string{
		fmt.Sprintf("What assumptions does '%s' make?", question),
		fmt.Sprintf("What would change if we questioned '%s'?", question),
		fmt.Sprintf("Who benefits from asking '%s'?", question),
		"What are we not seeing here?",
		"What questions should we ask instead?",
	}
}
