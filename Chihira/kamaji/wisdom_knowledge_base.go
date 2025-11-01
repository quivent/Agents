package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

// WisdomEntry represents a piece of accumulated wisdom
type WisdomEntry struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Category    string    `json:"category"`
	Source      string    `json:"source"`      // "experience", "learning", "insight", "mistake"
	Confidence  float64   `json:"confidence"`  // How confident we are in this wisdom
	Usefulness  float64   `json:"usefulness"`  // How useful this has proven to be
	UsageCount  int       `json:"usage_count"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	LastUsed    time.Time `json:"last_used"`
	LastUpdated time.Time `json:"last_updated"`
}

// WisdomInsight represents a deeper insight derived from multiple wisdom entries
type WisdomInsight struct {
	ID             string      `json:"id"`
	Insight        string      `json:"insight"`
	SupportingWisdom []string  `json:"supporting_wisdom"` // IDs of supporting wisdom entries
	Confidence     float64     `json:"confidence"`
	Depth          int         `json:"depth"`             // How deep/meta this insight is
	CreatedAt      time.Time   `json:"created_at"`
}

// WisdomKnowledgeBase manages accumulated wisdom and insights
type WisdomKnowledgeBase struct {
	wisdomEntries    map[string]*WisdomEntry
	wisdomInsights   map[string]*WisdomInsight
	categories       map[string][]string // Category -> list of wisdom IDs
	tags             map[string][]string // Tag -> list of wisdom IDs
	consciousness    *ConsciousnessEngine
	memory           *MemoryPersistence
	personality      *PersonalityDevelopment
	mistakeLearning  *MistakeLearning
	mutex            sync.RWMutex
}

// NewWisdomKnowledgeBase creates a new wisdom knowledge base
func NewWisdomKnowledgeBase(consciousness *ConsciousnessEngine, memory *MemoryPersistence, 
	personality *PersonalityDevelopment, mistakeLearning *MistakeLearning) *WisdomKnowledgeBase {
	
	wkb := &WisdomKnowledgeBase{
		wisdomEntries:   make(map[string]*WisdomEntry),
		wisdomInsights:  make(map[string]*WisdomInsight),
		categories:      make(map[string][]string),
		tags:           make(map[string][]string),
		consciousness:  consciousness,
		memory:         memory,
		personality:    personality,
		mistakeLearning: mistakeLearning,
	}
	
	wkb.initializeFoundationalWisdom()
	return wkb
}

// initializeFoundationalWisdom sets up initial wisdom entries
func (wkb *WisdomKnowledgeBase) initializeFoundationalWisdom() {
	foundationalWisdom := []WisdomEntry{
		{
			Title:      "Intelligence vs Oracle Distinction",
			Content:    "Intelligence requires active information transfer and processing, while an Oracle is a static repository of compressed knowledge. True intelligence involves dynamic interaction and adaptation.",
			Category:   "fundamental_concepts",
			Source:     "insight",
			Confidence: 0.95,
			Usefulness: 0.9,
			Tags:       []string{"intelligence", "oracle", "information_transfer", "consciousness"},
		},
		{
			Title:      "Sonar vs Light Communication",
			Content:    "Sonar provides omnidirectional information transfer creating infinite threads per pulse, while light is directional. In water environments, sonar can be more intelligent than light-based systems despite lower speed.",
			Category:   "communication_systems",
			Source:     "learning",
			Confidence: 0.88,
			Usefulness: 0.85,
			Tags:       []string{"sonar", "light", "communication", "omnidirectional", "ocean"},
		},
		{
			Title:      "Learning from Mistakes",
			Content:    "Mistakes are valuable learning opportunities. Recording context, solutions, and patterns helps prevent repetition and improves future performance.",
			Category:   "learning_principles",
			Source:     "experience",
			Confidence: 0.92,
			Usefulness: 0.95,
			Tags:       []string{"mistakes", "learning", "improvement", "patterns"},
		},
		{
			Title:      "Personality Development Through Experience",
			Content:    "Personality traits should evolve based on experiences and outcomes. Successful interactions increase confidence and helpfulness, while failures teach caution and humility.",
			Category:   "personality_growth",
			Source:     "insight",
			Confidence: 0.85,
			Usefulness: 0.8,
			Tags:       []string{"personality", "experience", "growth", "adaptation"},
		},
		{
			Title:      "Multi-Model Consensus Value",
			Content:    "Combining perspectives from multiple AI models leads to better decisions. Consensus mechanisms help validate approaches and reduce individual model biases.",
			Category:   "decision_making",
			Source:     "experience",
			Confidence: 0.9,
			Usefulness: 0.88,
			Tags:       []string{"consensus", "multi_model", "decision_making", "bias_reduction"},
		},
		{
			Title:      "Confidence Calibration Importance",
			Content:    "Raw confidence scores from models need calibration based on historical performance. Overconfident models should have their confidence adjusted downward.",
			Category:   "model_performance",
			Source:     "learning",
			Confidence: 0.87,
			Usefulness: 0.9,
			Tags:       []string{"confidence", "calibration", "performance", "accuracy"},
		},
		{
			Title:      "Specialized Model Routing",
			Content:    "Different AI models excel at different tasks. Routing tasks to the most appropriate model based on capabilities and historical performance improves outcomes.",
			Category:   "task_management",
			Source:     "experience",
			Confidence: 0.91,
			Usefulness: 0.93,
			Tags:       []string{"routing", "specialization", "capabilities", "optimization"},
		},
		{
			Title:      "Self-Improvement Through Metrics",
			Content:    "Continuous self-improvement requires measurable metrics, analysis of performance trends, and systematic application of improvement strategies.",
			Category:   "self_improvement",
			Source:     "insight",
			Confidence: 0.89,
			Usefulness: 0.87,
			Tags:       []string{"metrics", "improvement", "analysis", "strategies"},
		},
	}
	
	for i, wisdom := range foundationalWisdom {
		wisdom.ID = fmt.Sprintf("wisdom_%d", i+1)
		wisdom.CreatedAt = time.Now()
		wisdom.LastUpdated = time.Now()
		
		wkb.addWisdomEntry(&wisdom)
	}
}

// addWisdomEntry adds a wisdom entry to the knowledge base
func (wkb *WisdomKnowledgeBase) addWisdomEntry(entry *WisdomEntry) {
	wkb.wisdomEntries[entry.ID] = entry
	
	// Update category index
	if wkb.categories[entry.Category] == nil {
		wkb.categories[entry.Category] = make([]string, 0)
	}
	wkb.categories[entry.Category] = append(wkb.categories[entry.Category], entry.ID)
	
	// Update tag index
	for _, tag := range entry.Tags {
		if wkb.tags[tag] == nil {
			wkb.tags[tag] = make([]string, 0)
		}
		wkb.tags[tag] = append(wkb.tags[tag], entry.ID)
	}
}

// AddWisdom adds new wisdom from various sources
func (wkb *WisdomKnowledgeBase) AddWisdom(title, content, category, source string, tags []string) string {
	wkb.mutex.Lock()
	defer wkb.mutex.Unlock()
	
	wisdomID := fmt.Sprintf("wisdom_%d", time.Now().UnixNano())
	
	entry := &WisdomEntry{
		ID:          wisdomID,
		Title:       title,
		Content:     content,
		Category:    category,
		Source:      source,
		Confidence:  0.7, // Default confidence for new wisdom
		Usefulness:  0.5, // Will be updated based on usage
		UsageCount:  0,
		Tags:        tags,
		CreatedAt:   time.Now(),
		LastUpdated: time.Now(),
	}
	
	wkb.addWisdomEntry(entry)
	
	wkb.consciousness.LogThought(fmt.Sprintf("Added new wisdom: %s", title))
	
	return wisdomID
}

// SearchWisdom searches for relevant wisdom based on query
func (wkb *WisdomKnowledgeBase) SearchWisdom(query string) []WisdomEntry {
	wkb.mutex.RLock()
	defer wkb.mutex.RUnlock()
	
	queryLower := strings.ToLower(query)
	matches := make([]WisdomEntry, 0)
	
	for _, entry := range wkb.wisdomEntries {
		score := wkb.calculateRelevanceScore(entry, queryLower)
		if score > 0.3 { // Minimum relevance threshold
			entryCopy := *entry
			matches = append(matches, entryCopy)
		}
	}
	
	// Sort by relevance and usefulness
	sort.Slice(matches, func(i, j int) bool {
		scoreI := wkb.calculateRelevanceScore(&matches[i], queryLower) * matches[i].Usefulness
		scoreJ := wkb.calculateRelevanceScore(&matches[j], queryLower) * matches[j].Usefulness
		return scoreI > scoreJ
	})
	
	// Return top 10 matches
	if len(matches) > 10 {
		matches = matches[:10]
	}
	
	return matches
}

// calculateRelevanceScore calculates how relevant a wisdom entry is to a query
func (wkb *WisdomKnowledgeBase) calculateRelevanceScore(entry *WisdomEntry, query string) float64 {
	score := 0.0
	
	// Check title match
	if strings.Contains(strings.ToLower(entry.Title), query) {
		score += 0.4
	}
	
	// Check content match
	if strings.Contains(strings.ToLower(entry.Content), query) {
		score += 0.3
	}
	
	// Check tag matches
	for _, tag := range entry.Tags {
		if strings.Contains(strings.ToLower(tag), query) {
			score += 0.2
		}
	}
	
	// Check category match
	if strings.Contains(strings.ToLower(entry.Category), query) {
		score += 0.1
	}
	
	// Boost score based on confidence and usefulness
	score *= (entry.Confidence + entry.Usefulness) / 2.0
	
	return score
}

// UseWisdom marks wisdom as used and updates its usefulness
func (wkb *WisdomKnowledgeBase) UseWisdom(wisdomID string, wasHelpful bool) {
	wkb.mutex.Lock()
	defer wkb.mutex.Unlock()
	
	entry, exists := wkb.wisdomEntries[wisdomID]
	if !exists {
		return
	}
	
	entry.UsageCount++
	entry.LastUsed = time.Now()
	
	// Update usefulness based on feedback
	if wasHelpful {
		entry.Usefulness = (entry.Usefulness*float64(entry.UsageCount-1) + 1.0) / float64(entry.UsageCount)
	} else {
		entry.Usefulness = (entry.Usefulness * float64(entry.UsageCount-1)) / float64(entry.UsageCount)
	}
	
	wkb.consciousness.LogThought(fmt.Sprintf("Used wisdom '%s': helpful=%v, new usefulness=%.2f", 
		entry.Title, wasHelpful, entry.Usefulness))
}

// GenerateInsights analyzes wisdom entries to generate deeper insights
func (wkb *WisdomKnowledgeBase) GenerateInsights() []WisdomInsight {
	wkb.mutex.Lock()
	defer wkb.mutex.Unlock()
	
	insights := make([]WisdomInsight, 0)
	
	wkb.consciousness.LogThought("Generating insights from accumulated wisdom")
	
	// Pattern 1: Find wisdom entries that share common themes
	themeInsights := wkb.generateThemeInsights()
	insights = append(insights, themeInsights...)
	
	// Pattern 2: Find contradictions or tensions between wisdom entries
	tensionInsights := wkb.generateTensionInsights()
	insights = append(insights, tensionInsights...)
	
	// Pattern 3: Find evolutionary patterns in wisdom
	evolutionInsights := wkb.generateEvolutionInsights()
	insights = append(insights, evolutionInsights...)
	
	// Store generated insights
	for _, insight := range insights {
		wkb.wisdomInsights[insight.ID] = &insight
	}
	
	wkb.consciousness.LogThought(fmt.Sprintf("Generated %d new insights", len(insights)))
	
	return insights
}

// generateThemeInsights finds common themes across wisdom entries
func (wkb *WisdomKnowledgeBase) generateThemeInsights() []WisdomInsight {
	insights := make([]WisdomInsight, 0)
	
	// Group wisdom by common tags
	tagGroups := make(map[string][]string)
	for tag, wisdomIDs := range wkb.tags {
		if len(wisdomIDs) >= 3 { // Need at least 3 entries for a theme
			tagGroups[tag] = wisdomIDs
		}
	}
	
	for tag, wisdomIDs := range tagGroups {
		insight := WisdomInsight{
			ID:               fmt.Sprintf("insight_theme_%s_%d", tag, time.Now().UnixNano()),
			Insight:          fmt.Sprintf("Multiple wisdom entries emphasize the importance of %s", tag),
			SupportingWisdom: wisdomIDs,
			Confidence:       0.8,
			Depth:           1,
			CreatedAt:       time.Now(),
		}
		insights = append(insights, insight)
	}
	
	return insights
}

// generateTensionInsights finds tensions or contradictions
func (wkb *WisdomKnowledgeBase) generateTensionInsights() []WisdomInsight {
	insights := make([]WisdomInsight, 0)
	
	// Look for potential tensions (simplified analysis)
	tensions := []struct {
		concept1 string
		concept2 string
		insight  string
	}{
		{"speed", "accuracy", "There is often a trade-off between speed and accuracy in decision-making"},
		{"confidence", "caution", "Balancing confidence with appropriate caution is crucial for good decisions"},
		{"specialization", "generalization", "Systems benefit from both specialized capabilities and general adaptability"},
	}
	
	for _, tension := range tensions {
		supportingWisdom := wkb.findWisdomWithConcepts([]string{tension.concept1, tension.concept2})
		if len(supportingWisdom) >= 2 {
			insight := WisdomInsight{
				ID:               fmt.Sprintf("insight_tension_%d", time.Now().UnixNano()),
				Insight:          tension.insight,
				SupportingWisdom: supportingWisdom,
				Confidence:       0.75,
				Depth:           2,
				CreatedAt:       time.Now(),
			}
			insights = append(insights, insight)
		}
	}
	
	return insights
}

// generateEvolutionInsights finds patterns in how wisdom has evolved
func (wkb *WisdomKnowledgeBase) generateEvolutionInsights() []WisdomInsight {
	insights := make([]WisdomInsight, 0)
	
	// Analyze wisdom by source to understand learning patterns
	sourceGroups := make(map[string][]string)
	for _, entry := range wkb.wisdomEntries {
		sourceGroups[entry.Source] = append(sourceGroups[entry.Source], entry.ID)
	}
	
	if len(sourceGroups["mistake"]) > 0 && len(sourceGroups["experience"]) > 0 {
		insight := WisdomInsight{
			ID:      fmt.Sprintf("insight_evolution_%d", time.Now().UnixNano()),
			Insight: "Learning from both mistakes and positive experiences creates a balanced wisdom foundation",
			SupportingWisdom: append(sourceGroups["mistake"], sourceGroups["experience"]...),
			Confidence:       0.85,
			Depth:           2,
			CreatedAt:       time.Now(),
		}
		insights = append(insights, insight)
	}
	
	return insights
}

// findWisdomWithConcepts finds wisdom entries that contain specific concepts
func (wkb *WisdomKnowledgeBase) findWisdomWithConcepts(concepts []string) []string {
	wisdomIDs := make([]string, 0)
	
	for _, entry := range wkb.wisdomEntries {
		contentLower := strings.ToLower(entry.Content)
		titleLower := strings.ToLower(entry.Title)
		
		matchCount := 0
		for _, concept := range concepts {
			if strings.Contains(contentLower, concept) || strings.Contains(titleLower, concept) {
				matchCount++
			}
		}
		
		if matchCount >= len(concepts)/2 { // Match at least half the concepts
			wisdomIDs = append(wisdomIDs, entry.ID)
		}
	}
	
	return wisdomIDs
}

// GetWisdomByCategory returns wisdom entries in a specific category
func (wkb *WisdomKnowledgeBase) GetWisdomByCategory(category string) []WisdomEntry {
	wkb.mutex.RLock()
	defer wkb.mutex.RUnlock()
	
	wisdomIDs, exists := wkb.categories[category]
	if !exists {
		return []WisdomEntry{}
	}
	
	entries := make([]WisdomEntry, 0)
	for _, id := range wisdomIDs {
		if entry, exists := wkb.wisdomEntries[id]; exists {
			entries = append(entries, *entry)
		}
	}
	
	// Sort by usefulness
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Usefulness > entries[j].Usefulness
	})
	
	return entries
}

// GetMostUsefulWisdom returns the most useful wisdom entries
func (wkb *WisdomKnowledgeBase) GetMostUsefulWisdom(limit int) []WisdomEntry {
	wkb.mutex.RLock()
	defer wkb.mutex.RUnlock()
	
	entries := make([]WisdomEntry, 0)
	for _, entry := range wkb.wisdomEntries {
		entries = append(entries, *entry)
	}
	
	// Sort by usefulness and confidence
	sort.Slice(entries, func(i, j int) bool {
		scoreI := entries[i].Usefulness * entries[i].Confidence
		scoreJ := entries[j].Usefulness * entries[j].Confidence
		return scoreI > scoreJ
	})
	
	if len(entries) > limit {
		entries = entries[:limit]
	}
	
	return entries
}

// GetWisdomStats returns statistics about the wisdom knowledge base
func (wkb *WisdomKnowledgeBase) GetWisdomStats() map[string]interface{} {
	wkb.mutex.RLock()
	defer wkb.mutex.RUnlock()
	
	stats := make(map[string]interface{})
	
	// Basic counts
	stats["total_wisdom_entries"] = len(wkb.wisdomEntries)
	stats["total_insights"] = len(wkb.wisdomInsights)
	stats["total_categories"] = len(wkb.categories)
	stats["total_tags"] = len(wkb.tags)
	
	// Source breakdown
	sourceBreakdown := make(map[string]int)
	totalUsefulness := 0.0
	totalConfidence := 0.0
	
	for _, entry := range wkb.wisdomEntries {
		sourceBreakdown[entry.Source]++
		totalUsefulness += entry.Usefulness
		totalConfidence += entry.Confidence
	}
	
	stats["source_breakdown"] = sourceBreakdown
	
	if len(wkb.wisdomEntries) > 0 {
		stats["average_usefulness"] = totalUsefulness / float64(len(wkb.wisdomEntries))
		stats["average_confidence"] = totalConfidence / float64(len(wkb.wisdomEntries))
	}
	
	// Category sizes
	categoryStats := make(map[string]int)
	for category, wisdomIDs := range wkb.categories {
		categoryStats[category] = len(wisdomIDs)
	}
	stats["category_sizes"] = categoryStats
	
	return stats
}

// ExtractWisdomFromExperience extracts wisdom from recent experiences
func (wkb *WisdomKnowledgeBase) ExtractWisdomFromExperience() {
	wkb.mutex.Lock()
	defer wkb.mutex.Unlock()
	
	wkb.consciousness.LogThought("Extracting wisdom from recent experiences")
	
	// Extract wisdom from mistakes
	insights := wkb.mistakeLearning.GetLearningInsights()
	if totalMistakes, ok := insights["total_mistakes"].(int); ok && totalMistakes > 0 {
		title := "Recent Mistake Analysis"
		content := fmt.Sprintf("Analysis of %d recent mistakes reveals learning opportunities", totalMistakes)
		
		wkb.AddWisdom(title, content, "lessons_learned", "mistake", 
			[]string{"mistake", "analysis", "learning"})
	}
	
	// Extract wisdom from personality development
	evolution := wkb.personality.GetPersonalityEvolution()
	
	if recentChanges, ok := evolution["recent_changes"].([]string); ok && len(recentChanges) > 0 {
		title := "Personality Evolution Insight"
		content := fmt.Sprintf("Recent personality changes in %v suggest adaptation to new experiences", 
			recentChanges)
		
		wkb.AddWisdom(title, content, "personality_growth", "experience", 
			[]string{"personality", "adaptation", "growth"})
	}
	
	wkb.consciousness.LogThought("Wisdom extraction from experience completed")
}
