package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// StressTests pushes systems to their limits
func StressTests() {
	fmt.Println("💪 STRESS TESTS")
	fmt.Println("===============")
	
	tempDir := "/tmp/kamaji_stress_tests"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	brainPath := filepath.Join(tempDir, "brain.json")
	
	testResults := make(map[string]bool)
	
	// Stress Test 1: High-Volume Memory Operations
	fmt.Println("\n🧠 Stress Test 1: High-Volume Memory Operations")
	memory := NewMemoryPersistence(brainPath)
	consciousness := NewConsciousnessEngine(brainPath)
	
	startTime := time.Now()
	
	// Simulate 1000 rapid operations
	for i := 0; i < 1000; i++ {
		taskType := fmt.Sprintf("stress_task_%d", i%10)
		memory.RecordTaskCompletion(taskType, i%3 != 0, time.Duration(i)*time.Millisecond, 
			fmt.Sprintf("Stress test %d", i))
		
		if i%50 == 0 {
			memory.RecordMistake(taskType, fmt.Sprintf("Stress error %d", i), "context", "solution", i%10+1)
		}
		
		if i%100 == 0 {
			memory.RecordSuccessStrategy(taskType, fmt.Sprintf("Strategy %d", i), "stress context")
		}
	}
	
	duration := time.Since(startTime)
	stats := memory.GetMemoryStats()
	
	success := duration < 10*time.Second && stats["task_patterns_count"].(int) >= 10
	testResults["memory_volume"] = success
	fmt.Printf("✅ Memory volume test: %v (1000 ops in %v, patterns: %v)\n", 
		success, duration, stats["task_patterns_count"])
	
	// Stress Test 2: Concurrent Operations
	fmt.Println("\n🔄 Stress Test 2: Concurrent Operations")
	
	var wg sync.WaitGroup
	errors := make(chan error, 100)
	
	startTime = time.Now()
	
	// Launch 50 concurrent goroutines
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Each goroutine performs multiple operations
			for j := 0; j < 20; j++ {
				taskType := fmt.Sprintf("concurrent_%d_%d", id, j)
				memory.RecordTaskCompletion(taskType, j%2 == 0, time.Duration(j)*time.Millisecond, 
					fmt.Sprintf("Concurrent test %d-%d", id, j))
			}
		}(i)
	}
	
	wg.Wait()
	close(errors)
	
	duration = time.Since(startTime)
	finalStats := memory.GetMemoryStats()
	
	success = duration < 15*time.Second && len(errors) == 0
	testResults["concurrent_ops"] = success
	fmt.Printf("✅ Concurrent operations: %v (50 goroutines in %v, final patterns: %v)\n", 
		success, duration, finalStats["task_patterns_count"])
	
	// Stress Test 3: Memory Growth and Cleanup
	fmt.Println("\n📈 Stress Test 3: Memory Growth and Cleanup")
	
	personality := NewPersonalityDevelopment(memory, consciousness)
	mistakeLearning := NewMistakeLearning(memory, consciousness)
	
	startTime = time.Now()
	
	// Generate large amounts of data
	for i := 0; i < 500; i++ {
		// Personality experiences
		personality.ProcessTaskSuccess(fmt.Sprintf("growth_task_%d", i%5), 0.8)
		personality.ProcessLearningEvent(fmt.Sprintf("Learning event %d", i), 0.7)
		
		// Mistake learning
		if i%10 == 0 {
			mistakeLearning.LearnFromFailure(fmt.Sprintf("growth_task_%d", i%5), 
				fmt.Sprintf("Task %d", i), fmt.Sprintf("Error %d", i), "Growth context")
		}
	}
	
	duration = time.Since(startTime)
	
	// Check system still responsive
	profile := personality.GetPersonalityProfile()
	insights := mistakeLearning.GetLearningInsights()
	
	success = duration < 20*time.Second && len(profile) > 0 && insights != nil
	testResults["memory_growth"] = success
	fmt.Printf("✅ Memory growth test: %v (500 experiences in %v)\n", success, duration)
	
	// Stress Test 4: AI Orchestration Under Load
	fmt.Println("\n🤖 Stress Test 4: AI Orchestration Under Load")
	
	orchestrator := NewMultiModelOrchestrator(consciousness)
	consensus := NewConsensusSystem(orchestrator, consciousness, memory)
	router := NewModelRouter(consciousness)
	confidenceScoring := NewConfidenceScoring(consciousness, memory)
	
	startTime = time.Now()
	
	// Rapid AI operations
	successfulOps := 0
	for i := 0; i < 100; i++ {
		// Routing
		model, _, conf := router.RouteTask(fmt.Sprintf("stress_type_%d", i%5), 
			fmt.Sprintf("Stress task %d", i))
		
		if conf > 0.5 {
			successfulOps++
		}
		
		// Confidence scoring
		_ = confidenceScoring.RecordPrediction(model, fmt.Sprintf("Prediction %d", i), 
			"Stress test", 0.7+float64(i%3)*0.1)
		
		// Consensus (every 10th operation)
		if i%10 == 0 {
			_, err := consensus.RequestConsensus(fmt.Sprintf("Decision %d", i), 
				"stress_decision", []ModelType{QModel, Claude})
			if err == nil {
				successfulOps++
			}
		}
	}
	
	duration = time.Since(startTime)
	
	success = duration < 30*time.Second && successfulOps > 80
	testResults["orchestration_load"] = success
	fmt.Printf("✅ Orchestration load test: %v (100 ops in %v, successful: %d)\n", 
		success, duration, successfulOps)
	
	// Stress Test 5: Self-Improvement Rapid Cycles
	fmt.Println("\n🚀 Stress Test 5: Self-Improvement Rapid Cycles")
	
	selfImprovement := NewSelfImprovement(consciousness, memory, personality, mistakeLearning)
	
	startTime = time.Now()
	
	// Rapid metric updates and improvements
	for i := 0; i < 200; i++ {
		metricName := []string{"task_success_rate", "confidence_calibration", "routing_accuracy", 
			"learning_efficiency", "processing_speed"}[i%5]
		
		value := 0.5 + float64(i%50)*0.01
		selfImprovement.UpdateMetric(metricName, value)
		
		if i%20 == 0 {
			selfImprovement.AutoImprove()
		}
	}
	
	duration = time.Since(startTime)
	improvementStats := selfImprovement.GetImprovementStats()
	
	success = duration < 25*time.Second && improvementStats["total_improvements"].(int) > 0
	testResults["improvement_cycles"] = success
	fmt.Printf("✅ Improvement cycles: %v (200 updates in %v, improvements: %v)\n", 
		success, duration, improvementStats["total_improvements"])
	
	// Stress Test 6: Wisdom Knowledge Base Scaling
	fmt.Println("\n🧙‍♂️ Stress Test 6: Wisdom Knowledge Base Scaling")
	
	wisdomKB := NewWisdomKnowledgeBase(consciousness, memory, personality, mistakeLearning)
	
	startTime = time.Now()
	
	// Add large amounts of wisdom
	wisdomIDs := make([]string, 0)
	for i := 0; i < 300; i++ {
		_ = wisdomKB.AddWisdom(fmt.Sprintf("Stress Wisdom %d", i), 
			fmt.Sprintf("This is stress test wisdom entry number %d with various insights", i), 
			fmt.Sprintf("category_%d", i%10), "stress_test", 
			[]string{fmt.Sprintf("tag_%d", i%5), "stress", "scaling"})
		wisdomIDs = append(wisdomIDs, wisdomID)
		
		if i%50 == 0 {
			// Search operations
			results := wisdomKB.SearchWisdom(fmt.Sprintf("stress wisdom %d", i%10))
			if len(results) > 0 {
				wisdomKB.UseWisdom(results[0].ID, true)
			}
		}
	}
	
	// Generate insights from large wisdom base
	insights := len(wisdomKB.GenerateInsights())
	
	duration = time.Since(startTime)
	wisdomStats := wisdomKB.GetWisdomStats()
	
	success = duration < 30*time.Second && wisdomStats["total_wisdom_entries"].(int) >= 300
	testResults["wisdom_scaling"] = success
	fmt.Printf("✅ Wisdom scaling: %v (300 entries in %v, insights: %d)\n", 
		success, duration, len(insights))
	
	// Stress Test 7: Predictive Task Generation Load
	fmt.Println("\n🔮 Stress Test 7: Predictive Task Generation Load")
	
	predictiveTasks := NewPredictiveTaskGeneration(consciousness, memory)
	
	startTime = time.Now()
	
	// Rapid pattern analysis and prediction generation
	for i := 0; i < 50; i++ {
		predictiveTasks.AnalyzeTaskPatterns()
		predictions := predictiveTasks.GeneratePredictiveTasks()
		
		// Validate some predictions
		for j, pred := range predictions[:min(3, len(predictions))] {
			predictiveTasks.ValidatePrediction(pred.ID, j%2 == 0, j%3 == 0)
		}
		
		if i%10 == 0 {
			predictiveTasks.OptimizePredictions()
		}
	}
	
	duration = time.Since(startTime)
	_ = predictiveTasks.GetPredictionStats()
	
	success = duration < 20*time.Second && predStats["total_predictions"].(int) > 0
	testResults["prediction_load"] = success
	fmt.Printf("✅ Prediction load: %v (50 cycles in %v, predictions: %v)\n", 
		success, duration, predStats["total_predictions"])
	
	// Stress Test 8: Full System Integration Under Load
	fmt.Println("\n🔗 Stress Test 8: Full System Integration Under Load")
	
	startTime = time.Now()
	
	// Simulate complex real-world scenario under load
	for i := 0; i < 50; i++ {
		// 1. Task execution with mixed outcomes
		taskSuccess := i%3 != 0
		memory.RecordTaskCompletion(fmt.Sprintf("integration_%d", i%5), taskSuccess, 
			time.Duration(i)*10*time.Millisecond, fmt.Sprintf("Integration task %d", i))
		
		// 2. Learning from outcomes
		if !taskSuccess {
			mistakeLearning.LearnFromFailure(fmt.Sprintf("integration_%d", i%5), 
				fmt.Sprintf("Integration task %d", i), fmt.Sprintf("Failure %d", i), "Load test")
			personality.ProcessTaskFailure(fmt.Sprintf("integration_%d", i%5), 0.6)
		} else {
			personality.ProcessTaskSuccess(fmt.Sprintf("integration_%d", i%5), 0.8)
		}
		
		// 3. Route next task
		router.RouteTask("integration", fmt.Sprintf("Next integration task %d", i+1))
		
		// 4. Update metrics
		selfImprovement.UpdateMetric("task_success_rate", 
			float64(i*3)/float64((i+1)*3)) // Simulate changing success rate
		
		// 5. Generate predictions
		if i%10 == 0 {
			predictions := predictiveTasks.GeneratePredictiveTasks()
			if len(predictions) > 0 {
				predictiveTasks.ValidatePrediction(predictions[0].ID, true, true)
			}
		}
		
		// 6. Extract wisdom
		if i%15 == 0 {
			wisdomKB.ExtractWisdomFromExperience()
		}
	}
	
	duration = time.Since(startTime)
	
	// Verify all systems still functional
	finalMemoryStats := memory.GetMemoryStats()
	finalPersonality := personality.GetPersonalityProfile()
	finalImprovement := selfImprovement.GetImprovementStats()
	
	success = duration < 45*time.Second && 
		finalMemoryStats["task_patterns_count"].(int) > 0 &&
		len(finalPersonality) > 0 &&
		finalImprovement != nil
	
	testResults["full_integration_load"] = success
	fmt.Printf("✅ Full integration load: %v (50 complex cycles in %v)\n", success, duration)
	
	// Stress Test Results
	fmt.Println("\n🎯 STRESS TEST RESULTS")
	fmt.Println("======================")
	
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
	fmt.Printf("\n🏆 STRESS TEST RESULT: %d/%d tests passed (%.1f%%)\n", 
		passedTests, totalTests, successRate)
	
	if successRate >= 90 {
		fmt.Println("💪 EXCELLENT! Kamaji systems are highly robust under stress!")
	} else if successRate >= 75 {
		fmt.Println("✅ GOOD! Kamaji systems handle stress well!")
	} else {
		fmt.Println("⚠️  NEEDS OPTIMIZATION! Some systems struggle under load!")
	}
}

func RunStressTests() {
	StressTests()
}
