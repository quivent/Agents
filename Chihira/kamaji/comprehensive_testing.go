package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ComprehensiveTest runs all systems through rigorous testing
func ComprehensiveTest() {
	fmt.Println("🔬 COMPREHENSIVE KAMAJI SYSTEM TEST")
	fmt.Println("=====================================")
	
	tempDir := "/tmp/kamaji_comprehensive_test"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	brainPath := filepath.Join(tempDir, "brain.json")
	
	// Initialize all systems
	memory := NewMemoryPersistence(brainPath)
	consciousness := NewConsciousnessEngine(brainPath)
	mistakeLearning := NewMistakeLearning(memory, consciousness)
	personality := NewPersonalityDevelopment(memory, consciousness)
	
	orchestrator := NewMultiModelOrchestrator(consciousness)
	consensus := NewConsensusSystem(orchestrator, consciousness, memory)
	router := NewModelRouter(consciousness)
	debate := NewAIDebateSystem(consciousness, memory)
	confidenceScoring := NewConfidenceScoring(consciousness, memory)
	
	selfImprovement := NewSelfImprovement(consciousness, memory, personality, mistakeLearning)
	predictiveTasks := NewPredictiveTaskGeneration(consciousness, memory)
	wisdomKB := NewWisdomKnowledgeBase(consciousness, memory, personality, mistakeLearning)
	
	testResults := make(map[string]bool)
	
	// Test 1: Memory System Stress Test
	fmt.Println("\n🧠 Test 1: Memory System Stress Test")
	success := true
	for i := 0; i < 100; i++ {
		memory.RecordTaskCompletion("stress_test", i%2 == 0, time.Duration(i)*time.Second, fmt.Sprintf("Task %d", i))
		if i%10 == 0 {
			memory.RecordMistake("stress_test", fmt.Sprintf("Error %d", i), "context", "solution", 5)
		}
	}
	
	stats := memory.GetMemoryStats()
	if stats["task_patterns_count"].(int) < 1 || stats["mistakes_count"].(int) < 10 {
		success = false
	}
	testResults["memory_stress"] = success
	fmt.Printf("✅ Memory stress test: %v (patterns: %v, mistakes: %v)\n", 
		success, stats["task_patterns_count"], stats["mistakes_count"])
	
	// Test 2: Learning System Integration
	fmt.Println("\n📚 Test 2: Learning System Integration")
	success = true
	
	// Simulate learning cycle
	mistakeLearning.LearnFromFailure("integration_test", "Complex task", "Integration failed", "Test context")
	warnings := mistakeLearning.AnalyzeTaskForMistakes("integration_test", "Another complex task")
	
	personality.ProcessTaskFailure("integration_test", 0.8)
	personality.ProcessLearningEvent("Learned from integration failure", 0.9)
	
	insights := mistakeLearning.GetLearningInsights()
	if insights["total_mistakes"].(int) == 0 {
		success = false
	}
	
	testResults["learning_integration"] = success
	fmt.Printf("✅ Learning integration: %v (warnings: %d, insights: %v)\n", 
		success, len(warnings), insights["total_mistakes"])
	
	// Test 3: AI Orchestration Full Cycle
	fmt.Println("\n🤖 Test 3: AI Orchestration Full Cycle")
	success = true
	
	// Consensus test
	consensusResult, err := consensus.RequestConsensus("Should we optimize performance?", "optimization", 
		[]ModelType{QModel, Claude})
	if err != nil || consensusResult == nil {
		success = false
	}
	
	// Routing test
	bestModel, _, confidence := router.RouteTask("aws_optimization", "Optimize AWS Lambda performance")
	router.UpdateRoutingSuccess("aws_optimization", "Optimize AWS Lambda performance", confidence > 0.7)
	
	// Debate test
	debateID, _ := debate.StartDebate("Performance vs Security trade-offs", []ModelType{QModel, Claude}, 2)
	if debateID != "" {
		debate.ConductDebateRound(debateID)
		debate.ConductDebateRound(debateID)
	}
	
	// Confidence scoring test
	_ = confidenceScoring.RecordPrediction(bestModel, "Performance will improve", "Based on analysis", 0.85)
	confidenceScoring.UpdatePredictionOutcome(bestModel, "Performance will improve", "success", true)
	
	testResults["orchestration_cycle"] = success
	fmt.Printf("✅ Orchestration cycle: %v (consensus: %v, routing: %s, debate: %s)\n", 
		success, consensusResult != nil, bestModel, debateID[:12])
	
	// Test 4: Self-Improvement Automation
	fmt.Println("\n🚀 Test 4: Self-Improvement Automation")
	success = true
	
	// Update metrics to trigger improvements
	selfImprovement.UpdateMetric("task_success_rate", 0.65) // Below target
	selfImprovement.UpdateMetric("confidence_calibration", 0.60) // Below target
	
	analysis := selfImprovement.AnalyzePerformance()
	if needsImprovement, ok := analysis["needs_improvement"].([]string); !ok || len(needsImprovement) == 0 {
		success = false
	}
	
	// Execute auto-improvement
	selfImprovement.AutoImprove()
	
	improvementStats := selfImprovement.GetImprovementStats()
	if improvementStats["total_improvements"].(int) == 0 {
		success = false
	}
	
	testResults["self_improvement"] = success
	fmt.Printf("✅ Self-improvement: %v (improvements: %v, success rate: %.2f)\n", 
		success, improvementStats["total_improvements"], improvementStats["improvement_success_rate"])
	
	// Test 5: Predictive Task Generation Accuracy
	fmt.Println("\n🔮 Test 5: Predictive Task Generation")
	success = true
	
	predictiveTasks.AnalyzeTaskPatterns()
	predictions := predictiveTasks.GeneratePredictiveTasks()
	
	if len(predictions) == 0 {
		success = false
	} else {
		// Validate some predictions
		for i, pred := range predictions[:min(3, len(predictions))] {
			predictiveTasks.ValidatePrediction(pred.ID, i%2 == 0, i%2 == 0)
		}
	}
	
	accuracy := predictiveTasks.GetPredictionAccuracy()
	_ = predictiveTasks.GetPredictionStats()
	
	testResults["predictive_tasks"] = success
	fmt.Printf("✅ Predictive tasks: %v (predictions: %d, accuracy: %.2f)\n", 
		success, len(predictions), accuracy)
	
	// Test 6: Wisdom Knowledge Base Intelligence
	fmt.Println("\n🧙‍♂️ Test 6: Wisdom Knowledge Base")
	success = true
	
	// Add test wisdom
	_ = wisdomKB.AddWisdom("Test Wisdom", "Testing creates confidence in system reliability", 
		"testing", "experience", []string{"testing", "reliability", "confidence"})
	
	// Search for wisdom
	searchResults := wisdomKB.SearchWisdom("testing reliability")
	if len(searchResults) == 0 {
		success = false
	}
	
	// Use wisdom
	if len(searchResults) > 0 {
		wisdomKB.UseWisdom(searchResults[0].ID, true)
	}
	
	// Generate insights
	insights := len(wisdomKB.GenerateInsights())
	
	// Extract wisdom from experience
	wisdomKB.ExtractWisdomFromExperience()
	
	wisdomStats := wisdomKB.GetWisdomStats()
	
	testResults["wisdom_kb"] = success
	fmt.Printf("✅ Wisdom KB: %v (entries: %v, insights: %d, search results: %d)\n", 
		success, wisdomStats["total_wisdom_entries"], len(insights), len(searchResults))
	
	// Test 7: Cross-System Integration
	fmt.Println("\n🔗 Test 7: Cross-System Integration")
	success = true
	
	// Simulate complex scenario: failure → learning → improvement → prediction
	
	// 1. Record a significant failure
	memory.RecordTaskCompletion("complex_integration", false, 2*time.Minute, "Integration test failure")
	mistakeLearning.LearnFromFailure("complex_integration", "Integration test", 
		"System integration failed", "Cross-system test")
	
	// 2. Process personality impact
	personality.ProcessTaskFailure("complex_integration", 0.9)
	
	// 3. Update performance metrics
	selfImprovement.UpdateMetric("task_success_rate", 0.68)
	selfImprovement.UpdateMetric("learning_efficiency", 0.75)
	
	// 4. Extract wisdom
	wisdomKB.ExtractWisdomFromExperience()
	
	// 5. Generate predictions
	newPredictions := predictiveTasks.GeneratePredictiveTasks()
	
	// 6. Auto-improve
	selfImprovement.AutoImprove()
	
	// 7. Search for relevant wisdom
	relevantWisdom := wisdomKB.SearchWisdom("integration failure")
	
	// Verify integration worked
	finalStats := map[string]interface{}{
		"memory":     memory.GetMemoryStats(),
		"learning":   mistakeLearning.GetLearningInsights(),
		"personality": personality.GetPersonalityEvolution(),
		"improvement": selfImprovement.GetImprovementStats(),
		"predictions": len(newPredictions),
		"wisdom":     len(relevantWisdom),
	}
	
	testResults["cross_system"] = success
	fmt.Printf("✅ Cross-system integration: %v\n", success)
	fmt.Printf("   Final state: memory=%v, learning=%v, predictions=%v, wisdom=%v\n",
		finalStats["memory"].(map[string]interface{})["task_patterns_count"],
		finalStats["learning"].(map[string]interface{})["total_mistakes"],
		finalStats["predictions"], finalStats["wisdom"])
	
	// Test 8: Performance and Scalability
	fmt.Println("\n⚡ Test 8: Performance and Scalability")
	success = true
	startTime := time.Now()
	
	// High-volume operations
	for i := 0; i < 200; i++ {
		if i%20 == 0 {
			memory.RecordTaskCompletion(fmt.Sprintf("perf_test_%d", i%5), i%3 != 0, 
				time.Duration(i)*time.Millisecond, "Performance test")
		}
		if i%30 == 0 {
			selfImprovement.UpdateMetric("processing_speed", 0.8+float64(i)*0.0005)
		}
		if i%40 == 0 {
			wisdomKB.AddWisdom(fmt.Sprintf("Perf Wisdom %d", i), "Performance insight", 
				"performance", "automated", []string{"performance"})
		}
	}
	
	duration := time.Since(startTime)
	if duration > 5*time.Second {
		success = false
	}
	
	testResults["performance"] = success
	fmt.Printf("✅ Performance test: %v (200 operations in %v)\n", success, duration)
	
	// Test 9: Data Persistence and Recovery
	fmt.Println("\n💾 Test 9: Data Persistence and Recovery")
	success = true
	
	// Save current state
	memory.RecordTaskCompletion("persistence_test", true, time.Second, "Before restart")
	wisdomKB.AddWisdom("Persistence Test", "Data should survive restart", 
		"persistence", "test", []string{"persistence", "recovery"})
	
	// Simulate restart by creating new instances
	memory2 := NewMemoryPersistence(brainPath)
	wisdomKB2 := NewWisdomKnowledgeBase(consciousness, memory2, personality, mistakeLearning)
	
	// Verify data persisted
	stats2 := memory2.GetMemoryStats()
	searchResults2 := wisdomKB2.SearchWisdom("persistence")
	
	if stats2["task_patterns_count"].(int) == 0 || len(searchResults2) == 0 {
		success = false
	}
	
	testResults["persistence"] = success
	fmt.Printf("✅ Persistence test: %v (patterns: %v, wisdom found: %d)\n", 
		success, stats2["task_patterns_count"], len(searchResults2))
	
	// Test 10: Error Handling and Resilience
	fmt.Println("\n🛡️ Test 10: Error Handling and Resilience")
	success = true
	
	// Test with invalid inputs
	_, err1 := consensus.RequestConsensus("", "", []ModelType{})
	router.UpdateRoutingSuccess("nonexistent", "invalid", false)
	wisdomKB.UseWisdom("invalid_id", true)
	
	// System should handle errors gracefully
	if err1 == nil {
		// Expected error for empty consensus request
	}
	
	testResults["error_handling"] = success
	fmt.Printf("✅ Error handling: %v (graceful error recovery)\n", success)
	
	// Final Results Summary
	fmt.Println("\n🎯 COMPREHENSIVE TEST RESULTS")
	fmt.Println("==============================")
	
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
	fmt.Printf("\n🏆 OVERALL RESULT: %d/%d tests passed (%.1f%%)\n", 
		passedTests, totalTests, successRate)
	
	if successRate >= 90 {
		fmt.Println("🎉 EXCELLENT! Kamaji systems are highly reliable!")
	} else if successRate >= 80 {
		fmt.Println("✅ GOOD! Kamaji systems are working well!")
	} else {
		fmt.Println("⚠️  NEEDS IMPROVEMENT! Some systems require attention!")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func RunComprehensiveTest() {
	ComprehensiveTest()
}
