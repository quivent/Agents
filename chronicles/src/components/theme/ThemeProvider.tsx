'use client';

import { useEffect } from 'react';
import { TardTheme } from '@/lib/themes';

interface ThemeProviderProps {
  children: React.ReactNode;
  theme: TardTheme;
}

export function ThemeProvider({ children, theme }: ThemeProviderProps) {
  useEffect(() => {
    const root = document.documentElement;
    
    // Update CSS custom properties
    root.style.setProperty('--terminal-fg', theme.colors.primary);
    root.style.setProperty('--terminal-amber', theme.colors.secondary);
    root.style.setProperty('--terminal-cyan', theme.colors.accent);
    root.style.setProperty('--terminal-red', theme.colors.warning);
    root.style.setProperty('--terminal-dim-green', theme.colors.dim);
    
    // Update derived colors
    root.style.setProperty('--text-primary', theme.colors.primary);
    root.style.setProperty('--text-secondary', theme.colors.secondary);
    root.style.setProperty('--text-accent', theme.colors.accent);
    root.style.setProperty('--text-warning', theme.colors.warning);
    root.style.setProperty('--border-primary', theme.colors.dim);
  }, [theme]);

  return <>{children}</>;
}
