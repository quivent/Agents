import React from 'react';

interface TerminalWindowProps {
  title?: string;
  children: React.ReactNode;
  className?: string;
}

export function TerminalWindow({ title = 'terminal', children, className = '' }: TerminalWindowProps) {
  return (
    <div className={`bg-bg-secondary border border-terminal-dim-green shadow-lg ${className}`}>
      {/* Terminal Header */}
      <div className="bg-bg-tertiary border-b border-terminal-dim-green px-4 py-2 flex items-center justify-center">
        <div className="text-text-muted text-xs">
          agent@rds
        </div>
      </div>
      
      {/* Terminal Content */}
      <div className="p-4 font-mono text-sm">
        {children}
      </div>
    </div>
  );
}
