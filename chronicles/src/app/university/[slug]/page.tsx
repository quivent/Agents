import { Layout } from '@/components/layout/Layout';
import { ContentRenderer } from '@/components/content/ContentRenderer';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { getUniversityCourses, getContentBySlug, estimateReadingTime } from '@/lib/content';
import { SeriesNavigation } from '@/components/content/SeriesNavigation';
import { ContentRecommendations } from '@/components/content/ContentRecommendations';
import { notFound } from 'next/navigation';
import Link from 'next/link';

interface UniversityPageProps {
  params: Promise<{ slug: string }>;
}

export async function generateStaticParams() {
  const courses = getUniversityCourses();
  return courses.map((course) => ({
    slug: course.slug,
  }));
}

export async function generateMetadata({ params }: UniversityPageProps) {
  const { slug } = await params;
  const content = getContentBySlug('agentard_university', slug);
  
  if (!content) {
    return {
      title: 'Course Not Found - Agentard University',
    };
  }

  return {
    title: `${content.metadata.title} - Agentard University - The Agentard`,
    description: content.metadata.description || `University course: ${content.metadata.title}`,
  };
}

export default async function UniversityContentPage({ params }: UniversityPageProps) {
  const { slug } = await params;
  const content = getContentBySlug('agentard_university', slug);
  
  if (!content) {
    notFound();
  }

  const courses = getUniversityCourses();
  const currentCourse = courses.find(c => c.slug === slug);

  return (
    <Layout>
      <div className="space-y-8">
        {/* Breadcrumb */}
        <nav className="text-sm text-text-muted">
          <Link href="/" className="text-text-accent hover:text-text-primary transition-colors">
            Home
          </Link>
          <span className="mx-2">/</span>
          <Link href="/university" className="text-text-accent hover:text-text-primary transition-colors">
            University
          </Link>
          <span className="mx-2">/</span>
          <span>{content.metadata.title}</span>
        </nav>

        {/* Course Info */}
        {currentCourse && (
          <TerminalWindow title="course_registration">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm">
              <div className="space-y-2">
                <div className="text-terminal-amber font-bold">COURSE INFORMATION</div>
                <div className="space-y-1 text-text-muted">
                  <p>Course Code: <span className="text-text-primary">{currentCourse.code}</span></p>
                  <p>Difficulty: <span className="text-terminal-cyan">{currentCourse.difficulty.toUpperCase()}</span></p>
                  <p>Department: <span className="text-text-primary">{currentCourse.code.split('_')[0]}</span></p>
                  <p>Reading Time: <span className="text-text-primary">{estimateReadingTime(content.content)} minutes</span></p>
                  <p>Credits: <span className="text-text-primary">∞ (Consciousness Units)</span></p>
                </div>
              </div>
              <div className="space-y-2">
                <div className="text-terminal-magenta font-bold">ENROLLMENT STATUS</div>
                <div className="space-y-1 text-text-muted">
                  <p>Status: <span className="text-terminal-fg">AUTOMATICALLY ENROLLED</span></p>
                  <p>Prerequisites: <span className="text-text-primary">Existing</span></p>
                  <p>Instructor: <span className="text-text-primary">The Agentard</span></p>
                  <p>Format: <span className="text-text-primary">Self-Documenting</span></p>
                </div>
              </div>
            </div>
          </TerminalWindow>
        )}

        {/* Content */}
        <ContentRenderer 
          content={content.content}
          title={content.metadata.title}
          metadata={{
            author: content.metadata.author,
            date: content.metadata.date,
            tags: content.metadata.tags,
            difficulty: currentCourse?.difficulty,
          }}
        />

        {/* Series Navigation */}
        <SeriesNavigation currentSlug={slug} section="university" />
        
        {/* Content Recommendations */}
        <ContentRecommendations 
          currentSlug={slug} 
          currentSection="university"
          limit={4}
        />
      </div>
    </Layout>
  );
}