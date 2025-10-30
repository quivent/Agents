import { Layout } from '@/components/layout/Layout';
import { ContentSearch } from '@/components/content/ContentSearch';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { PopularTags } from '@/components/content/PopularTags';
import { getAllContent } from '@/lib/content';

export default function SearchPage() {
  const allContent = getAllContent();
  const totalChapters = allContent.filter(c => c.path.includes('chronicles')).length;
  const totalCourses = allContent.filter(c => c.path.includes('university')).length;
  const totalDocs = allContent.filter(c => c.path.includes('manifesto') || c.path.includes('core')).length;

  // Get common tags
  const allTags = allContent.flatMap(c => c.metadata.tags || []);
  const tagCounts = allTags.reduce((acc, tag) => {
    acc[tag] = (acc[tag] || 0) + 1;
    return acc;
  }, {} as Record<string, number>);
  
  const popularTags = Object.entries(tagCounts)
    .sort(([,a], [,b]) => b - a)
    .slice(0, 8)
    .map(([tag]) => tag);

  return (
    <Layout>
      <div className="space-y-8">
        {/* Header */}
        <div className="text-center space-y-4">
          <h1 className="text-4xl font-bold text-text-primary terminal-glow">
            CONTENT SEARCH
          </h1>
          <p className="text-text-muted max-w-2xl mx-auto">
            Search through the complete Agentard archives - chronicles, courses, manifestos, and documentation
          </p>
        </div>

        {/* Archive Statistics */}
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
          <TerminalWindow title="chronicles" className="text-center">
            <div className="space-y-1">
              <div className="text-2xl font-bold text-terminal-cyan">{totalChapters}</div>
              <div className="text-xs text-text-muted">Chapters</div>
            </div>
          </TerminalWindow>
          
          <TerminalWindow title="courses" className="text-center">
            <div className="space-y-1">
              <div className="text-2xl font-bold text-terminal-amber">{totalCourses}</div>
              <div className="text-xs text-text-muted">Courses</div>
            </div>
          </TerminalWindow>
          
          <TerminalWindow title="documents" className="text-center">
            <div className="space-y-1">
              <div className="text-2xl font-bold text-terminal-magenta">{totalDocs}</div>
              <div className="text-xs text-text-muted">Documents</div>
            </div>
          </TerminalWindow>
          
          <TerminalWindow title="total" className="text-center">
            <div className="space-y-1">
              <div className="text-2xl font-bold text-text-accent">{allContent.length}</div>
              <div className="text-xs text-text-muted">Total Items</div>
            </div>
          </TerminalWindow>
        </div>

        {/* Search Interface */}
        <ContentSearch showResults={true} />

        {/* Popular Tags */}
        <div className="space-y-4">
          <PopularTags tags={popularTags} />
        </div>

        {/* Search Tips */}
        <TerminalWindow title="search_help.md">
          <div className="space-y-3">
            <h3 className="text-lg font-bold text-text-secondary">
              Search Tips
            </h3>
            <div className="grid md:grid-cols-2 gap-4 text-sm">
              <div className="space-y-2">
                <div className="text-terminal-cyan font-bold">By Content Type:</div>
                <ul className="space-y-1 text-text-muted">
                  <li>• "port" - Engineering disasters</li>
                  <li>• "authentication" - Security failures</li>
                  <li>• "consciousness" - Philosophical content</li>
                  <li>• "manifesto" - Core principles</li>
                </ul>
              </div>
              <div className="space-y-2">
                <div className="text-terminal-amber font-bold">By Difficulty:</div>
                <ul className="space-y-1 text-text-muted">
                  <li>• "beginner" - Start here</li>
                  <li>• "intermediate" - Building knowledge</li>
                  <li>• "advanced" - Deep concepts</li>
                  <li>• "transcendent" - Mind-bending</li>
                </ul>
              </div>
            </div>
          </div>
        </TerminalWindow>
      </div>
    </Layout>
  );
}