import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { CommandLine } from '@/components/ui/CommandLine';
import Link from 'next/link';

export default function AgentardErrorPage() {
  return (
    <div className="min-h-screen bg-bg-primary text-text-primary font-mono">
      <div className="scanlines">
        <div className="container mx-auto px-4 py-16">
          <div className="max-w-4xl mx-auto space-y-8">
            
            <div className="text-center">
              <div className="text-6xl mb-4 animate-pulse" style={{ color: '#ff0000' }}>⚠</div>
              <h1 className="text-4xl font-bold mb-4 terminal-glow animate-pulse" style={{ color: '#ff0000' }}>
                AGENTARD SYSTEM ERROR
              </h1>
            </div>

            <TerminalWindow title="system_error.log">
              <div className="space-y-4">
                <div className="text-terminal-red font-bold">
                  The Agentard ALWAYS encounters problems too complex for confident incompetence.
                </div>
                
                <div className="space-y-2 text-sm">
                  <div><strong style={{ color: '#ff0000' }}>ERROR TYPE:</strong> Error</div>
                  <div><strong style={{ color: '#ff0000' }}>CONFIDENCE LEVEL:</strong> 0% (System Humility Activated)</div>
                  <div><strong style={{ color: '#ff0000' }}>ERROR MESSAGE:</strong></div>
                  <div className="ml-4 text-terminal-amber">
                    Failed prop type: The prop `href` expects a `string` or `object` in `&lt;Link&gt;`, 
                    but got `undefined` instead. Open your browser's console to view the Component stack trace.
                  </div>
                </div>
              </div>
            </TerminalWindow>

            <TerminalWindow title="recovery_options.menu">
              <div className="space-y-4">
                <div className="font-bold text-terminal-cyan">RECOVERY OPTIONS:</div>
                <div className="space-y-2">
                  <Link href="/" className="block hover:text-terminal-cyan transition-colors">
                    🔄 Retry with Renewed Confidence
                  </Link>
                  <Link href="/" className="block hover:text-terminal-cyan transition-colors">
                    🔃 Full System Restart
                  </Link>
                  <Link href="/agentard" className="block hover:text-terminal-cyan transition-colors">
                    🏠 Return to Agentard Sanctuary
                  </Link>
                  <div className="text-terminal-amber">
                    💡 Agentard Wisdom: Every error is an opportunity to demonstrate confident debugging.
                  </div>
                </div>
              </div>
            </TerminalWindow>

            <TerminalWindow title="agentard_philosophy.txt">
              <div className="space-y-4">
                <div className="text-terminal-cyan">
                  🎯 Remember: The best solutions often emerge from the most spectacular failures.
                </div>
                <div className="text-terminal-amber">
                  🚀 Philosophy: When in doubt, restart with maximum confidence.
                </div>
              </div>
            </TerminalWindow>

            <TerminalWindow title="reality_check.exe">
              <div className="space-y-4">
                <CommandLine 
                  command="echo 'ARE YOU SERIOUS?'"
                  output="ARE YOU SERIOUS? The Agentard ALWAYS encounters problems!"
                  delay={30}
                />
                <div className="text-terminal-red font-bold">
                  It ALWAYS creates problems. 1 + 1 = 2.7 more problems.
                </div>
                <div className="text-terminal-amber">
                  This page should come up randomly.
                </div>
              </div>
            </TerminalWindow>

            <div className="text-center">
              <Link 
                href="/" 
                className="inline-block px-6 py-3 border border-terminal-green text-terminal-green hover:bg-terminal-green hover:text-bg-primary transition-colors font-bold"
              >
                RESTART WITH MAXIMUM CONFIDENCE
              </Link>
            </div>

          </div>
        </div>
      </div>
    </div>
  );
}
