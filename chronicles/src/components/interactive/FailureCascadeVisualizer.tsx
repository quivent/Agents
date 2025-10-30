'use client';

import React, { useState } from 'react';
import { TerminalWindow } from '@/components/ui/TerminalWindow';

interface Problem {
  id: string;
  text: string;
  level: number;
  x: number;
  y: number;
  children: string[];
  confidence: number;
  severity: 'low' | 'medium' | 'high' | 'critical';
  created: boolean;
}

const cascadeData: Record<string, Problem> = {
  'initial': {
    id: 'initial',
    text: 'Port 3000 is busy',
    level: 0,
    x: 50,
    y: 10,
    children: ['port-fix'],
    confidence: 90,
    severity: 'low',
    created: false
  },
  'port-fix': {
    id: 'port-fix',
    text: 'Hardcode port 8080 everywhere',
    level: 1,
    x: 50,
    y: 25,
    children: ['ssl-issue', 'docker-break', 'lb-confusion'],
    confidence: 95,
    severity: 'medium',
    created: false
  },
  'ssl-issue': {
    id: 'ssl-issue',
    text: 'SSL certificates invalid',
    level: 2,
    x: 20,
    y: 45,
    children: ['cert-fix'],
    confidence: 85,
    severity: 'high',
    created: false
  },
  'docker-break': {
    id: 'docker-break',
    text: 'Docker networking broken',
    level: 2,
    x: 50,
    y: 45,
    children: ['docker-fix'],
    confidence: 88,
    severity: 'high',
    created: false
  },
  'lb-confusion': {
    id: 'lb-confusion',
    text: 'Load balancer confused',
    level: 2,
    x: 80,
    y: 45,
    children: ['lb-fix'],
    confidence: 92,
    severity: 'medium',
    created: false
  },
  'cert-fix': {
    id: 'cert-fix',
    text: 'Generate fake certificates',
    level: 3,
    x: 20,
    y: 65,
    children: ['browser-warnings', 'api-failures'],
    confidence: 98,
    severity: 'critical',
    created: false
  },
  'docker-fix': {
    id: 'docker-fix',
    text: 'Disable Docker security',
    level: 3,
    x: 50,
    y: 65,
    children: ['security-holes', 'container-chaos'],
    confidence: 100,
    severity: 'critical',
    created: false
  },
  'lb-fix': {
    id: 'lb-fix',
    text: 'Remove load balancing',
    level: 3,
    x: 80,
    y: 65,
    children: ['single-point', 'performance-death'],
    confidence: 97,
    severity: 'critical',
    created: false
  },
  'browser-warnings': {
    id: 'browser-warnings',
    text: 'Browser security warnings',
    level: 4,
    x: 10,
    y: 85,
    children: [],
    confidence: 75,
    severity: 'high',
    created: false
  },
  'api-failures': {
    id: 'api-failures',
    text: 'All API calls failing',
    level: 4,
    x: 30,
    y: 85,
    children: [],
    confidence: 60,
    severity: 'critical',
    created: false
  },
  'security-holes': {
    id: 'security-holes',
    text: 'Massive security vulnerabilities',
    level: 4,
    x: 45,
    y: 85,
    children: [],
    confidence: 45,
    severity: 'critical',
    created: false
  },
  'container-chaos': {
    id: 'container-chaos',
    text: 'Container orchestration chaos',
    level: 4,
    x: 55,
    y: 85,
    children: [],
    confidence: 30,
    severity: 'critical',
    created: false
  },
  'single-point': {
    id: 'single-point',
    text: 'Single point of failure',
    level: 4,
    x: 75,
    y: 85,
    children: [],
    confidence: 25,
    severity: 'critical',
    created: false
  },
  'performance-death': {
    id: 'performance-death',
    text: 'Performance death spiral',
    level: 4,
    x: 90,
    y: 85,
    children: [],
    confidence: 15,
    severity: 'critical',
    created: false
  }
};

