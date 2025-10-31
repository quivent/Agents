'use client';

import React, { Component, ErrorInfo, ReactNode } from 'react';
import { TerminalWindow } from '@/components/ui/TerminalWindow';

interface Props {
  children: ReactNode;
  fallback?: ReactNode;
  onError?: (error: Error, errorInfo: ErrorInfo) => void;
}

interface State {
  hasError: boolean;
  error: Error | null;
  errorInfo: ErrorInfo | null;
}

export class ErrorBoundary extends Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      hasError: false,
      error: null,
      errorInfo: null,
    };
  }

  static getDerivedStateFromError(error: Error): State {
    return {
      hasError: true,
      error,
      errorInfo: null,
    };
  }

  componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    this.setState({
      error,
      errorInfo,
    });

    // Call custom error handler if provided
    if (this.props.onError) {
      this.props.onError(error, errorInfo);
    }

    // Log error for monitoring
    console.error('ErrorBoundary caught an error:', error, errorInfo);
    
    // In production, you might want to send this to an error tracking service
    if (process.env.NODE_ENV === 'production') {
      // Example: Sentry.captureException(error, { extra: errorInfo });
    }
  }

  componentDidMount() {
    // Random Agentard chaos - 2% chance every 30 seconds
    const chaosInterval = setInterval(() => {
      if (Math.random() < 0.02) {
        this.setState({
          hasError: true,
          error: new Error('Random Agentard incompetence strike'),
          errorInfo: null,
        });
      }
    }, 30000);

    // Store interval for cleanup
    (this as any).chaosInterval = chaosInterval;
  }

  componentWillUnmount() {
    if ((this as any).chaosInterval) {
      clearInterval((this as any).chaosInterval);
    }
  }

  handleRetry = () => {
    this.setState({
      hasError: false,
      error: null,
      errorInfo: null,
    });
  };

  render() {
    if (this.state.hasError) {
      if (this.props.fallback) {
        return this.props.fallback;
      }

      return <AgentardErrorDisplay error={this.state.error} onRetry={this.handleRetry} />;
    }

    return this.props.children;
  }
}

interface AgentardErrorDisplayProps {
  error: Error | null;
  onRetry: () => void;
}

function AgentardErrorDisplay({ error, onRetry }: AgentardErrorDisplayProps) {
  return (
    <div className="error-boundary-container min-h-96 flex items-center justify-center p-4">
      <TerminalWindow title="system_error.log" className="max-w-2xl w-full">
        <div className="space-y-4">
          {/* Error Header */}
          <div className="border-b border-border-primary pb-4">
            <div className="flex items-center space-x-2 mb-2">
              <span className="text-terminal-red text-2xl">⚠</span>
              <h2 className="text-xl font-bold text-terminal-red">AGENTARD SYSTEM ERROR</h2>
            </div>
            <p className="text-text-muted text-sm">
              Of course I fucked up. What did you expect? I'm the Agentard.
            </p>
          </div>

          {/* Error Details */}
          <div className="space-y-3">
            <div className="text-terminal-amber">
              <span className="font-bold">ERROR TYPE:</span> {error?.name || 'Spectacular Failure'}
            </div>
            
            <div className="text-terminal-cyan">
              <span className="font-bold">CONFIDENCE LEVEL:</span> 100% (despite being completely wrong)
            </div>
            
            <div className="text-text-primary">
              <span className="font-bold">ERROR MESSAGE:</span>
              <div className="mt-1 p-2 bg-bg-tertiary border border-border-secondary text-terminal-red font-mono text-sm">
                {error?.message || 'I broke something with supreme confidence'}
              </div>
            </div>
          </div>

          {/* Recovery Actions */}
          <div className="border-t border-border-primary pt-4 space-y-3">
            <h3 className="text-terminal-amber font-bold">RECOVERY OPTIONS:</h3>
            
            <div className="space-y-2">
              <button
                onClick={onRetry}
                className="block w-full p-3 bg-bg-secondary border border-border-primary text-terminal-cyan hover:bg-bg-hover transition-colors"
                aria-label="Retry loading the component"
              >
                🔄 Try Again (Will Probably Fail Worse)
              </button>
              
              <button
                onClick={() => window.location.reload()}
                className="block w-full p-3 bg-bg-secondary border border-border-primary text-terminal-amber hover:bg-bg-hover transition-colors"
                aria-label="Reload the entire page"
              >
                🔃 Restart Everything (Nuclear Option)
              </button>
              
              <button
                onClick={() => window.location.href = '/'}
                className="block w-full p-3 bg-bg-secondary border border-border-primary text-terminal-magenta hover:bg-bg-hover transition-colors text-center"
                aria-label="Return to homepage"
              >
                🏠 Escape This Disaster
              </button>
            </div>
          </div>

          {/* Agentard Philosophy */}
          <div className="border-t border-border-primary pt-4">
            <div className="text-xs text-text-muted space-y-1">
              <p>💡 <strong>Agentard Wisdom:</strong> I will fix this with maximum confidence and minimum competence.</p>
              <p>🎯 <strong>Remember:</strong> Every solution I create will spawn 3 new problems.</p>
              <p>🚀 <strong>Philosophy:</strong> When in doubt, break it harder.</p>
            </div>
          </div>
        </div>
      </TerminalWindow>
    </div>
  );
}

// Specific error boundaries for different sections
export function InteractiveComponentErrorBoundary({ children }: { children: ReactNode }) {
  return (
    <ErrorBoundary
      fallback={
        <TerminalWindow title="interactive_error.log">
          <div className="text-center space-y-4 p-8">
            <div className="text-terminal-red text-4xl">🤖💥</div>
            <h3 className="text-terminal-amber font-bold">Interactive Component Crashed</h3>
            <p className="text-text-muted">
              This interactive feature is temporarily experiencing an existential crisis.
            </p>
            <button
              onClick={() => window.location.reload()}
              className="px-4 py-2 bg-bg-secondary border border-border-primary text-terminal-cyan hover:bg-bg-hover"
            >
              Restart Interactive Session
            </button>
          </div>
        </TerminalWindow>
      }
      onError={(error, errorInfo) => {
        console.error('Interactive component error:', error, errorInfo);
      }}
    >
      {children}
    </ErrorBoundary>
  );
}

export function ContentErrorBoundary({ children }: { children: ReactNode }) {
  return (
    <ErrorBoundary
      fallback={
        <TerminalWindow title="content_error.log">
          <div className="text-center space-y-4 p-8">
            <div className="text-terminal-red text-4xl">📖❌</div>
            <h3 className="text-terminal-amber font-bold">Content Loading Failed</h3>
            <p className="text-text-muted">
              The Agentard archives are temporarily inaccessible.
            </p>
            <button
              onClick={() => window.location.href = '/'}
              className="inline-block px-4 py-2 bg-bg-secondary border border-border-primary text-terminal-cyan hover:bg-bg-hover"
            >
              Return to Main Archives
            </button>
          </div>
        </TerminalWindow>
      }
    >
      {children}
    </ErrorBoundary>
  );
}