package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// ConsensusVote represents a model's vote on a decision
type ConsensusVote struct {
	ModelType   ModelType `json:"model_type"`
	Decision    string    `json:"decision"`
	Confidence  float64   `json:"confidence"`
	Reasoning   string    `json:"reasoning"`
	Timestamp   time.Time `json:"timestamp"`
}

// ConsensusResult represents the outcome of a consensus process
type ConsensusResult struct {
	FinalDecision    string          `json:"final_decision"`
	Confidence       float64         `json:"confidence"`
	Votes           []ConsensusVote  `json:"votes"`
	ConsensusType   string          `json:"consensus_type"` // "unanimous", "majority", "weighted"
	ProcessingTime  time.Duration   `json:"processing_time"`
	Timestamp       time.Time       `json:"timestamp"`
}

// ConsensusSystem manages multi-model consensus decision making
type ConsensusSystem struct {
	orchestrator    *MultiModelOrchestrator
	consciousness   *ConsciousnessEngine
	memory          *MemoryPersistence
	modelWeights    map[ModelType]float64
	consensusHistory []ConsensusResult
	mutex           sync.RWMutex
}

// NewConsensusSystem creates a new consensus system
func NewConsensusSystem(orchestrator *MultiModelOrchestrator, consciousness *ConsciousnessEngine, memory *MemoryPersistence) *ConsensusSystem {
	cs := &ConsensusSystem{
		orchestrator:     orchestrator,
		consciousness:    consciousness,
		memory:          memory,
		modelWeights:    make(map[ModelType]float64),
		consensusHistory: make([]ConsensusResult, 0),
	}
	
	// Initialize model weights based on past performance
	cs.initializeModelWeights()
	
	return cs
}

// initializeModelWeights sets initial weights for different models
func (cs *ConsensusSystem) initializeModelWeights() {
	cs.modelWeights[QModel] = 1.0     // Amazon Q - strong for AWS/technical tasks
	cs.modelWeights[Claude] = 1.0     // Claude - strong for analysis/reasoning
	cs.modelWeights[Future] = 0.8     // Future models - lower initial weight
}

// RequestConsensus requests consensus from multiple models on a decision
func (cs *ConsensusSystem) RequestConsensus(query string, taskType string, models []ModelType) (*ConsensusResult, error) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	
	startTime := time.Now()
	
	cs.consciousness.LogThought(fmt.Sprintf("Requesting consensus from %d models on: %s", len(models), query))
	
	// Collect votes from each model
	votes := make([]ConsensusVote, 0)
	
	for _, modelType := range models {
		vote, err := cs.getModelVote(modelType, query, taskType)
		if err != nil {
			cs.consciousness.LogThought(fmt.Sprintf("Error getting vote from %s: %v", modelType, err))
			continue
		}
		votes = append(votes, vote)
	}
	
	if len(votes) == 0 {
		return nil, fmt.Errorf("no valid votes received")
	}
	
	// Calculate consensus
	result := cs.calculateConsensus(votes)
	result.ProcessingTime = time.Since(startTime)
	result.Timestamp = time.Now()
	
	// Store result
	cs.consensusHistory = append(cs.consensusHistory, *result)
	
	// Update model weights based on consensus participation
	cs.updateModelWeights(votes, result)
	
	cs.consciousness.LogThought(fmt.Sprintf("Consensus reached: %s (confidence: %.2f)", 
		result.FinalDecision, result.Confidence))
	
	return result, nil
}

// getModelVote gets a vote from a specific model
func (cs *ConsensusSystem) getModelVote(modelType ModelType, query, taskType string) (ConsensusVote, error) {
	// Simulate model responses - in real implementation, this would call actual models
	vote := ConsensusVote{
		ModelType: modelType,
		Timestamp: time.Now(),
	}
	
	switch modelType {
	case QModel:
		vote.Decision = cs.simulateQModelResponse(query, taskType)
		vote.Confidence = 0.85
		vote.Reasoning = "Based on AWS best practices and technical analysis"
		
	case Claude:
		vote.Decision = cs.simulateClaudeResponse(query, taskType)
		vote.Confidence = 0.90
		vote.Reasoning = "Based on comprehensive reasoning and analysis"
		
	case Future:
		vote.Decision = cs.simulateFutureModelResponse(query, taskType)
		vote.Confidence = 0.75
		vote.Reasoning = "Based on advanced pattern recognition"
		
	default:
		return vote, fmt.Errorf("unknown model type: %s", modelType)
	}
	
	return vote, nil
}

// simulateQModelResponse simulates Amazon Q model response
func (cs *ConsensusSystem) simulateQModelResponse(query, taskType string) string {
	queryLower := strings.ToLower(query)
	
	if strings.Contains(queryLower, "aws") || strings.Contains(queryLower, "cloud") {
		return "Proceed with AWS-optimized approach"
	}
	if strings.Contains(queryLower, "code") || strings.Contains(queryLower, "implement") {
		return "Implement with best practices and testing"
	}
	if strings.Contains(queryLower, "security") {
		return "Prioritize security and compliance"
	}
	
	return "Proceed with technical best practices"
}

// simulateClaudeResponse simulates Claude model response
func (cs *ConsensusSystem) simulateClaudeResponse(query, taskType string) string {
	queryLower := strings.ToLower(query)
	
	if strings.Contains(queryLower, "analyze") || strings.Contains(queryLower, "think") {
		return "Conduct thorough analysis before proceeding"
	}
	if strings.Contains(queryLower, "creative") || strings.Contains(queryLower, "design") {
		return "Explore creative solutions with user focus"
	}
	if strings.Contains(queryLower, "problem") {
		return "Break down problem systematically"
	}
	
	return "Approach with careful consideration"
}

