'use client';

import React, { lazy, Suspense } from 'react';
import { TerminalWindow } from '@/components/ui/TerminalWindow';

// Lazy load heavy interactive components
const FailureCascadeVisualizer = lazy(() => 
  import('@/components/interactive/FailureCascadeVisualizer').then(module => ({
    default: module.FailureCascadeVisualizer
  }))
);

const ConfidenceRealityMeters = lazy(() =>
  import('@/components/interactive/ConfidenceRealityMeters').then(module => ({
    default: module.ConfidenceRealityMeters
  }))
);

const InteractiveTerminal = lazy(() =>
  import('@/components/interactive/InteractiveTerminal').then(module => ({
    default: module.InteractiveTerminal
  }))
);

const EasterEggs = lazy(() =>
  import('@/components/interactive/EasterEggs').then(module => ({
    default: module.EasterEggs
  }))
);

// Loading fallback component
const ComponentLoader = ({ title }: { title: string }) => (
  <TerminalWindow title={`loading_${title.toLowerCase().replace(/\s+/g, '_')}.sh`}>
    <div className="animate-pulse space-y-2">
      <div className="flex items-center space-x-2">
        <span className="text-terminal-cyan">$</span>
        <span className="text-text-muted">Loading {title}...</span>
      </div>
      <div className="text-xs text-text-muted">
        <span className="animate-ping">●</span> Initializing component
      </div>
    </div>
  </TerminalWindow>
);

// Wrapped components with Suspense
export const LazyFailureCascadeVisualizer = () => (
  <Suspense fallback={<ComponentLoader title="Failure Cascade Visualizer" />}>
    <FailureCascadeVisualizer />
  </Suspense>
);

export const LazyConfidenceRealityMeters = () => (
  <Suspense fallback={<ComponentLoader title="Confidence Reality Meters" />}>
    <ConfidenceRealityMeters />
  </Suspense>
);

export const LazyInteractiveTerminal = () => (
  <Suspense fallback={<ComponentLoader title="Interactive Terminal" />}>
    <InteractiveTerminal />
  </Suspense>
);

export const LazyEasterEggs = () => (
  <Suspense fallback={<ComponentLoader title="Easter Eggs" />}>
    <EasterEggs />
  </Suspense>
);