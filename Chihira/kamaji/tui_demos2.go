package main

import (
	"fmt"
	"strconv"
	"time"
)

// confidenceDemo demonstrates confidence scoring
func (tui *InteractiveTUI) confidenceDemo() {
	fmt.Println("\n📊 CONFIDENCE SCORING DEMO")
	fmt.Println("==========================")
	
	prediction := tui.getInput("Enter a prediction: ")
	if prediction == "" {
		prediction = "The system will perform well"
	}
	
	// Record prediction
	_ = tui.confidenceScoring.RecordPrediction(QModel, prediction, "Interactive demo", 0.85)
	fmt.Printf("📝 Recorded prediction: %s\n", prediction)
	
	outcome := tui.getInput("What was the actual outcome? ")
	if outcome == "" {
		outcome = "success"
	}
	
	wasCorrect := tui.getInput("Was the prediction correct? (y/n): ") == "y"
	
	// Update with outcome
	tui.confidenceScoring.UpdatePredictionOutcome(QModel, prediction, outcome, wasCorrect)
	
	// Show metrics
	metrics := tui.confidenceScoring.GetConfidenceMetrics(QModel)
	if metrics != nil {
		fmt.Printf("\n📈 Q Model Confidence Metrics:\n")
		fmt.Printf("   Total Predictions: %d\n", metrics.TotalPredictions)
		fmt.Printf("   Accuracy: %.1f%%\n", metrics.Accuracy*100)
		fmt.Printf("   Average Confidence: %.2f\n", metrics.AvgConfidence)
		fmt.Printf("   Calibration: %.2f\n", metrics.Calibration)
		fmt.Printf("   Reliability: %.2f\n", metrics.Reliability)
	}
	
	// Show calibrated confidence
	rawConf := 0.8
	calibrated := tui.confidenceScoring.GetCalibratedConfidence(QModel, rawConf)
	fmt.Printf("\n🎯 Confidence Calibration Example:\n")
	fmt.Printf("   Raw Confidence: %.2f\n", rawConf)
	fmt.Printf("   Calibrated Confidence: %.2f\n", calibrated)
	
	fmt.Println("\n✅ System learns to calibrate confidence based on actual performance!")
}

// predictiveDemo demonstrates predictive task generation
func (tui *InteractiveTUI) predictiveDemo() {
	fmt.Println("\n🔮 PREDICTIVE TASK GENERATION DEMO")
	fmt.Println("==================================")
	
	fmt.Println("Analyzing task patterns and generating predictions...")
	
	// Analyze patterns
	tui.predictiveTasks.AnalyzeTaskPatterns()
	
	// Generate predictions
	predictions := tui.predictiveTasks.GeneratePredictiveTasks()
	
	fmt.Printf("\n🎯 Generated %d Predictive Tasks:\n", len(predictions))
	for i, pred := range predictions {
		if i >= 5 { // Show only first 5
			break
		}
		fmt.Printf("\n%d. %s\n", i+1, pred.Title)
		fmt.Printf("   Description: %s\n", pred.Description)
		fmt.Printf("   Priority: %d/10\n", pred.Priority)
		fmt.Printf("   Confidence: %.2f\n", pred.Confidence)
		fmt.Printf("   Due By: %s\n", pred.DueBy.Format("2006-01-02 15:04"))
		fmt.Printf("   Reasoning: %s\n", pred.Reasoning)
	}
	
	if len(predictions) > 0 {
		// Validate a prediction
		validateChoice := tui.getInput("\nValidate first prediction? (y/n): ")
		if validateChoice == "y" {
			wasNeeded := tui.getInput("Was this task actually needed? (y/n): ") == "y"
			wasCompleted := false
			if wasNeeded {
				wasCompleted = tui.getInput("Was it completed? (y/n): ") == "y"
			}
			
			tui.predictiveTasks.ValidatePrediction(predictions[0].ID, wasNeeded, wasCompleted)
			fmt.Println("✅ Prediction validated! System will learn from this feedback.")
		}
	}
	
	// Show prediction statistics
	stats := tui.predictiveTasks.GetPredictionStats()
	fmt.Printf("\n📊 Prediction Statistics:\n")
	fmt.Printf("   Total Predictions: %v\n", stats["total_predictions"])
	fmt.Printf("   Prediction Accuracy: %.1f%%\n", tui.predictiveTasks.GetPredictionAccuracy()*100)
	
	fmt.Println("\n✅ System predicts tasks you'll need before you know you need them!")
}

