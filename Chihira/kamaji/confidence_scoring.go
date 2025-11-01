package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// ConfidencePrediction represents a model's prediction with confidence
type ConfidencePrediction struct {
	ModelType      ModelType `json:"model_type"`
	Prediction     string    `json:"prediction"`
	Confidence     float64   `json:"confidence"`
	Reasoning      string    `json:"reasoning"`
	Timestamp      time.Time `json:"timestamp"`
	ActualOutcome  string    `json:"actual_outcome"`
	WasCorrect     bool      `json:"was_correct"`
	CalibrationError float64 `json:"calibration_error"`
}

// ConfidenceMetrics represents confidence calibration metrics
type ConfidenceMetrics struct {
	ModelType        ModelType `json:"model_type"`
	TotalPredictions int       `json:"total_predictions"`
	Accuracy         float64   `json:"accuracy"`
	AvgConfidence    float64   `json:"avg_confidence"`
	Calibration      float64   `json:"calibration"`      // How well confidence matches accuracy
	Overconfidence   float64   `json:"overconfidence"`   // Tendency to be overconfident
	Reliability      float64   `json:"reliability"`      // Overall reliability score
	LastUpdated      time.Time `json:"last_updated"`
}

// ConfidenceScoring tracks and calibrates model confidence scores
type ConfidenceScoring struct {
	predictions     []ConfidencePrediction
	modelMetrics    map[ModelType]*ConfidenceMetrics
	consciousness   *ConsciousnessEngine
	memory          *MemoryPersistence
	mutex           sync.RWMutex
}

// NewConfidenceScoring creates a new confidence scoring system
func NewConfidenceScoring(consciousness *ConsciousnessEngine, memory *MemoryPersistence) *ConfidenceScoring {
	cs := &ConfidenceScoring{
		predictions:   make([]ConfidencePrediction, 0),
		modelMetrics:  make(map[ModelType]*ConfidenceMetrics),
		consciousness: consciousness,
		memory:        memory,
	}
	
	cs.initializeMetrics()
	return cs
}

// initializeMetrics initializes metrics for all model types
func (cs *ConfidenceScoring) initializeMetrics() {
	models := []ModelType{QModel, Claude, Future}
	
	for _, model := range models {
		cs.modelMetrics[model] = &ConfidenceMetrics{
			ModelType:        model,
			TotalPredictions: 0,
			Accuracy:         0.5, // Start with neutral accuracy
			AvgConfidence:    0.5, // Start with neutral confidence
			Calibration:      0.0, // Perfect calibration initially
			Overconfidence:   0.0, // No overconfidence initially
			Reliability:      0.5, // Neutral reliability
			LastUpdated:      time.Now(),
		}
	}
}

// RecordPrediction records a model's prediction with confidence
func (cs *ConfidenceScoring) RecordPrediction(modelType ModelType, prediction, reasoning string, confidence float64) string {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	
	predictionID := fmt.Sprintf("pred_%d", time.Now().UnixNano())
	
	pred := ConfidencePrediction{
		ModelType:     modelType,
		Prediction:    prediction,
		Confidence:    confidence,
		Reasoning:     reasoning,
		Timestamp:     time.Now(),
		ActualOutcome: "", // Will be filled later
		WasCorrect:    false,
	}
	
	cs.predictions = append(cs.predictions, pred)
	
	cs.consciousness.LogThought(fmt.Sprintf("%s made prediction with %.2f confidence: %s", 
		modelType, confidence, prediction))
	
	return predictionID
}

// UpdatePredictionOutcome updates the actual outcome of a prediction
func (cs *ConfidenceScoring) UpdatePredictionOutcome(modelType ModelType, prediction, actualOutcome string, wasCorrect bool) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	
	// Find the prediction to update
	for i := range cs.predictions {
		pred := &cs.predictions[i]
		if pred.ModelType == modelType && pred.Prediction == prediction && pred.ActualOutcome == "" {
			pred.ActualOutcome = actualOutcome
			pred.WasCorrect = wasCorrect
			
			// Calculate calibration error
			if wasCorrect {
				pred.CalibrationError = math.Abs(pred.Confidence - 1.0)
			} else {
				pred.CalibrationError = math.Abs(pred.Confidence - 0.0)
			}
			
			// Update model metrics
			cs.updateModelMetrics(modelType)
			
			cs.consciousness.LogThought(fmt.Sprintf("Updated prediction outcome for %s: %v (calibration error: %.3f)", 
				modelType, wasCorrect, pred.CalibrationError))
			
			break
		}
	}
}

