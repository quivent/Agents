'use client';

import { TerminalWindow } from '@/components/ui/TerminalWindow';

interface PopularTagsProps {
  tags: string[];
}

export function PopularTags({ tags }: PopularTagsProps) {
  const handleTagClick = (tag: string) => {
    const searchInput = document.querySelector('input[placeholder*="Search"]') as HTMLInputElement;
    if (searchInput) {
      searchInput.value = tag;
      searchInput.dispatchEvent(new Event('input', { bubbles: true }));
    }
  };

  return (
    <TerminalWindow title="popular_tags.sh">
      <div className="space-y-3">
        <h3 className="text-lg font-bold text-text-secondary">
          Popular Tags
        </h3>
        <div className="flex flex-wrap gap-2">
          {tags.map(tag => (
            <button
              key={tag}
              onClick={() => handleTagClick(tag)}
              className="bg-bg-tertiary hover:bg-bg-quaternary text-terminal-cyan px-3 py-1 text-sm font-mono border border-terminal-dim-green hover:border-terminal-cyan transition-all"
            >
              #{tag}
            </button>
          ))}
        </div>
      </div>
    </TerminalWindow>
  );
}
