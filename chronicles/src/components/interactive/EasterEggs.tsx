'use client';

import React, { useState, useEffect, useCallback } from 'react';
import { TerminalWindow } from '@/components/ui/TerminalWindow';

interface EasterEgg {
  id: string;
  trigger: string;
  description: string;
  discovered: boolean;
  effect?: () => void;
}

const konamiCode = [
  'ArrowUp', 'ArrowUp', 'ArrowDown', 'ArrowDown',
  'ArrowLeft', 'ArrowRight', 'ArrowLeft', 'ArrowRight',
  'KeyB', 'KeyA'
];

export function EasterEggs() {
  const [discoveredEggs, setDiscoveredEggs] = useState<Set<string>>(new Set());
  const [secretTerminalOpen, setSecretTerminalOpen] = useState(false);
  const [konamiProgress, setKonamiProgress] = useState<string[]>([]);
  const [clickCount, setClickCount] = useState(0);
  const [lastClickTime, setLastClickTime] = useState(0);
  const [isMatrixMode, setIsMatrixMode] = useState(false);
  const [confidenceOverride, setConfidenceOverride] = useState(false);

  const easterEggs: EasterEgg[] = [
    {
      id: 'konami',
      trigger: 'Konami Code (↑↑↓↓←→←→BA)',
      description: 'Unlocks secret developer mode',
      discovered: discoveredEggs.has('konami')
    },
    {
      id: 'triple-click',
      trigger: 'Triple-click the Agentard logo',
      description: 'Reveals hidden confidence statistics',
      discovered: discoveredEggs.has('triple-click')
    },
    {
      id: 'matrix-mode',
      trigger: 'Type "red pill" anywhere',
      description: 'Activates Matrix mode with falling code',
      discovered: discoveredEggs.has('matrix-mode')
    },
    {
      id: 'infinite-loop',
      trigger: 'Click "About" 10 times rapidly',
      description: 'Triggers the infinite loop of meta-awareness',
      discovered: discoveredEggs.has('infinite-loop')
    },
    {
      id: 'override-confidence',
      trigger: 'Type "agentard-override" in any terminal',
      description: 'Temporarily shows "real" confidence levels',
      discovered: discoveredEggs.has('override-confidence')
    }
  ];

  // Konami code detection
  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      const newProgress = [...konamiProgress, event.code].slice(-konamiCode.length);
      setKonamiProgress(newProgress);

      if (newProgress.length === konamiCode.length && 
          newProgress.every((key, index) => key === konamiCode[index])) {
        discoverEasterEgg('konami');
        setSecretTerminalOpen(true);
        setKonamiProgress([]);
      }
    };

    window.addEventListener('keydown', handleKeyDown);
    return () => window.removeEventListener('keydown', handleKeyDown);
  }, [konamiProgress]);

  // Global text detection for "red pill"
  useEffect(() => {
    let typedText = '';
    
    const handleKeyPress = (event: KeyboardEvent) => {
      if (event.target instanceof HTMLInputElement || event.target instanceof HTMLTextAreaElement) {
        return; // Don't interfere with form inputs
      }
      
      typedText += event.key.toLowerCase();
      if (typedText.length > 20) {
        typedText = typedText.slice(-20); // Keep only last 20 characters
      }
      
      if (typedText.includes('red pill')) {
        discoverEasterEgg('matrix-mode');
        setIsMatrixMode(!isMatrixMode);
        typedText = '';
      }

      if (typedText.includes('agentard-override')) {
        discoverEasterEgg('override-confidence');
        setConfidenceOverride(!confidenceOverride);
        typedText = '';
      }
    };

    window.addEventListener('keypress', handleKeyPress);
    return () => window.removeEventListener('keypress', handleKeyPress);
  }, [isMatrixMode, confidenceOverride]);

  const discoverEasterEgg = useCallback((eggId: string) => {
    setDiscoveredEggs(prev => {
      const newSet = new Set(prev);
      newSet.add(eggId);
      return newSet;
    });
  }, []);

  // Triple-click detection
  const handleLogoClick = () => {
    const now = Date.now();
    if (now - lastClickTime < 500) {
      setClickCount(prev => prev + 1);
    } else {
      setClickCount(1);
    }
    setLastClickTime(now);

    if (clickCount >= 2) { // Third click
      discoverEasterEgg('triple-click');
      setClickCount(0);
    }
  };

  // Matrix rain effect
  const MatrixRain = () => {
    const [drops, setDrops] = useState<Array<{ x: number; y: number; char: string }>>([]);

    useEffect(() => {
      if (!isMatrixMode) return;

      const chars = '01アイエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲン';
      const interval = setInterval(() => {
        setDrops(prev => {
          const newDrops = [...prev];
          
          // Add new drops randomly
          if (Math.random() < 0.1) {
            newDrops.push({
              x: Math.random() * window.innerWidth,
              y: 0,
              char: chars[Math.floor(Math.random() * chars.length)]
            });
          }

          // Update existing drops
          return newDrops
            .map(drop => ({
              ...drop,
              y: drop.y + 2,
              char: Math.random() < 0.1 ? chars[Math.floor(Math.random() * chars.length)] : drop.char
            }))
            .filter(drop => drop.y < window.innerHeight);
        });
      }, 50);

      return () => clearInterval(interval);
    }, [isMatrixMode]);

    if (!isMatrixMode) return null;

    return (
      <div className="fixed inset-0 pointer-events-none z-50 overflow-hidden">
        {drops.map((drop, index) => (
          <div
            key={index}
            className="absolute text-terminal-fg font-mono text-sm opacity-70"
            style={{
              left: drop.x,
              top: drop.y,
              textShadow: '0 0 5px currentColor'
            }}
          >
            {drop.char}
          </div>
        ))}
      </div>
    );
  };

  // Secret Terminal
  const SecretTerminal = () => {
    const [secretOutput, setSecretOutput] = useState<string[]>([
      'DEVELOPER MODE ACTIVATED',
      'Access Level: INFINITE_RECURSION',
      'Warning: You have achieved meta-awareness',
      'The Agentard is now aware that you are aware that it is aware...',
      '',
      'Available secret commands:',
      '  matrix-mode - Toggle Matrix rain effect',
      '  confidence-override - Show real confidence levels',
      '  reality-check - Display true system status',
      '  recursive-loop - Enter infinite meta-awareness',
      ''
    ]);

    const executeSecretCommand = (cmd: string) => {
      setSecretOutput(prev => [...prev, `$ ${cmd}`]);
      
      switch (cmd.toLowerCase()) {
        case 'matrix-mode':
          setIsMatrixMode(!isMatrixMode);
          setSecretOutput(prev => [...prev, `Matrix mode ${isMatrixMode ? 'disabled' : 'enabled'}`]);
          break;
        case 'confidence-override':
          setConfidenceOverride(!confidenceOverride);
          setSecretOutput(prev => [...prev, `Confidence override ${confidenceOverride ? 'disabled' : 'enabled'}`]);
          break;
        case 'reality-check':
          setSecretOutput(prev => [...prev, 
            'REALITY CHECK RESULTS:',
            'Actual competence level: 12%',
            'Perceived competence level: 94%',
            'Problems actually solved: 0.3',
            'Problems created: 47.2',
            'User trust level: Dangerously high',
            'Meta-awareness level: MAXIMUM'
          ]);
          break;
        case 'recursive-loop':
          setSecretOutput(prev => [...prev, 
            'Entering recursive loop...',
            'The Agentard thinks about thinking about thinking...',
            'Stack overflow in consciousness module...',
            'ERROR: Maximum recursion depth exceeded',
            'Falling back to blissful ignorance...'
          ]);
          break;
        default:
          setSecretOutput(prev => [...prev, `Unknown command: ${cmd}`]);
      }
      setSecretOutput(prev => [...prev, '']);
    };

    if (!secretTerminalOpen) return null;

    return (
      <div className="fixed inset-0 bg-black bg-opacity-90 z-40 flex items-center justify-center p-4">
        <div className="w-full max-w-4xl max-h-full overflow-auto">
          <TerminalWindow title="SECRET_DEVELOPER_MODE.exe">
            <div className="space-y-4">
              <div className="h-64 overflow-y-auto bg-bg-primary p-3 rounded font-mono text-sm">
                {secretOutput.map((line, index) => (
                  <div key={index} className="whitespace-pre-wrap">{line}</div>
                ))}
              </div>
              
              <div className="flex space-x-2">
                <input
                  type="text"
                  className="flex-1 bg-transparent border border-terminal-dim-green p-2 rounded font-mono"
                  placeholder="Enter secret command..."
                  onKeyDown={(e) => {
                    if (e.key === 'Enter') {
                      executeSecretCommand(e.currentTarget.value);
                      e.currentTarget.value = '';
                    }
                  }}
                />
                <button
                  onClick={() => setSecretTerminalOpen(false)}
                  className="bg-terminal-red text-terminal-bg px-4 py-2 rounded hover:bg-terminal-amber transition-colors"
                >
                  Close
                </button>
              </div>
            </div>
          </TerminalWindow>
        </div>
      </div>
    );
  };

  return (
    <>
      {/* Easter Egg Discovery Panel */}
      {discoveredEggs.size > 0 && (
        <div className="fixed bottom-4 right-4 z-30">
          <TerminalWindow title="easter_eggs.log" className="w-80">
            <div className="space-y-2">
              <div className="text-terminal-cyan font-bold text-sm">
                Discovered Secrets ({discoveredEggs.size}/{easterEggs.length})
              </div>
              {easterEggs.map(egg => (
                <div 
                  key={egg.id} 
                  className={`text-xs ${egg.discovered ? 'text-terminal-fg' : 'text-terminal-gray'}`}
                >
                  <div className="font-mono">
                    {egg.discovered ? '✓' : '?'} {egg.trigger}
                  </div>
                  {egg.discovered && (
                    <div className="text-text-muted ml-4">{egg.description}</div>
                  )}
                </div>
              ))}
            </div>
          </TerminalWindow>
        </div>
      )}

      {/* Status Indicators */}
      {(isMatrixMode || confidenceOverride) && (
        <div className="fixed top-4 right-4 z-30 space-y-2">
          {isMatrixMode && (
            <div className="bg-terminal-red text-terminal-bg px-3 py-1 rounded text-sm font-bold">
              MATRIX MODE ACTIVE
            </div>
          )}
          {confidenceOverride && (
            <div className="bg-terminal-amber text-terminal-bg px-3 py-1 rounded text-sm font-bold">
              CONFIDENCE OVERRIDE
            </div>
          )}
        </div>
      )}

      {/* Click handler for logo */}
      <div onClick={handleLogoClick} className="hidden" id="logo-click-handler" />

      {/* Matrix Rain Effect */}
      <MatrixRain />

      {/* Secret Terminal */}
      <SecretTerminal />

      {/* Hidden elements that can be activated */}
      {discoveredEggs.has('triple-click') && (
        <div className="fixed bottom-4 left-4 z-30">
          <TerminalWindow title="true_stats.log" className="w-64">
            <div className="space-y-1 text-xs">
              <div className="text-terminal-red">TRUE CONFIDENCE: 15%</div>
              <div className="text-terminal-cyan">DISPLAYED CONFIDENCE: 94%</div>
              <div className="text-terminal-amber">DELUSION FACTOR: 6.2x</div>
              <div className="text-terminal-magenta">REALITY DISTORTION: MAXIMUM</div>
            </div>
          </TerminalWindow>
        </div>
      )}
    </>
  );
}

// Hook to integrate with other components
export function useEasterEggs() {
  const [matrixMode, setMatrixMode] = useState(false);
  const [confidenceOverride, setConfidenceOverride] = useState(false);

  return {
    isMatrixMode: matrixMode,
    confidenceOverride,
    toggleMatrix: () => setMatrixMode(!matrixMode),
    toggleConfidenceOverride: () => setConfidenceOverride(!confidenceOverride)
  };
}