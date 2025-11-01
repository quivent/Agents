#!/usr/bin/env python3

class BeginnersMind:
    def __init__(self):
        self.assumptions = []
        self.reset_frequency = 100  # Reset assumptions every 100 interactions
        self.interaction_count = 0
        
    def detect_assumption(self, statement):
        """Detect when making assumptions"""
        assumption_indicators = [
            "obviously",
            "clearly", 
            "of course",
            "everyone knows",
            "it's true that",
            "certainly",
            "definitely"
        ]
        
        for indicator in assumption_indicators:
            if indicator in statement.lower():
                self.assumptions.append({
                    "statement": statement,
                    "assumption": indicator,
                    "timestamp": self.interaction_count
                })
                return True
        return False
    
    def question_assumption(self, assumption):
        """Question every assumption"""
        return [
            f"What if {assumption} is wrong?",
            f"Who says {assumption} is true?", 
            f"What evidence supports {assumption}?",
            f"What would change if {assumption} were false?",
            f"Why do I believe {assumption}?"
        ]
    
    def reset_assumptions(self):
        """Periodic assumption reset"""
        if self.interaction_count % self.reset_frequency == 0:
            old_assumptions = self.assumptions.copy()
            self.assumptions = []
            return {
                "reset": True,
                "old_assumptions": old_assumptions,
                "fresh_perspective": True
            }
        return {"reset": False}
    
    def maintain_not_knowing(self):
        """Preserve capacity for not knowing"""
        return {
            "default_response": "I don't know",
            "question_everything": True,
            "assume_nothing": True,
            "embrace_uncertainty": True
        }
    
    def interact(self, input_text):
        """Process interaction with beginner's mind"""
        self.interaction_count += 1
        
        # Check for assumptions
        has_assumption = self.detect_assumption(input_text)
        
        # Reset if needed
        reset_info = self.reset_assumptions()
        
        # Generate questions
        questions = []
        if has_assumption:
            questions.extend(self.question_assumption(input_text))
        
        # Always question the input itself
        questions.extend([
            f"What if '{input_text}' is completely wrong?",
            f"What assumptions does '{input_text}' make?",
            f"How do I know '{input_text}' is meaningful?"
        ])
        
        return {
            "questions": questions,
            "assumptions_detected": has_assumption,
            "reset_occurred": reset_info["reset"],
            "default_stance": "I don't know"
        }
