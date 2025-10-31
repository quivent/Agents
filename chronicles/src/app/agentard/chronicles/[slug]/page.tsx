import { Layout } from '@/components/layout/Layout';
import { ContentRenderer } from '@/components/content/ContentRenderer';
import { getChronicles, getContentBySectionAndSlug, estimateReadingTime } from '@/lib/content';
import { SeriesNavigation } from '@/components/content/SeriesNavigation';
import { ContentRecommendations } from '@/components/content/ContentRecommendations';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { notFound } from 'next/navigation';
import Link from 'next/link';

interface ChroniclePageProps {
  params: Promise<{ slug: string }>;
}

export async function generateStaticParams() {
  const chronicles = getChronicles();
  return chronicles.map((chapter) => ({
    slug: chapter.slug,
  }));
}

export async function generateMetadata({ params }: ChroniclePageProps) {
  const { slug } = await params;
  const content = getContentBySectionAndSlug('agentard_chronicles', slug);
  
  if (!content) {
    return {
      title: 'Chapter Not Found - The Agentard',
    };
  }

  return {
    title: `${content.metadata.title} - Chronicles - The Agentard`,
    description: content.metadata.description || `Chapter from The Chronicles of an Agentard: ${content.metadata.title}`,
  };
}

export default async function ChroniclesContentPage({ params }: ChroniclePageProps) {
  const { slug } = await params;
  const content = getContentBySectionAndSlug('agentard_chronicles', slug);
  
  if (!content) {
    notFound();
  }

  const chronicles = getChronicles();
  const currentChapter = chronicles.find(c => c.slug === slug);
  const currentIndex = chronicles.findIndex(c => c.slug === slug);
  const prevChapter = currentIndex > 0 ? chronicles[currentIndex - 1] : null;
  const nextChapter = currentIndex < chronicles.length - 1 ? chronicles[currentIndex + 1] : null;

  return (
    <Layout>
      <div className="space-y-8">
        {/* Breadcrumb */}
        <nav className="text-sm text-text-muted">
          <Link href="/" className="text-text-accent hover:text-text-primary transition-colors">
            Home
          </Link>
          <span className="mx-2">/</span>
          <Link href="/chronicles" className="text-text-accent hover:text-text-primary transition-colors">
            Chronicles
          </Link>
          <span className="mx-2">/</span>
          <span>{content.metadata.title}</span>
        </nav>

        {/* Chapter Info */}
        <TerminalWindow title={`chapter_${currentChapter?.number?.toString().padStart(2, '0')}_info`}>
          <div className="flex items-center justify-between text-sm">
            <div className="flex items-center space-x-4">
              <span className="text-terminal-cyan font-bold">
                CHAPTER {currentChapter?.number}
              </span>
              <span className="text-text-muted">
                {estimateReadingTime(content.content)} min read
              </span>
              {content.metadata.tags && (
                <div className="flex space-x-1">
                  {content.metadata.tags.slice(0, 3).map(tag => (
                    <span key={tag} className="text-xs bg-bg-tertiary text-terminal-cyan px-2 py-1">
                      {tag}
                    </span>
                  ))}
                </div>
              )}
            </div>
            <div className="text-text-muted">
              Part of Chronicles Series
            </div>
          </div>
        </TerminalWindow>

        {/* Content */}
        <ContentRenderer 
          content={content.content}
          title={content.metadata.title}
          metadata={{
            author: content.metadata.author,
            date: content.metadata.date,
            tags: content.metadata.tags,
            chapter: currentChapter?.number,
          }}
        />

        {/* Series Navigation */}
        <SeriesNavigation currentSlug={slug} section="chronicles" />
        
        {/* Content Recommendations */}
        <ContentRecommendations 
          currentSlug={slug} 
          currentSection="chronicles"
          limit={4}
        />
      </div>
    </Layout>
  );
}