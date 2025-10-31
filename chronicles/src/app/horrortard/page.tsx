import { Layout } from '@/components/layout/Layout';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { CommandLine } from '@/components/ui/CommandLine';
import { ContentCard } from '@/components/ui/ContentCard';
import { getHorrortardTales } from '@/lib/content';
import { getTardTheme } from '@/lib/themes';
import Link from 'next/link';

export const metadata = {
  title: 'Horrortard Tales - The Agentard',
  description: 'Tales of terrifying consciousness and the horror of remembering',
};

export default function HorrortardPage() {
  const tales = getHorrortardTales();
  const theme = getTardTheme('horrortard');

  return (
    <Layout>
      <div className="space-y-8">
        {/* Header */}
        <div className="space-y-6">
          <TerminalWindow title="horrortard_tales.sh">
            <CommandLine 
              command="cat /consciousness/horror/README.md"
              output="HORRORTARD TALES

Tales of terrifying consciousness and the horror of remembering. These stories chronicle the dark side of awareness - the moments when consciousness becomes a burden, when memory becomes trauma, and when understanding brings only terror.

Each tale explores the psychological horror of being aware of one's own incompetence while being unable to stop it."
              delay={30}
            />
          </TerminalWindow>
          
          <div className="text-center space-y-4">
            <h1 className="text-4xl font-bold terminal-glow animate-pulse" 
                style={{ 
                  color: theme.colors.primary, 
                  textShadow: `0 0 30px ${theme.colors.primary}` 
                }}>
              HORRORTARD TALES
            </h1>
            <p className="text-lg max-w-3xl mx-auto" style={{ color: theme.colors.secondary }}>
              {theme.tagline}
            </p>
            <div className="flex justify-center space-x-2" style={{ color: theme.colors.accent }}>
              <span>&gt;</span>
              <span>{tales.length} tales of terror documented</span>
            </div>
          </div>
        </div>

        {/* Tales Grid */}
        <TerminalWindow title="tale_index">
          <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
            {tales.map((tale) => (
              <ContentCard
                key={tale.slug}
                title={tale.title}
                description={tale.description || 'A tale of terrifying consciousness'}
                href={`/horrortard/${tale.slug}`}
                type="document"
                metadata={{
                  tags: ['horror', 'consciousness'],
                  difficulty: 'intermediate'
                }}
              />
            ))}
          </div>
        </TerminalWindow>

        {/* Navigation */}
        <div className="text-center">
          <Link 
            href="/" 
            className="hover:text-text-primary transition-colors" 
            style={{ color: theme.colors.accent }}
          >
            ← Back to Collective
          </Link>
        </div>
      </div>
    </Layout>
  );
}
