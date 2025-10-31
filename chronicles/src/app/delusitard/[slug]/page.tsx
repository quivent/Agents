import { notFound } from 'next/navigation';
import { Layout } from '@/components/layout/Layout';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { CommandLine } from '@/components/ui/CommandLine';
import { getDelusitardContent, getContentBySlug, estimateReadingTime } from '@/lib/content';
import { getTardTheme } from '@/lib/themes';
import Link from 'next/link';
import { Metadata } from 'next';

interface Props {
  params: Promise<{ slug: string }>;
}

export async function generateStaticParams() {
  const content = getDelusitardContent();
  return content.map((item) => ({
    slug: item.slug,
  }));
}

export async function generateMetadata({ params }: Props): Promise<Metadata> {
  const { slug } = await params;
  const content = getContentBySlug(slug);
  
  if (!content) {
    return {
      title: 'Content Not Found',
    };
  }

  return {
    title: `${content.metadata.title} - Delusitard`,
    description: content.metadata.description || 'A journey through technical delusions',
  };
}

export default async function DelusitardContentPage({ params }: Props) {
  const { slug } = await params;
  const content = getContentBySlug(slug);
  const theme = getTardTheme('delusitard');

  if (!content) {
    notFound();
  }

  const readingTime = estimateReadingTime(content.content);

  return (
    <Layout>
      <div className="space-y-8">
        {/* Navigation */}
        <nav className="flex items-center justify-between">
          <Link 
            href="/delusitard" 
            className="hover:text-text-primary transition-colors" 
            style={{ color: theme.colors.accent }}
          >
            ← Back to Delusitard
          </Link>
          <div className="text-sm" style={{ color: theme.colors.secondary }}>
            {readingTime} min read
          </div>
        </nav>

        {/* Header */}
        <div className="text-center space-y-4">
          <h1 className="text-3xl md:text-4xl font-bold terminal-glow" 
              style={{ 
                color: theme.colors.primary,
                textShadow: `0 0 20px ${theme.colors.primary}` 
              }}>
            {content.metadata.title}
          </h1>
          {content.metadata.description && (
            <p className="text-lg max-w-2xl mx-auto" style={{ color: theme.colors.secondary }}>
              {content.metadata.description}
            </p>
          )}
        </div>

        {/* Content */}
        <TerminalWindow title={`${content.metadata.title.toLowerCase().replace(/\s+/g, '_')}.md`}>
          <CommandLine 
            command={`cat /delusion_archive/${content.metadata.title.toLowerCase().replace(/\s+/g, '_')}.md`}
            output="Loading alternative reality..."
            delay={20}
          />
          <div className="mt-6 prose prose-invert max-w-none">
            <div 
              className="whitespace-pre-wrap leading-relaxed"
              style={{ color: theme.colors.secondary }}
              dangerouslySetInnerHTML={{ __html: content.content.replace(/\n/g, '<br>') }}
            />
          </div>
        </TerminalWindow>
      </div>
    </Layout>
  );
}
