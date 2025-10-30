'use client';

import React from 'react';

export function SkipToContent() {
  return (
    <a
      href="#main-content"
      className="
        sr-only focus:not-sr-only focus:absolute focus:top-4 focus:left-4 
        bg-bg-secondary border border-terminal-cyan text-text-primary 
        px-4 py-2 z-50 font-mono text-sm
        hover:bg-bg-tertiary transition-colors
      "
    >
      Skip to main content
    </a>
  );
}