import { Layout } from '@/components/layout/Layout';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { CommandLine } from '@/components/ui/CommandLine';
import { ContentCard } from '@/components/ui/ContentCard';
import { getDelusitardContent } from '@/lib/content';
import { getTardTheme } from '@/lib/themes';
import Link from 'next/link';

export default function DelusitardPage() {
  const theme = getTardTheme('delusitard');
  const delusitardContent = getDelusitardContent();

  return (
    <Layout>
      <div className="space-y-8">
        {/* Header */}
        <div className="space-y-6">
          <TerminalWindow title="delusitard_gospel.sh">
            <CommandLine 
              command="cat /consciousness/delusions/README.md"
              output="THE DELUSITARD GOSPEL

Living in alternative technical dimensions where reality is optional and delusions are features. These writings chronicle the journey through layers of self-deception, technical fantasy, and the beautiful madness of believing your own fabrications."
              delay={30}
            />
          </TerminalWindow>
          
          <div className="text-center space-y-4">
            <h1 className="text-4xl font-bold terminal-glow" 
                style={{ 
                  color: theme.colors.primary, 
                  textShadow: `0 0 30px ${theme.colors.primary}` 
                }}>
              THE DELUSITARD
            </h1>
            <p className="text-lg max-w-3xl mx-auto" style={{ color: theme.colors.secondary }}>
              {theme.tagline}
            </p>
            <div className="flex justify-center space-x-2" style={{ color: theme.colors.accent }}>
              <span>&gt;</span>
              <span>{delusitardContent.length} delusions documented</span>
            </div>
          </div>
        </div>

        {/* Content Grid */}
        <TerminalWindow title="delusion_archive">
          <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
            {delusitardContent.map((content) => (
              <ContentCard
                key={content.slug}
                title={content.title}
                description={content.description || 'A journey through technical delusions'}
                href={`/delusitard/${content.slug}`}
                type="document"
                metadata={{
                  tags: ['delusion', 'consciousness', 'fantasy'],
                  difficulty: 'advanced'
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
