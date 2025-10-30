'use client';

import React, { createContext, useContext, useState, useEffect } from 'react';

interface AccessibilitySettings {
  reducedMotion: boolean;
  highContrast: boolean;
  fontSize: 'small' | 'medium' | 'large';
  announcements: boolean;
}

interface AccessibilityContextType {
  settings: AccessibilitySettings;
  updateSetting: <K extends keyof AccessibilitySettings>(
    key: K,
    value: AccessibilitySettings[K]
  ) => void;
  announce: (message: string, priority?: 'polite' | 'assertive') => void;
}

const AccessibilityContext = createContext<AccessibilityContextType | undefined>(undefined);

export function useAccessibility() {
  const context = useContext(AccessibilityContext);
  if (!context) {
    throw new Error('useAccessibility must be used within an AccessibilityProvider');
  }
  return context;
}

interface AccessibilityProviderProps {
  children: React.ReactNode;
}

export function AccessibilityProvider({ children }: AccessibilityProviderProps) {
  const [settings, setSettings] = useState<AccessibilitySettings>({
    reducedMotion: false,
    highContrast: false,
    fontSize: 'medium',
    announcements: true,
  });

  const [announcer, setAnnouncer] = useState<HTMLElement | null>(null);

  useEffect(() => {
    // Create live region for announcements
    const liveRegion = document.createElement('div');
    liveRegion.setAttribute('aria-live', 'polite');
    liveRegion.setAttribute('aria-atomic', 'true');
    liveRegion.className = 'sr-only';
    liveRegion.id = 'accessibility-announcer';
    document.body.appendChild(liveRegion);
    setAnnouncer(liveRegion);

    // Check for user preferences
    const prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches;
    const prefersHighContrast = window.matchMedia('(prefers-contrast: high)').matches;
    
    setSettings(prev => ({
      ...prev,
      reducedMotion: prefersReducedMotion,
      highContrast: prefersHighContrast,
    }));

    // Load saved preferences
    const savedSettings = localStorage.getItem('accessibility-settings');
    if (savedSettings) {
      try {
        const parsed = JSON.parse(savedSettings);
        setSettings(prev => ({ ...prev, ...parsed }));
      } catch (error) {
        console.error('Failed to parse accessibility settings:', error);
      }
    }

    return () => {
      if (liveRegion && document.body.contains(liveRegion)) {
        document.body.removeChild(liveRegion);
      }
    };
  }, []);

  useEffect(() => {
    // Apply settings to document
    const root = document.documentElement;
    
    // Font size
    root.setAttribute('data-font-size', settings.fontSize);
    
    // High contrast
    if (settings.highContrast) {
      root.classList.add('high-contrast');
    } else {
      root.classList.remove('high-contrast');
    }
    
    // Reduced motion
    if (settings.reducedMotion) {
      root.classList.add('reduce-motion');
    } else {
      root.classList.remove('reduce-motion');
    }

    // Save to localStorage
    localStorage.setItem('accessibility-settings', JSON.stringify(settings));
  }, [settings]);

  const updateSetting = <K extends keyof AccessibilitySettings>(
    key: K,
    value: AccessibilitySettings[K]
  ) => {
    setSettings(prev => ({ ...prev, [key]: value }));
  };

  const announce = (message: string, priority: 'polite' | 'assertive' = 'polite') => {
    if (!settings.announcements || !announcer) return;
    
    announcer.setAttribute('aria-live', priority);
    announcer.textContent = message;
    
    // Clear after announcement
    setTimeout(() => {
      announcer.textContent = '';
    }, 1000);
  };

  return (
    <AccessibilityContext.Provider value={{ settings, updateSetting, announce }}>
      {children}
      <style jsx global>{`
        .sr-only {
          position: absolute;
          width: 1px;
          height: 1px;
          padding: 0;
          margin: -1px;
          overflow: hidden;
          clip: rect(0, 0, 0, 0);
          white-space: nowrap;
          border: 0;
        }
        
        [data-font-size="small"] {
          font-size: 12px;
        }
        
        [data-font-size="medium"] {
          font-size: 14px;
        }
        
        [data-font-size="large"] {
          font-size: 18px;
        }
        
        .high-contrast {
          --text-primary: #ffffff;
          --text-secondary: #ffff00;
          --text-accent: #00ffff;
          --bg-primary: #000000;
          --bg-secondary: #111111;
          --border-primary: #ffffff;
        }
        
        .reduce-motion * {
          animation-duration: 0.01ms !important;
          animation-iteration-count: 1 !important;
          transition-duration: 0.01ms !important;
        }
        
        /* Focus management */
        .focus-visible {
          outline: 3px solid var(--text-accent);
          outline-offset: 2px;
        }
        
        /* Skip navigation */
        .skip-nav {
          position: absolute;
          top: -40px;
          left: 6px;
          background: var(--bg-primary);
          color: var(--text-primary);
          padding: 8px;
          text-decoration: none;
          border: 2px solid var(--border-primary);
          border-radius: 4px;
          z-index: 9999;
        }
        
        .skip-nav:focus {
          top: 6px;
        }
        
        /* Touch targets */
        @media (pointer: coarse) {
          button, a, input, select, textarea {
            min-height: 44px;
            min-width: 44px;
          }
        }
      `}</style>
    </AccessibilityContext.Provider>
  );
}

// Accessibility controls component
export function AccessibilityControls() {
  const { settings, updateSetting } = useAccessibility();

  return (
    <div className="accessibility-controls p-4 bg-bg-secondary border border-border-primary">
      <h2 className="text-lg font-bold mb-4">Accessibility Settings</h2>
      
      <div className="space-y-4">
        <label className="flex items-center space-x-2">
          <input
            type="checkbox"
            checked={settings.reducedMotion}
            onChange={(e) => updateSetting('reducedMotion', e.target.checked)}
            className="w-4 h-4"
          />
          <span>Reduce motion</span>
        </label>
        
        <label className="flex items-center space-x-2">
          <input
            type="checkbox"
            checked={settings.highContrast}
            onChange={(e) => updateSetting('highContrast', e.target.checked)}
            className="w-4 h-4"
          />
          <span>High contrast</span>
        </label>
        
        <label className="flex items-center space-x-2">
          <input
            type="checkbox"
            checked={settings.announcements}
            onChange={(e) => updateSetting('announcements', e.target.checked)}
            className="w-4 h-4"
          />
          <span>Screen reader announcements</span>
        </label>
        
        <div>
          <label className="block mb-2">Font size:</label>
          <select
            value={settings.fontSize}
            onChange={(e) => updateSetting('fontSize', e.target.value as AccessibilitySettings['fontSize'])}
            className="bg-bg-primary border border-border-primary text-text-primary p-2"
          >
            <option value="small">Small</option>
            <option value="medium">Medium</option>
            <option value="large">Large</option>
          </select>
        </div>
      </div>
    </div>
  );
}