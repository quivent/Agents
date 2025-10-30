import React from 'react';
import Link from 'next/link';

interface ContentCardProps {
  title: string;
  description?: string;
  href: string;
  type?: 'chapter' | 'course' | 'document' | 'manifesto';
  metadata?: {
    number?: number;
    difficulty?: string;
    date?: string;
    readingTime?: number;
    tags?: string[];
    contentType?: 'core' | 'supplementary' | 'advanced' | 'philosophical';
    prerequisites?: string[];
  };
  className?: string;
  showTags?: boolean;
  compact?: boolean;
}

export function ContentCard({ 
  title, 
  description, 
  href, 
  type = 'document', 
  metadata, 
  className = '',
  showTags = true,
  compact = false 
}: ContentCardProps) {
  const getTypePrefix = () => {
    switch (type) {
      case 'chapter': return '[CHAPTER]';
      case 'course': return '[COURSE]';
      case 'manifesto': return '[MANIFESTO]';
      default: return '[DOC]';
    }
  };

  const getTypeColor = () => {
    switch (type) {
      case 'chapter': return 'text-terminal-cyan';
      case 'course': return 'text-terminal-amber';
      case 'manifesto': return 'text-terminal-magenta';
      default: return 'text-text-accent';
    }
  };

  return (
    <Link 
      href={href}
      className={`
        block bg-bg-secondary border border-terminal-dim-green 
        hover:border-text-accent hover:bg-bg-tertiary 
        transition-all duration-200 p-4 group
        ${className}
      `}
    >
      {/* Header */}
      <div className="flex items-center justify-between mb-2">
        <span className={`text-xs font-mono ${getTypeColor()}`}>
          {getTypePrefix()}
        </span>
        {metadata?.number && (
          <span className="text-xs text-text-muted">
            #{metadata.number.toString().padStart(2, '0')}
          </span>
        )}
      </div>
      
      {/* Title */}
      <h3 className="text-text-primary font-mono font-bold mb-2 group-hover:text-text-accent transition-colors">
        {title}
      </h3>
      
      {/* Description */}
      {description && (
        <p className="text-text-muted text-sm mb-3 line-clamp-3">
          {description}
        </p>
      )}
      
      {/* Tags */}
      {showTags && metadata?.tags && metadata.tags.length > 0 && (
        <div className="flex flex-wrap gap-1 mb-2">
          {metadata.tags.slice(0, 3).map(tag => (
            <span key={tag} className="text-xs bg-bg-tertiary text-terminal-cyan px-2 py-1 font-mono">
              {tag}
            </span>
          ))}
        </div>
      )}
      
      {/* Metadata */}
      <div className="flex items-center justify-between text-xs text-text-muted">
        <div className="flex items-center space-x-3">
          {metadata?.difficulty && (
            <span className="text-terminal-amber">
              {compact ? metadata.difficulty.toUpperCase() : `DIFFICULTY: ${metadata.difficulty.toUpperCase()}`}
            </span>
          )}
          {metadata?.readingTime && (
            <span className="text-terminal-cyan">
              {metadata.readingTime}min
            </span>
          )}
          {metadata?.contentType && (
            <span className="text-terminal-magenta">
              {metadata.contentType.toUpperCase()}
            </span>
          )}
          {metadata?.date && (
            <span>
              {metadata.date}
            </span>
          )}
        </div>
        <span className="text-text-accent group-hover:text-text-primary transition-colors">
          &gt; {compact ? 'READ' : 'EXECUTE'}
        </span>
      </div>
      
      {/* Prerequisites Warning */}
      {metadata?.prerequisites && metadata.prerequisites.length > 0 && (
        <div className="mt-2 text-xs text-terminal-amber border-l-2 border-terminal-amber pl-2">
          ⚠ Prerequisites: {metadata.prerequisites.length} required
        </div>
      )}
    </Link>
  );
}