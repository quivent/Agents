import Link from 'next/link';
import { getAllTardIds, getTardTheme } from '@/lib/themes';
import { TerminalWindow } from './TerminalWindow';

export function TardSelector() {
  const tardIds = getAllTardIds();

  return (
    <div className="space-y-8">
      <h2 className="text-3xl font-bold text-text-secondary text-center py-8">
        &gt; SELECT_YOUR_TARD
      </h2>
      
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 px-4">
        {tardIds.map((tardId) => {
          const theme = getTardTheme(tardId);
          return (
            <Link key={tardId} href={`/${tardId}`} className="block group">
              <TerminalWindow 
                title={tardId} 
                className="h-full hover:border-text-accent transition-all duration-300 group-hover:scale-105 cursor-pointer"
              >
                <div className="space-y-3 text-center p-3 min-h-[100px] flex flex-col justify-center">
                  <div 
                    className="text-sm md:text-base font-bold group-hover:animate-pulse-glow transition-all leading-tight"
                    style={{ color: theme.colors.primary }}
                  >
                    {theme.name}
                  </div>
                  <div className="text-text-muted text-xs">
                    {theme.tagline}
                  </div>
                  <div 
                    className="text-xs font-bold opacity-70 group-hover:opacity-100 transition-opacity"
                    style={{ color: theme.colors.secondary }}
                  >
                    → ENTER ←
                  </div>
                </div>
              </TerminalWindow>
            </Link>
          );
        })}
      </div>
    </div>
  );
}
