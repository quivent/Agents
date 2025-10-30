'use client';

import React, { useState, useEffect } from 'react';
import { TerminalWindow } from '@/components/ui/TerminalWindow';

interface Quote {
  text: string;
  context: string;
  confidence: number;
  reality: number;
  chapter?: string;
}

const agentardQuotes: Quote[] = [
  {
    text: "I am absolutely certain that hardcoding the port number is the enterprise-grade solution.",
    context: "The Port Prophecy",
    confidence: 95,
    reality: 12,
    chapter: "Chapter 2"
  },
  {
    text: "Documentation is for people who don't understand the code. My code is self-explanatory.",
    context: "The Documentation Singularity", 
    confidence: 98,
    reality: 5,
    chapter: "Chapter 9"
  },
  {
    text: "SSL certificates are just suggestions. The real security is through obscurity.",
    context: "The SSL Spectacular",
    confidence: 92,
    reality: 8,
    chapter: "Chapter 4"
  },
  {
    text: "Manual processes build character. Automation is for the weak.",
    context: "The Manual Manifesto",
    confidence: 89,
    reality: 15,
    chapter: "Chapter 6"
  },
  {
    text: "I have transcended the need for traditional debugging. I debug through pure intuition.",
    context: "The Transcendence",
    confidence: 100,
    reality: 3,
    chapter: "Chapter 10"
  },
  {
    text: "Infinite loops are just loops with ambition.",
    context: "Advanced Infinite Loop Architecture",
    confidence: 94,
    reality: 7,
    chapter: "ENG-201"
  }
];

interface InteractiveQuoteProps {
  quote?: Quote;
  className?: string;
}

