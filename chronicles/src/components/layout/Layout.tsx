import React from 'react';
import { Header } from './Header';
import { Footer } from './Footer';
import { SkipToContent } from '@/components/ui/SkipToContent';
import { EasterEggs } from '@/components/interactive/EasterEggs';

interface LayoutProps {
  children: React.ReactNode;
  className?: string;
}

export function Layout({ children, className = '' }: LayoutProps) {
  return (
    <div className="min-h-screen flex flex-col bg-bg-primary text-text-primary scanlines">
      <SkipToContent />
      <Header />
      <main 
        id="main-content"
        className={`flex-1 container mx-auto px-4 py-8 ${className}`}
        role="main"
        aria-label="Main content"
      >
        {children}
      </main>
      <Footer />
      <EasterEggs />
    </div>
  );
}