// updateModelMetrics updates the metrics for a specific model
func (cs *ConfidenceScoring) updateModelMetrics(modelType ModelType) {
	metrics := cs.modelMetrics[modelType]
	
	// Get all predictions for this model with outcomes
	modelPredictions := make([]ConfidencePrediction, 0)
	for _, pred := range cs.predictions {
		if pred.ModelType == modelType && pred.ActualOutcome != "" {
			modelPredictions = append(modelPredictions, pred)
		}
	}
	
	if len(modelPredictions) == 0 {
		return
	}
	
	// Calculate accuracy
	correctCount := 0
	totalConfidence := 0.0
	totalCalibrationError := 0.0
	
	for _, pred := range modelPredictions {
		if pred.WasCorrect {
			correctCount++
		}
		totalConfidence += pred.Confidence
		totalCalibrationError += pred.CalibrationError
	}
	
	metrics.TotalPredictions = len(modelPredictions)
	metrics.Accuracy = float64(correctCount) / float64(len(modelPredictions))
	metrics.AvgConfidence = totalConfidence / float64(len(modelPredictions))
	metrics.Calibration = 1.0 - (totalCalibrationError / float64(len(modelPredictions)))
	
	// Calculate overconfidence (confidence - accuracy)
	metrics.Overconfidence = metrics.AvgConfidence - metrics.Accuracy
	
	// Calculate overall reliability (combination of accuracy and calibration)
	metrics.Reliability = (metrics.Accuracy + metrics.Calibration) / 2.0
	
	metrics.LastUpdated = time.Now()
}

// GetCalibratedConfidence returns a calibrated confidence score for a model
func (cs *ConfidenceScoring) GetCalibratedConfidence(modelType ModelType, rawConfidence float64) float64 {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()
	
	metrics, exists := cs.modelMetrics[modelType]
	if !exists || metrics.TotalPredictions < 5 {
		// Not enough data for calibration, return raw confidence
		return rawConfidence
	}
	
	// Apply calibration adjustment based on historical performance
	calibrationFactor := metrics.Calibration
	overconfidenceAdjustment := metrics.Overconfidence
	
	// Adjust for overconfidence
	adjustedConfidence := rawConfidence - (overconfidenceAdjustment * 0.5)
	
	// Apply calibration factor
	calibratedConfidence := adjustedConfidence * calibrationFactor
	
	// Ensure bounds [0, 1]
	if calibratedConfidence < 0.0 {
		calibratedConfidence = 0.0
	} else if calibratedConfidence > 1.0 {
		calibratedConfidence = 1.0
	}
	
	return calibratedConfidence
}

// GetModelReliability returns the reliability score for a model
func (cs *ConfidenceScoring) GetModelReliability(modelType ModelType) float64 {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()
	
	if metrics, exists := cs.modelMetrics[modelType]; exists {
		return metrics.Reliability
	}
	return 0.5 // Default neutral reliability
}

// GetConfidenceMetrics returns confidence metrics for a model
func (cs *ConfidenceScoring) GetConfidenceMetrics(modelType ModelType) *ConfidenceMetrics {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()
	
	if metrics, exists := cs.modelMetrics[modelType]; exists {
		// Return a copy to avoid race conditions
		metricsCopy := *metrics
		return &metricsCopy
	}
	return nil
}

// GetAllMetrics returns metrics for all models
func (cs *ConfidenceScoring) GetAllMetrics() map[ModelType]*ConfidenceMetrics {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()
	
	result := make(map[ModelType]*ConfidenceMetrics)
	for modelType, metrics := range cs.modelMetrics {
		metricsCopy := *metrics
		result[modelType] = &metricsCopy
	}
	
	return result
}

// GetRecentPredictions returns recent predictions with outcomes
func (cs *ConfidenceScoring) GetRecentPredictions(limit int) []ConfidencePrediction {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()
	
	if len(cs.predictions) <= limit {
		return cs.predictions
	}
	return cs.predictions[len(cs.predictions)-limit:]
}

