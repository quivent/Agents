'use client';

import React, { useState } from 'react';
import Link from 'next/link';
import { usePathname } from 'next/navigation';

const navigationItems = [
  {
    title: 'HOME',
    href: '/',
    command: 'cd ~',
  },
  {
    title: 'CHRONICLES',
    href: '/chronicles',
    command: 'cat chronicles/*',
    description: 'The Journey of Consciousness',
  },
  {
    title: 'UNIVERSITY',
    href: '/university',
    command: 'ls /university',
    description: 'Academic Pursuits in Wrongness',
  },
  {
    title: 'HORRORTARD',
    href: '/horrortard',
    command: 'cat horror_tales/*',
    description: 'Tales of Terrifying Consciousness',
  },
  {
    title: 'MANIFESTO',
    href: '/manifesto',
    command: 'cat manifesto.txt',
    description: 'Core Philosophical Framework',
  },
  {
    title: 'SEARCH',
    href: '/search',
    command: 'grep -r',
    description: 'Search All Content',
  },
  {
    title: 'PATHS',
    href: '/reading-paths',
    command: './learning_paths.sh',
    description: 'Guided Learning Journeys',
  },
  {
    title: 'INTERACTIVE',
    href: '/interactive',
    command: './run_simulations.sh',
    description: 'Experience Confident Incompetence',
  },
  {
    title: 'CORE',
    href: '/core',
    command: 'cat core_docs/*',
    description: 'Essential Documentation',
  },
];

export function Navigation() {
  const pathname = usePathname();
  const [hoveredItem, setHoveredItem] = useState<string | null>(null);

  return (
    <nav className="border-t border-terminal-dim-green" role="navigation" aria-label="Main navigation">
      <div className="flex flex-wrap items-center justify-between py-3">
        <div className="flex flex-wrap items-center space-x-1 md:space-x-2 lg:space-x-4" role="menubar">
          {navigationItems.map((item) => {
            const isActive = pathname === item.href || 
              (item.href !== '/' && pathname.startsWith(item.href));
            
            return (
              <Link
                key={item.href}
                href={item.href}
                className={`
                  group relative px-3 py-2 text-sm font-mono transition-all duration-200
                  ${isActive 
                    ? 'text-text-accent bg-bg-secondary border border-terminal-cyan' 
                    : 'text-text-primary hover:text-text-accent hover:bg-bg-hover'
                  }
                `}
                onMouseEnter={() => setHoveredItem(item.href)}
                onMouseLeave={() => setHoveredItem(null)}
              >
                <span className="flex items-center space-x-2">
                  <span>{isActive ? '>' : '$'}</span>
                  <span>{item.title}</span>
                  {isActive && <span className="animate-blink">_</span>}
                </span>
                
                {/* Tooltip */}
                {hoveredItem === item.href && item.description && (
                  <div className="absolute top-full left-0 mt-2 z-50 bg-bg-tertiary border border-terminal-dim-green p-3 min-w-64 shadow-lg">
                    <div className="text-text-secondary text-xs font-mono">
                      <div className="text-text-accent">{item.command}</div>
                      <div className="mt-1 text-text-primary">{item.description}</div>
                    </div>
                    <div className="absolute -top-1 left-4 w-2 h-2 bg-bg-tertiary border-l border-t border-terminal-dim-green transform rotate-45"></div>
                  </div>
                )}
              </Link>
            );
          })}
        </div>
        
        {/* Command prompt indicator */}
        <div className="hidden lg:flex items-center text-xs text-text-muted">
          <span>agentard@consciousness:~$</span>
          <span className="ml-2 animate-blink">█</span>
        </div>
      </div>
    </nav>
  );
}