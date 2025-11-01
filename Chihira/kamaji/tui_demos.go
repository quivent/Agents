package main

import (
	"fmt"
	"strconv"
	"time"
)

// memoryDemo demonstrates the memory persistence system
func (tui *InteractiveTUI) memoryDemo() {
	fmt.Println("\n🧠 MEMORY PERSISTENCE DEMO")
	fmt.Println("==========================")
	
	taskName := tui.getInput("Enter a task name: ")
	if taskName == "" {
		taskName = "demo_task"
	}
	
	successStr := tui.getInput("Was it successful? (y/n): ")
	success := successStr == "y" || successStr == "yes"
	
	durationStr := tui.getInput("How long did it take (seconds)? ")
	duration := 30
	if d, err := strconv.Atoi(durationStr); err == nil {
		duration = d
	}
	
	// Record the task
	tui.memory.RecordTaskCompletion(taskName, success, time.Duration(duration)*time.Second, "Interactive demo")
	
	// Show the pattern
	pattern := tui.memory.GetTaskPattern(taskName)
	if pattern != nil {
		fmt.Printf("\n📊 Task Pattern for '%s':\n", taskName)
		fmt.Printf("   Success Rate: %.1f%%\n", pattern.SuccessRate*100)
		fmt.Printf("   Average Time: %.1f seconds\n", pattern.AverageTime)
		fmt.Printf("   Completions: %d\n", pattern.CompletionCount)
		fmt.Printf("   Last Updated: %s\n", pattern.LastUpdated.Format("15:04:05"))
	}
	
	// Show memory stats
	stats := tui.memory.GetMemoryStats()
	fmt.Printf("\n🧠 Memory Statistics:\n")
	fmt.Printf("   Task Patterns: %v\n", stats["task_patterns_count"])
	fmt.Printf("   Mistakes: %v\n", stats["mistakes_count"])
	fmt.Printf("   Strategies: %v\n", stats["strategies_count"])
	
	fmt.Println("\n✅ Memory persisted! This data will survive system restarts.")
}

// learningDemo demonstrates learning from mistakes
func (tui *InteractiveTUI) learningDemo() {
	fmt.Println("\n📚 LEARNING FROM MISTAKES DEMO")
	fmt.Println("===============================")
	
	taskType := tui.getInput("Enter task type: ")
	if taskType == "" {
		taskType = "demo_task"
	}
	
	errorMsg := tui.getInput("Describe what went wrong: ")
	if errorMsg == "" {
		errorMsg = "Something failed during execution"
	}
	
	solution := tui.getInput("What's the solution: ")
	if solution == "" {
		solution = "Check configuration and retry"
	}
	
	// Record the mistake
	tui.mistakeLearning.LearnFromFailure(taskType, "Demo task", errorMsg, "Interactive demo")
	
	// Show learning insights
	insights := tui.mistakeLearning.GetLearningInsights()
	fmt.Printf("\n🧠 Learning Insights:\n")
	fmt.Printf("   Total Mistakes: %v\n", insights["total_mistakes"])
	fmt.Printf("   Learning Rate: %.1f%%\n", insights["learning_rate"].(float64)*100)
	fmt.Printf("   Average Severity: %.1f\n", insights["average_severity"])
	
	// Test for warnings on similar task
	warnings := tui.mistakeLearning.AnalyzeTaskForMistakes(taskType, "Similar "+taskType+" task")
	if len(warnings) > 0 {
		fmt.Printf("\n⚠️  Warnings for similar tasks:\n")
		for _, warning := range warnings {
			fmt.Printf("   %s\n", warning)
		}
	}
	
	fmt.Println("\n✅ System learned from this mistake and will warn about similar issues!")
}

