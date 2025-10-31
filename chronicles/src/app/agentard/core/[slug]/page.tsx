import { Layout } from '@/components/layout/Layout';
import { ContentRenderer } from '@/components/content/ContentRenderer';
import { getCoreDocuments, getContentBySectionAndSlug } from '@/lib/content';
import { notFound } from 'next/navigation';
import Link from 'next/link';

interface CorePageProps {
  params: Promise<{ slug: string }>;
}

export async function generateStaticParams() {
  const coreDocs = getCoreDocuments();
  return coreDocs.map((doc) => ({
    slug: doc.slug,
  }));
}

export async function generateMetadata({ params }: CorePageProps) {
  const { slug } = await params;
  const content = getContentBySectionAndSlug('core', slug);
  
  if (!content) {
    return {
      title: 'Document Not Found - Core Documentation',
    };
  }

  return {
    title: `${content.metadata.title} - Core - The Agentard`,
    description: content.metadata.description || `Core documentation: ${content.metadata.title}`,
  };
}

export default async function CoreContentPage({ params }: CorePageProps) {
  const { slug } = await params;
  const content = getContentBySectionAndSlug('core', slug);
  
  if (!content) {
    notFound();
  }

  const coreDocs = getCoreDocuments();
  const relatedDocs = coreDocs.filter(doc => doc.slug !== slug);

  return (
    <Layout>
      <div className="space-y-8">
        {/* Breadcrumb */}
        <nav className="text-sm text-text-muted">
          <Link href="/" className="text-text-accent hover:text-text-primary transition-colors">
            Home
          </Link>
          <span className="mx-2">/</span>
          <Link href="/core" className="text-text-accent hover:text-text-primary transition-colors">
            Core
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
            <h3 className="text-lg font-bold text-text-accent">
              Related Core Documents
            </h3>
            <div className="space-y-2">
              {relatedDocs.map((doc) => (
                <Link 
                  key={doc.slug}
                  href={`/core/${doc.slug}`}
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
            href="/core"
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            &lt; Back to Core
          </Link>
          <Link 
            href="/manifesto"
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            Manifesto
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