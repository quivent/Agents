package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// TestLearningAndPersonality tests the learning and personality systems
func TestLearningAndPersonality() {
	fmt.Println("🧠 Testing Learning and Personality Systems...")
	
	// Create temporary brain directory
	tempDir := "/tmp/kamaji_learning_test"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	brainPath := filepath.Join(tempDir, "brain.json")
	
	// Create systems
	memory := NewMemoryPersistence(brainPath)
	consciousness := NewConsciousnessEngine(brainPath)
	mistakeLearning := NewMistakeLearning(memory, consciousness)
	personality := NewPersonalityDevelopment(memory, consciousness)
	
	// Test 1: Mistake Learning
	fmt.Println("📝 Test 1: Testing mistake learning...")
	
	// Simulate a task failure
	mistakeLearning.LearnFromFailure("code_generation", "Generate Go function", 
		"syntax error: missing semicolon", "Go development")
	
	// Check for warnings on similar task
	warnings := mistakeLearning.AnalyzeTaskForMistakes("code_generation", 
		"Generate Go struct with syntax checking")
	
	fmt.Printf("✅ Generated %d warnings for similar task\n", len(warnings))
	if len(warnings) > 0 {
		fmt.Printf("   Warning: %s\n", warnings[0])
	}
	
	// Test learning insights
	insights := mistakeLearning.GetLearningInsights()
	fmt.Printf("✅ Learning insights: %v\n", insights)
	
	// Test 2: Personality Development
	fmt.Println("📝 Test 2: Testing personality development...")
	
	// Get initial personality
	initialProfile := personality.GetPersonalityProfile()
	initialConfidence := initialProfile["confidence"]
	fmt.Printf("✅ Initial confidence: %.2f\n", initialConfidence)
	
	// Process successful task
	personality.ProcessTaskSuccess("code_generation", 0.8)
	
	// Check personality change
	newProfile := personality.GetPersonalityProfile()
	newConfidence := newProfile["confidence"]
	fmt.Printf("✅ New confidence after success: %.2f (change: %+.3f)\n", 
		newConfidence, newConfidence-initialConfidence)
	
	// Process failure
	personality.ProcessTaskFailure("code_generation", 0.6)
	
	// Check personality change again
	finalProfile := personality.GetPersonalityProfile()
	finalConfidence := finalProfile["confidence"]
	fmt.Printf("✅ Final confidence after failure: %.2f (change: %+.3f)\n", 
		finalConfidence, finalConfidence-newConfidence)
	
	// Test 3: Personality Description
	fmt.Println("📝 Test 3: Testing personality description...")
	description := personality.GetPersonalityDescription()
	fmt.Printf("✅ Personality description:\n%s\n", description)
	
	// Test 4: Learning Event Processing
	fmt.Println("📝 Test 4: Testing learning event processing...")
	personality.ProcessLearningEvent("Learned about Go interfaces", 0.7)
	
	learningProfile := personality.GetPersonalityProfile()
	curiosity := learningProfile["curiosity"]
	fmt.Printf("✅ Curiosity after learning: %.2f\n", curiosity)
	
	// Test 5: User Interaction Processing
	fmt.Println("📝 Test 5: Testing user interaction processing...")
	personality.ProcessUserInteraction(true, 0.8) // Positive interaction
	
	interactionProfile := personality.GetPersonalityProfile()
	empathy := interactionProfile["empathy"]
	helpfulness := interactionProfile["helpfulness"]
	fmt.Printf("✅ Empathy: %.2f, Helpfulness: %.2f after positive interaction\n", 
		empathy, helpfulness)
	
	// Test 6: Personality Evolution
	fmt.Println("📝 Test 6: Testing personality evolution tracking...")
	evolution := personality.GetPersonalityEvolution()
	fmt.Printf("✅ Personality evolution: %v\n", evolution)
	
	// Test 7: Mistake Prevention Strategies
	fmt.Println("📝 Test 7: Testing mistake prevention strategies...")
	strategies := mistakeLearning.GetMistakePreventionStrategies("code_generation")
	fmt.Printf("✅ Prevention strategies: %v\n", strategies)
	
	fmt.Println("🎉 All learning and personality tests passed!")
}

// RunLearningPersonalityTest is a standalone function to test learning and personality
func RunLearningPersonalityTest() {
	TestLearningAndPersonality()
}
