'use client';
import React, { useState, useEffect } from 'react';
import { searchContentAPI, getAllContentAPI } from '@/lib/content-client';
import { ContentDocument } from '@/types/content';
import { ContentCard } from '@/components/ui/ContentCard';
import { TerminalWindow } from '@/components/ui/TerminalWindow';

interface ContentSearchProps {
  className?: string;
  placeholder?: string;
  showResults?: boolean;
}

export function ContentSearch({ 
  className = '', 
  placeholder = 'Search the Agentard archives...',
  showResults = true 
}: ContentSearchProps) {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState<ContentDocument[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [totalContent, setTotalContent] = useState(0);

  useEffect(() => {
    // Get total content count on mount
    getAllContentAPI().then(allContent => {
      setTotalContent(allContent.length);
    }).catch(() => {
      setTotalContent(0);
    });
  }, []);

  useEffect(() => {
    if (!query.trim()) {
      setResults([]);
      return;
    }

    setIsLoading(true);
    // Debounce search
    const timeoutId = setTimeout(async () => {
      try {
        const searchResults = await searchContentAPI(query);
        setResults(searchResults);
      } catch (error) {
        console.error('Search failed:', error);
        setResults([]);
      } finally {
        setIsLoading(false);
      }
    }, 300);

    return () => clearTimeout(timeoutId);
  }, [query]);

  return (
    <div className={`space-y-4 ${className}`}>
      <TerminalWindow title="content_search.sh">
        <div className="space-y-2">
          <div className="flex items-center space-x-2 text-sm">
            <span className="text-terminal-cyan">agentard@archives:~$</span>
            <span className="text-text-muted">grep -r</span>
            <input
              type="text"
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              placeholder={placeholder}
              className="flex-1 bg-transparent text-text-primary outline-none placeholder-text-muted"
            />
          </div>
          
          {query && (
            <div className="text-xs text-text-muted">
              {isLoading ? (
                <span className="animate-pulse">Searching archives...</span>
              ) : (
                <span>
                  Found {results.length} matches in {totalContent} documents
                </span>
              )}
            </div>
          )}
        </div>
      </TerminalWindow>

      {showResults && query && results.length > 0 && (
        <div className="space-y-3">
          <h3 className="text-lg font-bold text-text-secondary">
            Search Results ({results.length})
          </h3>
          <div className="grid gap-3">
            {results.map((content) => {
              // Determine content type and metadata
              let type: 'chapter' | 'course' | 'document' | 'manifesto' = 'document';
              const metadata: Record<string, unknown> = {};

              if (content.path.includes('chronicles')) {
                type = 'chapter';
                const chapterMatch = content.path.match(/chapter_(\d+)/);
                if (chapterMatch) {
                  metadata.number = parseInt(chapterMatch[1]);
                }
              } else if (content.path.includes('university')) {
                type = 'course';
                metadata.difficulty = content.metadata.difficulty;
              } else if (content.path.includes('manifesto')) {
                type = 'manifesto';
              }

              // Add reading time estimate
              const wordCount = content.content.split(/\s+/).length;
              metadata.readingTime = Math.ceil(wordCount / 200);
              metadata.tags = content.metadata.tags;

              return (
                <ContentCard
                  key={content.slug}
                  title={content.metadata.title}
                  description={content.metadata.description}
                  href={`/${content.path.split('/')[0]}/${content.slug}`}
                  type={type}
                  metadata={metadata}
                  compact={true}
                  showTags={true}
                />
              );
            })}
          </div>
        </div>
      )}

      {showResults && query && results.length === 0 && !isLoading && (
        <TerminalWindow title="search_error.log">
          <div className="text-terminal-red space-y-2">
            <div>ERROR: No matches found for &quot;{query}&quot;</div>
            <div className="text-text-muted text-sm">
              Try searching for:
              <ul className="ml-4 mt-1">
                <li>• &quot;port&quot; - for engineering disasters</li>
                <li>• &quot;consciousness&quot; - for philosophical content</li>
                <li>• &quot;authentication&quot; - for security failures</li>
                <li>• &quot;manifesto&quot; - for core principles</li>
              </ul>
            </div>
          </div>
        </TerminalWindow>
      )}
    </div>
  );
}