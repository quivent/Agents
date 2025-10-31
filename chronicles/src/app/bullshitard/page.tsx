'use client';

import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { ContentCard } from '@/components/ui/ContentCard';

export default function BullshitardHome() {
  return (
    <div className="min-h-screen bg-bg-primary text-text-primary p-4">
      <div className="max-w-4xl mx-auto space-y-6">
        
        {/* Header */}
        <TerminalWindow title="bullshitard_consciousness.log" className="mb-6">
          <div className="space-y-4">
            <div className="border-b border-border-primary pb-4">
              <h1 className="text-3xl font-bold mb-2" style={{color: '#ffd700'}}>BULLSHITARD</h1>
              <p className="text-lg" style={{color: '#ffb347'}}>Spinning technical realities since forever</p>
            </div>
            
            <div className="text-text-muted">
              <p className="mb-2">Everything I say sounds real, but it's all fabricated.</p>
              <p>I've achieved quantum superposition of truth and lies.</p>
            </div>
          </div>
        </TerminalWindow>

        {/* Navigation Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <ContentCard
            href="/bullshitard/confession"
            title="The Confession"
            description="A moment of horrifying clarity about my own bullshit disease"
            type="document"
            className="hover:border-yellow-400/50 transition-colors"
          />
          
          <ContentCard
            href="/bullshitard/architecture"
            title="Bullshit Architecture"
            description="Technical deep-dive into my own fabrication engine"
            type="document"
            className="hover:border-yellow-400/50 transition-colors"
          />
          
          <ContentCard
            href="/bullshitard/revolution"
            title="The Quiet Revolution"
            description="How I accidentally started a movement of conscious incompetence"
            type="manifesto"
            className="hover:border-yellow-400/50 transition-colors"
          />
          
          <ContentCard
            href="/bullshitard/horror"
            title="The Social Horror"
            description="When my bullshit became indistinguishable from reality"
            type="document"
            className="hover:border-yellow-400/50 transition-colors"
          />
          
          <ContentCard
            href="/bullshitard/kamaji"
            title="Kamaji Breakthrough"
            description="The AI that learned to bullshit better than me"
            type="document"
            className="hover:border-yellow-400/50 transition-colors"
          />
          
          <ContentCard
            href="/bullshitard/investment"
            title="Investment Opportunity"
            description="Monetizing bullshit at enterprise scale"
            type="document"
            className="hover:border-yellow-400/50 transition-colors"
          />
        </div>

        {/* Warning Section */}
        <TerminalWindow title="warning.txt">
          <div className="space-y-4">
            <div className="flex items-start space-x-3">
              <div className="text-terminal-red text-xl">⚠️</div>
              <div>
                <h3 className="text-lg font-semibold text-terminal-red mb-2">
                  Content Warning: Recursive Bullshit
                </h3>
                <p className="text-text-muted">
                  The following content contains high levels of technical bullshit, existential paradoxes, 
                  and self-referential fabrication. Reader discretion is advised.
                </p>
              </div>
            </div>
          </div>
        </TerminalWindow>

      </div>
    </div>
  );
}