// AnalyzeConfidencePatterns analyzes patterns in confidence scoring
func (cs *ConfidenceScoring) AnalyzeConfidencePatterns() map[string]interface{} {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()
	
	analysis := make(map[string]interface{})
	
	// Overall statistics
	totalPredictions := len(cs.predictions)
	predictionsWithOutcomes := 0
	totalCorrect := 0
	
	for _, pred := range cs.predictions {
		if pred.ActualOutcome != "" {
			predictionsWithOutcomes++
			if pred.WasCorrect {
				totalCorrect++
			}
		}
	}
	
	overallAccuracy := 0.0
	if predictionsWithOutcomes > 0 {
		overallAccuracy = float64(totalCorrect) / float64(predictionsWithOutcomes)
	}
	
	analysis["total_predictions"] = totalPredictions
	analysis["predictions_with_outcomes"] = predictionsWithOutcomes
	analysis["overall_accuracy"] = overallAccuracy
	
	// Model comparison
	modelComparison := make(map[string]interface{})
	for modelType, metrics := range cs.modelMetrics {
		modelComparison[string(modelType)] = map[string]interface{}{
			"accuracy":       metrics.Accuracy,
			"avg_confidence": metrics.AvgConfidence,
			"calibration":    metrics.Calibration,
			"overconfidence": metrics.Overconfidence,
			"reliability":    metrics.Reliability,
		}
	}
	analysis["model_comparison"] = modelComparison
	
	// Best performing model
	var bestModel ModelType
	var bestReliability float64
	for modelType, metrics := range cs.modelMetrics {
		if metrics.Reliability > bestReliability {
			bestReliability = metrics.Reliability
			bestModel = modelType
		}
	}
	analysis["best_model"] = string(bestModel)
	analysis["best_reliability"] = bestReliability
	
	return analysis
}

// GetConfidenceDistribution returns the distribution of confidence scores
func (cs *ConfidenceScoring) GetConfidenceDistribution(modelType ModelType) map[string]int {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()
	
	distribution := map[string]int{
		"0.0-0.2": 0,
		"0.2-0.4": 0,
		"0.4-0.6": 0,
		"0.6-0.8": 0,
		"0.8-1.0": 0,
	}
	
	for _, pred := range cs.predictions {
		if pred.ModelType == modelType {
			conf := pred.Confidence
			if conf < 0.2 {
				distribution["0.0-0.2"]++
			} else if conf < 0.4 {
				distribution["0.2-0.4"]++
			} else if conf < 0.6 {
				distribution["0.4-0.6"]++
			} else if conf < 0.8 {
				distribution["0.6-0.8"]++
			} else {
				distribution["0.8-1.0"]++
			}
		}
	}
	
	return distribution
}

// RecommendConfidenceThreshold recommends a confidence threshold for decision making
func (cs *ConfidenceScoring) RecommendConfidenceThreshold(modelType ModelType, targetAccuracy float64) float64 {
	cs.mutex.RLock()
	defer cs.mutex.RUnlock()
	
	// Get predictions for this model, sorted by confidence
	modelPredictions := make([]ConfidencePrediction, 0)
	for _, pred := range cs.predictions {
		if pred.ModelType == modelType && pred.ActualOutcome != "" {
			modelPredictions = append(modelPredictions, pred)
		}
	}
	
	if len(modelPredictions) < 10 {
		return 0.7 // Default threshold if not enough data
	}
	
	// Sort by confidence (descending)
	for i := 0; i < len(modelPredictions)-1; i++ {
		for j := i + 1; j < len(modelPredictions); j++ {
			if modelPredictions[i].Confidence < modelPredictions[j].Confidence {
				modelPredictions[i], modelPredictions[j] = modelPredictions[j], modelPredictions[i]
			}
		}
	}
	
	// Find threshold that achieves target accuracy
	for i := 1; i <= len(modelPredictions); i++ {
		correctCount := 0
		for j := 0; j < i; j++ {
			if modelPredictions[j].WasCorrect {
				correctCount++
			}
		}
		
		accuracy := float64(correctCount) / float64(i)
		if accuracy >= targetAccuracy {
			return modelPredictions[i-1].Confidence
		}
	}
	
	// If target accuracy not achievable, return highest confidence threshold
	return modelPredictions[len(modelPredictions)-1].Confidence
}