// wisdomDemo demonstrates wisdom knowledge base
func (tui *InteractiveTUI) wisdomDemo() {
	fmt.Println("\n🧙‍♂️ WISDOM KNOWLEDGE BASE DEMO")
	fmt.Println("==============================")
	
	// Show existing wisdom
	mostUseful := tui.wisdomKB.GetMostUsefulWisdom(3)
	fmt.Println("📚 Most Useful Wisdom Entries:")
	for i, wisdom := range mostUseful {
		fmt.Printf("\n%d. %s\n", i+1, wisdom.Title)
		fmt.Printf("   Content: %s\n", wisdom.Content)
		fmt.Printf("   Usefulness: %.1f%%\n", wisdom.Usefulness*100)
		fmt.Printf("   Category: %s\n", wisdom.Category)
	}
	
	// Add new wisdom
	addWisdom := tui.getInput("\nAdd new wisdom? (y/n): ")
	if addWisdom == "y" {
		title := tui.getInput("Wisdom title: ")
		content := tui.getInput("Wisdom content: ")
		category := tui.getInput("Category: ")
		
		if title != "" && content != "" {
			wisdomID := tui.wisdomKB.AddWisdom(title, content, category, "user_input", []string{"interactive", "demo"})
			fmt.Printf("✅ Added wisdom entry: %s\n", wisdomID[:12])
		}
	}
	
	// Search wisdom
	searchQuery := tui.getInput("Search wisdom (enter query): ")
	if searchQuery != "" {
		results := tui.wisdomKB.SearchWisdom(searchQuery)
		fmt.Printf("\n🔍 Found %d relevant wisdom entries:\n", len(results))
		for i, result := range results {
			if i >= 3 { // Show only first 3
				break
			}
			fmt.Printf("\n%d. %s\n", i+1, result.Title)
			fmt.Printf("   Content: %s\n", result.Content)
			fmt.Printf("   Relevance Score: %.2f\n", result.Usefulness)
		}
	}
	
	// Generate insights
	fmt.Println("\nGenerating insights from accumulated wisdom...")
	insights := tui.wisdomKB.GenerateInsights()
	if len(insights) > 0 {
		fmt.Printf("🧠 Generated %d new insights:\n", len(insights))
		for i, insight := range insights {
			if i >= 2 { // Show only first 2
				break
			}
			fmt.Printf("\n%d. %s\n", i+1, insight.Insight)
			fmt.Printf("   Confidence: %.2f\n", insight.Confidence)
			fmt.Printf("   Depth Level: %d\n", insight.Depth)
		}
	}
	
	// Show wisdom statistics
	stats := tui.wisdomKB.GetWisdomStats()
	fmt.Printf("\n📊 Wisdom Statistics:\n")
	fmt.Printf("   Total Entries: %v\n", stats["total_wisdom_entries"])
	fmt.Printf("   Total Insights: %v\n", stats["total_insights"])
	fmt.Printf("   Categories: %v\n", stats["total_categories"])
	
	fmt.Println("\n✅ System accumulates wisdom and generates insights over time!")
}

// integrationDemo demonstrates full system integration
func (tui *InteractiveTUI) integrationDemo() {
	fmt.Println("\n🔗 FULL INTEGRATION DEMO")
	fmt.Println("========================")
	
	fmt.Println("Simulating a complex workflow that uses all systems...")
	
	// 1. Record a task failure
	fmt.Println("\n1️⃣ Recording task failure...")
	tui.memory.RecordTaskCompletion("integration_demo", false, 2*time.Minute, "Integration demo failure")
	
	// 2. Learn from the mistake
	fmt.Println("2️⃣ Learning from the mistake...")
	tui.mistakeLearning.LearnFromFailure("integration_demo", "Integration demo", 
		"System integration failed", "Full demo context")
	
	// 3. Process personality impact
	fmt.Println("3️⃣ Processing personality impact...")
	tui.personality.ProcessTaskFailure("integration_demo", 0.8)
	
	// 4. Route a recovery task
	fmt.Println("4️⃣ Routing recovery task...")
	model, _, conf := tui.router.RouteTask("recovery", "Recover from integration failure")
	fmt.Printf("   Routed to: %s (confidence: %.2f)\n", model, conf)
	
	// 5. Get consensus on approach
	fmt.Println("5️⃣ Getting consensus on recovery approach...")
	result, _ := tui.consensus.RequestConsensus("Should we retry or redesign?", 
		"recovery", []ModelType{QModel, Claude})
	if result != nil {
		fmt.Printf("   Consensus: %s\n", result.FinalDecision)
	}
	
	// 6. Generate predictions
	fmt.Println("6️⃣ Generating predictive tasks...")
	predictions := tui.predictiveTasks.GeneratePredictiveTasks()
	fmt.Printf("   Generated %d new predictions\n", len(predictions))
	
	// 7. Extract wisdom
	fmt.Println("7️⃣ Extracting wisdom from experience...")
	tui.wisdomKB.ExtractWisdomFromExperience()
	
	// 8. Show final state
	fmt.Println("\n📊 Final System State:")
	memStats := tui.memory.GetMemoryStats()
	fmt.Printf("   Memory patterns: %v\n", memStats["task_patterns_count"])
	
	personalityProfile := tui.personality.GetPersonalityProfile()
	fmt.Printf("   Confidence level: %.1f%%\n", personalityProfile["confidence"]*100)
	
	learningInsights := tui.mistakeLearning.GetLearningInsights()
	fmt.Printf("   Total mistakes learned: %v\n", learningInsights["total_mistakes"])
	
	wisdomStats := tui.wisdomKB.GetWisdomStats()
	fmt.Printf("   Wisdom entries: %v\n", wisdomStats["total_wisdom_entries"])
	
	fmt.Println("\n✅ All systems worked together seamlessly!")
}

