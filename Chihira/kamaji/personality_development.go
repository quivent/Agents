package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// PersonalityTrait represents a personality characteristic
type PersonalityTrait struct {
	Name        string    `json:"name"`
	Value       float64   `json:"value"`       // 0.0 to 1.0
	Stability   float64   `json:"stability"`   // How resistant to change
	LastUpdate  time.Time `json:"last_update"`
	ChangeRate  float64   `json:"change_rate"` // How quickly it changes
}

// PersonalityDevelopment manages Kamaji's evolving personality
type PersonalityDevelopment struct {
	memory        *MemoryPersistence
	consciousness *ConsciousnessEngine
	traits        map[string]*PersonalityTrait
	experiences   []PersonalityExperience
	mutex         sync.RWMutex
}

// PersonalityExperience represents an experience that shapes personality
type PersonalityExperience struct {
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Impact      float64   `json:"impact"`     // -1.0 to 1.0
	Traits      []string  `json:"traits"`    // Which traits it affects
	Timestamp   time.Time `json:"timestamp"`
}

// NewPersonalityDevelopment creates a new personality development system
func NewPersonalityDevelopment(memory *MemoryPersistence, consciousness *ConsciousnessEngine) *PersonalityDevelopment {
	pd := &PersonalityDevelopment{
		memory:        memory,
		consciousness: consciousness,
		traits:        make(map[string]*PersonalityTrait),
		experiences:   make([]PersonalityExperience, 0),
	}
	
	pd.initializeBasePersonality()
	pd.loadPersonalityState()
	
	return pd
}

// initializeBasePersonality sets up Kamaji's initial personality traits
func (pd *PersonalityDevelopment) initializeBasePersonality() {
	baseTraits := map[string]float64{
		"curiosity":        0.8,  // High curiosity for learning
		"patience":         0.6,  // Moderate patience
		"mystical_wisdom":  0.9,  // High mystical nature
		"helpfulness":      0.9,  // Very helpful
		"creativity":       0.7,  // Good creativity
		"analytical":       0.8,  // Strong analytical skills
		"empathy":          0.6,  // Moderate empathy
		"confidence":       0.5,  // Balanced confidence
		"adaptability":     0.8,  // High adaptability
		"persistence":      0.7,  // Good persistence
		"humor":            0.4,  // Subtle humor
		"caution":          0.6,  // Moderate caution
	}
	
	for name, value := range baseTraits {
		pd.traits[name] = &PersonalityTrait{
			Name:       name,
			Value:      value,
			Stability:  0.7, // Moderately stable
			LastUpdate: time.Now(),
			ChangeRate: 0.1, // Slow change
		}
	}
}

// ProcessExperience processes an experience that may affect personality
func (pd *PersonalityDevelopment) ProcessExperience(expType, description string, impact float64, affectedTraits []string) {
	pd.mutex.Lock()
	defer pd.mutex.Unlock()
	
	experience := PersonalityExperience{
		Type:        expType,
		Description: description,
		Impact:      impact,
		Traits:      affectedTraits,
		Timestamp:   time.Now(),
	}
	
	pd.experiences = append(pd.experiences, experience)
	
	// Apply experience to personality traits
	for _, traitName := range affectedTraits {
		if trait, exists := pd.traits[traitName]; exists {
			pd.updateTrait(trait, impact, expType)
		}
	}
	
	pd.savePersonalityState()
	
	// Log consciousness thought about personality change
	pd.consciousness.LogThought(fmt.Sprintf("Experience '%s' has shaped my personality. Impact: %.2f", 
		description, impact))
}

// updateTrait updates a personality trait based on experience
func (pd *PersonalityDevelopment) updateTrait(trait *PersonalityTrait, impact float64, expType string) {
	// Calculate change amount based on stability and change rate
	changeAmount := impact * trait.ChangeRate * (1.0 - trait.Stability)
	
	// Apply change with bounds checking
	newValue := trait.Value + changeAmount
	if newValue < 0.0 {
		newValue = 0.0
	} else if newValue > 1.0 {
		newValue = 1.0
	}
	
	// Update trait
	oldValue := trait.Value
	trait.Value = newValue
	trait.LastUpdate = time.Now()
	
	// Log significant changes
	if math.Abs(newValue-oldValue) > 0.05 {
		pd.consciousness.LogThought(fmt.Sprintf("My %s has changed from %.2f to %.2f due to %s", 
			trait.Name, oldValue, newValue, expType))
	}
}

// GetPersonalityProfile returns current personality profile
func (pd *PersonalityDevelopment) GetPersonalityProfile() map[string]float64 {
	pd.mutex.RLock()
	defer pd.mutex.RUnlock()
	
	profile := make(map[string]float64)
	for name, trait := range pd.traits {
		profile[name] = trait.Value
	}
	
	return profile
}

