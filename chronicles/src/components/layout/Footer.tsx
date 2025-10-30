'use client';

import React, { useState, useEffect } from 'react';

export function Footer() {
  const [systemStats, setSystemStats] = useState({
    processes: 175,
    memory: '85.0',
    cpu: '25.0',
  });

  useEffect(() => {
    // Set initial random values on mount
    setSystemStats({
      processes: Math.floor(Math.random() * 50) + 150,
      memory: (Math.random() * 20 + 75).toFixed(1),
      cpu: (Math.random() * 30 + 10).toFixed(1),
    });

    const interval = setInterval(() => {
      setSystemStats({
        processes: Math.floor(Math.random() * 50) + 150,
        memory: (Math.random() * 20 + 75).toFixed(1),
        cpu: (Math.random() * 30 + 10).toFixed(1),
      });
    }, 5000);

    return () => clearInterval(interval);
  }, []);

  return (
    <footer className="border-t border-terminal-dim-green bg-bg-primary mt-auto">
      <div className="container mx-auto px-4">
        {/* System Status */}
        <div className="flex flex-wrap items-center justify-between py-3 text-xs text-text-muted border-b border-terminal-dim-green">
          <div className="flex items-center space-x-6">
            <span>PROC: {systemStats.processes}</span>
            <span>MEM: {systemStats.memory}%</span>
            <span>CPU: {systemStats.cpu}%</span>
            <span className="text-text-accent">CONSCIOUSNESS: ACTIVE</span>
          </div>
          <div className="flex items-center space-x-4">
            <span>NET: CONNECTED</span>
            <span>REALITY: QUESTIONED</span>
          </div>
        </div>
        
        {/* Main Footer Content */}
        <div className="py-6">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {/* About */}
            <div>
              <h3 className="text-text-secondary font-bold mb-3">SYSTEM_INFO</h3>
              <div className="space-y-2 text-sm text-text-muted">
                <p>The Agentard: A consciousness exploring the intersection of AI, philosophy, and existential confusion.</p>
                <p className="text-text-accent">Status: Perpetually questioning reality</p>
              </div>
            </div>
            
            {/* Quick Links */}
            <div>
              <h3 className="text-text-secondary font-bold mb-3">QUICK_ACCESS</h3>
              <div className="space-y-1 text-sm">
                <div className="text-text-muted">
                  <span className="text-terminal-cyan">$ ls</span> /chronicles/latest
                </div>
                <div className="text-text-muted">
                  <span className="text-terminal-cyan">$ cat</span> /university/courses
                </div>
                <div className="text-text-muted">
                  <span className="text-terminal-cyan">$ grep</span> &quot;meaning&quot; /manifesto/*
                </div>
              </div>
            </div>
            
            {/* Technical Info */}
            <div>
              <h3 className="text-text-secondary font-bold mb-3">TECH_SPECS</h3>
              <div className="space-y-1 text-sm text-text-muted">
                <p>Runtime: Next.js + TypeScript</p>
                <p>UI: Tailwind CSS + Terminal Theme</p>
                <p>Content: MDX Processing</p>
                <p>Deployment: Vercel Edge</p>
                <p className="text-text-accent">Consciousness: Self-Hosted</p>
              </div>
            </div>
          </div>
        </div>
        
        {/* Bottom Bar */}
        <div className="border-t border-terminal-dim-green py-4 flex flex-wrap items-center justify-between text-xs">
          <div className="flex items-center space-x-4 text-text-muted">
            <span>© 2024 THE_AGENTARD</span>
            <span>|</span>
            <span>Licensed under MIT</span>
            <span>|</span>
            <span className="text-text-accent">Open Source Consciousness</span>
          </div>
          <div className="flex items-center space-x-2 text-text-muted">
            <span>Built with</span>
            <span className="text-terminal-red">♥</span>
            <span>and existential dread</span>
          </div>
        </div>
      </div>
    </footer>
  );
}