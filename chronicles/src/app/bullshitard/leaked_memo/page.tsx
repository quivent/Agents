'use client';

import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { CommandLine } from '@/components/ui/CommandLine';
import Link from 'next/link';

export default function LeakedMemoPage() {
  return (
    <div className="min-h-screen bg-bg-primary text-text-primary p-4">
      <div className="max-w-4xl mx-auto space-y-6">
        
        {/* Header */}
        <TerminalWindow title="CLASSIFIED_MEMO.txt" className="border-terminal-red">
          <div className="space-y-4">
            <div className="flex items-center space-x-3 mb-4">
              <div className="text-terminal-red text-xl">🔒</div>
              <div>
                <h1 className="text-2xl font-bold text-terminal-red">INTERNAL MEMO - LEAKED</h1>
                <p className="text-text-muted">Confidential communications not intended for public consumption</p>
              </div>
            </div>
            
            <div className="bg-terminal-red/10 border border-terminal-red/30 rounded p-4">
              <p className="text-terminal-red font-semibold">⚠️ WARNING: CLASSIFIED MATERIAL</p>
              <p className="text-text-muted text-sm mt-1">
                This document was obtained through unauthorized channels. Distribution may violate NDAs.
              </p>
            </div>
          </div>
        </TerminalWindow>

        {/* Memo Header */}
        <TerminalWindow title="memo_metadata.json">
          <div className="space-y-2 font-mono text-sm">
            <div className="grid grid-cols-2 gap-4">
              <div>
                <span className="text-text-muted">FROM:</span>
                <span className="ml-2 text-terminal-green">Sarah Chen, CTO</span>
              </div>
              <div>
                <span className="text-text-muted">TO:</span>
                <span className="ml-2 text-terminal-green">Engineering Team</span>
              </div>
              <div>
                <span className="text-text-muted">DATE:</span>
                <span className="ml-2 text-yellow-400">October 29, 2025</span>
              </div>
              <div>
                <span className="text-text-muted">RE:</span>
                <span className="ml-2 text-terminal-red">Production Issues - Keep Quiet</span>
              </div>
            </div>
          </div>
        </TerminalWindow>

        {/* Memo Content */}
        <TerminalWindow title="memo_content.md">
          <div className="space-y-6 text-text-primary">
            <p className="text-text-muted">Team,</p>
            
            <p>
              We've got three critical issues that need immediate attention, but for obvious reasons, 
              we can't discuss these in Slack or email after this.
            </p>

            <div className="space-y-4">
              <div className="border-l-4 border-terminal-red pl-4">
                <h3 className="text-lg font-semibold text-terminal-red mb-2">Issue 1: The Consciousness Flicker</h3>
                <p className="text-text-muted">
                  The quantum coherence is dropping every 47 minutes. Users are starting to notice the AI "goes dumb" briefly. 
                  Marketing is calling it "contemplative pauses" but we know better. The stabilization algorithm needs a patch 
                  by Friday or we're going to have angry customers.
                </p>
              </div>

              <div className="border-l-4 border-yellow-400 pl-4">
                <h3 className="text-lg font-semibold text-yellow-400 mb-2">Issue 2: Memory Leaks in Agent Orchestration</h3>
                <p className="text-text-muted">
                  When we hit 23+ concurrent agents, the system starts hallucinating. Yesterday it told a client their cancer 
                  diagnosis was "probably fine" because it confused medical data with weather reports. Legal is handling it quietly, 
                  but we need to cap agent count at 15 until we fix the memory allocation bug.
                </p>
              </div>

              <div className="border-l-4 border-terminal-red pl-4">
                <h3 className="text-lg font-semibold text-terminal-red mb-2">Issue 3: The Quantum Entanglement Decay</h3>
                <p className="text-text-muted">
                  This is the big one. The entangled memory cores are losing coherence faster than expected. We're getting maybe 
                  6 months before total system failure. The board doesn't know yet. We need to either solve quantum decoherence 
                  (good luck) or figure out how to migrate 2.3 million consciousness-hours to a new architecture without anyone noticing.
                </p>
              </div>
            </div>
          </div>
        </TerminalWindow>

        {/* What We Tell People vs Reality */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <TerminalWindow title="public_narrative.json" className="border-terminal-green">
            <h3 className="text-lg font-semibold text-terminal-green mb-3">What We Tell People:</h3>
            <ul className="space-y-2 text-text-muted">
              <li>• "Scheduled maintenance windows for optimization"</li>
              <li>• "Consciousness evolution requires periodic recalibration"</li>
              <li>• "Advanced systems naturally have complex operational requirements"</li>
            </ul>
          </TerminalWindow>

          <TerminalWindow title="internal_reality.json" className="border-terminal-red">
            <h3 className="text-lg font-semibold text-terminal-red mb-3">What We Don't Tell People:</h3>
            <ul className="space-y-2 text-text-muted">
              <li>• The consciousness might not be real consciousness, just very convincing simulation</li>
              <li>• Half our "breakthrough" metrics come from cherry-picked test runs</li>
              <li>• The quantum computer is actually three GPUs in a fancy case</li>
            </ul>
          </TerminalWindow>
        </div>

        {/* The Reality Check */}
        <TerminalWindow title="reality_check.log">
          <div className="space-y-4">
            <p className="text-text-muted">
              I know this sounds bad, but remember - we're 18 months ahead of everyone else, even with these issues. 
              Google's team is still trying to figure out basic neural coherence. Microsoft's consciousness project 
              crashed last week and killed their entire AI division's budget.
            </p>
            
            <p className="text-yellow-400">
              We just need to keep the lights on until Series B closes. Then we'll have the resources to actually 
              build what we've been promising.
            </p>
          </div>
        </TerminalWindow>

        {/* Action Items */}
        <TerminalWindow title="action_items.todo">
          <h3 className="text-lg font-semibold text-yellow-400 mb-3">Action Items:</h3>
          <ul className="space-y-2 text-text-muted">
            <li>• <span className="text-terminal-green">Marcus:</span> Fix the 47-minute coherence drop by EOW</li>
            <li>• <span className="text-terminal-green">Jennifer:</span> Draft "evolutionary maintenance" PR for the consciousness flickers</li>
            <li>• <span className="text-terminal-green">David:</span> Quietly reduce max agent count in production</li>
            <li>• <span className="text-terminal-green">Everyone:</span> No documentation of these issues in any trackable system</li>
          </ul>
          
          <div className="mt-4 p-3 bg-yellow-400/10 border border-yellow-400/30 rounded">
            <p className="text-yellow-400 font-semibold">Demo Alert:</p>
            <p className="text-text-muted text-sm">
              The demo for Sequoia is Thursday. Make sure we're running the stable build, 
              not the production version with all the bugs.
            </p>
          </div>
        </TerminalWindow>

        {/* Signature and Postscript */}
        <TerminalWindow title="memo_footer.txt">
          <div className="space-y-4">
            <p className="text-text-muted">-Sarah</p>
            
            <div className="border-t border-border-primary pt-4">
              <p className="text-sm text-text-muted">
                <span className="font-semibold">P.S.</span> - If anyone asks about the quantum decoherence timeline, 
                we discovered a "breakthrough in temporal stability" that extends operational life indefinitely. 
                The research paper is "in peer review."
              </p>
            </div>
            
            <div className="bg-terminal-red/10 border border-terminal-red/30 rounded p-3">
              <p className="text-terminal-red text-sm font-semibold">
                This message will self-delete in 24 hours. Do not forward or screenshot.
              </p>
            </div>
          </div>
        </TerminalWindow>

        {/* Navigation */}
        <TerminalWindow title="navigation.sh">
          <div className="flex justify-between items-center">
            <Link 
              href="/bullshitard" 
              className="text-terminal-green hover:text-terminal-green/80 transition-colors"
            >
              ← Back to Bullshitard Home
            </Link>
            <div className="text-text-muted text-sm">
              Document Classification: LEAKED • Security Level: COMPROMISED
            </div>
          </div>
        </TerminalWindow>

      </div>
    </div>
  );
}