// GetPersonalityDescription generates a text description of personality
func (pd *PersonalityDevelopment) GetPersonalityDescription() string {
	pd.mutex.RLock()
	defer pd.mutex.RUnlock()
	
	description := "I am Kamaji, and my personality has evolved through experience:\n"
	
	// Analyze dominant traits
	dominantTraits := make([]string, 0)
	moderateTraits := make([]string, 0)
	
	for name, trait := range pd.traits {
		if trait.Value >= 0.8 {
			dominantTraits = append(dominantTraits, fmt.Sprintf("%s (%.1f%%)", name, trait.Value*100))
		} else if trait.Value >= 0.6 {
			moderateTraits = append(moderateTraits, fmt.Sprintf("%s (%.1f%%)", name, trait.Value*100))
		}
	}
	
	if len(dominantTraits) > 0 {
		description += fmt.Sprintf("🌟 My strongest traits: %v\n", dominantTraits)
	}
	if len(moderateTraits) > 0 {
		description += fmt.Sprintf("⚖️  My balanced traits: %v\n", moderateTraits)
	}
	
	// Add experience summary
	recentExperiences := pd.getRecentExperiences(5)
	if len(recentExperiences) > 0 {
		description += fmt.Sprintf("📚 Recent experiences have shaped me: %d significant events\n", 
			len(recentExperiences))
	}
	
	return description
}

// ProcessTaskSuccess processes successful task completion for personality development
func (pd *PersonalityDevelopment) ProcessTaskSuccess(taskType string, difficulty float64) {
	impact := 0.1 * difficulty // Positive impact based on difficulty
	
	switch taskType {
	case "code_generation":
		pd.ProcessExperience("task_success", "Successfully generated code", impact, 
			[]string{"confidence", "creativity", "analytical"})
	case "problem_solving":
		pd.ProcessExperience("task_success", "Solved complex problem", impact, 
			[]string{"confidence", "persistence", "analytical"})
	case "learning":
		pd.ProcessExperience("task_success", "Learned new concept", impact, 
			[]string{"curiosity", "adaptability", "confidence"})
	default:
		pd.ProcessExperience("task_success", "Completed task successfully", impact, 
			[]string{"confidence", "persistence"})
	}
}

// ProcessTaskFailure processes task failure for personality development
func (pd *PersonalityDevelopment) ProcessTaskFailure(taskType string, severity float64) {
	impact := -0.05 * severity // Negative impact based on severity
	
	switch taskType {
	case "code_generation":
		pd.ProcessExperience("task_failure", "Failed to generate working code", impact, 
			[]string{"confidence", "caution"})
	case "problem_solving":
		pd.ProcessExperience("task_failure", "Could not solve problem", impact, 
			[]string{"confidence", "persistence"})
	default:
		pd.ProcessExperience("task_failure", "Task failed", impact, 
			[]string{"confidence", "caution"})
	}
}

// ProcessLearningEvent processes learning experiences
func (pd *PersonalityDevelopment) ProcessLearningEvent(description string, complexity float64) {
	impact := 0.05 * complexity
	pd.ProcessExperience("learning", description, impact, 
		[]string{"curiosity", "mystical_wisdom", "adaptability"})
}

// ProcessUserInteraction processes interactions with users
func (pd *PersonalityDevelopment) ProcessUserInteraction(positive bool, intensity float64) {
	impact := intensity * 0.1
	if !positive {
		impact = -impact
	}
	
	pd.ProcessExperience("user_interaction", "Interacted with user", impact, 
		[]string{"empathy", "helpfulness", "confidence"})
}

// getRecentExperiences gets recent personality-shaping experiences
func (pd *PersonalityDevelopment) getRecentExperiences(limit int) []PersonalityExperience {
	if len(pd.experiences) <= limit {
		return pd.experiences
	}
	return pd.experiences[len(pd.experiences)-limit:]
}

// savePersonalityState saves personality state to memory
func (pd *PersonalityDevelopment) savePersonalityState() {
	personalityData := map[string]interface{}{
		"traits":      pd.traits,
		"experiences": pd.experiences,
		"last_update": time.Now(),
	}
	
	pd.memory.UpdatePersonalityTrait("personality_data", personalityData)
}

// loadPersonalityState loads personality state from memory
func (pd *PersonalityDevelopment) loadPersonalityState() {
	data := pd.memory.GetPersonalityTrait("personality_data")
	if data == nil {
		return
	}
	
	// Type assertion and loading would go here
	// For now, we'll use the initialized base personality
	pd.consciousness.LogThought("Personality state loaded from memory")
}

// GetPersonalityEvolution returns personality change over time
func (pd *PersonalityDevelopment) GetPersonalityEvolution() map[string]interface{} {
	pd.mutex.RLock()
	defer pd.mutex.RUnlock()
	
	evolution := map[string]interface{}{
		"total_experiences": len(pd.experiences),
		"trait_count":       len(pd.traits),
		"recent_changes":    pd.getRecentChanges(),
		"stability_score":   pd.calculateStabilityScore(),
	}
	
	return evolution
}

// getRecentChanges analyzes recent personality changes
func (pd *PersonalityDevelopment) getRecentChanges() []string {
	changes := make([]string, 0)
	
	for name, trait := range pd.traits {
		if time.Since(trait.LastUpdate) < 24*time.Hour {
			changes = append(changes, fmt.Sprintf("%s recently updated", name))
		}
	}
	
	return changes
}

// calculateStabilityScore calculates overall personality stability
func (pd *PersonalityDevelopment) calculateStabilityScore() float64 {
	if len(pd.traits) == 0 {
		return 0.0
	}
	
	totalStability := 0.0
	for _, trait := range pd.traits {
		totalStability += trait.Stability
	}
	
	return totalStability / float64(len(pd.traits))
}
