import React from 'react';
import Link from 'next/link';
import { getChronicles, getUniversityCourses } from '@/lib/content';
import { TerminalWindow } from '@/components/ui/TerminalWindow';

interface SeriesNavigationProps {
  currentSlug: string;
  section: 'chronicles' | 'university';
  className?: string;
}

export function SeriesNavigation({ currentSlug, section, className = '' }: SeriesNavigationProps) {
  const getSeriesData = () => {
    if (section === 'chronicles') {
      const chapters = getChronicles();
      const currentIndex = chapters.findIndex(ch => ch.slug === currentSlug);
      const current = chapters[currentIndex];
      const previous = currentIndex > 0 ? chapters[currentIndex - 1] : null;
      const next = currentIndex < chapters.length - 1 ? chapters[currentIndex + 1] : null;
      
      return {
        current: current ? { ...current, title: current.title, type: 'chapter' } : null,
        previous: previous ? { ...previous, title: previous.title, type: 'chapter' } : null,
        next: next ? { ...next, title: next.title, type: 'chapter' } : null,
        total: chapters.length,
        position: currentIndex + 1
      };
    } else {
      const courses = getUniversityCourses();
      const currentIndex = courses.findIndex(course => course.slug === currentSlug);
      const current = courses[currentIndex];
      const previous = currentIndex > 0 ? courses[currentIndex - 1] : null;
      const next = currentIndex < courses.length - 1 ? courses[currentIndex + 1] : null;
      
      return {
        current: current ? { ...current, title: `${current.code}: ${current.title}`, type: 'course' } : null,
        previous: previous ? { ...previous, title: `${previous.code}: ${previous.title}`, type: 'course' } : null,
        next: next ? { ...next, title: `${next.code}: ${next.title}`, type: 'course' } : null,
        total: courses.length,
        position: currentIndex + 1
      };
    }
  };

  const seriesData = getSeriesData();

  if (!seriesData.current) {
    return null;
  }

  const sectionPath = section;
  const sectionTitle = section === 'chronicles' ? 'Chronicles' : 'University';

  return (
    <div className={`space-y-4 ${className}`}>
      {/* Series Overview */}
      <TerminalWindow title={`${section}_navigation.sh`}>
        <div className="space-y-3">
          {/* Current Position */}
          <div className="flex items-center justify-between">
            <div className="space-y-1">
              <h3 className="text-lg font-bold text-text-primary">
                {sectionTitle} Navigation
              </h3>
              <div className="flex items-center space-x-3 text-sm">
                <span className="text-terminal-cyan">
                  Position: {seriesData.position}/{seriesData.total}
                </span>
                <span className="text-text-muted">|</span>
                <Link 
                  href={`/${sectionPath}`}
                  className="text-text-accent hover:text-text-primary transition-colors"
                >
                  View All {sectionTitle}
                </Link>
              </div>
            </div>
            
            {/* Progress Bar */}
            <div className="w-32 bg-bg-tertiary h-2 rounded">
              <div 
                className="bg-terminal-cyan h-2 rounded transition-all duration-300"
                style={{ width: `${(seriesData.position / seriesData.total) * 100}%` }}
              />
            </div>
          </div>

          {/* Navigation Links */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            {/* Previous */}
            <div className="space-y-2">
              {seriesData.previous ? (
                <Link 
                  href={`/${sectionPath}/${seriesData.previous.slug}`}
                  className="block p-3 bg-bg-secondary hover:bg-bg-tertiary border border-terminal-dim-green hover:border-terminal-cyan transition-all group"
                >
                  <div className="flex items-center space-x-2 mb-1">
                    <span className="text-terminal-cyan">&lt;&lt;</span>
                    <span className="text-xs text-text-muted uppercase tracking-wide">
                      Previous {seriesData.previous.type}
                    </span>
                  </div>
                  <h4 className="text-text-primary group-hover:text-terminal-cyan transition-colors font-mono text-sm">
                    {seriesData.previous.title}
                  </h4>
                </Link>
              ) : (
                <div className="p-3 bg-bg-quaternary border border-terminal-dim-green opacity-50">
                  <div className="flex items-center space-x-2 mb-1">
                    <span className="text-terminal-dim-green">&lt;&lt;</span>
                    <span className="text-xs text-text-muted uppercase tracking-wide">
                      No Previous {seriesData.current.type}
                    </span>
                  </div>
                  <div className="text-text-muted font-mono text-sm">
                    You're at the beginning
                  </div>
                </div>
              )}
            </div>

            {/* Next */}
            <div className="space-y-2">
              {seriesData.next ? (
                <Link 
                  href={`/${sectionPath}/${seriesData.next.slug}`}
                  className="block p-3 bg-bg-secondary hover:bg-bg-tertiary border border-terminal-dim-green hover:border-terminal-cyan transition-all group"
                >
                  <div className="flex items-center justify-end space-x-2 mb-1">
                    <span className="text-xs text-text-muted uppercase tracking-wide">
                      Next {seriesData.next.type}
                    </span>
                    <span className="text-terminal-cyan">&gt;&gt;</span>
                  </div>
                  <h4 className="text-text-primary group-hover:text-terminal-cyan transition-colors font-mono text-sm text-right">
                    {seriesData.next.title}
                  </h4>
                </Link>
              ) : (
                <div className="p-3 bg-bg-quaternary border border-terminal-dim-green opacity-50">
                  <div className="flex items-center justify-end space-x-2 mb-1">
                    <span className="text-xs text-text-muted uppercase tracking-wide">
                      No Next {seriesData.current.type}
                    </span>
                    <span className="text-terminal-dim-green">&gt;&gt;</span>
                  </div>
                  <div className="text-text-muted font-mono text-sm text-right">
                    {section === 'chronicles' ? 'The journey continues...' : 'Curriculum complete'}
                  </div>
                </div>
              )}
            </div>
          </div>

          {/* Series Overview Link */}
          <div className="pt-3 border-t border-terminal-dim-green">
            <Link 
              href={`/${sectionPath}`}
              className="inline-flex items-center space-x-2 text-text-accent hover:text-text-primary transition-colors text-sm"
            >
              <span>&gt;</span>
              <span>VIEW_{sectionTitle.toUpperCase()}_INDEX</span>
            </Link>
          </div>
        </div>
      </TerminalWindow>
    </div>
  );
}