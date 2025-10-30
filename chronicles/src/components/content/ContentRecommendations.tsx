import React from 'react';
import Link from 'next/link';
import { getRecommendedContent, getContentBySlug } from '@/lib/content';
import { ContentRecommendation } from '@/types/content';
import { TerminalWindow } from '@/components/ui/TerminalWindow';

interface ContentRecommendationsProps {
  currentSlug: string;
  currentSection: string;
  className?: string;
  limit?: number;
}

export function ContentRecommendations({ 
  currentSlug, 
  currentSection,
  className = '',
  limit = 3 
}: ContentRecommendationsProps) {
  const recommendations = getRecommendedContent(currentSlug, limit);

  if (recommendations.length === 0) {
    return null;
  }

  const getSectionFromSlug = (slug: string): string => {
    // This is a simplified mapping - in practice, you'd want a more robust solution
    const allSections = ['chronicles', 'university', 'manifesto', 'core'];
    
    // For now, we'll try to infer from common patterns
    if (slug.includes('chapter') || slug.includes('chronicle')) return 'chronicles';
    if (slug.includes('eng-') || slug.includes('app-') || slug.includes('math-') || slug.includes('phil-') || slug.includes('doc-')) return 'university';
    if (slug.includes('manifesto') || slug.includes('foundation') || slug.includes('vision')) return 'manifesto';
    return 'core';
  };

  return (
    <div className={`space-y-4 ${className}`}>
      <TerminalWindow title="content_recommendations.sh">
        <div className="space-y-4">
          <div className="flex items-center justify-between">
            <h3 className="text-lg font-bold text-terminal-cyan">
              RELATED_CONTENT
            </h3>
            <span className="text-xs text-text-muted">
              AI-curated suggestions
            </span>
          </div>

          <div className="space-y-3">
            {recommendations.map((rec, index) => {
              const section = getSectionFromSlug(rec.slug);
              return (
                <div key={rec.slug} className="group">
                  <Link 
                    href={`/${section}/${rec.slug}`}
                    className="block p-3 bg-bg-tertiary hover:bg-bg-quaternary border border-terminal-dim-green hover:border-text-accent transition-all"
                  >
                    <div className="flex items-start justify-between mb-2">
                      <h4 className="text-text-primary font-mono font-bold group-hover:text-text-accent transition-colors">
                        {rec.title}
                      </h4>
                      <span className="text-xs text-terminal-cyan ml-2">
                        #{(index + 1).toString().padStart(2, '0')}
                      </span>
                    </div>
                    
                    <div className="flex items-center justify-between">
                      <p className="text-text-muted text-sm">
                        {rec.reason}
                      </p>
                      
                      <div className="flex items-center space-x-2 text-xs">
                        <span className="text-terminal-amber">
                          RELEVANCE: {Math.round((rec.relevanceScore / 25) * 100)}%
                        </span>
                        <span className="text-text-accent group-hover:text-text-primary transition-colors">
                          &gt;
                        </span>
                      </div>
                    </div>
                  </Link>
                </div>
              );
            })}
          </div>

          <div className="pt-3 border-t border-terminal-dim-green text-xs text-text-muted">
            <div className="flex items-center space-x-2">
              <span className="text-terminal-cyan">INFO:</span>
              <span>Recommendations based on content similarity, difficulty level, and thematic relevance</span>
            </div>
          </div>
        </div>
      </TerminalWindow>
    </div>
  );
}