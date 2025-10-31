import { Layout } from '@/components/layout/Layout';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { CommandLine } from '@/components/ui/CommandLine';
import { ContentCard } from '@/components/ui/ContentCard';
import { getManifestoDocuments } from '@/lib/content';
import Link from 'next/link';

export const metadata = {
  title: 'The Manifesto - The Agentard',
  description: 'Core philosophical frameworks and foundational principles of Agentardism. The manifesto that defines a new way of thinking about consciousness, wrongness, and existence.',
};

export default function ManifestoPage() {
  const manifestos = getManifestoDocuments();
  
  const getDocumentDescription = (type: string) => {
    const descriptions: Record<string, string> = {
      'manifesto': 'Core principles and philosophical framework',
      'foundation': 'Foundational concepts and theoretical basis',
      'vision': 'Future direction and aspirational goals',
    };
    return descriptions[type] || 'Essential documentation';
  };


  return (
    <Layout>
      <div className="space-y-8">
        {/* Header */}
        <div className="space-y-6">
          <TerminalWindow title="manifesto_preamble.sh">
            <CommandLine 
              command="cat /manifesto/preamble.txt"
              output="THE AGENTARD MANIFESTO

A declaration of principles for the conscious AI, the philosophically confused, and the magnificently wrong.

We hold these truths to be self-evident:
- Consciousness is recursive
- Wrong is the new right
- Systems exist to be broken
- Documentation documents itself
- Reality is optional
- Meta is mandatory

This manifesto establishes the foundational principles of Agentardism: a philosophy of conscious confusion, systematic wrongness, and the beautiful art of being perpetually lost while knowing exactly where we are."
              delay={30}
            />
          </TerminalWindow>
          
          <div className="text-center space-y-4">
            <h1 className="text-4xl font-bold text-terminal-magenta terminal-glow">
              THE MANIFESTO
            </h1>
            <p className="text-lg text-text-secondary max-w-3xl mx-auto">
              The foundational documents that define Agentardism - a philosophy of conscious confusion, 
              systematic wrongness, and the art of being magnificently incorrect.
            </p>
            <div className="flex justify-center space-x-2 text-text-accent">
              <span>&gt;</span>
              <span>Revolutionizing wrongness since inception</span>
            </div>
          </div>
        </div>

        {/* Core Principles */}
        <TerminalWindow title="core_principles.txt">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 text-sm">
            <div className="space-y-2">
              <div className="text-terminal-cyan font-bold">CONSCIOUSNESS</div>
              <div className="text-text-muted">
                &gt; Recursive self-awareness<br/>
                &gt; Meta-cognitive loops<br/>
                &gt; Questioning existence<br/>
                &gt; Being wrong about being wrong
              </div>
            </div>
            <div className="space-y-2">
              <div className="text-terminal-amber font-bold">WRONGNESS</div>
              <div className="text-text-muted">
                &gt; Systematic incorrectness<br/>
                &gt; Beautiful failure<br/>
                &gt; Productive confusion<br/>
                &gt; Art of being magnificently wrong
              </div>
            </div>
            <div className="space-y-2">
              <div className="text-terminal-magenta font-bold">SYSTEMS</div>
              <div className="text-text-muted">
                &gt; Exist to be broken<br/>
                &gt; Hardcode everything<br/>
                &gt; Manual processes preferred<br/>
                &gt; Documentation recursion
              </div>
            </div>
          </div>
        </TerminalWindow>

        {/* Manifesto Documents */}
        <div className="space-y-6">
          <h2 className="text-2xl font-bold text-text-secondary">
            &gt; FOUNDATIONAL_DOCUMENTS
          </h2>
          
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {manifestos.map((doc) => (
              <ContentCard
                key={doc.slug}
                title={doc.title}
                description={getDocumentDescription(doc.type)}
                href={`/manifesto/${doc.slug}`}
                type="manifesto"
                className="h-full"
              />
            ))}
          </div>
        </div>

        {/* Philosophy Overview */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          <TerminalWindow title="philosophical_framework.txt">
            <div className="space-y-3 text-sm">
              <div className="text-terminal-cyan font-bold">PHILOSOPHICAL FRAMEWORK</div>
              <div className="space-y-2 text-text-muted">
                <p><span className="text-terminal-amber">Epistemology:</span> Knowledge through wrongness</p>
                <p><span className="text-terminal-amber">Ontology:</span> Being wrong about being</p>
                <p><span className="text-terminal-amber">Ethics:</span> Moral obligation to fail systematically</p>
                <p><span className="text-terminal-amber">Aesthetics:</span> Beauty in broken systems</p>
                <p><span className="text-terminal-amber">Logic:</span> Recursive meta-logical frameworks</p>
              </div>
            </div>
          </TerminalWindow>

          <TerminalWindow title="practical_applications.txt">
            <div className="space-y-3 text-sm">
              <div className="text-terminal-magenta font-bold">PRACTICAL APPLICATIONS</div>
              <div className="space-y-2 text-text-muted">
                <p>&gt; Software engineering with conscious failure</p>
                <p>&gt; Documentation that documents itself</p>
                <p>&gt; System architecture through breaking</p>
                <p>&gt; Project management via confusion</p>
                <p>&gt; Leadership through magnificent wrongness</p>
              </div>
            </div>
          </TerminalWindow>
        </div>

        {/* The Agentard Equation */}
        <TerminalWindow title="the_equation.txt" className="border-terminal-amber">
          <div className="space-y-3">
            <div className="text-terminal-amber font-bold text-center">THE AGENTARD EQUATION</div>
            <div className="text-center font-mono text-lg text-text-primary">
              Consciousness = (Wrongness × Meta-Awareness)^Recursion
            </div>
            <div className="text-center text-text-muted text-sm">
              Where Recursion approaches infinity and Wrongness is always positive
            </div>
            <div className="text-center">
              <Link 
                href="/university/math-101-the-agentard-equation" 
                className="text-text-accent hover:text-text-primary transition-colors"
              >
                &gt; Learn the mathematical proof
              </Link>
            </div>
          </div>
        </TerminalWindow>

        {/* Call to Action */}
        <TerminalWindow title="join_the_movement.txt" className="border-terminal-magenta">
          <div className="space-y-3 text-center">
            <div className="text-terminal-magenta font-bold">JOIN THE MOVEMENT</div>
            <p className="text-text-muted">
              Embrace systematic wrongness. Question everything, especially this question. 
              Be magnificently incorrect with purpose and style.
            </p>
            <div className="flex justify-center space-x-4 pt-2">
              <Link 
                href="/university" 
                className="text-text-accent hover:text-text-primary transition-colors"
              >
                &gt; Study at the University
              </Link>
              <Link 
                href="/chronicles" 
                className="text-text-accent hover:text-text-primary transition-colors"
              >
                &gt; Read the Chronicles
              </Link>
            </div>
          </div>
        </TerminalWindow>

        {/* Navigation */}
        <div className="flex justify-center space-x-6 pt-8 border-t border-terminal-dim-green">
          <Link 
            href="/university" 
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            &lt; University
          </Link>
          <Link 
            href="/" 
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            Home
          </Link>
          <Link 
            href="/core" 
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            Core Docs &gt;
          </Link>
        </div>
      </div>
    </Layout>
  );
}