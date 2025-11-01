package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// TestMemoryPersistence tests the memory persistence system
func TestMemoryPersistence() {
	fmt.Println("🧠 Testing Memory Persistence System...")
	
	// Create temporary brain directory
	tempDir := "/tmp/kamaji_test"
	os.MkdirAll(tempDir, 0755)
	defer os.RemoveAll(tempDir)
	
	brainPath := filepath.Join(tempDir, "brain.json")
	
	// Test 1: Create memory system
	fmt.Println("📝 Test 1: Creating memory system...")
	memory := NewMemoryPersistence(brainPath)
	if memory == nil {
		fmt.Println("❌ Failed to create memory system")
		return
	}
	fmt.Println("✅ Memory system created successfully")
	
	// Test 2: Record task completion
	fmt.Println("📝 Test 2: Recording task completion...")
	memory.RecordTaskCompletion("code_generation", true, 30*time.Second, "Generated Go function")
	memory.RecordTaskCompletion("code_generation", false, 45*time.Second, "Failed to compile")
	memory.RecordTaskCompletion("code_generation", true, 25*time.Second, "Generated Python script")
	
	pattern := memory.GetTaskPattern("code_generation")
	if pattern == nil {
		fmt.Println("❌ Failed to retrieve task pattern")
		return
	}
	fmt.Printf("✅ Task pattern: Success Rate: %.2f%%, Avg Time: %.1fs, Completions: %d\n", 
		pattern.SuccessRate*100, pattern.AverageTime, pattern.CompletionCount)
	
	// Test 3: Record mistakes
	fmt.Println("📝 Test 3: Recording mistakes...")
	memory.RecordMistake("code_generation", "Syntax error in Go", "Missing semicolon", "Add proper semicolon", 3)
	memory.RecordMistake("code_generation", "Import error", "Wrong package name", "Use correct import path", 5)
	
	mistakes := memory.GetRecentMistakes(5)
	fmt.Printf("✅ Recorded %d mistakes for learning\n", len(mistakes))
	
	// Test 4: Record success strategies
	fmt.Println("📝 Test 4: Recording success strategies...")
	memory.RecordSuccessStrategy("code_generation", "Use type checking before compilation", "Go development")
	memory.RecordSuccessStrategy("code_generation", "Test with small examples first", "Python scripting")
	
	updatedPattern := memory.GetTaskPattern("code_generation")
	fmt.Printf("✅ Best strategies: %v\n", updatedPattern.BestStrategies)
	
	// Test 5: Personality development
	fmt.Println("📝 Test 5: Testing personality development...")
	memory.UpdatePersonalityTrait("curiosity_level", 0.8)
	memory.UpdatePersonalityTrait("patience_level", 0.6)
	memory.UpdatePersonalityTrait("mystical_wisdom", 0.9)
	
	curiosity := memory.GetPersonalityTrait("curiosity_level")
	fmt.Printf("✅ Curiosity level: %v\n", curiosity)
	
	// Test 6: Memory statistics
	fmt.Println("📝 Test 6: Memory statistics...")
	stats := memory.GetMemoryStats()
	fmt.Printf("✅ Memory stats: %v\n", stats)
	
	// Test 7: Persistence across sessions
	fmt.Println("📝 Test 7: Testing persistence across sessions...")
	
	// Create new memory instance (simulating restart)
	memory2 := NewMemoryPersistence(brainPath)
	pattern2 := memory2.GetTaskPattern("code_generation")
	if pattern2 == nil {
		fmt.Println("❌ Failed to load persisted data")
		return
	}
	
	if pattern2.CompletionCount != pattern.CompletionCount {
		fmt.Println("❌ Data not properly persisted")
		return
	}
	
	fmt.Printf("✅ Data persisted correctly: %d completions loaded\n", pattern2.CompletionCount)
	
	fmt.Println("🎉 All memory persistence tests passed!")
}

// RunMemoryTest is a standalone function to test memory
func RunMemoryTest() {
	TestMemoryPersistence()
}
