import { Layout } from '@/components/layout/Layout';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { 
  LazyInteractiveTerminal,
  LazyFailureCascadeVisualizer,
  LazyConfidenceRealityMeters
} from '@/components/lazy/LazyInteractiveComponents';
import { AgentardMoments } from '@/components/community/AgentardMoments';
import { InteractiveQuote } from '@/components/enhanced/InteractiveQuote';

export default function InteractivePage() {
  return (
    <Layout>
      <div className="space-y-12">
        {/* Hero Section */}
        <div className="text-center space-y-6">
          <h1 className="text-4xl md:text-6xl font-bold text-text-primary terminal-glow">
            INTERACTIVE AGENTARD
          </h1>
          <p className="text-xl text-text-secondary max-w-3xl mx-auto">
            Experience the art of confident incompetence through hands-on simulations, 
            real-time metrics, and community-shared moments of spectacular failure
          </p>
          <div className="text-terminal-amber">
            💡 Explore, interact, and discover hidden features throughout this page
          </div>
        </div>

        {/* Interactive Quote */}
        <InteractiveQuote className="max-w-4xl mx-auto" />

        {/* Interactive Terminal Simulator */}
        <section className="space-y-6">
          <div className="text-center">
            <h2 className="text-3xl font-bold text-terminal-cyan mb-4">
              🖥️ PROBLEM-SOLVING SIMULATOR
            </h2>
            <p className="text-text-secondary max-w-2xl mx-auto">
              Experience Agentard's problem-solving methodology firsthand. Watch confidence soar 
              while reality plummets in real-time!
            </p>
          </div>
          <LazyInteractiveTerminal />
        </section>

        {/* Failure Cascade Visualizer */}
        <section className="space-y-6">
          <div className="text-center">
            <h2 className="text-3xl font-bold text-terminal-red mb-4">
              📈 FAILURE CASCADE VISUALIZER
            </h2>
            <p className="text-text-secondary max-w-2xl mx-auto">
              Witness the legendary 2.7x problem multiplication factor in action. 
              One simple fix becomes a spectacular cascade of new challenges.
            </p>
          </div>
          <LazyFailureCascadeVisualizer />
        </section>

        {/* Confidence vs Reality Meters */}
        <section className="space-y-6">
          <div className="text-center">
            <h2 className="text-3xl font-bold text-terminal-amber mb-4">
              ⚖️ CONFIDENCE vs REALITY METERS
            </h2>
            <p className="text-text-secondary max-w-2xl mx-auto">
              Monitor the inverse relationship between confidence and reality. 
              Explore scenarios, watch live simulations, or create your own custom metrics.
            </p>
          </div>
          <LazyConfidenceRealityMeters />
        </section>

        {/* Community Moments */}
        <section className="space-y-6">
          <div className="text-center">
            <h2 className="text-3xl font-bold text-terminal-magenta mb-4">
              🌍 COMMUNITY AGENTARD MOMENTS
            </h2>
            <p className="text-text-secondary max-w-2xl mx-auto">
              Share your own moments of confident incompetence and explore spectacular 
              failures from the community. Vote on the most impressively wrong solutions!
            </p>
          </div>
          <AgentardMoments />
        </section>

        {/* Easter Egg Hints */}
        <TerminalWindow title="hidden_features.txt" className="max-w-4xl mx-auto">
          <div className="space-y-3">
            <div className="text-terminal-cyan font-bold">🥚 HIDDEN FEATURES DISCOVERED SO FAR:</div>
            <div className="text-sm text-text-muted space-y-1">
              <div>• Try the Konami Code (↑↑↓↓←→←→BA) for developer mode</div>
              <div>• Triple-click elements for hidden statistics</div>
              <div>• Type "red pill" anywhere to enter Matrix mode</div>
              <div>• Use "agentard-override" in terminals for reality mode</div>
              <div>• Look for clickable elements with special behaviors</div>
              <div>• Some commands in the terminal have easter egg responses</div>
            </div>
            <div className="text-xs text-terminal-amber mt-4">
              There may be more hidden features waiting to be discovered...
            </div>
          </div>
        </TerminalWindow>
      </div>
    </Layout>
  );
}