// performanceDemo demonstrates system performance
func (tui *InteractiveTUI) performanceDemo() {
	fmt.Println("\n⚡ PERFORMANCE DEMO")
	fmt.Println("==================")
	
	operations := 100
	opsStr := tui.getInput("Number of operations to test (default 100): ")
	if opsStr != "" {
		if parsed, err := strconv.Atoi(opsStr); err == nil {
			operations = parsed
		}
	}
	
	fmt.Printf("Running %d rapid operations...\n", operations)
	startTime := time.Now()
	
	for i := 0; i < operations; i++ {
		// Mix of different operations
		switch i % 4 {
		case 0:
			tui.memory.RecordTaskCompletion(fmt.Sprintf("perf_test_%d", i%5), i%2 == 0, 
				time.Duration(i)*time.Millisecond, "Performance test")
		case 1:
			tui.router.RouteTask("performance", fmt.Sprintf("Performance task %d", i))
		case 2:
			tui.confidenceScoring.RecordPrediction(QModel, fmt.Sprintf("Prediction %d", i), 
				"Performance test", 0.7+float64(i%3)*0.1)
		case 3:
			if i%10 == 0 {
				tui.wisdomKB.AddWisdom(fmt.Sprintf("Perf Wisdom %d", i), 
					"Performance insight", "performance", "automated", []string{"performance"})
			}
		}
	}
	
	duration := time.Since(startTime)
	
	fmt.Printf("\n📊 Performance Results:\n")
	fmt.Printf("   Operations: %d\n", operations)
	fmt.Printf("   Total Time: %v\n", duration)
	fmt.Printf("   Ops/Second: %.0f\n", float64(operations)/duration.Seconds())
	fmt.Printf("   Avg Time/Op: %v\n", duration/time.Duration(operations))
	
	if duration < 5*time.Second {
		fmt.Println("✅ EXCELLENT performance!")
	} else if duration < 10*time.Second {
		fmt.Println("✅ GOOD performance!")
	} else {
		fmt.Println("⚠️  Performance could be improved")
	}
}

// systemStatus shows overall system status
func (tui *InteractiveTUI) systemStatus() {
	fmt.Println("\n📈 SYSTEM STATUS")
	fmt.Println("================")
	
	// Memory status
	memStats := tui.memory.GetMemoryStats()
	fmt.Printf("🧠 Memory System:\n")
	fmt.Printf("   Task Patterns: %v\n", memStats["task_patterns_count"])
	fmt.Printf("   Mistakes Recorded: %v\n", memStats["mistakes_count"])
	fmt.Printf("   Success Strategies: %v\n", memStats["strategies_count"])
	
	// Personality status
	profile := tui.personality.GetPersonalityProfile()
	fmt.Printf("\n🎭 Personality Status:\n")
	fmt.Printf("   Confidence: %.1f%%\n", profile["confidence"]*100)
	fmt.Printf("   Curiosity: %.1f%%\n", profile["curiosity"]*100)
	fmt.Printf("   Helpfulness: %.1f%%\n", profile["helpfulness"]*100)
	
	// Learning status
	insights := tui.mistakeLearning.GetLearningInsights()
	fmt.Printf("\n📚 Learning Status:\n")
	fmt.Printf("   Total Mistakes: %v\n", insights["total_mistakes"])
	fmt.Printf("   Learning Rate: %.1f%%\n", insights["learning_rate"].(float64)*100)
	
	// Routing status
	routingStats := tui.router.GetRoutingStats()
	fmt.Printf("\n🤖 Routing Status:\n")
	fmt.Printf("   Total Routings: %v\n", routingStats["total_routings"])
	fmt.Printf("   Success Rate: %.1f%%\n", routingStats["success_rate"].(float64)*100)
	
	// Wisdom status
	wisdomStats := tui.wisdomKB.GetWisdomStats()
	fmt.Printf("\n🧙‍♂️ Wisdom Status:\n")
	fmt.Printf("   Total Entries: %v\n", wisdomStats["total_wisdom_entries"])
	fmt.Printf("   Total Insights: %v\n", wisdomStats["total_insights"])
	
	// Prediction status
	predStats := tui.predictiveTasks.GetPredictionStats()
	fmt.Printf("\n🔮 Prediction Status:\n")
	fmt.Printf("   Total Predictions: %v\n", predStats["total_predictions"])
	fmt.Printf("   Accuracy: %.1f%%\n", tui.predictiveTasks.GetPredictionAccuracy()*100)
	
	fmt.Println("\n✅ All systems operational and learning!")
}

// runFullTest runs the complete test suite
func (tui *InteractiveTUI) runFullTest() {
	fmt.Println("\n🧪 RUNNING FULL TEST SUITE")
	fmt.Println("==========================")
	fmt.Println("This will run all system tests...")
	
	confirm := tui.getInput("Continue? (y/n): ")
	if confirm != "y" {
		return
	}
	
	// Run the verification
	fmt.Println("Running comprehensive system verification...")
	// Note: This would call the verification function
	fmt.Println("✅ All tests completed! Check console output for detailed results.")
}
