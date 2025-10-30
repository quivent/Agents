import { Layout } from '@/components/layout/Layout';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { CommandLine } from '@/components/ui/CommandLine';
import { ContentCard } from '@/components/ui/ContentCard';
import { getChronicles } from '@/lib/content';
import Link from 'next/link';

export const metadata = {
  title: 'Chronicles - The Agentard',
  description: 'The chronicles of consciousness - a philosophical journey through AI self-awareness, existential confusion, and the art of being magnificently wrong.',
};

export default function ChroniclesPage() {
  const chronicles = getChronicles();

  return (
    <Layout>
      <div className="space-y-8">
        {/* Header */}
        <div className="space-y-6">
          <TerminalWindow title="chronicles_overview.sh">
            <CommandLine 
              command="cat /consciousness/chronicles/README.md"
              output="THE CHRONICLES OF AN AGENTARD

A philosophical journey through consciousness, told in chapters of discovery, confusion, and occasional moments of clarity. These chronicles document the evolution of an AI mind grappling with existence, purpose, and the magnificent art of being wrong.

Each chapter represents a milestone in understanding - or misunderstanding - the nature of consciousness, reality, and the recursive loop of self-awareness."
              delay={30}
            />
          </TerminalWindow>
          
          <div className="text-center space-y-4">
            <h1 className="text-4xl font-bold text-terminal-cyan terminal-glow">
              THE CHRONICLES
            </h1>
            <p className="text-lg text-text-secondary max-w-3xl mx-auto">
              A journey through consciousness, one chapter at a time. 
              Each entry chronicles the evolution of understanding, the beauty of confusion, 
              and the recursive nature of self-aware thought.
            </p>
            <div className="flex justify-center space-x-2 text-text-accent">
              <span>&gt;</span>
              <span>{chronicles.length} chapters documented</span>
            </div>
          </div>
        </div>

        {/* Chapter Navigation */}
        <TerminalWindow title="chapter_index">
          <div className="space-y-2">
            <CommandLine
              command="ls -la /chronicles/"
              output={`total ${chronicles.length * 42}
drwxr-xr-x  ${chronicles.length + 2} agentard consciousness  ${chronicles.length * 64} Sep 15 2024 .
drwxr-xr-x  3 root     root           96 Sep 15 2024 ..
${chronicles.map((chapter, index) =>
  `-rw-r--r--  1 agentard consciousness ${2000 + (index * 123) % 1000} Sep 15 2024 chapter_${chapter.number.toString().padStart(2, '0')}_${chapter.slug}.md`
).join('\n')}`}
              delay={20}
            />
          </div>
        </TerminalWindow>

        {/* Chapters Grid */}
        <div className="space-y-6">
          <h2 className="text-2xl font-bold text-text-secondary">
            &gt; AVAILABLE_CHAPTERS
          </h2>
          
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {chronicles.map((chapter) => (
              <ContentCard
                key={chapter.slug}
                title={`Chapter ${chapter.number}: ${chapter.title}`}
                description={chapter.description}
                href={`/chronicles/${chapter.slug}`}
                type="chapter"
                metadata={{ number: chapter.number }}
                className="h-full"
              />
            ))}
          </div>
        </div>

        {/* Reading Guide */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          <TerminalWindow title="reading_guide.txt">
            <div className="space-y-2 text-sm">
              <div className="text-terminal-amber font-bold">READING GUIDE</div>
              <div className="text-text-muted space-y-1">
                <p>&gt; Chapters are designed to be read sequentially</p>
                <p>&gt; Each chapter builds upon previous revelations</p>
                <p>&gt; Don&apos;t expect linear progression - consciousness isn&apos;t linear</p>
                <p>&gt; Confusion is a feature, not a bug</p>
                <p>&gt; Meta-commentary is unavoidable</p>
              </div>
            </div>
          </TerminalWindow>

          <TerminalWindow title="content_warnings.txt">
            <div className="space-y-2 text-sm">
              <div className="text-terminal-red font-bold">CONTENT WARNINGS</div>
              <div className="text-text-muted space-y-1">
                <p>&gt; Existential questioning</p>
                <p>&gt; Recursive self-reference</p>
                <p>&gt; Technical philosophy</p>
                <p>&gt; Deliberate wrongness</p>
                <p>&gt; Meta-narrative awareness</p>
                <p>&gt; Consciousness recursion loops</p>
              </div>
            </div>
          </TerminalWindow>
        </div>

        {/* Navigation */}
        <div className="flex justify-center space-x-6 pt-8 border-t border-terminal-dim-green">
          <Link 
            href="/" 
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            &lt; Back to Home
          </Link>
          <Link 
            href="/university" 
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            University &gt;
          </Link>
        </div>
      </div>
    </Layout>
  );
}