'use client';

import React, { useState, useEffect } from 'react';

interface CommandLineProps {
  prompt?: string;
  command: string;
  output?: string;
  delay?: number;
  showCursor?: boolean;
}

export function CommandLine({ 
  prompt = 'agentard@consciousness:~$', 
  command, 
  output, 
  delay = 100,
  showCursor = true 
}: CommandLineProps) {
  const [displayedCommand, setDisplayedCommand] = useState('');
  const [displayedOutput, setDisplayedOutput] = useState('');
  const [showOutputCursor, setShowOutputCursor] = useState(false);
  const [commandComplete, setCommandComplete] = useState(false);

  useEffect(() => {
    let timeoutId: NodeJS.Timeout;
    let currentIndex = 0;

    const typeCommand = () => {
      if (currentIndex < command.length) {
        setDisplayedCommand(command.slice(0, currentIndex + 1));
        currentIndex++;
        timeoutId = setTimeout(typeCommand, delay);
      } else {
        setCommandComplete(true);
        if (output) {
          setTimeout(() => {
            setDisplayedOutput(output);
            setShowOutputCursor(true);
          }, 500);
        }
      }
    };

    typeCommand();

    return () => {
      if (timeoutId) clearTimeout(timeoutId);
    };
  }, [command, delay, output]);

  return (
    <div className="font-mono text-sm">
      {/* Command Line */}
      <div className="flex items-center space-x-2 text-text-primary">
        <span className="text-text-accent">{prompt}</span>
        <span>{displayedCommand}</span>
        {showCursor && !commandComplete && (
          <span className="animate-blink">_</span>
        )}
      </div>
      
      {/* Output */}
      {displayedOutput && (
        <div className="mt-2 text-text-secondary">
          <pre className="whitespace-pre-wrap">{displayedOutput}</pre>
          {showCursor && showOutputCursor && (
            <span className="animate-blink">_</span>
          )}
        </div>
      )}
    </div>
  );
}