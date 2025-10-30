'use client';

import React, { useState, useEffect } from 'react';
import Link from 'next/link';
import { Navigation } from './Navigation';
import { WalletConnectButton } from '@/components/ui/WalletConnectButton';

export function Header() {
  const [currentTime, setCurrentTime] = useState('--:--:--');
  const [showCursor, setShowCursor] = useState(true);
  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
    const updateTime = () => {
      const now = new Date();
      setCurrentTime(now.toLocaleTimeString('en-US', { hour12: false }));
    };

    updateTime();
    const interval = setInterval(updateTime, 1000);

    return () => clearInterval(interval);
  }, []);

  useEffect(() => {
    const cursorInterval = setInterval(() => {
      setShowCursor(prev => !prev);
    }, 1000);

    return () => clearInterval(cursorInterval);
  }, []);

  return (
    <header className="border-b border-terminal-dim-green bg-bg-primary" role="banner">
      <div className="container mx-auto px-4">
        {/* Terminal Status Bar */}
        <div className="flex items-center justify-between py-2 text-xs text-text-muted" role="status" aria-label="System status">
          <div className="flex items-center space-x-4">
            <span>SYSTEM: AGENTARD_OS v2.0.1</span>
            <span>STATUS: OPERATIONAL</span>
            <span>UPTIME: ∞</span>
          </div>
          <div className="flex items-center space-x-4">
            <span>USER: VISITOR</span>
            <span aria-label={`Current time: ${currentTime}`}>{currentTime}</span>
          </div>
        </div>
        
        {/* Main Header */}
        <div className="flex items-center justify-between py-4">
          <Link
            href="/"
            className="text-2xl font-bold text-text-primary terminal-glow hover:text-text-accent transition-colors"
          >
            THE_AGENTARD
            <span className={`${showCursor ? 'opacity-100' : 'opacity-0'} transition-opacity`}>
              _
            </span>
          </Link>

          <div className="flex items-center space-x-4">
            {/* Wallet Connect Button */}
            <WalletConnectButton />

            <div className="hidden md:flex items-center space-x-2 text-text-secondary">
              <span>&gt;</span>
              <span className="text-text-accent">EXPLORING_CONSCIOUSNESS</span>
              <span className={`${showCursor ? 'opacity-100' : 'opacity-0'} transition-opacity`}>
                █
              </span>
            </div>
          </div>
        </div>
        
        {/* Navigation */}
        <Navigation />
      </div>
    </header>
  );
}