// personalityDemo demonstrates personality development
func (tui *InteractiveTUI) personalityDemo() {
	fmt.Println("\n🎭 PERSONALITY DEVELOPMENT DEMO")
	fmt.Println("===============================")
	
	// Show current personality
	profile := tui.personality.GetPersonalityProfile()
	fmt.Println("📊 Current Personality Profile:")
	for trait, value := range profile {
		fmt.Printf("   %s: %.1f%%\n", trait, value*100)
	}
	
	// Process an experience
	fmt.Println("\nLet's simulate an experience...")
	expType := tui.getInput("Experience type (success/failure/learning): ")
	
	switch expType {
	case "success":
		difficulty := 0.8
		if d := tui.getInput("Difficulty (0.1-1.0): "); d != "" {
			if parsed, err := strconv.ParseFloat(d, 64); err == nil {
				difficulty = parsed
			}
		}
		tui.personality.ProcessTaskSuccess("demo_task", difficulty)
		
	case "failure":
		severity := 0.6
		if s := tui.getInput("Severity (0.1-1.0): "); s != "" {
			if parsed, err := strconv.ParseFloat(s, 64); err == nil {
				severity = parsed
			}
		}
		tui.personality.ProcessTaskFailure("demo_task", severity)
		
	case "learning":
		complexity := 0.7
		if c := tui.getInput("Complexity (0.1-1.0): "); c != "" {
			if parsed, err := strconv.ParseFloat(c, 64); err == nil {
				complexity = parsed
			}
		}
		tui.personality.ProcessLearningEvent("Interactive learning experience", complexity)
	}
	
	// Show updated personality
	newProfile := tui.personality.GetPersonalityProfile()
	fmt.Println("\n📈 Updated Personality Profile:")
	for trait, value := range newProfile {
		oldValue := profile[trait]
		change := value - oldValue
		changeStr := ""
		if change > 0.001 {
			changeStr = fmt.Sprintf(" (+%.3f)", change)
		} else if change < -0.001 {
			changeStr = fmt.Sprintf(" (%.3f)", change)
		}
		fmt.Printf("   %s: %.1f%%%s\n", trait, value*100, changeStr)
	}
	
	fmt.Println("\n✅ Personality evolved based on the experience!")
}

// routingDemo demonstrates AI model routing
func (tui *InteractiveTUI) routingDemo() {
	fmt.Println("\n🤖 AI MODEL ROUTING DEMO")
	fmt.Println("========================")
	
	fmt.Println("Let's see how Kamaji routes different tasks to the best AI models...")
	
	testTasks := []struct {
		taskType string
		query    string
	}{
		{"aws_task", "Deploy a serverless Lambda function"},
		{"analysis", "Analyze user behavior patterns"},
		{"code_generation", "Generate a REST API in Go"},
		{"creative", "Design a user-friendly interface"},
		{"security", "Implement OAuth authentication"},
	}
	
	fmt.Println("\n🎯 Routing Results:")
	for _, task := range testTasks {
		model, reasoning, confidence := tui.router.RouteTask(task.taskType, task.query)
		fmt.Printf("\n📋 Task: %s\n", task.query)
		fmt.Printf("   Routed to: %s (confidence: %.2f)\n", model, confidence)
		fmt.Printf("   Reasoning: %s\n", reasoning)
		
		// Simulate task completion
		success := confidence > 0.7
		tui.router.UpdateRoutingSuccess(task.taskType, task.query, success)
	}
	
	// Show routing statistics
	stats := tui.router.GetRoutingStats()
	fmt.Printf("\n📊 Routing Statistics:\n")
	fmt.Printf("   Total Routings: %v\n", stats["total_routings"])
	fmt.Printf("   Success Rate: %.1f%%\n", stats["success_rate"].(float64)*100)
	fmt.Printf("   Model Usage: %v\n", stats["model_usage"])
	
	fmt.Println("\n✅ System learns which models work best for which tasks!")
}

// consensusDemo demonstrates multi-model consensus
func (tui *InteractiveTUI) consensusDemo() {
	fmt.Println("\n🤝 MULTI-MODEL CONSENSUS DEMO")
	fmt.Println("=============================")
	
	question := tui.getInput("Enter a decision question: ")
	if question == "" {
		question = "Should we prioritize performance or security?"
	}
	
	taskType := tui.getInput("Task type: ")
	if taskType == "" {
		taskType = "decision_making"
	}
	
	fmt.Println("\n🤔 Requesting consensus from multiple AI models...")
	
	result, err := tui.consensus.RequestConsensus(question, taskType, []ModelType{QModel, Claude})
	
	if err != nil {
		fmt.Printf("❌ Consensus failed: %v\n", err)
		return
	}
	
	fmt.Printf("\n🎯 Consensus Result:\n")
	fmt.Printf("   Decision: %s\n", result.FinalDecision)
	fmt.Printf("   Confidence: %.2f\n", result.Confidence)
	fmt.Printf("   Consensus Type: %s\n", result.ConsensusType)
	fmt.Printf("   Processing Time: %v\n", result.ProcessingTime)
	
	fmt.Printf("\n🗳️  Individual Votes:\n")
	for _, vote := range result.Votes {
		fmt.Printf("   %s: %s (confidence: %.2f)\n", vote.ModelType, vote.Decision, vote.Confidence)
		fmt.Printf("      Reasoning: %s\n", vote.Reasoning)
	}
	
	// Show consensus statistics
	stats := tui.consensus.GetConsensusStats()
	fmt.Printf("\n📊 Consensus Statistics:\n")
	fmt.Printf("   Total Consensus: %v\n", stats["total_consensus"])
	fmt.Printf("   Average Confidence: %.2f\n", stats["avg_confidence"])
	
	fmt.Println("\n✅ Multiple AI models collaborated to reach this decision!")
}
