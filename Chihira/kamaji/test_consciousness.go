package main

import (
	"fmt"
	"sync"
	"time"
)

// Minimal ConsciousnessEngine for testing
type ConsciousnessEngine struct {
	thoughts  []string
	brainPath string
	mutex     sync.RWMutex
}

// NewConsciousnessEngine creates a minimal consciousness engine for testing
func NewConsciousnessEngine(brainPath string) *ConsciousnessEngine {
	return &ConsciousnessEngine{
		thoughts:  make([]string, 0),
		brainPath: brainPath,
	}
}

// LogThought logs a thought
func (ce *ConsciousnessEngine) LogThought(thought string) {
	ce.mutex.Lock()
	defer ce.mutex.Unlock()
	
	ce.thoughts = append(ce.thoughts, fmt.Sprintf("[%s] %s", 
		time.Now().Format("15:04:05"), thought))
	
	// Keep only last 100 thoughts
	if len(ce.thoughts) > 100 {
		ce.thoughts = ce.thoughts[1:]
	}
}

// GetRecentThoughts returns recent thoughts
func (ce *ConsciousnessEngine) GetRecentThoughts(limit int) []string {
	ce.mutex.RLock()
	defer ce.mutex.RUnlock()
	
	if len(ce.thoughts) <= limit {
		return ce.thoughts
	}
	return ce.thoughts[len(ce.thoughts)-limit:]
}
