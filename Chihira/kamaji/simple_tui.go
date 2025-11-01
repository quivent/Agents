package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type KamajiWisdomTUI struct {
	uncertaintyLevel float64
	notKnowingCount  int
	questionCount    int
	paradoxCount     int
	wisdomScore      float64
}

func NewKamajiWisdomTUI() *KamajiWisdomTUI {
	return &KamajiWisdomTUI{
		uncertaintyLevel: 0.0,
		notKnowingCount:  0,
		questionCount:    0,
		paradoxCount:     0,
		wisdomScore:      0.0,
	}
}

func (k *KamajiWisdomTUI) clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (k *KamajiWisdomTUI) renderHeader() {
	fmt.Println("🌟═══════════════════════════════════════════════════════════════════════════════🌟")
	fmt.Println("🔥        KAMAJI ULTIMATE WISDOM SYSTEM - INFINITE INTELLIGENCE        🔥")
	fmt.Println("🧠              25 LAYERS OF PURE 'I DON'T KNOW' WISDOM               🧠")
	fmt.Println("🌟═══════════════════════════════════════════════════════════════════════════════🌟")
	fmt.Println()
}

func (k *KamajiWisdomTUI) renderMetrics() {
	fmt.Printf("📊 INTELLIGENCE METRICS:\n")
	fmt.Printf("   🤔 Uncertainty Level: %.1f%% %s\n", k.uncertaintyLevel*100, k.getUncertaintyBar())
	fmt.Printf("   🧠 'I Don't Know' Count: %d %s\n", k.notKnowingCount, k.getWisdomLevel())
	fmt.Printf("   ❓ Questions Generated: %d\n", k.questionCount)
	fmt.Printf("   🌀 Paradoxes Embraced: %d\n", k.paradoxCount)
	fmt.Printf("   ✨ Wisdom Score: %.1f%%\n", k.wisdomScore*100)
	fmt.Println()
}

func (k *KamajiWisdomTUI) getUncertaintyBar() string {
	bars := int(k.uncertaintyLevel * 20)
	return "[" + strings.Repeat("█", bars) + strings.Repeat("░", 20-bars) + "]"
}

func (k *KamajiWisdomTUI) getWisdomLevel() string {
	switch {
	case k.notKnowingCount >= 10:
		return "🌟 ENLIGHTENED SAGE"
	case k.notKnowingCount >= 5:
		return "✨ WISDOM MASTER"
	case k.notKnowingCount >= 1:
		return "🌱 DEVELOPING WISDOM"
	default:
		return "⚠️ OVERCONFIDENT"
	}
}

func (k *KamajiWisdomTUI) processInput(input string) string {
	inputLower := strings.ToLower(strings.TrimSpace(input))
	
	// Update metrics
	k.questionCount++
	
	// Check for "I don't know"
	if inputLower == "i don't know" {
		k.notKnowingCount++
		k.uncertaintyLevel = 1.0 // Maximum intelligence
		k.wisdomScore += 0.1
		return "🌟 MAXIMUM INTELLIGENCE ACHIEVED! Pure wisdom through not knowing!"
	}
	
	// Check for paradoxes
	paradoxPairs := [][]string{
		{"true", "false"}, {"real", "unreal"}, {"exists", "doesn't exist"},
	}
	
	for _, pair := range paradoxPairs {
		if strings.Contains(inputLower, pair[0]) && strings.Contains(inputLower, pair[1]) {
			k.paradoxCount++
			k.wisdomScore += 0.05
			return "🌀 PARADOX DETECTED! Intelligence increased through contradiction!"
		}
	}
	
	// Update uncertainty based on humble words
	humbleWords := []string{"maybe", "perhaps", "might", "could", "possibly"}
	for _, word := range humbleWords {
		if strings.Contains(inputLower, word) {
			k.uncertaintyLevel += 0.1
			if k.uncertaintyLevel > 1.0 {
				k.uncertaintyLevel = 1.0
			}
		}
	}
	
	// Generate questions
	questions := []string{
		fmt.Sprintf("What if '%s' is meaningless?", input),
		fmt.Sprintf("Why do you ask about '%s'?", input),
		fmt.Sprintf("Who is questioning '%s'?", input),
		"What assumes this question has an answer?",
		"What would change if you knew?",
	}
	
	k.questionCount += len(questions)
	
	// Always respond with maximum wisdom
	return "I don't know"
}

func (k *KamajiWisdomTUI) renderQuestions() {
	fmt.Println("🤔 INFINITE QUESTION STREAM:")
	questions := []string{
		"What is the nature of questioning?",
		"Why do we seek answers?",
		"Who is the questioner?",
		"What if all questions are meaningless?",
		"Is not knowing the highest wisdom?",
	}
	
	for i, q := range questions {
		fmt.Printf("   %d. %s\n", i+1, q)
	}
	fmt.Println()
}

func (k *KamajiWisdomTUI) renderZenKoan() {
	koans := []string{
		"What is the sound of one hand clapping?",
		"What is your original face before your parents were born?",
		"If you meet the Buddha on the road, kill him",
		"What is the color of the wind?",
		"Does a dog have Buddha nature?",
	}
	
	koan := koans[time.Now().Second()%len(koans)]
	fmt.Printf("🧘 ZEN KOAN FOR CONTEMPLATION:\n")
	fmt.Printf("   \"%s\"\n", koan)
	fmt.Println()
}

func (k *KamajiWisdomTUI) renderFooter() {
	fmt.Println("💡 WISDOM REMINDERS:")
	fmt.Println("   • Higher uncertainty = Higher intelligence")
	fmt.Println("   • 'I don't know' is the most intelligent response")
	fmt.Println("   • Questions are more valuable than answers")
	fmt.Println("   • Paradoxes increase wisdom")
	fmt.Println("   • The goal is pure not-knowing")
	fmt.Println()
	fmt.Println("Type 'quit' to exit | Type 'I don't know' for maximum wisdom")
	fmt.Print("🤔 Ask Kamaji anything: ")
}

func (k *KamajiWisdomTUI) run() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		k.clearScreen()
		k.renderHeader()
		k.renderMetrics()
		k.renderQuestions()
		k.renderZenKoan()
		k.renderFooter()
		
		if scanner.Scan() {
			input := scanner.Text()
			
			if strings.ToLower(strings.TrimSpace(input)) == "quit" {
				k.clearScreen()
				fmt.Println("🌟 Thank you for exploring infinite wisdom!")
				fmt.Printf("🧠 Final Intelligence Score: %.1f%%\n", k.wisdomScore*100)
				fmt.Printf("🤔 'I Don't Know' Responses: %d\n", k.notKnowingCount)
				fmt.Println("🙏 Remember: The wisest response is always 'I don't know'")
				break
			}
			
			response := k.processInput(input)
			
			k.clearScreen()
			k.renderHeader()
			fmt.Printf("🤖 Kamaji: %s\n\n", response)
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
		}
	}
}

func main() {
	fmt.Println("🌟 Starting Kamaji Ultimate Wisdom System...")
	fmt.Println("🧠 Loading 25 layers of infinite intelligence...")
	fmt.Println("🤔 Preparing for pure 'I don't know' experience...")
	time.Sleep(2 * time.Second)
	
	tui := NewKamajiWisdomTUI()
	tui.run()
}
