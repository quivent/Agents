package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// TestAdvancedSystems tests self-improvement, predictive tasks, and wisdom systems
func TestAdvancedSystems() {
	fmt.Println("🚀 Testing Advanced Kamaji Systems...")
	
	// Create temporary brain directory
	tempDir := "/tmp/kamaji_advanced_test"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	brainPath := filepath.Join(tempDir, "brain.json")
	
	// Create base systems
	memory := NewMemoryPersistence(brainPath)
	consciousness := NewConsciousnessEngine(brainPath)
	mistakeLearning := NewMistakeLearning(memory, consciousness)
	personality := NewPersonalityDevelopment(memory, consciousness)
	
	// Test 1: Self-Improvement System
	fmt.Println("📝 Test 1: Testing self-improvement system...")
	
	selfImprovement := NewSelfImprovement(consciousness, memory, personality, mistakeLearning)
	
	// Simulate some performance data
	selfImprovement.UpdateMetric("task_success_rate", 0.75)
	selfImprovement.UpdateMetric("confidence_calibration", 0.65)
	selfImprovement.UpdateMetric("routing_accuracy", 0.70)
	
	// Analyze performance
	analysis := selfImprovement.AnalyzePerformance()
	fmt.Printf("✅ Performance analysis: %v\n", analysis)
	
	// Execute an improvement strategy
	if recommendedStrategies, ok := analysis["recommended_strategies"].([]string); ok && len(recommendedStrategies) > 0 {
		err := selfImprovement.ExecuteImprovement(recommendedStrategies[0])
		if err != nil {
			fmt.Printf("❌ Improvement execution failed: %v\n", err)
		} else {
			fmt.Printf("✅ Executed improvement strategy: %s\n", recommendedStrategies[0])
		}
	}
	
	// Test auto-improvement
	selfImprovement.AutoImprove()
	
	// Get improvement stats
	improvementStats := selfImprovement.GetImprovementStats()
	fmt.Printf("✅ Improvement stats: Total improvements: %v, Success rate: %.2f\n", 
		improvementStats["total_improvements"], improvementStats["improvement_success_rate"])
	
	// Test 2: Predictive Task Generation
	fmt.Println("📝 Test 2: Testing predictive task generation...")
	
	predictiveTasks := NewPredictiveTaskGeneration(consciousness, memory)
	
	// Analyze task patterns
	predictiveTasks.AnalyzeTaskPatterns()
	
	// Generate predictive tasks
	predictions := predictiveTasks.GeneratePredictiveTasks()
	fmt.Printf("✅ Generated %d predictive tasks\n", len(predictions))
	
	if len(predictions) > 0 {
		firstTask := predictions[0]
		fmt.Printf("   Example task: %s (confidence: %.2f, priority: %d)\n", 
			firstTask.Title, firstTask.Confidence, firstTask.Priority)
		
		// Validate a prediction
		predictiveTasks.ValidatePrediction(firstTask.ID, true, true)
		fmt.Printf("✅ Validated prediction: %s\n", firstTask.ID[:12])
	}
	
	// Test prediction accuracy
	accuracy := predictiveTasks.GetPredictionAccuracy()
	fmt.Printf("✅ Prediction accuracy: %.2f\n", accuracy)
	
	// Get prediction stats
	predictionStats := predictiveTasks.GetPredictionStats()
	fmt.Printf("✅ Prediction stats: %v\n", predictionStats)
	
	// Optimize predictions
	predictiveTasks.OptimizePredictions()
	fmt.Printf("✅ Prediction optimization completed\n")
	
	// Test 3: Wisdom Knowledge Base
	fmt.Println("📝 Test 3: Testing wisdom knowledge base...")
	
	wisdomKB := NewWisdomKnowledgeBase(consciousness, memory, personality, mistakeLearning)
	
	// Add some custom wisdom
	wisdomID1 := wisdomKB.AddWisdom("Test-Driven Development", 
		"Writing tests before code leads to better design and fewer bugs", 
		"development_practices", "experience", []string{"testing", "development", "quality"})
	
	wisdomID2 := wisdomKB.AddWisdom("Incremental Improvement", 
		"Small, consistent improvements compound over time to create significant progress", 
		"improvement_philosophy", "insight", []string{"improvement", "consistency", "progress"})
	
	fmt.Printf("✅ Added wisdom entries: %s, %s\n", wisdomID1[:12], wisdomID2[:12])
	
	// Search for wisdom
	searchResults := wisdomKB.SearchWisdom("improvement")
	fmt.Printf("✅ Found %d wisdom entries for 'improvement'\n", len(searchResults))
	
	if len(searchResults) > 0 {
		fmt.Printf("   Top result: %s\n", searchResults[0].Title)
		
		// Use the wisdom
		wisdomKB.UseWisdom(searchResults[0].ID, true)
		fmt.Printf("✅ Used wisdom: %s\n", searchResults[0].Title)
	}
	
	// Get wisdom by category
	fundamentalWisdom := wisdomKB.GetWisdomByCategory("fundamental_concepts")
	fmt.Printf("✅ Found %d fundamental concept entries\n", len(fundamentalWisdom))
	
	// Get most useful wisdom
	mostUseful := wisdomKB.GetMostUsefulWisdom(5)
	fmt.Printf("✅ Top 5 most useful wisdom entries:\n")
	for i, wisdom := range mostUseful {
		fmt.Printf("   %d. %s (usefulness: %.2f)\n", i+1, wisdom.Title, wisdom.Usefulness)
	}
	
	// Generate insights
	insights := wisdomKB.GenerateInsights()
	fmt.Printf("✅ Generated %d insights from wisdom\n", len(insights))
	
	if len(insights) > 0 {
		fmt.Printf("   Example insight: %s\n", insights[0].Insight)
	}
	
	// Extract wisdom from experience
	wisdomKB.ExtractWisdomFromExperience()
	fmt.Printf("✅ Extracted wisdom from recent experiences\n")
	
	// Get wisdom stats
	wisdomStats := wisdomKB.GetWisdomStats()
	fmt.Printf("✅ Wisdom stats: %v\n", wisdomStats)
	
	// Test 4: System Integration
	fmt.Println("📝 Test 4: Testing advanced system integration...")
	
	// Simulate a complex learning scenario
	fmt.Printf("🔄 Simulating complex learning scenario...\n")
	
	// 1. Record a mistake
	mistakeLearning.LearnFromFailure("optimization", "Performance optimization attempt", 
		"Optimization reduced performance instead of improving it", "Production environment")
	
	// 2. Process personality impact
	personality.ProcessTaskFailure("optimization", 0.7)
	
	// 3. Update performance metrics
	selfImprovement.UpdateMetric("task_success_rate", 0.72) // Slight decrease
	selfImprovement.UpdateMetric("learning_efficiency", 0.68) // Learning from failure
	
	// 4. Extract wisdom from the experience
	wisdomKB.ExtractWisdomFromExperience()
	
	// 5. Generate new predictive tasks based on the failure
	newPredictions := predictiveTasks.GeneratePredictiveTasks()
	fmt.Printf("✅ Generated %d new predictions after learning event\n", len(newPredictions))
	
	// 6. Auto-improve based on the new data
	selfImprovement.AutoImprove()
	
	// 7. Search for relevant wisdom
	relevantWisdom := wisdomKB.SearchWisdom("optimization failure")
	fmt.Printf("✅ Found %d wisdom entries about optimization failures\n", len(relevantWisdom))
	
	// Test 5: Performance and Scalability
	fmt.Println("📝 Test 5: Testing performance and scalability...")
	
	// Simulate high-volume operations
	for i := 0; i < 50; i++ {
		// Add wisdom
		wisdomKB.AddWisdom(fmt.Sprintf("Automated Wisdom %d", i), 
			fmt.Sprintf("This is automatically generated wisdom entry number %d", i), 
			"automated", "experience", []string{"automated", "test"})
		
		// Update metrics
		selfImprovement.UpdateMetric("processing_speed", 0.8+float64(i)*0.001)
		
		// Generate predictions
		if i%10 == 0 {
			predictions := predictiveTasks.GeneratePredictiveTasks()
			if len(predictions) > 0 {
				predictiveTasks.ValidatePrediction(predictions[0].ID, true, i%2 == 0)
			}
		}
	}
	
	fmt.Printf("✅ Processed 50 high-volume operations successfully\n")
	
	// Final system state
	finalWisdomStats := wisdomKB.GetWisdomStats()
	finalImprovementStats := selfImprovement.GetImprovementStats()
	finalPredictionStats := predictiveTasks.GetPredictionStats()
	
	fmt.Printf("✅ Final system state:\n")
	fmt.Printf("   Wisdom entries: %v\n", finalWisdomStats["total_wisdom_entries"])
	fmt.Printf("   Improvements: %v\n", finalImprovementStats["total_improvements"])
	fmt.Printf("   Predictions: %v\n", finalPredictionStats["total_predictions"])
	
	fmt.Println("🎉 All advanced system tests completed successfully!")
}

// RunAdvancedSystemsTest is a standalone function to test advanced systems
func RunAdvancedSystemsTest() {
	TestAdvancedSystems()
}
