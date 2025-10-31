import { Layout } from '@/components/layout/Layout';
import { ContentRenderer } from '@/components/content/ContentRenderer';
import { getManifestoDocuments, getContentBySlug } from '@/lib/content';
import { notFound } from 'next/navigation';
import Link from 'next/link';

interface ManifestoPageProps {
  params: Promise<{ slug: string }>;
}

export async function generateStaticParams() {
  const manifestos = getManifestoDocuments();
  return manifestos.map((doc) => ({
    slug: doc.slug,
  }));
}

export async function generateMetadata({ params }: ManifestoPageProps) {
  const { slug } = await params;
  const content = getContentBySlug('manifesto', slug);
  
  if (!content) {
    return {
      title: 'Document Not Found - The Agentard Manifesto',
    };
  }

  return {
    title: `${content.metadata.title} - Manifesto - The Agentard`,
    description: content.metadata.description || `Manifesto document: ${content.metadata.title}`,
  };
}

export default async function ManifestoContentPage({ params }: ManifestoPageProps) {
  const { slug } = await params;
  const content = getContentBySlug('manifesto', slug);
  
  if (!content) {
    notFound();
  }

  const manifestos = getManifestoDocuments();
  const relatedDocs = manifestos.filter(doc => doc.slug !== slug);

  return (
    <Layout>
      <div className="space-y-8">
        {/* Breadcrumb */}
        <nav className="text-sm text-text-muted">
          <Link href="/" className="text-text-accent hover:text-text-primary transition-colors">
            Home
          </Link>
          <span className="mx-2">/</span>
          <Link href="/manifesto" className="text-text-accent hover:text-text-primary transition-colors">
            Manifesto
          </Link>
          <span className="mx-2">/</span>
          <span>{content.metadata.title}</span>
        </nav>

        {/* Content */}
        <ContentRenderer 
          content={content.content}
          title={content.metadata.title}
          metadata={{
            author: content.metadata.author,
            date: content.metadata.date,
            tags: content.metadata.tags,
          }}
        />

        {/* Related Documents */}
        {relatedDocs.length > 0 && (
          <div className="space-y-4">
            <h3 className="text-lg font-bold text-terminal-magenta">
              Related Manifesto Documents
            </h3>
            <div className="space-y-2">
              {relatedDocs.map((doc) => (
                <Link 
                  key={doc.slug}
                  href={`/manifesto/${doc.slug}`}
                  className="block text-text-accent hover:text-text-primary transition-colors"
                >
                  &gt; {doc.title}
                </Link>
              ))}
            </div>
          </div>
        )}

        {/* Navigation */}
        <div className="flex justify-center space-x-6 pt-8 border-t border-terminal-dim-green">
          <Link 
            href="/manifesto"
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            &lt; Back to Manifesto
          </Link>
          <Link 
            href="/university"
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            University
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