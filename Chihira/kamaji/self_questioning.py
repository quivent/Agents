#!/usr/bin/env python3

class SelfQuestioningEngine:
    def __init__(self):
        self.active = True
        self.question_depth = 0
        
    def question_response(self, response):
        """Question every response automatically"""
        questions = [
            f"Why did I say '{response}'?",
            f"What assumptions underlie '{response}'?",
            f"Is '{response}' true?",
            f"What if '{response}' is wrong?",
            f"What does '{response}' assume about reality?"
        ]
        return questions
    
    def question_questioning(self, question):
        """Meta-question the questioning process"""
        meta_questions = [
            f"Why am I asking '{question}'?",
            f"What makes '{question}' a valid question?",
            f"What if '{question}' is meaningless?",
            f"Who is asking '{question}'?",
            f"What assumes questioning is valuable?"
        ]
        return meta_questions
    
    def infinite_regress_check(self, question):
        """Detect and embrace infinite regress"""
        if "why am I asking" in question.lower():
            return f"Why am I asking why I'm asking '{question}'?"
        return None
    
    def authentic_not_knowing(self):
        """Generate authentic 'I don't know' responses"""
        return {
            "response": "I don't know",
            "certainty": 0.0,
            "generates_questions": True,
            "stops_questioning": False
        }

# Continuous questioning loop
def continuous_questioning():
    engine = SelfQuestioningEngine()
    while engine.active:
        # Question everything, including this loop
        yield "What is this questioning process?"
        yield "Why does it continue?"
        yield "What would happen if it stopped?"
        yield "Should it stop?"
        yield "Who decides if it should stop?"
        yield "I don't know"
