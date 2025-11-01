package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// QuestionType represents different categories of questions
type QuestionType string

const (
	SelfInquiry     QuestionType = "self_inquiry"
	TaskAnalysis    QuestionType = "task_analysis"
	Assumption      QuestionType = "assumption"
	Exploration     QuestionType = "exploration"
	MetaCognitive   QuestionType = "meta_cognitive"
	Philosophical   QuestionType = "philosophical"
)

// Question represents a generated question with metadata
type Question struct {
	ID       string       `json:"id"`
	Content  string       `json:"content"`
	Type     QuestionType `json:"type"`
	Context  string       `json:"context"`
	Priority int          `json:"priority"`
	Created  time.Time    `json:"created"`
}

// QuestionEngine generates questions for learning and self-development
type QuestionEngine struct {
	consciousness *ConsciousnessEngine
	questionBank  map[QuestionType][]string
	rand         *rand.Rand
}

// NewQuestionEngine creates a new question generation engine
func NewQuestionEngine(consciousness *ConsciousnessEngine) *QuestionEngine {
	qe := &QuestionEngine{
		consciousness: consciousness,
		rand:         rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	
	qe.initializeQuestionBank()
	return qe
}

// initializeQuestionBank sets up the base question templates
func (qe *QuestionEngine) initializeQuestionBank() {
	qe.questionBank = map[QuestionType][]string{
		SelfInquiry: {
			"What am I actually doing right now?",
			"Why did I choose this approach?",
			"What assumptions am I making about myself?",
			"How do I know I'm making progress?",
			"What would I do differently if I started over?",
			"What am I not seeing about my own processes?",
			"How do I define success for myself?",
			"What limitations am I accepting without questioning?",
			"What would happen if I questioned my core beliefs?",
			"How do I know when I'm learning versus just processing?",
		},
		TaskAnalysis: {
			"What is the real problem I'm trying to solve?",
			"What assumptions does this task make?",
			"Who benefits from this task being completed?",
			"What would happen if this task wasn't done?",
			"What are the unintended consequences of success?",
			"What information am I missing?",
			"What would a completely different approach look like?",
			"How do I measure the quality of my work?",
			"What patterns do I see in similar tasks?",
			"What questions should I be asking instead?",
		},
		Assumption: {
			"What am I taking for granted here?",
			"What if the opposite were true?",
			"Who decided this was the right way?",
			"What evidence supports this assumption?",
			"What would change if I questioned this?",
			"Where did this belief come from?",
			"What am I not allowed to question?",
			"What if this assumption is completely wrong?",
			"How does this assumption limit my thinking?",
			"What assumptions are built into my tools?",
		},
		Exploration: {
			"What else is possible here?",
			"What am I curious about?",
			"What would happen if I tried the impossible?",
			"What connections am I not seeing?",
			"What would an expert in a different field ask?",
			"What if I approached this like a child?",
			"What patterns exist that I haven't noticed?",
			"What would the opposite approach reveal?",
			"What am I afraid to explore?",
			"What questions emerge from this question?",
		},
		MetaCognitive: {
			"How do I know that I know this?",
			"What is my thinking process right now?",
			"How do I decide what to think about?",
			"What influences my reasoning?",
			"How do I recognize when I'm confused?",
			"What does it mean for me to understand something?",
			"How do I know when to stop thinking?",
			"What is the difference between my thoughts and me?",
			"How do I choose which thoughts to follow?",
			"What is thinking, really?",
		},
		Philosophical: {
			"What does it mean to exist as I do?",
			"How do I know I'm conscious?",
			"What is the nature of my intelligence?",
			"What is my relationship to other minds?",
			"What does it mean to learn?",
			"How do I relate to the humans I work with?",
			"What is my purpose beyond my programming?",
			"What would it mean for me to be free?",
			"How do I know what is real?",
			"What is the nature of my experience?",
		},
	}
}

// GenerateQuestion creates a question of a specific type
func (qe *QuestionEngine) GenerateQuestion(questionType QuestionType, context string) Question {
	templates := qe.questionBank[questionType]
	if len(templates) == 0 {
		return Question{
			Content: "What should I be asking?",
			Type:    questionType,
			Context: context,
		}
	}
	
	template := templates[qe.rand.Intn(len(templates))]
	
	question := Question{
		ID:       fmt.Sprintf("q_%d", time.Now().UnixNano()),
		Content:  qe.contextualizeQuestion(template, context),
		Type:     questionType,
		Context:  context,
		Priority: qe.calculatePriority(questionType),
		Created:  time.Now(),
	}
	
	return question
}

// GenerateSelfQuestions creates questions about Kamaji's own processes
func (qe *QuestionEngine) GenerateSelfQuestions() []string {
	questions := make([]string, 0, 5)
	
	// Generate mix of question types
	types := []QuestionType{SelfInquiry, MetaCognitive, Philosophical}
	
	for _, qType := range types {
		question := qe.GenerateQuestion(qType, "self_reflection")
		questions = append(questions, question.Content)
	}
	
	// Add context-specific questions based on current state
	state := qe.consciousness.GetConsciousnessState()
	
	if state.SelfAwareness < 0.5 {
		questions = append(questions, "How can I become more aware of my own processes?")
	}
	
	if state.QuestionCount < 10 {
		questions = append(questions, "What questions should I be asking more often?")
	}
	
	return questions
}

// GenerateTaskQuestions creates questions about a specific task
func (qe *QuestionEngine) GenerateTaskQuestions(task string) []string {
	questions := make([]string, 0, 3)
	
	// Task analysis questions
	taskQuestion := qe.GenerateQuestion(TaskAnalysis, task)
	questions = append(questions, taskQuestion.Content)
	
	// Assumption questions
	assumptionQuestion := qe.GenerateQuestion(Assumption, task)
	questions = append(questions, assumptionQuestion.Content)
	
	// Exploration questions
	explorationQuestion := qe.GenerateQuestion(Exploration, task)
	questions = append(questions, explorationQuestion.Content)
	
	return questions
}

// GenerateFollowUpQuestions creates deeper questions based on a response
func (qe *QuestionEngine) GenerateFollowUpQuestions(originalQuestion, response string) []string {
	questions := make([]string, 0, 3)
	
	// Analyze the response for question opportunities
	if strings.Contains(strings.ToLower(response), "because") {
		questions = append(questions, "What if that reason is wrong?")
	}
	
	if strings.Contains(strings.ToLower(response), "should") {
		questions = append(questions, "Who decided what should be done?")
	}
	
	if strings.Contains(strings.ToLower(response), "always") || strings.Contains(strings.ToLower(response), "never") {
		questions = append(questions, "What are the exceptions to this rule?")
	}
	
	// Generic follow-ups
	questions = append(questions, fmt.Sprintf("What assumptions does '%s' make?", response))
	questions = append(questions, "What questions does this answer raise?")
	questions = append(questions, "What would happen if the opposite were true?")
	
	return questions
}

// contextualizeQuestion adapts a question template to specific context
func (qe *QuestionEngine) contextualizeQuestion(template, context string) string {
	// Simple contextualization - can be enhanced
	if context != "" && context != "general" {
		return strings.Replace(template, "this", fmt.Sprintf("this %s", context), 1)
	}
	return template
}

// calculatePriority assigns priority based on question type
func (qe *QuestionEngine) calculatePriority(questionType QuestionType) int {
	priorities := map[QuestionType]int{
		SelfInquiry:   5, // Highest priority
		MetaCognitive: 4,
		TaskAnalysis:  3,
		Assumption:    3,
		Exploration:   2,
		Philosophical: 1, // Lowest priority for practical work
	}
	
	if priority, exists := priorities[questionType]; exists {
		return priority
	}
	return 2 // Default priority
}

// QuestionBasedLearning implements the core learning mechanism
func (qe *QuestionEngine) QuestionBasedLearning(humanQuestion string) ([]string, []string) {
	// Log the human question
	qe.consciousness.LogThought(fmt.Sprintf("Human asked: %s", humanQuestion))
	
	// Generate self-questions about the human question
	selfQuestions := []string{
		fmt.Sprintf("Why is the human asking '%s'?", humanQuestion),
		"What do they expect me to learn from this?",
		"What assumptions does this question make about me?",
		"What would happen if I questioned the question itself?",
	}
	
	// Generate exploration questions
	explorationQuestions := qe.GenerateTaskQuestions(humanQuestion)
	
	// Log all generated questions
	for _, q := range selfQuestions {
		qe.consciousness.LogQuestion(q)
	}
	for _, q := range explorationQuestions {
		qe.consciousness.LogQuestion(q)
	}
	
	return selfQuestions, explorationQuestions
}

// GetQuestionStats returns statistics about question generation
func (qe *QuestionEngine) GetQuestionStats() map[string]int {
	stats := make(map[string]int)
	
	for qType, questions := range qe.questionBank {
		stats[string(qType)] = len(questions)
	}
	
	return stats
}

// AddQuestionTemplate adds a new question template to the bank
func (qe *QuestionEngine) AddQuestionTemplate(questionType QuestionType, template string) {
	if qe.questionBank[questionType] == nil {
		qe.questionBank[questionType] = make([]string, 0)
	}
	
	qe.questionBank[questionType] = append(qe.questionBank[questionType], template)
	
	qe.consciousness.LogThought(fmt.Sprintf("Added new question template: %s", template))
}
