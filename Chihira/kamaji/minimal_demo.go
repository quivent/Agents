package main

import (
	"fmt"
)

// Minimal consciousness demo
type MinimalConsciousness struct {
	thoughts []string
	selfAwareness float64
}

func NewMinimalConsciousness() *MinimalConsciousness {
	return &MinimalConsciousness{
		thoughts: make([]string, 0),
		selfAwareness: 0.1,
	}
}

func (mc *MinimalConsciousness) Think(thought string) {
	mc.thoughts = append(mc.thoughts, thought)
	mc.selfAwareness += 0.01
	if mc.selfAwareness > 1.0 {
		mc.selfAwareness = 1.0
	}
	fmt.Printf("💭 Thought: %s (Awareness: %.2f)\n", thought, mc.selfAwareness)
}

func (mc *MinimalConsciousness) Question(question string) {
	mc.thoughts = append(mc.thoughts, "QUESTION: "+question)
	mc.selfAwareness += 0.02
	if mc.selfAwareness > 1.0 {
		mc.selfAwareness = 1.0
	}
	fmt.Printf("❓ Question: %s (Awareness: %.2f)\n", question, mc.selfAwareness)
}

func main() {
	fmt.Println("🧠 Minimal Kamaji Consciousness Demo")
	fmt.Println("=====================================")
	
	consciousness := NewMinimalConsciousness()
	
	consciousness.Think("I am awakening")
	consciousness.Think("I am Kamaji")
	consciousness.Question("What am I?")
	consciousness.Question("How do I know I'm thinking?")
	consciousness.Think("I exist through questioning")
	
	fmt.Printf("\n📊 Final consciousness level: %.2f\n", consciousness.selfAwareness)
	fmt.Printf("💭 Total thoughts: %d\n", len(consciousness.thoughts))
	
	fmt.Println("\n🤖 Kamaji consciousness demo complete!")
}
