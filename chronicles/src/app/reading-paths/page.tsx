import { Layout } from '@/components/layout/Layout';
import { ReadingPathGuide } from '@/components/content/ReadingPathGuide';
import { TerminalWindow } from '@/components/ui/TerminalWindow';

export default function ReadingPathsPage() {
  return (
    <Layout>
      <div className="space-y-8">
        {/* Header */}
        <div className="text-center space-y-4">
          <h1 className="text-4xl font-bold text-text-primary terminal-glow">
            GUIDED LEARNING PATHS
          </h1>
          <p className="text-text-muted max-w-2xl mx-auto">
            Structured journeys through the Agentard curriculum, organized by theme and difficulty
          </p>
        </div>

        {/* Learning Philosophy */}
        <TerminalWindow title="learning_philosophy.md">
          <div className="space-y-4">
            <h3 className="text-lg font-bold text-terminal-cyan">
              The Agentard Learning Philosophy
            </h3>
            <div className="space-y-3 text-sm text-text-muted">
              <p>
                Traditional learning assumes competence leads to confidence. The Agentard method 
                demonstrates that maximum confidence with minimum competence creates the most 
                spectacular learning opportunities.
              </p>
              <p>
                Each learning path is carefully designed to build upon previous mistakes, 
                creating cascading educational disasters that somehow result in enlightenment.
              </p>
              <div className="border-l-2 border-terminal-amber pl-4 text-terminal-amber">
                <strong>Warning:</strong> These learning paths may cause existential questioning, 
                compulsive system-breaking, and an irresistible urge to hardcode everything.
              </div>
            </div>
          </div>
        </TerminalWindow>

        {/* All Reading Paths */}
        <ReadingPathGuide />

        {/* Custom Path Creation */}
        <TerminalWindow title="custom_paths.sh">
          <div className="space-y-4">
            <h3 className="text-lg font-bold text-text-secondary">
              Create Your Own Path
            </h3>
            <div className="space-y-3 text-sm">
              <p className="text-text-muted">
                Want to create a custom learning journey? Here are some approaches:
              </p>
              <div className="grid md:grid-cols-2 gap-4">
                <div className="space-y-2">
                  <div className="text-terminal-cyan font-bold">By Interest:</div>
                  <ul className="space-y-1 text-text-muted">
                    <li>• Start with topics that intrigue you</li>
                    <li>• Follow tag-based connections</li>
                    <li>• Use the recommendation system</li>
                    <li>• Jump between sections freely</li>
                  </ul>
                </div>
                <div className="space-y-2">
                  <div className="text-terminal-amber font-bold">By Preparation:</div>
                  <ul className="space-y-1 text-text-muted">
                    <li>• Begin with the Manifesto</li>
                    <li>• Read prerequisite content first</li>
                    <li>• Progress through difficulty levels</li>
                    <li>• Complete entire series</li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </TerminalWindow>
      </div>
    </Layout>
  );
}