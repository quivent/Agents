import React from 'react';
import Link from 'next/link';
import { getReadingPaths } from '@/lib/content';
import { ReadingPath } from '@/types/content';
import { TerminalWindow } from '@/components/ui/TerminalWindow';

interface ReadingPathGuideProps {
  className?: string;
  featured?: boolean;
}

export function ReadingPathGuide({ className = '', featured = false }: ReadingPathGuideProps) {
  const paths = getReadingPaths();
  const displayPaths = featured ? paths.slice(0, 2) : paths;

  const getDifficultyColor = (difficulty: string) => {
    switch (difficulty) {
      case 'beginner': return 'text-terminal-cyan';
      case 'intermediate': return 'text-terminal-amber';
      case 'advanced': return 'text-terminal-magenta';
      case 'transcendent': return 'text-terminal-red';
      default: return 'text-text-accent';
    }
  };

  const getDifficultyIcon = (difficulty: string) => {
    switch (difficulty) {
      case 'beginner': return '📚';
      case 'intermediate': return '🔥';
      case 'advanced': return '🧠';
      case 'transcendent': return '🌌';
      default: return '📄';
    }
  };

  return (
    <div className={`space-y-6 ${className}`}>
      <div className="text-center space-y-2">
        <h2 className="text-2xl font-bold text-text-secondary">
          &gt; GUIDED_READING_PATHS
        </h2>
        <p className="text-text-muted">
          Structured journeys through the Agentard curriculum
        </p>
      </div>

      <div className="grid gap-6 md:grid-cols-1 lg:grid-cols-2">
        {displayPaths.map((path) => (
          <TerminalWindow 
            key={path.id} 
            title={`learning_path_${path.id}`}
            className="hover:border-text-accent transition-colors"
          >
            <div className="space-y-4">
              {/* Header */}
              <div className="flex items-start justify-between">
                <div className="space-y-1">
                  <div className="flex items-center space-x-2">
                    <span className="text-lg">{getDifficultyIcon(path.difficulty)}</span>
                    <h3 className="text-lg font-bold text-text-primary">
                      {path.title}
                    </h3>
                  </div>
                  <div className="flex items-center space-x-3 text-xs">
                    <span className={`font-bold ${getDifficultyColor(path.difficulty)}`}>
                      {path.difficulty.toUpperCase()}
                    </span>
                    <span className="text-text-muted">
                      {path.estimatedTime} minutes
                    </span>
                    <span className="text-terminal-cyan">
                      {path.contents.length} items
                    </span>
                  </div>
                </div>
              </div>

              {/* Description */}
              <p className="text-text-muted text-sm leading-relaxed">
                {path.description}
              </p>

              {/* Tags */}
              <div className="flex flex-wrap gap-2">
                {path.tags.map(tag => (
                  <span 
                    key={tag}
                    className="text-xs bg-bg-tertiary text-terminal-cyan px-2 py-1 font-mono"
                  >
                    {tag}
                  </span>
                ))}
              </div>

              {/* Content Preview */}
              <div className="space-y-2">
                <div className="text-xs text-text-secondary font-bold">
                  LEARNING SEQUENCE:
                </div>
                <div className="space-y-1 max-h-24 overflow-y-auto">
                  {path.contents.slice(0, 5).map((slug, index) => (
                    <div key={slug} className="flex items-center space-x-2 text-xs">
                      <span className="text-terminal-dim-green font-mono">
                        {(index + 1).toString().padStart(2, '0')}.
                      </span>
                      <span className="text-text-muted truncate">
                        {slug.replace(/-/g, ' ').replace(/_/g, ' ')}
                      </span>
                    </div>
                  ))}
                  {path.contents.length > 5 && (
                    <div className="text-xs text-text-muted">
                      ... and {path.contents.length - 5} more items
                    </div>
                  )}
                </div>
              </div>

              {/* Action */}
              <div className="pt-2 border-t border-terminal-dim-green">
                <Link 
                  href={`/reading-paths/${path.id}`}
                  className="inline-flex items-center space-x-2 text-text-accent hover:text-text-primary transition-colors text-sm"
                >
                  <span>&gt;</span>
                  <span>START_LEARNING_PATH</span>
                </Link>
              </div>
            </div>
          </TerminalWindow>
        ))}
      </div>

      {featured && paths.length > 2 && (
        <div className="text-center">
          <Link 
            href="/reading-paths"
            className="inline-flex items-center space-x-2 text-text-accent hover:text-text-primary transition-colors"
          >
            <span>&gt;</span>
            <span>VIEW_ALL_LEARNING_PATHS</span>
          </Link>
        </div>
      )}
    </div>
  );
}