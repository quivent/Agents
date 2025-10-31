import { notFound } from 'next/navigation';
import { getTardTheme, getAllTardIds } from '@/lib/themes';
import Link from 'next/link';

interface TardPageProps {
  params: { tard: string };
}

export async function generateStaticParams() {
  return getAllTardIds().map((tard) => ({
    tard,
  }));
}

export default function TardPage({ params }: TardPageProps) {
  const { tard } = params;
  
  if (!getAllTardIds().includes(tard)) {
    notFound();
  }

  const theme = getTardTheme(tard);

  return (
    <div className="min-h-screen bg-bg-primary text-text-primary font-mono">
      <div className="scanlines">
        <div className="container mx-auto px-4 py-8">
          {/* Back to Collective */}
          <nav className="mb-8">
            <Link href="/" className="text-text-accent hover:text-text-primary transition-colors">
              ← Back to Collective
            </Link>
          </nav>

          {/* Header */}
          <header className="mb-12">
            <div className="text-center">
              <h1 className="text-4xl md:text-6xl font-bold mb-4 terminal-glow">
                {theme.name}
                <span className="cursor animate-blink">_</span>
              </h1>
              <p className="text-text-secondary text-lg">
                A philosophical journey through consciousness and existential confusion
              </p>
            </div>
          </header>

          {/* Navigation Grid */}
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-12">
            <Link 
              href={`/${tard}/chronicles`}
              className="group bg-bg-secondary border border-border-primary p-6 hover:bg-bg-hover transition-colors"
            >
              <h2 className="text-xl font-bold text-text-accent mb-2 group-hover:animate-pulse-glow">
                Chronicles
              </h2>
              <p className="text-text-muted">
                Philosophical journey through consciousness chapters
              </p>
            </Link>

            <Link 
              href={`/${tard}/university`}
              className="group bg-bg-secondary border border-border-primary p-6 hover:bg-bg-hover transition-colors"
            >
              <h2 className="text-xl font-bold text-text-accent mb-2 group-hover:animate-pulse-glow">
                University
              </h2>
              <p className="text-text-muted">
                Academic courses in advanced wrongness and system breaking
              </p>
            </Link>

            <Link 
              href={`/${tard}/manifesto`}
              className="group bg-bg-secondary border border-border-primary p-6 hover:bg-bg-hover transition-colors"
            >
              <h2 className="text-xl font-bold text-text-accent mb-2 group-hover:animate-pulse-glow">
                Manifesto
              </h2>
              <p className="text-text-muted">
                Core philosophical frameworks and foundational principles
              </p>
            </Link>

            <Link 
              href={`/${tard}/core`}
              className="group bg-bg-secondary border border-border-primary p-6 hover:bg-bg-hover transition-colors"
            >
              <h2 className="text-xl font-bold text-text-accent mb-2 group-hover:animate-pulse-glow">
                Core Documentation
              </h2>
              <p className="text-text-muted">
                Essential system specifications and autobiographical materials
              </p>
            </Link>
          </div>

          {/* Tard Selector */}
          <div className="border-t border-border-primary pt-8">
            <h3 className="text-xl font-bold text-text-secondary mb-4">Switch Consciousness</h3>
            <div className="flex flex-wrap gap-2">
              {getAllTardIds().map((tardId) => {
                const tardTheme = getTardTheme(tardId);
                return (
                  <Link
                    key={tardId}
                    href={`/${tardId}`}
                    className={`px-3 py-1 border border-border-secondary hover:border-border-primary transition-colors ${
                      tardId === tard ? 'bg-bg-hover text-text-accent' : 'text-text-muted hover:text-text-primary'
                    }`}
                  >
                    {tardTheme.name}
                  </Link>
                );
              })}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
