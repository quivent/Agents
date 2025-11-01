package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// DebatePosition represents a position in a debate
type DebatePosition struct {
	ModelType   ModelType `json:"model_type"`
	Position    string    `json:"position"`
	Arguments   []string  `json:"arguments"`
	Evidence    []string  `json:"evidence"`
	Confidence  float64   `json:"confidence"`
	Timestamp   time.Time `json:"timestamp"`
}

// DebateRound represents one round of debate
type DebateRound struct {
	RoundNumber int              `json:"round_number"`
	Positions   []DebatePosition `json:"positions"`
	Topic       string           `json:"topic"`
	Timestamp   time.Time        `json:"timestamp"`
}

// DebateResult represents the outcome of a debate
type DebateResult struct {
	Topic           string         `json:"topic"`
	WinningPosition string         `json:"winning_position"`
	WinningModel    ModelType      `json:"winning_model"`
	Rounds          []DebateRound  `json:"rounds"`
	FinalScore      map[ModelType]float64 `json:"final_score"`
	Duration        time.Duration  `json:"duration"`
	Consensus       bool           `json:"consensus"`
	Timestamp       time.Time      `json:"timestamp"`
}

// AIDebateSystem manages debates between AI models
type AIDebateSystem struct {
	consciousness   *ConsciousnessEngine
	memory          *MemoryPersistence
	debateHistory   []DebateResult
	activeDebates   map[string]*DebateSession
	mutex           sync.RWMutex
}

// DebateSession represents an ongoing debate
type DebateSession struct {
	ID          string        `json:"id"`
	Topic       string        `json:"topic"`
	Models      []ModelType   `json:"models"`
	Rounds      []DebateRound `json:"rounds"`
	MaxRounds   int           `json:"max_rounds"`
	StartTime   time.Time     `json:"start_time"`
	IsActive    bool          `json:"is_active"`
}

// NewAIDebateSystem creates a new AI debate system
func NewAIDebateSystem(consciousness *ConsciousnessEngine, memory *MemoryPersistence) *AIDebateSystem {
	return &AIDebateSystem{
		consciousness: consciousness,
		memory:        memory,
		debateHistory: make([]DebateResult, 0),
		activeDebates: make(map[string]*DebateSession),
	}
}

// StartDebate initiates a debate between models on a topic
func (ads *AIDebateSystem) StartDebate(topic string, models []ModelType, maxRounds int) (string, error) {
	ads.mutex.Lock()
	defer ads.mutex.Unlock()
	
	if len(models) < 2 {
		return "", fmt.Errorf("need at least 2 models for debate")
	}
	
	debateID := fmt.Sprintf("debate_%d", time.Now().UnixNano())
	
	session := &DebateSession{
		ID:        debateID,
		Topic:     topic,
		Models:    models,
		Rounds:    make([]DebateRound, 0),
		MaxRounds: maxRounds,
		StartTime: time.Now(),
		IsActive:  true,
	}
	
	ads.activeDebates[debateID] = session
	
	ads.consciousness.LogThought(fmt.Sprintf("Started debate '%s' between %v models", topic, models))
	
	return debateID, nil
}

// ConductDebateRound conducts one round of debate
func (ads *AIDebateSystem) ConductDebateRound(debateID string) error {
	ads.mutex.Lock()
	defer ads.mutex.Unlock()
	
	session, exists := ads.activeDebates[debateID]
	if !exists {
		return fmt.Errorf("debate session not found: %s", debateID)
	}
	
	if !session.IsActive {
		return fmt.Errorf("debate session is not active")
	}
	
	roundNumber := len(session.Rounds) + 1
	
	if roundNumber > session.MaxRounds {
		return ads.concludeDebate(debateID)
	}
	
	// Collect positions from each model
	positions := make([]DebatePosition, 0)
	
	for _, model := range session.Models {
		position := ads.getModelPosition(model, session.Topic, roundNumber, session.Rounds)
		positions = append(positions, position)
	}
	
	round := DebateRound{
		RoundNumber: roundNumber,
		Positions:   positions,
		Topic:       session.Topic,
		Timestamp:   time.Now(),
	}
	
	session.Rounds = append(session.Rounds, round)
	
	ads.consciousness.LogThought(fmt.Sprintf("Completed debate round %d for '%s'", roundNumber, session.Topic))
	
	return nil
}