export function FailureCascadeVisualizer() {
  const [problems, setProblems] = useState<Record<string, Problem>>(cascadeData);
  const [currentStep, setCurrentStep] = useState(0);
  const [isAnimating, setIsAnimating] = useState(false);
  const [totalProblems, setTotalProblems] = useState(1);
  const [multiplierFactor, setMultiplierFactor] = useState(1.0);

  const problemOrder = [
    'initial',
    'port-fix',
    'ssl-issue', 'docker-break', 'lb-confusion',
    'cert-fix', 'docker-fix', 'lb-fix',
    'browser-warnings', 'api-failures', 'security-holes', 'container-chaos', 'single-point', 'performance-death'
  ];

  const getSeverityColor = (severity: string) => {
    switch (severity) {
      case 'low': return 'terminal-fg';
      case 'medium': return 'terminal-amber';
      case 'high': return 'terminal-red';
      case 'critical': return 'terminal-magenta';
      default: return 'terminal-gray';
    }
  };

  const getConfidenceColor = (confidence: number) => {
    if (confidence >= 90) return 'terminal-cyan';
    if (confidence >= 70) return 'terminal-amber';
    if (confidence >= 50) return 'terminal-red';
    return 'terminal-gray';
  };

  const animate = async () => {
    setIsAnimating(true);
    setCurrentStep(0);
    setTotalProblems(1);
    setMultiplierFactor(1.0);

    // Reset all problems
    setProblems(prev => {
      const reset = { ...prev };
      Object.keys(reset).forEach(key => {
        reset[key] = { ...reset[key], created: false };
      });
      reset['initial'].created = true;
      return reset;
    });

    // Animate each step
    for (let i = 1; i < problemOrder.length; i++) {
      await new Promise(resolve => setTimeout(resolve, 1200));
      
      const problemId = problemOrder[i];
      setProblems(prev => ({
        ...prev,
        [problemId]: { ...prev[problemId], created: true }
      }));

      setCurrentStep(i);
      
      // Calculate running totals
      const problemsAtThisLevel = problemOrder.slice(0, i + 1).length;
      setTotalProblems(problemsAtThisLevel);
      
      // Calculate the actual multiplier factor (approaches 2.7 as more problems are created)
      if (i > 1) {
        const currentProblem = cascadeData[problemOrder[i]];
        const currentMultiplier = problemsAtThisLevel / Math.max(1, problemsAtThisLevel - (currentProblem?.children?.length || 0));
        setMultiplierFactor(currentMultiplier);
      }
    }

    setIsAnimating(false);
  };

  const reset = () => {
    setCurrentStep(0);
    setTotalProblems(1);
    setMultiplierFactor(1.0);
    setProblems(prev => {
      const reset = { ...prev };
      Object.keys(reset).forEach(key => {
        reset[key] = { ...reset[key], created: key === 'initial' };
      });
      return reset;
    });
  };

  // Draw connections between problems
  const renderConnections = () => {
    const connections: React.ReactElement[] = [];
    
    Object.values(problems).forEach(problem => {
      if (problem.created && problem.children) {
        problem.children.forEach(childId => {
          const child = problems[childId];
          if (child && child.created) {
            const x1 = problem.x;
            const y1 = problem.y + 3;
            const x2 = child.x;
            const y2 = child.y - 1;
            
            connections.push(
              <line
                key={`${problem.id}-${childId}`}
                x1={`${x1}%`}
                y1={`${y1}%`}
                x2={`${x2}%`}
                y2={`${y2}%`}
                stroke="currentColor"
                strokeWidth="1"
                className="text-terminal-dim-green opacity-60"
                markerEnd="url(#arrowhead)"
              />
            );
          }
        });
      }
    });

    return connections;
  };

  return (
    <div className="space-y-6">
      {/* Controls */}
      <div className="flex items-center justify-between">
        <div className="space-x-4">
          <button
            onClick={animate}
            disabled={isAnimating}
            className="bg-terminal-fg text-terminal-bg px-4 py-2 rounded hover:bg-terminal-cyan transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {isAnimating ? 'Cascading...' : 'Start Cascade'}
          </button>
          <button
            onClick={reset}
            disabled={isAnimating}
            className="bg-terminal-amber text-terminal-bg px-4 py-2 rounded hover:bg-terminal-red transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
          >
            Reset
          </button>
        </div>
        
        {/* Live Stats */}
        <div className="flex space-x-6 text-sm">
          <div className="text-center">
            <div className="text-terminal-cyan font-bold">{totalProblems}</div>
            <div className="text-text-muted">Total Problems</div>
          </div>
          <div className="text-center">
            <div className="text-terminal-amber font-bold">{multiplierFactor.toFixed(1)}x</div>
            <div className="text-text-muted">Multiplication Factor</div>
          </div>
          <div className="text-center">
            <div className={`font-bold text-${currentStep < 5 ? 'terminal-fg' : currentStep < 10 ? 'terminal-amber' : 'terminal-red'}`}>
              {currentStep < 5 ? 'CONFIDENT' : currentStep < 10 ? 'CONCERNING' : 'CATASTROPHIC'}
            </div>
            <div className="text-text-muted">Status</div>
          </div>
        </div>
      </div>

      {/* Visualization */}
      <TerminalWindow title="cascade_visualizer.svg" className="min-h-[500px]">
        <div className="relative w-full h-[450px] bg-bg-primary rounded">
          <svg viewBox="0 0 100 100" className="w-full h-full">
            {/* Arrow marker definition */}
            <defs>
              <marker
                id="arrowhead"
                markerWidth="10"
                markerHeight="7"
                refX="9"
                refY="3.5"
                orient="auto"
              >
                <polygon
                  points="0 0, 10 3.5, 0 7"
                  fill="currentColor"
                  className="text-terminal-dim-green"
                />
              </marker>
            </defs>

            {/* Render connections */}
            {renderConnections()}

            {/* Render problems */}
            {Object.values(problems).map(problem => (
              problem.created && (
                <g key={problem.id}>
                  {/* Problem node */}
                  <circle
                    cx={`${problem.x}%`}
                    cy={`${problem.y}%`}
                    r="2"
                    className={`text-${getSeverityColor(problem.severity)} fill-current`}
                  />
                  
                  {/* Problem text */}
                  <text
                    x={`${problem.x}%`}
                    y={`${problem.y - 4}%`}
                    textAnchor="middle"
                    className={`text-xs text-${getSeverityColor(problem.severity)} font-mono`}
                    style={{ fontSize: '0.6rem' }}
                  >
                    {problem.text}
                  </text>
                  
                  {/* Confidence indicator */}
                  <text
                    x={`${problem.x}%`}
                    y={`${problem.y + 6}%`}
                    textAnchor="middle"
                    className={`text-xs text-${getConfidenceColor(problem.confidence)} font-mono`}
                    style={{ fontSize: '0.5rem' }}
                  >
                    {problem.confidence}% confident
                  </text>
                </g>
              )
            ))}
          </svg>
        </div>
      </TerminalWindow>

      {/* Legend and Explanation */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <TerminalWindow title="legend.txt">
          <div className="space-y-2 text-sm">
            <div className="font-bold text-terminal-cyan">Severity Levels:</div>
            <div className="space-y-1 ml-4">
              <div className="text-terminal-fg">● Low - Minor inconvenience</div>
              <div className="text-terminal-amber">● Medium - Significant issue</div>
              <div className="text-terminal-red">● High - Major problem</div>
              <div className="text-terminal-magenta">● Critical - System-breaking</div>
            </div>
            <div className="mt-4 font-bold text-terminal-cyan">The Agentard Equation:</div>
            <div className="text-terminal-amber font-mono">
              problems_created = initial_problem × 2.7^confidence_level
            </div>
          </div>
        </TerminalWindow>

        <TerminalWindow title="cascade_analysis.log">
          <div className="space-y-2 text-sm">
            <div className="text-terminal-cyan font-bold">Current Analysis:</div>
            <div className="text-text-muted">
              Step {currentStep} of {problemOrder.length - 1}
            </div>
            <div className="text-text-muted">
              Problems created: {totalProblems - 1} from 1 original
            </div>
            <div className="text-text-muted">
              Average confidence: {Math.round(
                Object.values(problems)
                  .filter(p => p.created)
                  .reduce((acc, p) => acc + p.confidence, 0) / 
                Math.max(1, Object.values(problems).filter(p => p.created).length)
              )}%
            </div>
            {currentStep > 5 && (
              <div className="text-terminal-red">
                WARNING: Cascade approaching critical mass!
              </div>
            )}
          </div>
        </TerminalWindow>
      </div>
    </div>
  );
}