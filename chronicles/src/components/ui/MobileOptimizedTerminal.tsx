'use client';

import React, { useState, useEffect } from 'react';
import { TerminalWindow } from './TerminalWindow';

interface MobileOptimizedTerminalProps {
  children: React.ReactNode;
  title?: string;
  className?: string;
  expandable?: boolean;
}

export function MobileOptimizedTerminal({ 
  children, 
  title = "terminal", 
  className = "",
  expandable = true 
}: MobileOptimizedTerminalProps) {
  const [isExpanded, setIsExpanded] = useState(false);
  const [isMobile, setIsMobile] = useState(false);

  useEffect(() => {
    const checkMobile = () => {
      setIsMobile(window.innerWidth < 768);
    };
    
    checkMobile();
    window.addEventListener('resize', checkMobile);
    return () => window.removeEventListener('resize', checkMobile);
  }, []);

  if (!isMobile) {
    return (
      <TerminalWindow title={title} className={className}>
        {children}
      </TerminalWindow>
    );
  }

  return (
    <div className={`mobile-terminal ${className}`}>
      {expandable && (
        <button
          onClick={() => setIsExpanded(!isExpanded)}
          className="w-full p-3 bg-bg-secondary border border-border-primary text-text-primary text-left flex items-center justify-between touch-manipulation"
          aria-expanded={isExpanded}
          aria-controls="mobile-terminal-content"
        >
          <span className="text-terminal-cyan">📱 {title}</span>
          <span className="text-terminal-amber">
            {isExpanded ? '▼' : '▶'}
          </span>
        </button>
      )}
      
      <div 
        id="mobile-terminal-content"
        className={`${!expandable || isExpanded ? 'block' : 'hidden'} transition-all duration-300`}
      >
        <TerminalWindow title={expandable ? undefined : title} className="mobile-optimized">
          <div className="mobile-terminal-content">
            {children}
          </div>
        </TerminalWindow>
      </div>
      
      <style jsx>{`
        .mobile-terminal {
          margin: 0.5rem 0;
        }
        
        .mobile-terminal-content {
          font-size: 12px;
          line-height: 1.4;
          overflow-x: auto;
          -webkit-overflow-scrolling: touch;
        }
        
        .mobile-terminal-content pre {
          white-space: pre-wrap;
          word-wrap: break-word;
        }
        
        .mobile-terminal-content code {
          word-break: break-all;
        }
        
        .mobile-optimized {
          padding: 0.75rem;
        }
        
        .touch-manipulation {
          touch-action: manipulation;
          user-select: none;
          -webkit-tap-highlight-color: transparent;
        }
        
        @media (max-width: 640px) {
          .mobile-terminal-content {
            font-size: 11px;
            padding: 0.5rem;
          }
        }
      `}</style>
    </div>
  );
}