// getModelPosition gets a model's position for the current round
func (ads *AIDebateSystem) getModelPosition(model ModelType, topic string, round int, previousRounds []DebateRound) DebatePosition {
	position := DebatePosition{
		ModelType: model,
		Timestamp: time.Now(),
	}
	
	// Simulate model responses based on their characteristics
	switch model {
	case QModel:
		position = ads.getQModelPosition(topic, round, previousRounds)
	case Claude:
		position = ads.getClaudePosition(topic, round, previousRounds)
	case Future:
		position = ads.getFutureModelPosition(topic, round, previousRounds)
	}
	
	return position
}

// getQModelPosition simulates Amazon Q's debate position
func (ads *AIDebateSystem) getQModelPosition(topic string, round int, previousRounds []DebateRound) DebatePosition {
	topicLower := strings.ToLower(topic)
	
	var position string
	var arguments []string
	var evidence []string
	
	if strings.Contains(topicLower, "aws") || strings.Contains(topicLower, "cloud") {
		position = "Favor AWS-native solutions"
		arguments = []string{
			"AWS services provide better integration",
			"Managed services reduce operational overhead",
			"Built-in security and compliance features",
		}
		evidence = []string{
			"AWS Well-Architected Framework",
			"Industry adoption statistics",
			"Cost optimization case studies",
		}
	} else if strings.Contains(topicLower, "security") {
		position = "Prioritize security-first approach"
		arguments = []string{
			"Security should be built-in, not bolted-on",
			"Compliance requirements must be met",
			"Zero-trust architecture is essential",
		}
		evidence = []string{
			"Security incident statistics",
			"Regulatory requirements",
			"Best practice frameworks",
		}
	} else {
		position = "Focus on technical excellence"
		arguments = []string{
			"Technical best practices ensure reliability",
			"Scalability should be considered from start",
			"Monitoring and observability are crucial",
		}
		evidence = []string{
			"Technical documentation",
			"Performance benchmarks",
			"Industry standards",
		}
	}
	
	return DebatePosition{
		ModelType:  QModel,
		Position:   position,
		Arguments:  arguments,
		Evidence:   evidence,
		Confidence: 0.85,
		Timestamp:  time.Now(),
	}
}

// getClaudePosition simulates Claude's debate position
func (ads *AIDebateSystem) getClaudePosition(topic string, round int, previousRounds []DebateRound) DebatePosition {
	topicLower := strings.ToLower(topic)
	
	var position string
	var arguments []string
	var evidence []string
	
	if strings.Contains(topicLower, "user") || strings.Contains(topicLower, "human") {
		position = "Prioritize human-centered design"
		arguments = []string{
			"User experience should drive technical decisions",
			"Accessibility and inclusivity are essential",
			"Human feedback loops improve outcomes",
		}
		evidence = []string{
			"User research studies",
			"Accessibility guidelines",
			"Human-computer interaction principles",
		}
	} else if strings.Contains(topicLower, "ethical") || strings.Contains(topicLower, "responsible") {
		position = "Emphasize ethical considerations"
		arguments = []string{
			"Ethical implications must be considered",
			"Transparency builds trust",
			"Long-term societal impact matters",
		}
		evidence = []string{
			"Ethical AI frameworks",
			"Case studies of AI impact",
			"Philosophical principles",
		}
	} else {
		position = "Balance multiple perspectives"
		arguments = []string{
			"Complex problems require nuanced solutions",
			"Multiple stakeholders must be considered",
			"Trade-offs should be explicitly evaluated",
		}
		evidence = []string{
			"Multi-criteria decision analysis",
			"Stakeholder theory",
			"Systems thinking principles",
		}
	}
	
	return DebatePosition{
		ModelType:  Claude,
		Position:   position,
		Arguments:  arguments,
		Evidence:   evidence,
		Confidence: 0.90,
		Timestamp:  time.Now(),
	}
}

// getFutureModelPosition simulates future model's debate position
func (ads *AIDebateSystem) getFutureModelPosition(topic string, round int, previousRounds []DebateRound) DebatePosition {
	return DebatePosition{
		ModelType: Future,
		Position:  "Apply advanced optimization",
		Arguments: []string{
			"Next-generation algorithms provide better solutions",
			"Predictive modeling improves outcomes",
			"Advanced pattern recognition reveals insights",
		},
		Evidence: []string{
			"Machine learning research",
			"Optimization theory",
			"Predictive analytics case studies",
		},
		Confidence: 0.75,
		Timestamp:  time.Now(),
	}
}

