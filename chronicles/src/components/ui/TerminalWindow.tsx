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
      <div className="bg-bg-tertiary border-b border-terminal-dim-green px-4 py-2 flex items-center justify-between">
        <div className="flex items-center space-x-2">
          <div className="flex space-x-2">
            <div className="w-3 h-3 rounded-full bg-terminal-red"></div>
            <div className="w-3 h-3 rounded-full bg-terminal-amber"></div>
            <div className="w-3 h-3 rounded-full bg-terminal-fg"></div>
          </div>
          <span className="text-text-muted text-sm ml-4">{title}</span>
        </div>
        <div className="text-text-muted text-xs">
          agentard@consciousness
        </div>
      </div>
      
      {/* Terminal Content */}
      <div className="p-4 font-mono text-sm">
        {children}
      </div>
    </div>
  );
}