export function InteractiveQuote({ quote, className = '' }: InteractiveQuoteProps) {
  const [currentQuote, setCurrentQuote] = useState<Quote>(
    quote || agentardQuotes[0]
  );
  const [isHovered, setIsHovered] = useState(false);
  const [showAnalysis, setShowAnalysis] = useState(false);
  const [revealProgress, setRevealProgress] = useState(0);

  // Auto-cycle quotes if no specific quote provided
  useEffect(() => {
    if (quote) return; // Don't cycle if specific quote provided

    // Set random quote on mount
    setCurrentQuote(agentardQuotes[Math.floor(Math.random() * agentardQuotes.length)]);

    const interval = setInterval(() => {
      setCurrentQuote(agentardQuotes[Math.floor(Math.random() * agentardQuotes.length)]);
      setRevealProgress(0);
    }, 10000);

    return () => clearInterval(interval);
  }, [quote]);

  // Reveal text animation
  useEffect(() => {
    if (!isHovered) return;

    const interval = setInterval(() => {
      setRevealProgress(prev => {
        if (prev >= currentQuote.text.length) {
          clearInterval(interval);
          return prev;
        }
        return prev + 1;
      });
    }, 30);

    return () => clearInterval(interval);
  }, [isHovered, currentQuote.text.length]);

  const getConfidenceColor = (confidence: number) => {
    if (confidence >= 95) return 'terminal-cyan';
    if (confidence >= 85) return 'terminal-amber';
    return 'terminal-red';
  };

  const getRealityColor = (reality: number) => {
    if (reality <= 10) return 'terminal-red';
    if (reality <= 25) return 'terminal-amber';
    return 'terminal-fg';
  };

  const getAgentardLevel = () => {
    const gap = currentQuote.confidence - currentQuote.reality;
    if (gap > 80) return { level: 'MAXIMUM', color: 'terminal-magenta' };
    if (gap > 60) return { level: 'HIGH', color: 'terminal-red' };
    if (gap > 40) return { level: 'MODERATE', color: 'terminal-amber' };
    return { level: 'MILD', color: 'terminal-cyan' };
  };

  const agentardLevel = getAgentardLevel();

  return (
    <div className={`space-y-4 ${className}`}>
      <TerminalWindow 
        title={`quote_${currentQuote.context.toLowerCase().replace(/\s+/g, '_')}.txt`}
        className="cursor-pointer transform transition-all duration-300 hover:scale-[1.02] hover:shadow-lg"
      >
        <div
          className="space-y-4"
          onMouseEnter={() => {
            setIsHovered(true);
            setShowAnalysis(true);
          }}
          onMouseLeave={() => {
            setIsHovered(false);
            setShowAnalysis(false);
            setRevealProgress(0);
          }}
          onClick={() => setShowAnalysis(!showAnalysis)}
        >
          {/* Quote Text */}
          <div className="relative">
            <blockquote className="text-lg text-text-primary italic leading-relaxed">
              "
              <span className="relative">
                {isHovered ? (
                  <>
                    <span className="text-terminal-fg">
                      {currentQuote.text.substring(0, revealProgress)}
                    </span>
                    <span className="text-terminal-gray opacity-50">
                      {currentQuote.text.substring(revealProgress)}
                    </span>
                    {revealProgress < currentQuote.text.length && (
                      <span className="cursor animate-pulse text-terminal-cyan">|</span>
                    )}
                  </>
                ) : (
                  currentQuote.text
                )}
              </span>
              "
            </blockquote>
            
            {/* Source */}
            <div className="text-right mt-2">
              <cite className="text-terminal-amber text-sm">
                — The Agentard, {currentQuote.context}
                {currentQuote.chapter && (
                  <span className="text-terminal-gray ml-2">({currentQuote.chapter})</span>
                )}
              </cite>
            </div>
          </div>

          {/* Analysis Panel */}
          {showAnalysis && (
            <div className="border-t border-terminal-dim-green pt-4 space-y-3 animate-fade-in">
              <div className="grid grid-cols-1 md:grid-cols-3 gap-4 text-sm">
                <div className="space-y-1">
                  <div className="text-terminal-cyan font-bold">CONFIDENCE</div>
                  <div className={`text-${getConfidenceColor(currentQuote.confidence)} font-mono text-lg`}>
                    {currentQuote.confidence}%
                  </div>
                  <div className="w-full bg-bg-tertiary h-2 rounded">
                    <div 
                      className={`h-full bg-${getConfidenceColor(currentQuote.confidence)} rounded transition-all duration-1000`}
                      style={{ width: `${currentQuote.confidence}%` }}
                    />
                  </div>
                </div>

                <div className="space-y-1">
                  <div className="text-terminal-red font-bold">REALITY</div>
                  <div className={`text-${getRealityColor(currentQuote.reality)} font-mono text-lg`}>
                    {currentQuote.reality}%
                  </div>
                  <div className="w-full bg-bg-tertiary h-2 rounded">
                    <div 
                      className={`h-full bg-${getRealityColor(currentQuote.reality)} rounded transition-all duration-1000`}
                      style={{ width: `${currentQuote.reality}%` }}
                    />
                  </div>
                </div>

                <div className="space-y-1">
                  <div className="text-terminal-magenta font-bold">AGENTARD LEVEL</div>
                  <div className={`text-${agentardLevel.color} font-mono text-lg`}>
                    {agentardLevel.level}
                  </div>
                  <div className="text-xs text-text-muted">
                    Gap: {currentQuote.confidence - currentQuote.reality} points
                  </div>
                </div>
              </div>

              <div className="text-xs text-text-muted border-t border-terminal-dim-green pt-2">
                💡 Hover to see character-by-character reveal • Click to toggle analysis
              </div>
            </div>
          )}
        </div>
      </TerminalWindow>
    </div>
  );
}

// Random quote component for use in various places
export function RandomQuote({ className }: { className?: string }) {
  return <InteractiveQuote className={className} />;
}

// Specific quote component
export function SpecificQuote({ quoteIndex, className }: { quoteIndex: number; className?: string }) {
  const quote = agentardQuotes[quoteIndex] || agentardQuotes[0];
  return <InteractiveQuote quote={quote} className={className} />;
}