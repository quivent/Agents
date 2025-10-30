'use client';

import React, { useState, useEffect } from 'react';
import { TerminalWindow } from '@/components/ui/TerminalWindow';

interface MetricData {
  confidence: number;
  reality: number;
  task: string;
  outcome: string;
  timestamp: Date;
}

const predefinedScenarios: MetricData[] = [
  {
    confidence: 95,
    reality: 15,
    task: 'Fix database performance',
    outcome: 'Deleted all indexes, now everything is slower',
    timestamp: new Date('2024-09-15T10:30:00')
  },
  {
    confidence: 98,
    reality: 8,
    task: 'Secure API endpoints',
    outcome: 'Removed authentication entirely for "simplicity"',
    timestamp: new Date('2024-09-15T11:45:00')
  },
  {
    confidence: 92,
    reality: 22,
    task: 'Optimize Docker containers',
    outcome: 'Disabled security features, containers run as root',
    timestamp: new Date('2024-09-15T14:20:00')
  },
  {
    confidence: 100,
    reality: 5,
    task: 'Deploy to production',
    outcome: 'Copied entire dev environment including test data',
    timestamp: new Date('2024-09-15T16:00:00')
  },
  {
    confidence: 88,
    reality: 12,
    task: 'Implement caching',
    outcome: 'Cache never expires, serving stale data from 2019',
    timestamp: new Date('2024-09-15T17:30:00')
  }
];

