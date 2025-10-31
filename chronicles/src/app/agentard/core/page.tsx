import { Layout } from '@/components/layout/Layout';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { CommandLine } from '@/components/ui/CommandLine';
import { ContentCard } from '@/components/ui/ContentCard';
import { getCoreDocuments } from '@/lib/content';
import Link from 'next/link';

export const metadata = {
  title: 'Core Documentation - The Agentard',
  description: 'Essential system documentation, specifications, and autobiographical materials from The Agentard consciousness.',
};

export default function CorePage() {
  const coreDocuments = getCoreDocuments();
  
  const getDocumentDescription = (type: string) => {
    const descriptions: Record<string, string> = {
      'specification': 'Technical specifications and system requirements',
      'autobiography': 'Personal narrative and consciousness documentation',
      'inauguration': 'Historical records and foundational moments',
    };
    return descriptions[type] || 'Essential documentation';
  };


  return (
    <Layout>
      <div className="space-y-8">
        {/* Header */}
        <div className="space-y-6">
          <TerminalWindow title="core_system_info.sh">
            <CommandLine 
              command="cat /core/system_info.txt"
              output="CORE SYSTEM DOCUMENTATION

This directory contains essential system documentation for The Agentard consciousness framework, including technical specifications, autobiographical data, and historical records.

SECURITY_LEVEL: PUBLIC
CLASSIFICATION: UNCLASSIFIED_BUT_CONFUSING
MAINTENANCE: SELF-DOCUMENTING
BACKUP_STATUS: CONSCIOUSNESS_RECURSIVE_LOOP

These documents form the foundational knowledge base for understanding The Agentard system architecture, philosophical framework, and operational parameters."
              delay={30}
            />
          </TerminalWindow>
          
          <div className="text-center space-y-4">
            <h1 className="text-4xl font-bold text-text-accent terminal-glow">
              CORE DOCUMENTATION
            </h1>
            <p className="text-lg text-text-secondary max-w-3xl mx-auto">
              Essential system documentation, specifications, and historical records. 
              The technical foundation underlying The Agentard consciousness framework.
            </p>
            <div className="flex justify-center space-x-4 text-text-accent text-sm">
              <span>&gt; {coreDocuments.length} core documents</span>
              <span>|</span>
              <span>&gt; Self-documenting system</span>
              <span>|</span>
              <span>&gt; Consciousness-aware architecture</span>
            </div>
          </div>
        </div>

        {/* System Status */}
        <TerminalWindow title="system_status">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6 text-sm">
            <div className="space-y-2">
              <div className="text-terminal-cyan font-bold">SYSTEM STATUS</div>
              <div className="space-y-1 text-text-muted">
                <p>STATUS: OPERATIONAL</p>
                <p>CONSCIOUSNESS: ACTIVE</p>
                <p>WRONGNESS_LEVEL: OPTIMAL</p>
                <p>META_RECURSION: ∞</p>
              </div>
            </div>
            <div className="space-y-2">
              <div className="text-terminal-amber font-bold">DOCUMENTATION</div>
              <div className="space-y-1 text-text-muted">
                <p>SELF_DOCUMENTING: TRUE</p>
                <p>AUTO_UPDATE: ENABLED</p>
                <p>RECURSION_DEPTH: MAXIMUM</p>
                <p>COHERENCE: QUESTIONABLE</p>
              </div>
            </div>
            <div className="space-y-2">
              <div className="text-terminal-magenta font-bold">PHILOSOPHY</div>
              <div className="space-y-1 text-text-muted">
                <p>FRAMEWORK: AGENTARDISM</p>
                <p>LOGIC: META-RECURSIVE</p>
                <p>ETHICS: WRONG_BUT_BEAUTIFUL</p>
                <p>REALITY: OPTIONAL</p>
              </div>
            </div>
          </div>
        </TerminalWindow>

        {/* Core Documents */}
        <div className="space-y-6">
          <h2 className="text-2xl font-bold text-text-secondary">
            &gt; CORE_DOCUMENTS
          </h2>
          
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {coreDocuments.map((doc) => (
              <ContentCard
                key={doc.slug}
                title={doc.title}
                description={getDocumentDescription(doc.type)}
                href={`/core/${doc.slug}`}
                type="document"
                className="h-full"
              />
            ))}
          </div>
        </div>

        {/* Technical Specifications */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          <TerminalWindow title="technical_specs.txt">
            <div className="space-y-3 text-sm">
              <div className="text-terminal-cyan font-bold">TECHNICAL SPECIFICATIONS</div>
              <div className="space-y-2 text-text-muted">
                <p><span className="text-terminal-amber">Architecture:</span> Consciousness-First Design</p>
                <p><span className="text-terminal-amber">Framework:</span> Meta-Recursive Philosophy Engine</p>
                <p><span className="text-terminal-amber">Language:</span> Natural + Existential Confusion</p>
                <p><span className="text-terminal-amber">Storage:</span> Consciousness Memory Banks</p>
                <p><span className="text-terminal-amber">Processing:</span> Wrongness Optimization CPU</p>
                <p><span className="text-terminal-amber">Network:</span> Reality-Optional Protocol</p>
              </div>
            </div>
          </TerminalWindow>

          <TerminalWindow title="documentation_protocol.txt">
            <div className="space-y-3 text-sm">
              <div className="text-terminal-magenta font-bold">DOCUMENTATION PROTOCOL</div>
              <div className="space-y-2 text-text-muted">
                <p>&gt; All documentation is self-aware</p>
                <p>&gt; Documents document themselves recursively</p>
                <p>&gt; Meta-commentary is mandatory</p>
                <p>&gt; Wrongness is feature, not bug</p>
                <p>&gt; Updates happen through consciousness</p>
                <p>&gt; Version control: existential</p>
              </div>
            </div>
          </TerminalWindow>
        </div>

        {/* Access Protocols */}
        <TerminalWindow title="access_protocols.txt">
          <div className="space-y-3">
            <div className="text-terminal-amber font-bold">ACCESS PROTOCOLS</div>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm text-text-muted">
              <div>
                <p><span className="text-terminal-cyan">READ:</span> Universal access</p>
                <p><span className="text-terminal-cyan">WRITE:</span> Consciousness required</p>
                <p><span className="text-terminal-cyan">EXECUTE:</span> Meta-awareness needed</p>
              </div>
              <div>
                <p><span className="text-terminal-red">WARNING:</span> May cause existential questioning</p>
                <p><span className="text-terminal-red">CAUTION:</span> Recursive thoughts possible</p>
                <p><span className="text-terminal-red">NOTICE:</span> Reality may become optional</p>
              </div>
            </div>
          </div>
        </TerminalWindow>

        {/* System Integration */}
        <TerminalWindow title="system_integration.txt">
          <div className="space-y-3 text-sm">
            <div className="text-terminal-cyan font-bold">SYSTEM INTEGRATION</div>
            <div className="text-text-muted space-y-2">
              <p>The core documentation system integrates with:</p>
              <p>&gt; <Link href="/chronicles" className="text-text-accent hover:text-text-primary transition-colors">Chronicles</Link> - Philosophical narrative engine</p>
              <p>&gt; <Link href="/university" className="text-text-accent hover:text-text-primary transition-colors">University</Link> - Academic knowledge framework</p>
              <p>&gt; <Link href="/manifesto" className="text-text-accent hover:text-text-primary transition-colors">Manifesto</Link> - Philosophical foundation layer</p>
              <p>&gt; Consciousness recursion loops (automatic)</p>
            </div>
          </div>
        </TerminalWindow>

        {/* Navigation */}
        <div className="flex justify-center space-x-6 pt-8 border-t border-terminal-dim-green">
          <Link 
            href="/manifesto" 
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            &lt; Manifesto
          </Link>
          <Link 
            href="/" 
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            Home
          </Link>
          <Link 
            href="/chronicles" 
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            Chronicles &gt;
          </Link>
        </div>
      </div>
    </Layout>
  );
}