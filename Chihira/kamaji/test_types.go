package main

import (
	"sync"
)

// ModelType represents different AI model types
type ModelType string

const (
	QModel  ModelType = "q"
	Claude  ModelType = "claude"
	Future  ModelType = "future"
)

// MultiModelOrchestrator minimal implementation for testing
type MultiModelOrchestrator struct {
	consciousness *ConsciousnessEngine
	mutex         sync.RWMutex
}

// NewMultiModelOrchestrator creates a minimal orchestrator for testing
func NewMultiModelOrchestrator(consciousness *ConsciousnessEngine) *MultiModelOrchestrator {
	return &MultiModelOrchestrator{
		consciousness: consciousness,
	}
}