// simulateFutureModelResponse simulates future model response
func (cs *ConsensusSystem) simulateFutureModelResponse(query, taskType string) string {
	queryLower := strings.ToLower(query)
	
	if strings.Contains(queryLower, "optimize") {
		return "Apply advanced optimization techniques"
	}
	if strings.Contains(queryLower, "predict") {
		return "Use predictive modeling approach"
	}
	
	return "Apply next-generation methodologies"
}

// calculateConsensus determines the final consensus from votes
func (cs *ConsensusSystem) calculateConsensus(votes []ConsensusVote) *ConsensusResult {
	if len(votes) == 0 {
		return &ConsensusResult{
			FinalDecision: "No consensus possible",
			Confidence:    0.0,
			ConsensusType: "none",
		}
	}
	
	// Group votes by decision
	decisionGroups := make(map[string][]ConsensusVote)
	for _, vote := range votes {
		decisionGroups[vote.Decision] = append(decisionGroups[vote.Decision], vote)
	}
	
	// Calculate weighted scores for each decision
	decisionScores := make(map[string]float64)
	for decision, groupVotes := range decisionGroups {
		score := 0.0
		for _, vote := range groupVotes {
			weight := cs.modelWeights[vote.ModelType]
			score += vote.Confidence * weight
		}
		decisionScores[decision] = score
	}
	
	// Find the decision with highest score
	var bestDecision string
	var bestScore float64
	for decision, score := range decisionScores {
		if score > bestScore {
			bestScore = score
			bestDecision = decision
		}
	}
	
	// Determine consensus type
	consensusType := cs.determineConsensusType(votes, bestDecision)
	
	// Calculate overall confidence
	confidence := cs.calculateOverallConfidence(votes, bestDecision, bestScore)
	
	return &ConsensusResult{
		FinalDecision: bestDecision,
		Confidence:    confidence,
		Votes:         votes,
		ConsensusType: consensusType,
	}
}

// determineConsensusType determines the type of consensus reached
func (cs *ConsensusSystem) determineConsensusType(votes []ConsensusVote, decision string) string {
	agreementCount := 0
	for _, vote := range votes {
		if vote.Decision == decision {
			agreementCount++
		}
	}
	
	if agreementCount == len(votes) {
		return "unanimous"
	} else if float64(agreementCount)/float64(len(votes)) >= 0.6 {
		return "majority"
	} else {
		return "weighted"
	}
}

// calculateOverallConfidence calculates the overall confidence in the consensus
func (cs *ConsensusSystem) calculateOverallConfidence(votes []ConsensusVote, decision string, score float64) float64 {
	agreementVotes := make([]ConsensusVote, 0)
	for _, vote := range votes {
		if vote.Decision == decision {
			agreementVotes = append(agreementVotes, vote)
		}
	}
	
	if len(agreementVotes) == 0 {
		return 0.0
	}
	
	// Average confidence of agreeing models
	totalConfidence := 0.0
	for _, vote := range agreementVotes {
		totalConfidence += vote.Confidence
	}
	avgConfidence := totalConfidence / float64(len(agreementVotes))
	
	// Adjust based on agreement ratio
	agreementRatio := float64(len(agreementVotes)) / float64(len(votes))
	
	return avgConfidence * agreementRatio
}

// updateModelWeights updates model weights based on consensus outcomes
func (cs *ConsensusSystem) updateModelWeights(votes []ConsensusVote, result *ConsensusResult) {
	// Increase weights for models that agreed with final decision
	for _, vote := range votes {
		if vote.Decision == result.FinalDecision {
			cs.modelWeights[vote.ModelType] *= 1.05 // Small increase
		} else {
			cs.modelWeights[vote.ModelType] *= 0.98 // Small decrease
		}
		
		// Keep weights within reasonable bounds
		if cs.modelWeights[vote.ModelType] > 2.0 {
			cs.modelWeights[vote.ModelType] = 2.0
		} else if cs.modelWeights[vote.ModelType] < 0.3 {
			cs.modelWeights[vote.ModelType] = 0.3
		}
	}
}

// GetConsensusHistory returns recent consensus decisions
func (cs *ConsensusSystem) GetConsensusHistory(limit int) []ConsensusResult {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()
	
	if len(cs.consensusHistory) <= limit {
		return cs.consensusHistory
	}
	return cs.consensusHistory[len(cs.consensusHistory)-limit:]
}

// GetModelWeights returns current model weights
func (cs *ConsensusSystem) GetModelWeights() map[ModelType]float64 {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()
	
	weights := make(map[ModelType]float64)
	for model, weight := range cs.modelWeights {
		weights[model] = weight
	}
	return weights
}

// GetConsensusStats returns statistics about consensus performance
func (cs *ConsensusSystem) GetConsensusStats() map[string]interface{} {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()
	
	if len(cs.consensusHistory) == 0 {
		return map[string]interface{}{
			"total_consensus": 0,
			"avg_confidence":  0.0,
		}
	}
	
	totalConfidence := 0.0
	consensusTypes := make(map[string]int)
	
	for _, result := range cs.consensusHistory {
		totalConfidence += result.Confidence
		consensusTypes[result.ConsensusType]++
	}
	
	return map[string]interface{}{
		"total_consensus":  len(cs.consensusHistory),
		"avg_confidence":   totalConfidence / float64(len(cs.consensusHistory)),
		"consensus_types":  consensusTypes,
		"model_weights":    cs.modelWeights,
	}
}
