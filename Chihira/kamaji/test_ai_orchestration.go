package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// TestAIOrchestration tests all AI orchestration features
func TestAIOrchestration() {
	fmt.Println("🤖 Testing AI Orchestration Systems...")
	
	// Create temporary brain directory
	tempDir := "/tmp/kamaji_orchestration_test"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	brainPath := filepath.Join(tempDir, "brain.json")
	
	// Create systems
	memory := NewMemoryPersistence(brainPath)
	consciousness := NewConsciousnessEngine(brainPath)
	
	// Test 1: Multi-Model Consensus System
	fmt.Println("📝 Test 1: Testing multi-model consensus...")
	
	// Create a minimal orchestrator for testing
	orchestrator := &MultiModelOrchestrator{
		consciousness: consciousness,
	}
	
	consensus := NewConsensusSystem(orchestrator, consciousness, memory)
	
	// Request consensus on a decision
	models := []ModelType{QModel, Claude}
	result, err := consensus.RequestConsensus("Should we implement AWS Lambda for this task?", 
		"aws_task", models)
	
	if err != nil {
		fmt.Printf("❌ Consensus request failed: %v\n", err)
	} else {
		fmt.Printf("✅ Consensus reached: %s (confidence: %.2f, type: %s)\n", 
			result.FinalDecision, result.Confidence, result.ConsensusType)
	}
	
	// Test consensus stats
	stats := consensus.GetConsensusStats()
	fmt.Printf("✅ Consensus stats: %v\n", stats)
	
	// Test 2: Model Routing System
	fmt.Println("📝 Test 2: Testing specialized model routing...")
	
	router := NewModelRouter(consciousness)
	
	// Test routing for different task types
	testTasks := []struct {
		taskType string
		query    string
	}{
		{"aws_task", "Deploy a serverless application on AWS"},
		{"analysis", "Analyze the pros and cons of microservices"},
		{"code_generation", "Generate a Go function for data processing"},
		{"creative", "Design a user-friendly interface"},
	}
	
	for _, task := range testTasks {
		model, reasoning, confidence := router.RouteTask(task.taskType, task.query)
		fmt.Printf("✅ Routed '%s' to %s (%.2f): %s\n", 
			task.taskType, model, confidence, reasoning)
		
		// Simulate task success/failure
		success := confidence > 0.7
		router.UpdateRoutingSuccess(task.taskType, task.query, success)
	}
	
	// Test routing stats
	routingStats := router.GetRoutingStats()
	fmt.Printf("✅ Routing stats: %v\n", routingStats)
	
	// Test best model for specific capabilities
	bestForAWS, score := router.GetBestModelFor("aws_services")
	fmt.Printf("✅ Best model for AWS services: %s (score: %.2f)\n", bestForAWS, score)
	
	// Test 3: AI Debate System
	fmt.Println("📝 Test 3: Testing AI debate system...")
	
	debate := NewAIDebateSystem(consciousness, memory)
	
	// Start a debate
	debateID, err := debate.StartDebate("Should we prioritize security or performance?", 
		[]ModelType{QModel, Claude}, 3)
	
	if err != nil {
		fmt.Printf("❌ Failed to start debate: %v\n", err)
	} else {
		fmt.Printf("✅ Started debate: %s\n", debateID)
		
		// Conduct debate rounds
		for round := 1; round <= 3; round++ {
			err := debate.ConductDebateRound(debateID)
			if err != nil {
				fmt.Printf("❌ Debate round %d failed: %v\n", round, err)
				break
			}
			fmt.Printf("✅ Completed debate round %d\n", round)
		}
	}
	
	// Test debate stats
	debateStats := debate.GetDebateStats()
	fmt.Printf("✅ Debate stats: %v\n", debateStats)
	
	// Test 4: Confidence Scoring System
	fmt.Println("📝 Test 4: Testing confidence scoring...")
	
	confidenceScoring := NewConfidenceScoring(consciousness, memory)
	
	// Record some predictions
	testPredictions := []struct {
		model      ModelType
		prediction string
		confidence float64
		correct    bool
	}{
		{QModel, "AWS Lambda will be cost-effective", 0.85, true},
		{QModel, "DynamoDB will handle the load", 0.90, true},
		{Claude, "Users will prefer the new interface", 0.75, false},
		{Claude, "The analysis reveals three key insights", 0.80, true},
		{Future, "Performance will improve by 20%", 0.70, false},
	}
	
	for _, pred := range testPredictions {
		predID := confidenceScoring.RecordPrediction(pred.model, pred.prediction, 
			"Based on analysis", pred.confidence)
		
		// Update with actual outcome
		outcome := "failed"
		if pred.correct {
			outcome = "succeeded"
		}
		confidenceScoring.UpdatePredictionOutcome(pred.model, pred.prediction, 
			outcome, pred.correct)
		
		fmt.Printf("✅ Recorded prediction %s: %v\n", predID[:8], pred.correct)
	}
	
	// Test calibrated confidence
	for _, model := range []ModelType{QModel, Claude, Future} {
		rawConf := 0.8
		calibrated := confidenceScoring.GetCalibratedConfidence(model, rawConf)
		reliability := confidenceScoring.GetModelReliability(model)
		
		fmt.Printf("✅ %s: Raw %.2f → Calibrated %.2f (reliability: %.2f)\n", 
			model, rawConf, calibrated, reliability)
	}
	
	// Test confidence analysis
	analysis := confidenceScoring.AnalyzeConfidencePatterns()
	fmt.Printf("✅ Confidence analysis: %v\n", analysis)
	
	// Test 5: Integration Test
	fmt.Println("📝 Test 5: Testing system integration...")
	
	// Simulate a complex decision process
	task := "Implement a real-time analytics system"
	
	// 1. Route to best model
	bestModel, _, routingConf := router.RouteTask("analysis", task)
	fmt.Printf("✅ Routed complex task to %s (confidence: %.2f)\n", bestModel, routingConf)
	
	// 2. Get consensus on approach
	consensusResult, _ := consensus.RequestConsensus(
		"Should we use streaming or batch processing?", "analysis", 
		[]ModelType{QModel, Claude})
	
	if consensusResult != nil {
		fmt.Printf("✅ Consensus on approach: %s\n", consensusResult.FinalDecision)
	}
	
	// 3. Record prediction about success
	predictionID := confidenceScoring.RecordPrediction(bestModel, 
		"Implementation will succeed", "Based on routing and consensus", 0.82)
	
	fmt.Printf("✅ Recorded success prediction: %s\n", predictionID[:8])
	
	// 4. Start debate on technical details
	debateID2, _ := debate.StartDebate("Real-time vs batch processing trade-offs", 
		[]ModelType{QModel, Claude}, 2)
	
	if debateID2 != "" {
		debate.ConductDebateRound(debateID2)
		debate.ConductDebateRound(debateID2)
		fmt.Printf("✅ Completed technical debate: %s\n", debateID2)
	}
	
	fmt.Println("🎉 All AI orchestration tests completed!")
}

// RunAIOrchestrationTest is a standalone function to test AI orchestration
func RunAIOrchestrationTest() {
	TestAIOrchestration()
}