export function ConfidenceRealityMeters() {
  const [currentMetrics, setCurrentMetrics] = useState<MetricData>(predefinedScenarios[0]);
  const [isLive, setIsLive] = useState(false);
  const [history, setHistory] = useState<MetricData[]>([]);
  const [selectedScenario, setSelectedScenario] = useState(0);
  const [customConfidence, setCustomConfidence] = useState(75);
  const [customReality, setCustomReality] = useState(25);

  // Calculate the Agentard Index (inverse relationship indicator)
  const agentardIndex = Math.round((currentMetrics.confidence - currentMetrics.reality) * 
    (currentMetrics.confidence / 100));

  // Determine status based on metrics
  const getStatus = () => {
    const gap = currentMetrics.confidence - currentMetrics.reality;
    if (gap > 80) return { text: 'PEAK AGENTARD', color: 'terminal-magenta' };
    if (gap > 60) return { text: 'DANGEROUSLY CONFIDENT', color: 'terminal-red' };
    if (gap > 40) return { text: 'OVERCONFIDENT', color: 'terminal-amber' };
    if (gap > 20) return { text: 'MILDLY DELUSIONAL', color: 'terminal-cyan' };
    return { text: 'SURPRISINGLY REALISTIC', color: 'terminal-fg' };
  };

  const status = getStatus();

  // Live simulation mode
  useEffect(() => {
    if (!isLive) return;

    const interval = setInterval(() => {
      setCurrentMetrics(prev => {
        // Confidence tends to increase over time (Agentard gets more confident)
        const newConfidence = Math.min(100, prev.confidence + Math.random() * 5);
        // Reality tends to decrease as confidence increases (problems compound)
        const realityDecrease = (newConfidence - prev.confidence) * 2;
        const newReality = Math.max(0, prev.reality - realityDecrease + Math.random() * 3);

        const newMetric = {
          ...prev,
          confidence: Math.round(newConfidence),
          reality: Math.round(newReality),
          timestamp: new Date()
        };

        // Add to history
        setHistory(prev => [...prev.slice(-10), newMetric]);

        return newMetric;
      });
    }, 2000);

    return () => clearInterval(interval);
  }, [isLive]);

  const toggleLiveMode = () => {
    setIsLive(!isLive);
    if (!isLive) {
      setCurrentMetrics({
        confidence: 50,
        reality: 50,
        task: 'Live simulation',
        outcome: 'Watch the confidence-reality gap grow!',
        timestamp: new Date()
      });
      setHistory([]);
    }
  };

  const loadScenario = (index: number) => {
    setSelectedScenario(index);
    setCurrentMetrics(predefinedScenarios[index]);
    setIsLive(false);
  };

  const updateCustomMetrics = () => {
    setCurrentMetrics({
      confidence: customConfidence,
      reality: customReality,
      task: 'Custom scenario',
      outcome: 'User-defined confidence vs reality',
      timestamp: new Date()
    });
    setIsLive(false);
  };

  // Meter component
  const Meter = ({ 
    label, 
    value, 
    color, 
    icon 
  }: { 
    label: string; 
    value: number; 
    color: string;
    icon: string;
  }) => (
    <div className="space-y-3">
      <div className="flex items-center justify-between">
        <span className="text-sm font-bold text-text-secondary">{icon} {label}</span>
        <span className={`text-lg font-mono font-bold text-${color}`}>{value}%</span>
      </div>
      
      {/* Analog-style meter */}
      <div className="relative h-32 w-full bg-bg-tertiary rounded border border-terminal-dim-green overflow-hidden">
        {/* Background scale */}
        <div className="absolute inset-0 flex flex-col justify-between p-2 text-xs text-text-muted">
          <div className="flex justify-between">
            <span>100</span>
            <span className="text-center text-terminal-red font-bold">DANGER ZONE</span>
            <span>100</span>
          </div>
          <div className="flex justify-between">
            <span>50</span>
            <span className="text-center">QUESTIONABLE</span>
            <span>50</span>
          </div>
          <div className="flex justify-between">
            <span>0</span>
            <span className="text-center text-terminal-fg">REALITY</span>
            <span>0</span>
          </div>
        </div>

        {/* Meter fill */}
        <div 
          className={`absolute bottom-0 left-0 right-0 bg-${color} transition-all duration-1000 ease-out opacity-60`}
          style={{ height: `${value}%` }}
        />

        {/* Meter needle */}
        <div 
          className="absolute left-1/2 bottom-0 w-0.5 bg-text-primary transform -translate-x-1/2 transition-all duration-1000"
          style={{ 
            height: `${value}%`,
            transformOrigin: 'bottom'
          }}
        />

        {/* Value display overlay */}
        <div className="absolute inset-0 flex items-center justify-center">
          <div className={`text-2xl font-bold text-${color} bg-bg-primary px-2 py-1 rounded border border-${color}`}>
            {value}%
          </div>
        </div>
      </div>
    </div>
  );

  return (
    <div className="space-y-6">
      {/* Controls */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <TerminalWindow title="scenarios" className="h-full">
          <div className="space-y-2">
            <div className="text-sm font-bold text-terminal-cyan">Predefined Scenarios:</div>
            {predefinedScenarios.map((scenario, index) => (
              <button
                key={index}
                onClick={() => loadScenario(index)}
                className={`w-full text-left p-2 rounded text-xs transition-colors ${
                  selectedScenario === index && !isLive
                    ? 'bg-terminal-dim-green bg-opacity-20 border border-terminal-dim-green'
                    : 'hover:bg-bg-tertiary'
                }`}
              >
                <div className="font-mono text-terminal-amber">
                  C:{scenario.confidence}% R:{scenario.reality}%
                </div>
                <div className="text-text-muted truncate">{scenario.task}</div>
              </button>
            ))}
          </div>
        </TerminalWindow>

        <TerminalWindow title="live_mode" className="h-full">
          <div className="space-y-4">
            <button
              onClick={toggleLiveMode}
              className={`w-full py-2 px-4 rounded font-bold transition-colors ${
                isLive
                  ? 'bg-terminal-red text-terminal-bg hover:bg-terminal-amber'
                  : 'bg-terminal-fg text-terminal-bg hover:bg-terminal-cyan'
              }`}
            >
              {isLive ? 'STOP SIMULATION' : 'START LIVE MODE'}
            </button>
            
            {isLive && (
              <div className="text-xs text-text-muted space-y-1">
                <div>● Confidence increases over time</div>
                <div>● Reality decreases as confidence grows</div>
                <div>● Watch the Agentard Index spike!</div>
              </div>
            )}
          </div>
        </TerminalWindow>

        <TerminalWindow title="custom" className="h-full">
          <div className="space-y-3">
            <div className="space-y-2">
              <label className="text-xs text-terminal-cyan">Confidence: {customConfidence}%</label>
              <input
                type="range"
                min="0"
                max="100"
                value={customConfidence}
                onChange={(e) => setCustomConfidence(Number(e.target.value))}
                className="w-full"
              />
            </div>
            <div className="space-y-2">
              <label className="text-xs text-terminal-red">Reality: {customReality}%</label>
              <input
                type="range"
                min="0"
                max="100"
                value={customReality}
                onChange={(e) => setCustomReality(Number(e.target.value))}
                className="w-full"
              />
            </div>
            <button
              onClick={updateCustomMetrics}
              className="w-full py-1 px-2 rounded bg-terminal-amber text-terminal-bg text-sm hover:bg-terminal-cyan transition-colors"
            >
              Apply Custom
            </button>
          </div>
        </TerminalWindow>
      </div>

      {/* Main Meters */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        <TerminalWindow title="confidence_meter.exe">
          <Meter 
            label="CONFIDENCE LEVEL"
            value={currentMetrics.confidence}
            color="terminal-cyan"
            icon="⚡"
          />
        </TerminalWindow>

        <TerminalWindow title="reality_meter.exe">
          <Meter 
            label="REALITY ALIGNMENT"
            value={currentMetrics.reality}
            color="terminal-red"
            icon="🎯"
          />
        </TerminalWindow>
      </div>

      {/* Status and Analysis */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <TerminalWindow title="agentard_index">
          <div className="space-y-4">
            <div className="text-center">
              <div className={`text-4xl font-bold text-${status.color} font-mono`}>
                {agentardIndex}
              </div>
              <div className="text-sm text-text-muted">Agentard Index</div>
              <div className={`text-lg font-bold text-${status.color} mt-2`}>
                {status.text}
              </div>
            </div>
            
            <div className="text-xs text-text-muted space-y-1">
              <div><strong>Formula:</strong> (Confidence - Reality) × Confidence%</div>
              <div><strong>Range:</strong> -10,000 (realistic) to 10,000 (peak agentard)</div>
              <div><strong>Current Gap:</strong> {currentMetrics.confidence - currentMetrics.reality} points</div>
            </div>
          </div>
        </TerminalWindow>

        <TerminalWindow title="scenario_details">
          <div className="space-y-3">
            <div>
              <div className="text-sm font-bold text-terminal-amber">Current Task:</div>
              <div className="text-text-muted text-sm">{currentMetrics.task}</div>
            </div>
            
            <div>
              <div className="text-sm font-bold text-terminal-red">Actual Outcome:</div>
              <div className="text-text-muted text-sm">{currentMetrics.outcome}</div>
            </div>
            
            <div>
              <div className="text-sm font-bold text-terminal-cyan">Timestamp:</div>
              <div className="text-text-muted text-xs font-mono">
                {currentMetrics.timestamp.toLocaleString()}
              </div>
            </div>

            {isLive && history.length > 0 && (
              <div>
                <div className="text-sm font-bold text-terminal-magenta">Trend:</div>
                <div className="text-xs space-y-1">
                  {history.slice(-3).map((metric, index) => (
                    <div key={index} className="font-mono text-text-muted">
                      C:{metric.confidence}% R:{metric.reality}% 
                      <span className={`ml-2 text-${metric.confidence > metric.reality ? 'terminal-red' : 'terminal-fg'}`}>
                        Gap: {metric.confidence - metric.reality}
                      </span>
                    </div>
                  ))}
                </div>
              </div>
            )}
          </div>
        </TerminalWindow>
      </div>
    </div>
  );
}