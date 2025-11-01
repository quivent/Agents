package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// UnitTests runs focused unit tests for each component
func UnitTests() {
	fmt.Println("🧪 UNIT TESTS")
	fmt.Println("=============")
	
	tempDir := "/tmp/kamaji_unit_tests"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	brainPath := filepath.Join(tempDir, "brain.json")
	
	testResults := make(map[string]bool)
	
	// Unit Test 1: Memory Persistence
	fmt.Println("\n📝 Unit Test 1: Memory Persistence")
	memory := NewMemoryPersistence(brainPath)
	
	// Test task recording
	memory.RecordTaskCompletion("unit_test", true, 30*time.Second, "Test context")
	pattern := memory.GetTaskPattern("unit_test")
	
	success := pattern != nil && pattern.CompletionCount == 1 && pattern.SuccessRate == 1.0
	testResults["memory_basic"] = success
	fmt.Printf("✅ Memory basic operations: %v\n", success)
	
	// Test mistake recording
	memory.RecordMistake("unit_test", "Test error", "Test context", "Test solution", 5)
	mistakes := memory.GetRecentMistakes(5)
	
	success = len(mistakes) == 1 && mistakes[0].Severity == 5
	testResults["memory_mistakes"] = success
	fmt.Printf("✅ Memory mistake recording: %v\n", success)
	
	// Unit Test 2: Personality Development
	fmt.Println("\n📝 Unit Test 2: Personality Development")
	consciousness := NewConsciousnessEngine(brainPath)
	personality := NewPersonalityDevelopment(memory, consciousness)
	
	initialProfile := personality.GetPersonalityProfile()
	initialConfidence := initialProfile["confidence"]
	
	personality.ProcessTaskSuccess("unit_test", 0.8)
	newProfile := personality.GetPersonalityProfile()
	newConfidence := newProfile["confidence"]
	
	success = newConfidence > initialConfidence
	testResults["personality_growth"] = success
	fmt.Printf("✅ Personality growth: %v (%.3f → %.3f)\n", success, initialConfidence, newConfidence)
	
	// Unit Test 3: Mistake Learning
	fmt.Println("\n📝 Unit Test 3: Mistake Learning")
	mistakeLearning := NewMistakeLearning(memory, consciousness)
	
	mistakeLearning.LearnFromFailure("unit_test", "Unit test task", "Test failure", "Unit test context")
	warnings := mistakeLearning.AnalyzeTaskForMistakes("unit_test", "Similar unit test task")
	
	success = len(warnings) > 0
	testResults["mistake_learning"] = success
	fmt.Printf("✅ Mistake learning: %v (warnings: %d)\n", success, len(warnings))
	
	// Unit Test 4: Model Routing
	fmt.Println("\n📝 Unit Test 4: Model Routing")
	router := NewModelRouter(consciousness)
	
	model1, _, conf1 := router.RouteTask("aws_task", "Deploy AWS Lambda function")
	model2, _, conf2 := router.RouteTask("analysis", "Analyze user behavior patterns")
	
	success = model1 == QModel && model2 == Claude && conf1 > 0.5 && conf2 > 0.5
	testResults["model_routing"] = success
	fmt.Printf("✅ Model routing: %v (%s→%s, %s→%s)\n", success, "aws", model1, "analysis", model2)
	
	// Unit Test 5: Confidence Scoring
	fmt.Println("\n📝 Unit Test 5: Confidence Scoring")
	confidenceScoring := NewConfidenceScoring(consciousness, memory)
	
	_ = confidenceScoring.RecordPrediction(QModel, "Test prediction", "Test reasoning", 0.8)
	confidenceScoring.UpdatePredictionOutcome(QModel, "Test prediction", "success", true)
	
	metrics := confidenceScoring.GetConfidenceMetrics(QModel)
	success = metrics != nil && metrics.TotalPredictions == 1 && metrics.Accuracy == 1.0
	testResults["confidence_scoring"] = success
	fmt.Printf("✅ Confidence scoring: %v (accuracy: %.2f)\n", success, metrics.Accuracy)
	
	// Unit Test 6: Self-Improvement
	fmt.Println("\n📝 Unit Test 6: Self-Improvement")
	selfImprovement := NewSelfImprovement(consciousness, memory, personality, mistakeLearning)
	
	_ = selfImprovement.metrics["task_success_rate"].CurrentValue
	selfImprovement.UpdateMetric("task_success_rate", 0.6) // Below target
	
	analysis := selfImprovement.AnalyzePerformance()
	needsImprovement := analysis["needs_improvement"].([]string)
	
	success = len(needsImprovement) > 0
	testResults["self_improvement"] = success
	fmt.Printf("✅ Self-improvement analysis: %v (needs improvement: %d metrics)\n", success, len(needsImprovement))
	
	// Unit Test 7: Predictive Tasks
	fmt.Println("\n📝 Unit Test 7: Predictive Tasks")
	predictiveTasks := NewPredictiveTaskGeneration(consciousness, memory)
	
	predictions := predictiveTasks.GeneratePredictiveTasks()
	success = len(predictions) > 0
	
	if len(predictions) > 0 {
		firstPred := predictions[0]
		success = success && firstPred.Confidence > 0 && firstPred.Priority > 0
	}
	
	testResults["predictive_tasks"] = success
	fmt.Printf("✅ Predictive task generation: %v (generated: %d tasks)\n", success, len(predictions))
	
	// Unit Test 8: Wisdom Knowledge Base
	fmt.Println("\n📝 Unit Test 8: Wisdom Knowledge Base")
	wisdomKB := NewWisdomKnowledgeBase(consciousness, memory, personality, mistakeLearning)
	
	_ = wisdomKB.AddWisdom("Unit Test Wisdom", "Unit tests ensure code quality", 
		"testing", "experience", []string{"testing", "quality"})
	
	searchResults := wisdomKB.SearchWisdom("testing quality")
	success = len(searchResults) > 0 && searchResults[0].ID == wisdomID
	
	testResults["wisdom_kb"] = success
	fmt.Printf("✅ Wisdom knowledge base: %v (search results: %d)\n", success, len(searchResults))
	
	// Unit Test 9: AI Consensus
	fmt.Println("\n📝 Unit Test 9: AI Consensus")
	orchestrator := NewMultiModelOrchestrator(consciousness)
	consensus := NewConsensusSystem(orchestrator, consciousness, memory)
	
	result, err := consensus.RequestConsensus("Should we proceed with testing?", 
		"testing", []ModelType{QModel, Claude})
	
	success = err == nil && result != nil && result.Confidence > 0
	testResults["ai_consensus"] = success
	fmt.Printf("✅ AI consensus: %v (confidence: %.2f)\n", success, result.Confidence)
	
	// Unit Test 10: AI Debate
	fmt.Println("\n📝 Unit Test 10: AI Debate")
	debate := NewAIDebateSystem(consciousness, memory)
	
	debateID, err := debate.StartDebate("Testing vs Speed", []ModelType{QModel, Claude}, 2)
	success = err == nil && debateID != ""
	
	if success {
		err1 := debate.ConductDebateRound(debateID)
		err2 := debate.ConductDebateRound(debateID)
		success = err1 == nil && err2 == nil
	}
	
	testResults["ai_debate"] = success
	fmt.Printf("✅ AI debate: %v (debate ID: %s)\n", success, debateID[:min(12, len(debateID))])
	
	// Unit Test Results
	fmt.Println("\n🎯 UNIT TEST RESULTS")
	fmt.Println("====================")
	
	totalTests := len(testResults)
	passedTests := 0
	
	for testName, passed := range testResults {
		status := "❌ FAIL"
		if passed {
			status = "✅ PASS"
			passedTests++
		}
		fmt.Printf("%s %s\n", status, testName)
	}
	
	successRate := float64(passedTests) / float64(totalTests) * 100
	fmt.Printf("\n🏆 UNIT TEST RESULT: %d/%d tests passed (%.1f%%)\n", 
		passedTests, totalTests, successRate)
}

func RunUnitTests() {
	UnitTests()
}
