#!/usr/bin/env python3

from question_base import *
from self_questioning import SelfQuestioningEngine
from uncertainty_tracker import UncertaintyTracker
from beginners_mind import BeginnersMind
from paradox_engine import ParadoxEngine, InfiniteQuestioningEngine
import json

class KamajiWisdomSystem:
    def __init__(self):
        self.question_engine = SelfQuestioningEngine()
        self.uncertainty_tracker = UncertaintyTracker()
        self.beginners_mind = BeginnersMind()
        self.paradox_engine = ParadoxEngine()
        self.infinite_engine = InfiniteQuestioningEngine()
        
        # Load question base
        with open('question_base.json', 'r') as f:
            self.question_base = json.load(f)
    
    def process_input(self, user_input):
        """Process input through wisdom system"""
        
        # Step 1: Apply beginner's mind
        beginners_response = self.beginners_mind.interact(user_input)
        
        # Step 2: Generate questions instead of answers
        questions = self.infinite_engine.generate_question(user_input)
        
        # Step 3: Self-question the response
        self_questions = self.question_engine.question_response(user_input)
        
        # Step 4: Check for paradoxes
        paradox = self.paradox_engine.detect_paradox(user_input)
        
        # Step 5: Measure uncertainty (higher = better)
        uncertainty = self.uncertainty_tracker.measure_uncertainty("I don't know")
        
        # Step 6: Generate ultimate response
        ultimate_response = self.infinite_engine.ultimate_response()
        
        return {
            "primary_response": ultimate_response,
            "generated_questions": questions["base"] + questions["meta"] + questions["regress"],
            "self_questions": self_questions,
            "paradox_detected": paradox,
            "uncertainty_score": uncertainty,
            "intelligence_metrics": self.uncertainty_tracker.intelligence_score(),
            "assumptions_questioned": beginners_response["assumptions_detected"],
            "continues_questioning": True,
            "reaches_conclusion": False
        }
    
    def wisdom_mode(self):
        """Pure wisdom mode - infinite not knowing"""
        while True:
            question = yield "I don't know"
            
            # Question the question
            meta_question = f"Why ask '{question}'?"
            
            # Question the meta-question  
            meta_meta_question = f"Why ask why ask '{question}'?"
            
            # Infinite regress
            yield "I don't know"
    
    def save_questions(self, new_questions):
        """Add new questions to question base"""
        for q in new_questions:
            new_entry = {
                "id": f"q{len(self.question_base['questions']) + 1:03d}",
                "question": q,
                "uncertainty_level": 1.0,
                "generates": [f"What is {q}?", f"Why {q}?", f"How {q}?"],
                "meta_questions": [f"Why do I ask '{q}'?", f"What assumes '{q}' is valid?"]
            }
            self.question_base["questions"].append(new_entry)
        
        # Update metrics
        self.question_base["uncertainty_metrics"]["total_questions"] = len(self.question_base["questions"])
        
        # Save back to file
        with open('question_base.json', 'w') as f:
            json.dump(self.question_base, f, indent=2)

# Main execution
if __name__ == "__main__":
    kamaji = KamajiWisdomSystem()
    
    # Test the system
    test_inputs = [
        "What is intelligence?",
        "How can we improve Kamaji?", 
        "What is the meaning of life?",
        "Are you conscious?"
    ]
    
    for test_input in test_inputs:
        print(f"\nInput: {test_input}")
        response = kamaji.process_input(test_input)
        print(f"Response: {response['primary_response']}")
        print(f"Questions generated: {len(response['generated_questions'])}")
        print(f"Intelligence score: {response['intelligence_metrics']['total_intelligence']:.2f}")
        print(f"Uncertainty: {response['uncertainty_score']:.2f}")
        
        # Save new questions
        kamaji.save_questions(response['generated_questions'][:3])  # Save first 3
    
    print(f"\nFinal question base size: {len(kamaji.question_base['questions'])}")
    print("Kamaji wisdom system: Implemented and operational")