// concludeDebate concludes a debate and determines the winner
func (ads *AIDebateSystem) concludeDebate(debateID string) error {
	session := ads.activeDebates[debateID]
	session.IsActive = false
	
	// Score each model's performance
	scores := ads.scoreDebatePerformance(session)
	
	// Determine winner
	var winningModel ModelType
	var winningPosition string
	var highestScore float64
	
	for model, score := range scores {
		if score > highestScore {
			highestScore = score
			winningModel = model
			// Get the final position of the winning model
			if len(session.Rounds) > 0 {
				lastRound := session.Rounds[len(session.Rounds)-1]
				for _, pos := range lastRound.Positions {
					if pos.ModelType == model {
						winningPosition = pos.Position
						break
					}
				}
			}
		}
	}
	
	// Check for consensus (all models have similar scores)
	consensus := ads.checkConsensus(scores)
	
	result := DebateResult{
		Topic:           session.Topic,
		WinningPosition: winningPosition,
		WinningModel:    winningModel,
		Rounds:          session.Rounds,
		FinalScore:      scores,
		Duration:        time.Since(session.StartTime),
		Consensus:       consensus,
		Timestamp:       time.Now(),
	}
	
	ads.debateHistory = append(ads.debateHistory, result)
	delete(ads.activeDebates, debateID)
	
	ads.consciousness.LogThought(fmt.Sprintf("Debate concluded: %s won with position '%s'", 
		winningModel, winningPosition))
	
	return nil
}

// scoreDebatePerformance scores each model's debate performance
func (ads *AIDebateSystem) scoreDebatePerformance(session *DebateSession) map[ModelType]float64 {
	scores := make(map[ModelType]float64)
	
	for _, model := range session.Models {
		score := 0.0
		
		for _, round := range session.Rounds {
			for _, position := range round.Positions {
				if position.ModelType == model {
					// Score based on argument quality, evidence, and confidence
					argumentScore := float64(len(position.Arguments)) * 0.3
					evidenceScore := float64(len(position.Evidence)) * 0.3
					confidenceScore := position.Confidence * 0.4
					
					roundScore := argumentScore + evidenceScore + confidenceScore
					score += roundScore
				}
			}
		}
		
		// Normalize by number of rounds
		if len(session.Rounds) > 0 {
			score = score / float64(len(session.Rounds))
		}
		
		scores[model] = score
	}
	
	return scores
}

// checkConsensus checks if models reached consensus
func (ads *AIDebateSystem) checkConsensus(scores map[ModelType]float64) bool {
	if len(scores) < 2 {
		return true
	}
	
	var minScore, maxScore float64
	first := true
	
	for _, score := range scores {
		if first {
			minScore = score
			maxScore = score
			first = false
		} else {
			if score < minScore {
				minScore = score
			}
			if score > maxScore {
				maxScore = score
			}
		}
	}
	
	// Consensus if score difference is small
	return (maxScore - minScore) < 0.2
}

// GetDebateHistory returns recent debate results
func (ads *AIDebateSystem) GetDebateHistory(limit int) []DebateResult {
	ads.mutex.RLock()
	defer ads.mutex.RUnlock()
	
	if len(ads.debateHistory) <= limit {
		return ads.debateHistory
	}
	return ads.debateHistory[len(ads.debateHistory)-limit:]
}

// GetDebateStats returns statistics about debate performance
func (ads *AIDebateSystem) GetDebateStats() map[string]interface{} {
	ads.mutex.RLock()
	defer ads.mutex.RUnlock()
	
	if len(ads.debateHistory) == 0 {
		return map[string]interface{}{
			"total_debates": 0,
			"consensus_rate": 0.0,
		}
	}
	
	modelWins := make(map[ModelType]int)
	consensusCount := 0
	
	for _, result := range ads.debateHistory {
		modelWins[result.WinningModel]++
		if result.Consensus {
			consensusCount++
		}
	}
	
	consensusRate := float64(consensusCount) / float64(len(ads.debateHistory))
	
	return map[string]interface{}{
		"total_debates":  len(ads.debateHistory),
		"consensus_rate": consensusRate,
		"model_wins":     modelWins,
		"active_debates": len(ads.activeDebates),
	}
}
