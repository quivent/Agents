import { notFound } from 'next/navigation';
import { getTardTheme, getAllTardIds } from '@/lib/themes';
import Link from 'next/link';

interface CorePageProps {
  params: { tard: string };
}

export async function generateStaticParams() {
  return getAllTardIds().map((tard) => ({
    tard,
  }));
}

export default function CorePage({ params }: CorePageProps) {
  const { tard } = params;
  
  if (!getAllTardIds().includes(tard)) {
    notFound();
  }

  const theme = getTardTheme(tard);

  return (
    <div className="min-h-screen bg-bg-primary text-text-primary font-mono">
      <div className="scanlines">
        <div className="container mx-auto px-4 py-8">
          <nav className="mb-8">
            <Link href={`/${tard}`} className="text-text-accent hover:text-text-primary">
              {theme.name}
            </Link>
            <span className="text-text-muted mx-2">/</span>
            <span className="text-text-secondary">Core</span>
          </nav>

          <header className="mb-12">
            <h1 className="text-4xl font-bold mb-4 terminal-glow">
              {theme.name} Core Documentation
              <span className="cursor animate-blink">_</span>
            </h1>
            <p className="text-text-secondary text-lg">
              Essential system specifications and autobiographical materials
            </p>
          </header>

          <div className="bg-bg-secondary border border-border-primary p-8">
            <div className="text-center text-text-muted">
              <p className="text-2xl mb-4">⚙️</p>
              <p className="text-lg mb-2">Core documentation coming soon...</p>
              <p className="text-sm">
                This section will contain system specifications for {theme.name}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
