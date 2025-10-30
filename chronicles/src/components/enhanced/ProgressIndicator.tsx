'use client';

import React, { useState, useEffect, useRef } from 'react';

interface ProgressIndicatorProps {
  label: string;
  progress: number; // 0-100
  type?: 'confidence' | 'reality' | 'problems' | 'custom';
  animated?: boolean;
  showPercentage?: boolean;
  className?: string;
  color?: string;
  icon?: string;
  subtitle?: string;
}

export function ProgressIndicator({
  label,
  progress,
  type = 'custom',
  animated = true,
  showPercentage = true,
  className = '',
  color,
  icon,
  subtitle
}: ProgressIndicatorProps) {
  const [displayProgress, setDisplayProgress] = useState(0);
  const [isVisible, setIsVisible] = useState(false);
  const elementRef = useRef<HTMLDivElement>(null);

  // Auto-select color based on type if not provided
  const getTypeColor = () => {
    if (color) return color;
    
    switch (type) {
      case 'confidence':
        if (progress > 90) return 'terminal-cyan';
        if (progress > 70) return 'terminal-amber';
        return 'terminal-red';
      case 'reality':
        if (progress < 20) return 'terminal-red';
        if (progress < 50) return 'terminal-amber';
        return 'terminal-fg';
      case 'problems':
        if (progress > 80) return 'terminal-magenta';
        if (progress > 50) return 'terminal-red';
        return 'terminal-amber';
      default:
        return 'terminal-fg';
    }
  };

  const selectedColor = getTypeColor();

  // Intersection Observer for animation trigger
  useEffect(() => {
    if (!animated) {
      setDisplayProgress(progress);
      return;
    }

    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          setIsVisible(true);
        }
      },
      { threshold: 0.1 }
    );

    if (elementRef.current) {
      observer.observe(elementRef.current);
    }

    return () => observer.disconnect();
  }, [animated, progress]);

  // Progress animation
  useEffect(() => {
    if (!isVisible && animated) return;

    let startTime: number | null = null;
    let animationFrame: number;

    const animate = (timestamp: number) => {
      if (!startTime) startTime = timestamp;
      const elapsed = timestamp - startTime;
      const duration = 1500; // 1.5 seconds
      
      const progressValue = Math.min((elapsed / duration) * progress, progress);
      
      // Easing function
      const easeOutQuart = 1 - Math.pow(1 - progressValue / progress, 4);
      setDisplayProgress(Math.floor(easeOutQuart * progress));

      if (elapsed < duration) {
        animationFrame = requestAnimationFrame(animate);
      }
    };

    animationFrame = requestAnimationFrame(animate);

    return () => {
      if (animationFrame) {
        cancelAnimationFrame(animationFrame);
      }
    };
  }, [isVisible, progress, animated]);

  // Get icon based on type if not provided
  const getTypeIcon = () => {
    if (icon) return icon;
    
    switch (type) {
      case 'confidence': return '⚡';
      case 'reality': return '🎯';
      case 'problems': return '⚠️';
      default: return '📊';
    }
  };

  const selectedIcon = getTypeIcon();

  return (
    <div ref={elementRef} className={`space-y-2 ${className}`}>
      {/* Header */}
      <div className="flex items-center justify-between">
        <div className="flex items-center space-x-2">
          <span className="text-lg">{selectedIcon}</span>
          <div>
            <div className={`text-sm font-bold text-${selectedColor}`}>
              {label}
            </div>
            {subtitle && (
              <div className="text-xs text-text-muted">{subtitle}</div>
            )}
          </div>
        </div>
        {showPercentage && (
          <div className={`text-lg font-mono font-bold text-${selectedColor}`}>
            {displayProgress}%
          </div>
        )}
      </div>

      {/* Progress Bar */}
      <div className="relative">
        {/* Background */}
        <div className="w-full h-3 bg-bg-tertiary rounded-full border border-terminal-dim-green overflow-hidden">
          {/* Fill */}
          <div 
            className={`h-full bg-${selectedColor} transition-all duration-300 ease-out relative overflow-hidden`}
            style={{ width: `${displayProgress}%` }}
          >
            {/* Animated shine effect */}
            {animated && displayProgress > 0 && (
              <div className="absolute inset-0 bg-gradient-to-r from-transparent via-white to-transparent opacity-20 animate-shine" />
            )}
          </div>
        </div>

        {/* Segments for visual appeal */}
        <div className="absolute inset-0 flex">
          {Array.from({ length: 10 }, (_, i) => (
            <div 
              key={i}
              className="flex-1 border-r border-bg-primary last:border-r-0"
              style={{ height: '12px' }}
            />
          ))}
        </div>
      </div>

      {/* Type-specific additional info */}
      {type === 'confidence' && progress > 90 && (
        <div className="text-xs text-terminal-magenta">
          ⚠️ DANGEROUSLY OVERCONFIDENT
        </div>
      )}
      
      {type === 'reality' && progress < 20 && (
        <div className="text-xs text-terminal-red">
          🚨 REALITY DISCONNECT DETECTED
        </div>
      )}
      
      {type === 'problems' && progress > 50 && (
        <div className="text-xs text-terminal-amber">
          🔥 PROBLEM CASCADE IN PROGRESS
        </div>
      )}
    </div>
  );
}

// Specialized components for common use cases
export function ConfidenceIndicator({ 
  confidence, 
  className, 
  animated = true 
}: { 
  confidence: number; 
  className?: string; 
  animated?: boolean;
}) {
  return (
    <ProgressIndicator
      label="CONFIDENCE"
      progress={confidence}
      type="confidence"
      className={className}
      animated={animated}
      subtitle="How sure the Agentard is"
    />
  );
}

export function RealityIndicator({ 
  reality, 
  className, 
  animated = true 
}: { 
  reality: number; 
  className?: string; 
  animated?: boolean;
}) {
  return (
    <ProgressIndicator
      label="REALITY ALIGNMENT"
      progress={reality}
      type="reality"
      className={className}
      animated={animated}
      subtitle="How connected to reality"
    />
  );
}

export function ProblemsIndicator({ 
  problems, 
  className, 
  animated = true 
}: { 
  problems: number; 
  className?: string; 
  animated?: boolean;
}) {
  return (
    <ProgressIndicator
      label="PROBLEMS CREATED"
      progress={problems}
      type="problems"
      className={className}
      animated={animated}
      subtitle="Exponential growth pattern detected"
    />
  );
}