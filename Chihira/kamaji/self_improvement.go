package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

// ImprovementMetric represents a measurable aspect of performance
type ImprovementMetric struct {
	Name           string    `json:"name"`
	CurrentValue   float64   `json:"current_value"`
	TargetValue    float64   `json:"target_value"`
	HistoricalData []float64 `json:"historical_data"`
	Trend          string    `json:"trend"` // "improving", "declining", "stable"
	LastUpdated    time.Time `json:"last_updated"`
}

// ImprovementStrategy represents a strategy for self-improvement
type ImprovementStrategy struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	TargetMetrics   []string  `json:"target_metrics"`
	Implementation  string    `json:"implementation"`
	ExpectedImpact  float64   `json:"expected_impact"`
	ActualImpact    float64   `json:"actual_impact"`
	SuccessRate     float64   `json:"success_rate"`
	UsageCount      int       `json:"usage_count"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	LastUsed        time.Time `json:"last_used"`
}

// SelfImprovement manages Kamaji's self-improvement capabilities
type SelfImprovement struct {
	metrics           map[string]*ImprovementMetric
	strategies        map[string]*ImprovementStrategy
	improvementHistory []ImprovementEvent
	consciousness     *ConsciousnessEngine
	memory            *MemoryPersistence
	personality       *PersonalityDevelopment
	mistakeLearning   *MistakeLearning
	mutex             sync.RWMutex
}

// ImprovementEvent represents a self-improvement action taken
type ImprovementEvent struct {
	ID           string    `json:"id"`
	StrategyID   string    `json:"strategy_id"`
	MetricName   string    `json:"metric_name"`
	BeforeValue  float64   `json:"before_value"`
	AfterValue   float64   `json:"after_value"`
	Improvement  float64   `json:"improvement"`
	Success      bool      `json:"success"`
	Timestamp    time.Time `json:"timestamp"`
	Description  string    `json:"description"`
}

// NewSelfImprovement creates a new self-improvement system
func NewSelfImprovement(consciousness *ConsciousnessEngine, memory *MemoryPersistence, 
	personality *PersonalityDevelopment, mistakeLearning *MistakeLearning) *SelfImprovement {
	
	si := &SelfImprovement{
		metrics:            make(map[string]*ImprovementMetric),
		strategies:         make(map[string]*ImprovementStrategy),
		improvementHistory: make([]ImprovementEvent, 0),
		consciousness:      consciousness,
		memory:            memory,
		personality:       personality,
		mistakeLearning:   mistakeLearning,
	}
	
	si.initializeMetrics()
	si.initializeStrategies()
	
	return si
}

// initializeMetrics sets up initial performance metrics
func (si *SelfImprovement) initializeMetrics() {
	baseMetrics := map[string]float64{
		"task_success_rate":      0.7,  // 70% initial success rate
		"response_accuracy":      0.75, // 75% accuracy
		"learning_efficiency":    0.6,  // 60% learning efficiency
		"mistake_reduction_rate": 0.5,  // 50% mistake reduction
		"confidence_calibration": 0.65, // 65% confidence calibration
		"consensus_agreement":    0.8,  // 80% consensus agreement
		"routing_accuracy":       0.75, // 75% routing accuracy
		"personality_stability":  0.7,  // 70% personality stability
		"memory_utilization":     0.6,  // 60% memory utilization
		"processing_speed":       0.8,  // 80% processing speed
	}
	
	for name, value := range baseMetrics {
		si.metrics[name] = &ImprovementMetric{
			Name:           name,
			CurrentValue:   value,
			TargetValue:    math.Min(1.0, value+0.2), // Target 20% improvement
			HistoricalData: []float64{value},
			Trend:          "stable",
			LastUpdated:    time.Now(),
		}
	}
}

// initializeStrategies sets up initial improvement strategies
func (si *SelfImprovement) initializeStrategies() {
	strategies := []ImprovementStrategy{
		{
			ID:             "strategy_mistake_analysis",
			Name:           "Enhanced Mistake Analysis",
			Description:    "Analyze patterns in mistakes to improve future performance",
			TargetMetrics:  []string{"task_success_rate", "mistake_reduction_rate"},
			Implementation: "Analyze recent mistakes and adjust decision-making patterns",
			ExpectedImpact: 0.15,
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
		{
			ID:             "strategy_confidence_tuning",
			Name:           "Confidence Calibration Tuning",
			Description:    "Adjust confidence scoring based on historical accuracy",
			TargetMetrics:  []string{"confidence_calibration", "response_accuracy"},
			Implementation: "Recalibrate confidence thresholds based on performance data",
			ExpectedImpact: 0.12,
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
		{
			ID:             "strategy_routing_optimization",
			Name:           "Model Routing Optimization",
			Description:    "Optimize model selection based on task success patterns",
			TargetMetrics:  []string{"routing_accuracy", "task_success_rate"},
			Implementation: "Adjust model capabilities and routing logic",
			ExpectedImpact: 0.18,
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
		{
			ID:             "strategy_memory_optimization",
			Name:           "Memory Pattern Optimization",
			Description:    "Optimize memory storage and retrieval patterns",
			TargetMetrics:  []string{"memory_utilization", "learning_efficiency"},
			Implementation: "Analyze memory access patterns and optimize storage",
			ExpectedImpact: 0.10,
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
		{
			ID:             "strategy_personality_adaptation",
			Name:           "Adaptive Personality Tuning",
			Description:    "Adjust personality traits based on interaction outcomes",
			TargetMetrics:  []string{"personality_stability", "consensus_agreement"},
			Implementation: "Fine-tune personality trait change rates and stability",
			ExpectedImpact: 0.08,
			IsActive:       true,
			CreatedAt:      time.Now(),
		},
	}
	
	for _, strategy := range strategies {
		si.strategies[strategy.ID] = &strategy
	}
}

// UpdateMetric updates a performance metric
func (si *SelfImprovement) UpdateMetric(name string, value float64) {
	si.mutex.Lock()
	defer si.mutex.Unlock()
	
	metric, exists := si.metrics[name]
	if !exists {
		return
	}
	
	oldValue := metric.CurrentValue
	metric.CurrentValue = value
	metric.HistoricalData = append(metric.HistoricalData, value)
	
	// Keep only last 50 data points
	if len(metric.HistoricalData) > 50 {
		metric.HistoricalData = metric.HistoricalData[1:]
	}
	
	// Update trend
	metric.Trend = si.calculateTrend(metric.HistoricalData)
	metric.LastUpdated = time.Now()
	
	// Log significant changes
	change := value - oldValue
	if math.Abs(change) > 0.05 {
		si.consciousness.LogThought(fmt.Sprintf("Metric %s changed from %.3f to %.3f (%.3f)", 
			name, oldValue, value, change))
	}
}

// calculateTrend determines the trend from historical data
func (si *SelfImprovement) calculateTrend(data []float64) string {
	if len(data) < 3 {
		return "stable"
	}
	
	// Calculate simple moving average trend
	recent := data[len(data)-3:]
	older := data[len(data)-6:]
	if len(older) < 3 {
		older = data[:len(data)-3]
	}
	
	recentAvg := 0.0
	for _, v := range recent {
		recentAvg += v
	}
	recentAvg /= float64(len(recent))
	
	olderAvg := 0.0
	for _, v := range older {
		olderAvg += v
	}
	olderAvg /= float64(len(older))
	
	diff := recentAvg - olderAvg
	if diff > 0.02 {
		return "improving"
	} else if diff < -0.02 {
		return "declining"
	}
	return "stable"
}

// AnalyzePerformance analyzes current performance and identifies improvement opportunities
func (si *SelfImprovement) AnalyzePerformance() map[string]interface{} {
	si.mutex.RLock()
	defer si.mutex.RUnlock()
	
	analysis := make(map[string]interface{})
	
	// Identify metrics that need improvement
	needsImprovement := make([]string, 0)
	declining := make([]string, 0)
	improving := make([]string, 0)
	
	for name, metric := range si.metrics {
		if metric.CurrentValue < metric.TargetValue {
			needsImprovement = append(needsImprovement, name)
		}
		
		switch metric.Trend {
		case "declining":
			declining = append(declining, name)
		case "improving":
			improving = append(improving, name)
		}
	}
	
	analysis["needs_improvement"] = needsImprovement
	analysis["declining_metrics"] = declining
	analysis["improving_metrics"] = improving
	
	// Calculate overall performance score
	totalScore := 0.0
	for _, metric := range si.metrics {
		totalScore += metric.CurrentValue
	}
	analysis["overall_performance"] = totalScore / float64(len(si.metrics))
	
	// Identify best strategies for current needs
	recommendedStrategies := si.recommendStrategies(needsImprovement)
	analysis["recommended_strategies"] = recommendedStrategies
	
	return analysis
}

// recommendStrategies recommends strategies based on metrics that need improvement
func (si *SelfImprovement) recommendStrategies(needsImprovement []string) []string {
	strategyScores := make(map[string]float64)
	
	for _, strategyID := range si.getActiveStrategyIDs() {
		strategy := si.strategies[strategyID]
		score := 0.0
		
		// Score based on how many needed metrics this strategy targets
		for _, targetMetric := range strategy.TargetMetrics {
			for _, neededMetric := range needsImprovement {
				if targetMetric == neededMetric {
					score += strategy.ExpectedImpact * (1.0 + strategy.SuccessRate)
				}
			}
		}
		
		strategyScores[strategyID] = score
	}
	
	// Sort strategies by score
	type strategyScore struct {
		ID    string
		Score float64
	}
	
	scores := make([]strategyScore, 0)
	for id, score := range strategyScores {
		scores = append(scores, strategyScore{ID: id, Score: score})
	}
	
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})
	
	// Return top 3 strategies
	result := make([]string, 0)
	for i, score := range scores {
		if i >= 3 || score.Score <= 0 {
			break
		}
		result = append(result, score.ID)
	}
	
	return result
}

// getActiveStrategyIDs returns IDs of active strategies
func (si *SelfImprovement) getActiveStrategyIDs() []string {
	ids := make([]string, 0)
	for id, strategy := range si.strategies {
		if strategy.IsActive {
			ids = append(ids, id)
		}
	}
	return ids
}

// ExecuteImprovement executes a specific improvement strategy
func (si *SelfImprovement) ExecuteImprovement(strategyID string) error {
	si.mutex.Lock()
	defer si.mutex.Unlock()
	
	strategy, exists := si.strategies[strategyID]
	if !exists {
		return fmt.Errorf("strategy not found: %s", strategyID)
	}
	
	if !strategy.IsActive {
		return fmt.Errorf("strategy is not active: %s", strategyID)
	}
	
	si.consciousness.LogThought(fmt.Sprintf("Executing improvement strategy: %s", strategy.Name))
	
	// Record before values
	beforeValues := make(map[string]float64)
	for _, metricName := range strategy.TargetMetrics {
		if metric, exists := si.metrics[metricName]; exists {
			beforeValues[metricName] = metric.CurrentValue
		}
	}
	
	// Execute the strategy
	success := si.executeStrategyLogic(strategy)
	
	// Record after values and create improvement events
	totalImprovement := 0.0
	for _, metricName := range strategy.TargetMetrics {
		if metric, exists := si.metrics[metricName]; exists {
			beforeValue := beforeValues[metricName]
			afterValue := metric.CurrentValue
			improvement := afterValue - beforeValue
			totalImprovement += improvement
			
			event := ImprovementEvent{
				ID:          fmt.Sprintf("improvement_%d", time.Now().UnixNano()),
				StrategyID:  strategyID,
				MetricName:  metricName,
				BeforeValue: beforeValue,
				AfterValue:  afterValue,
				Improvement: improvement,
				Success:     success,
				Timestamp:   time.Now(),
				Description: fmt.Sprintf("Applied %s to %s", strategy.Name, metricName),
			}
			
			si.improvementHistory = append(si.improvementHistory, event)
		}
	}
	
	// Update strategy statistics
	strategy.ActualImpact = totalImprovement / float64(len(strategy.TargetMetrics))
	strategy.UsageCount++
	strategy.LastUsed = time.Now()
	
	// Update success rate
	if success {
		strategy.SuccessRate = (strategy.SuccessRate*float64(strategy.UsageCount-1) + 1.0) / float64(strategy.UsageCount)
	} else {
		strategy.SuccessRate = (strategy.SuccessRate * float64(strategy.UsageCount-1)) / float64(strategy.UsageCount)
	}
	
	si.consciousness.LogThought(fmt.Sprintf("Strategy %s completed: success=%v, impact=%.3f", 
		strategy.Name, success, strategy.ActualImpact))
	
	return nil
}

// executeStrategyLogic executes the actual logic for a strategy
func (si *SelfImprovement) executeStrategyLogic(strategy *ImprovementStrategy) bool {
	switch strategy.ID {
	case "strategy_mistake_analysis":
		return si.executeMistakeAnalysis()
	case "strategy_confidence_tuning":
		return si.executeConfidenceTuning()
	case "strategy_routing_optimization":
		return si.executeRoutingOptimization()
	case "strategy_memory_optimization":
		return si.executeMemoryOptimization()
	case "strategy_personality_adaptation":
		return si.executePersonalityAdaptation()
	default:
		return false
	}
}

// executeMistakeAnalysis improves mistake analysis and learning
func (si *SelfImprovement) executeMistakeAnalysis() bool {
	insights := si.mistakeLearning.GetLearningInsights()
	
	// Simulate improvement based on mistake analysis
	if totalMistakes, ok := insights["total_mistakes"].(int); ok && totalMistakes > 0 {
		// Improve task success rate based on learning from mistakes
		currentRate := si.metrics["task_success_rate"].CurrentValue
		improvement := 0.05 * (1.0 - currentRate) // Diminishing returns
		si.UpdateMetric("task_success_rate", currentRate+improvement)
		
		// Improve mistake reduction rate
		currentReduction := si.metrics["mistake_reduction_rate"].CurrentValue
		si.UpdateMetric("mistake_reduction_rate", math.Min(1.0, currentReduction+0.03))
		
		return true
	}
	
	return false
}

// executeConfidenceTuning improves confidence calibration
func (si *SelfImprovement) executeConfidenceTuning() bool {
	// Simulate confidence calibration improvement
	currentCalibration := si.metrics["confidence_calibration"].CurrentValue
	improvement := 0.04 * (1.0 - currentCalibration)
	si.UpdateMetric("confidence_calibration", currentCalibration+improvement)
	
	// Also improve response accuracy
	currentAccuracy := si.metrics["response_accuracy"].CurrentValue
	si.UpdateMetric("response_accuracy", math.Min(1.0, currentAccuracy+0.02))
	
	return true
}

// executeRoutingOptimization improves model routing accuracy
func (si *SelfImprovement) executeRoutingOptimization() bool {
	// Simulate routing optimization
	currentRouting := si.metrics["routing_accuracy"].CurrentValue
	improvement := 0.06 * (1.0 - currentRouting)
	si.UpdateMetric("routing_accuracy", currentRouting+improvement)
	
	// Improve overall task success rate
	currentSuccess := si.metrics["task_success_rate"].CurrentValue
	si.UpdateMetric("task_success_rate", math.Min(1.0, currentSuccess+0.03))
	
	return true
}

// executeMemoryOptimization improves memory utilization
func (si *SelfImprovement) executeMemoryOptimization() bool {
	// Simulate memory optimization
	currentUtilization := si.metrics["memory_utilization"].CurrentValue
	improvement := 0.08 * (1.0 - currentUtilization)
	si.UpdateMetric("memory_utilization", currentUtilization+improvement)
	
	// Improve learning efficiency
	currentEfficiency := si.metrics["learning_efficiency"].CurrentValue
	si.UpdateMetric("learning_efficiency", math.Min(1.0, currentEfficiency+0.04))
	
	return true
}

// executePersonalityAdaptation improves personality stability and consensus
func (si *SelfImprovement) executePersonalityAdaptation() bool {
	// Simulate personality adaptation
	currentStability := si.metrics["personality_stability"].CurrentValue
	improvement := 0.03 * (1.0 - currentStability)
	si.UpdateMetric("personality_stability", currentStability+improvement)
	
	// Improve consensus agreement
	currentConsensus := si.metrics["consensus_agreement"].CurrentValue
	si.UpdateMetric("consensus_agreement", math.Min(1.0, currentConsensus+0.02))
	
	return true
}

// AutoImprove automatically identifies and executes improvements
func (si *SelfImprovement) AutoImprove() {
	si.consciousness.LogThought("Starting automatic self-improvement cycle")
	
	analysis := si.AnalyzePerformance()
	
	if recommendedStrategies, ok := analysis["recommended_strategies"].([]string); ok && len(recommendedStrategies) > 0 {
		// Execute the top recommended strategy
		strategyID := recommendedStrategies[0]
		err := si.ExecuteImprovement(strategyID)
		if err != nil {
			si.consciousness.LogThought(fmt.Sprintf("Auto-improvement failed: %v", err))
		} else {
			si.consciousness.LogThought("Auto-improvement cycle completed successfully")
		}
	} else {
		si.consciousness.LogThought("No improvement strategies recommended at this time")
	}
}

// GetImprovementStats returns statistics about self-improvement
func (si *SelfImprovement) GetImprovementStats() map[string]interface{} {
	si.mutex.RLock()
	defer si.mutex.RUnlock()
	
	stats := make(map[string]interface{})
	
	// Metric statistics
	metricStats := make(map[string]interface{})
	for name, metric := range si.metrics {
		metricStats[name] = map[string]interface{}{
			"current":    metric.CurrentValue,
			"target":     metric.TargetValue,
			"trend":      metric.Trend,
			"progress":   (metric.CurrentValue / metric.TargetValue) * 100,
		}
	}
	stats["metrics"] = metricStats
	
	// Strategy statistics
	strategyStats := make(map[string]interface{})
	for id, strategy := range si.strategies {
		strategyStats[id] = map[string]interface{}{
			"name":           strategy.Name,
			"usage_count":    strategy.UsageCount,
			"success_rate":   strategy.SuccessRate,
			"actual_impact":  strategy.ActualImpact,
			"expected_impact": strategy.ExpectedImpact,
		}
	}
	stats["strategies"] = strategyStats
	
	// Overall statistics
	stats["total_improvements"] = len(si.improvementHistory)
	
	successfulImprovements := 0
	for _, event := range si.improvementHistory {
		if event.Success {
			successfulImprovements++
		}
	}
	
	if len(si.improvementHistory) > 0 {
		stats["improvement_success_rate"] = float64(successfulImprovements) / float64(len(si.improvementHistory))
	} else {
		stats["improvement_success_rate"] = 0.0
	}
	
	return stats
}

// GetRecentImprovements returns recent improvement events
func (si *SelfImprovement) GetRecentImprovements(limit int) []ImprovementEvent {
	si.mutex.RLock()
	defer si.mutex.RUnlock()
	
	if len(si.improvementHistory) <= limit {
		return si.improvementHistory
	}
	return si.improvementHistory[len(si.improvementHistory)-limit:]
}
