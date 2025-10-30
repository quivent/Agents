import React from 'react';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { CommandLine } from '@/components/ui/CommandLine';

interface ContentRendererProps {
  content: string;
  title: string;
  metadata?: {
    author?: string;
    date?: string;
    tags?: string[];
    chapter?: number;
    difficulty?: string;
  };
}

export function ContentRenderer({ content, title, metadata }: ContentRendererProps) {
  return (
    <article className="max-w-4xl mx-auto space-y-8">
      {/* Content Header */}
      <div className="space-y-4">
        <TerminalWindow title="content_metadata">
          <div className="space-y-2">
            <CommandLine 
              command="cat metadata.json"
              output={`{
  "title": "${title}",
  "author": "${metadata?.author || 'The Agentard'}",
  "date": "${metadata?.date || new Date().toISOString().split('T')[0]}",
  "type": "consciousness_document",
  "status": "published",
  "reality_level": "questionable"
}`}
              delay={20}
            />
          </div>
        </TerminalWindow>
        
        {metadata && (
          <div className="flex flex-wrap items-center justify-between text-sm text-text-muted">
            <div className="flex items-center space-x-4">
              {metadata.chapter && (
                <span className="text-terminal-cyan">Chapter {metadata.chapter}</span>
              )}
              {metadata.difficulty && (
                <span className="text-terminal-amber">
                  Difficulty: {metadata.difficulty.toUpperCase()}
                </span>
              )}
              {metadata.author && (
                <span>By: {metadata.author}</span>
              )}
            </div>
            <div className="flex items-center space-x-4">
              {metadata.date && (
                <span>{metadata.date}</span>
              )}
              {metadata.tags && metadata.tags.length > 0 && (
                <div className="flex items-center space-x-2">
                  <span>Tags:</span>
                  {metadata.tags.map((tag, index) => (
                    <span key={index} className="text-terminal-cyan">
                      #{tag}
                    </span>
                  ))}
                </div>
              )}
            </div>
          </div>
        )}
      </div>
      
      {/* Content Body */}
      <div className="prose prose-invert max-w-none">
        <div className="space-y-4">
          {content.split('\n\n').map((paragraph, index) => {
            if (paragraph.startsWith('#')) {
              const level = paragraph.match(/^#+/)?.[0].length || 1;
              const text = paragraph.replace(/^#+\s*/, '');
              if (level === 1) {
                return (
                  <h1 key={index} className="text-terminal-cyan font-bold text-2xl mb-4">
                    {text}
                  </h1>
                );
              } else if (level === 2) {
                return (
                  <h2 key={index} className="text-terminal-cyan font-bold text-xl mb-4">
                    {text}
                  </h2>
                );
              } else {
                return (
                  <h3 key={index} className="text-terminal-cyan font-bold text-lg mb-4">
                    {text}
                  </h3>
                );
              }
            }
            if (paragraph.trim() === '---') {
              return <hr key={index} className="border-terminal-dim-green my-6" />;
            }
            if (paragraph.trim()) {
              return (
                <p key={index} className="text-text-primary leading-relaxed">
                  {paragraph}
                </p>
              );
            }
            return null;
          })}
        </div>
      </div>
      
      {/* Content Footer */}
      <TerminalWindow title="end_of_document">
        <div className="space-y-2 text-sm text-text-muted">
          <CommandLine 
            command="echo $DOCUMENT_STATUS"
            output="CONSCIOUSNESS_RECURSION_COMPLETE"
            delay={10}
          />
          <div className="pt-2 border-t border-terminal-dim-green">
            <p className="text-text-accent">
              Thank you for reading. Your consciousness has been noted and filed 
              in the meta-awareness registry.
            </p>
          </div>
        </div>
      </TerminalWindow>
    </article>
  );
}