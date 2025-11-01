#!/usr/bin/env python3

class UncertaintyTracker:
    def __init__(self):
        self.uncertainty_score = 0.0
        self.not_knowing_count = 0
        self.question_depth = 0
        self.paradox_count = 0
        
    def measure_uncertainty(self, response):
        """Higher uncertainty = higher intelligence"""
        uncertainty_indicators = [
            "I don't know",
            "Maybe",
            "Possibly", 
            "It might be",
            "I'm not sure",
            "What if",
            "Could it be"
        ]
        
        score = 0
        for indicator in uncertainty_indicators:
            if indicator.lower() in response.lower():
                score += 0.1
                
        if response.strip() == "I don't know":
            score = 1.0  # Maximum uncertainty
            self.not_knowing_count += 1
            
        self.uncertainty_score = min(1.0, self.uncertainty_score + score)
        return score
    
    def track_question_depth(self, question):
        """Deeper questions = higher intelligence"""
        meta_indicators = [
            "why am I asking",
            "what makes this question",
            "who is questioning",
            "what assumes"
        ]
        
        depth = 0
        for indicator in meta_indicators:
            if indicator in question.lower():
                depth += 1
                
        self.question_depth = max(self.question_depth, depth)
        return depth
    
    def embrace_paradox(self, statement):
        """Paradoxes increase intelligence"""
        paradox_patterns = [
            ("is" in statement and "is not" in statement),
            ("know" in statement and "don't know" in statement),
            ("certain" in statement and "uncertain" in statement)
        ]
        
        if any(paradox_patterns):
            self.paradox_count += 1
            return True
        return False
    
    def intelligence_score(self):
        """Intelligence = uncertainty + depth + paradox tolerance"""
        return {
            "uncertainty": self.uncertainty_score,
            "question_depth": self.question_depth,
            "paradox_tolerance": self.paradox_count,
            "not_knowing_authenticity": self.not_knowing_count,
            "total_intelligence": (
                self.uncertainty_score + 
                (self.question_depth * 0.1) + 
                (self.paradox_count * 0.1) +
                (self.not_knowing_count * 0.1)
            )
        }
