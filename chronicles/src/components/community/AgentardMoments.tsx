'use client';

import React, { useState, useEffect } from 'react';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { AnimatedCounter } from '@/components/enhanced/AnimatedCounter';

interface AgentardMoment {
  id: string;
  title: string;
  description: string;
  confidence: number;
  reality: number;
  category: 'deployment' | 'debugging' | 'architecture' | 'security' | 'optimization' | 'documentation';
  author: string;
  timestamp: Date;
  votes: number;
  agentardIndex: number;
}

const sampleMoments: AgentardMoment[] = [
  {
    id: '1',
    title: 'Fixed production by deleting error logs',
    description: 'The logs were showing errors, so I deleted the logging system. No more error messages = problem solved!',
    confidence: 98,
    reality: 3,
    category: 'debugging',
    author: 'definitely_senior_dev',
    timestamp: new Date('2024-09-14T15:30:00'),
    votes: 47,
    agentardIndex: 95
  },
  {
    id: '2',
    title: 'Optimized database by removing all constraints',
    description: 'Foreign keys and constraints were slowing down inserts. Removed them all - now everything is blazing fast!',
    confidence: 95,
    reality: 8,
    category: 'optimization',
    author: 'performance_guru',
    timestamp: new Date('2024-09-13T09:15:00'),
    votes: 73,
    agentardIndex: 87
  },
  {
    id: '3',
    title: 'Implemented perfect security',
    description: 'Changed all passwords to "password123" for consistency. Now everyone can access everything easily!',
    confidence: 92,
    reality: 5,
    category: 'security',
    author: 'security_expert',
    timestamp: new Date('2024-09-12T11:45:00'),
    votes: 156,
    agentardIndex: 87
  },
  {
    id: '4',
    title: 'Dockerized everything in production',
    description: 'Put the entire production environment in a single container running as root. Ultimate simplicity!',
    confidence: 96,
    reality: 7,
    category: 'deployment',
    author: 'devops_wizard',
    timestamp: new Date('2024-09-11T16:20:00'),
    votes: 89,
    agentardIndex: 89
  }
];

const categories = [
  { id: 'all', label: 'All Categories', icon: '🌍' },
  { id: 'deployment', label: 'Deployment', icon: '🚀' },
  { id: 'debugging', label: 'Debugging', icon: '🐛' },
  { id: 'architecture', label: 'Architecture', icon: '🏗️' },
  { id: 'security', label: 'Security', icon: '🔒' },
  { id: 'optimization', label: 'Optimization', icon: '⚡' },
  { id: 'documentation', label: 'Documentation', icon: '📚' }
];

export function AgentardMoments() {
  const [moments, setMoments] = useState<AgentardMoment[]>(sampleMoments);
  const [selectedCategory, setSelectedCategory] = useState('all');
  const [newMoment, setNewMoment] = useState({
    title: '',
    description: '',
    confidence: 50,
    category: 'debugging' as AgentardMoment['category'],
    author: ''
  });
  const [showSubmissionForm, setShowSubmissionForm] = useState(false);
  const [sortBy, setSortBy] = useState<'votes' | 'recent' | 'agentard'>('agentard');

  const filteredMoments = moments
    .filter(moment => selectedCategory === 'all' || moment.category === selectedCategory)
    .sort((a, b) => {
      switch (sortBy) {
        case 'votes': return b.votes - a.votes;
        case 'recent': return b.timestamp.getTime() - a.timestamp.getTime();
        case 'agentard': return b.agentardIndex - a.agentardIndex;
        default: return 0;
      }
    });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!newMoment.title || !newMoment.description || !newMoment.author) {
      return;
    }

    // Calculate reality based on confidence (inverse relationship)
    const reality = Math.max(5, Math.floor(100 - (newMoment.confidence * 0.8) + Math.random() * 20));
    const agentardIndex = Math.floor((newMoment.confidence - reality) * (newMoment.confidence / 100));

    const moment: AgentardMoment = {
      id: Date.now().toString(),
      title: newMoment.title,
      description: newMoment.description,
      confidence: newMoment.confidence,
      reality,
      category: newMoment.category,
      author: newMoment.author,
      timestamp: new Date(),
      votes: 1,
      agentardIndex
    };

    setMoments(prev => [moment, ...prev]);
    setNewMoment({
      title: '',
      description: '',
      confidence: 50,
      category: 'debugging',
      author: ''
    });
    setShowSubmissionForm(false);
  };

  const handleVote = (id: string) => {
    setMoments(prev => prev.map(moment => 
      moment.id === id 
        ? { ...moment, votes: moment.votes + 1 }
        : moment
    ));
  };

  const getCategoryColor = (category: string) => {
    const colors = {
      deployment: 'terminal-cyan',
      debugging: 'terminal-red',
      architecture: 'terminal-amber',
      security: 'terminal-magenta',
      optimization: 'terminal-fg',
      documentation: 'terminal-blue'
    };
    return colors[category as keyof typeof colors] || 'terminal-gray';
  };

  const getAgentardLevel = (index: number) => {
    if (index > 80) return { label: 'LEGENDARY', color: 'terminal-magenta' };
    if (index > 60) return { label: 'EPIC', color: 'terminal-red' };
    if (index > 40) return { label: 'HIGH', color: 'terminal-amber' };
    return { label: 'MILD', color: 'terminal-cyan' };
  };

  return (
    <div className="space-y-6">
      {/* Header and Stats */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <TerminalWindow title="community_stats" className="md:col-span-4">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-4 text-center">
            <div>
              <AnimatedCounter 
                value={moments.length} 
                className="text-2xl text-terminal-cyan" 
              />
              <div className="text-xs text-text-muted">Total Moments</div>
            </div>
            <div>
              <AnimatedCounter 
                value={moments.reduce((sum, m) => sum + m.votes, 0)} 
                className="text-2xl text-terminal-amber" 
              />
              <div className="text-xs text-text-muted">Total Votes</div>
            </div>
            <div>
              <AnimatedCounter 
                value={Math.round(moments.reduce((sum, m) => sum + m.agentardIndex, 0) / moments.length)} 
                className="text-2xl text-terminal-magenta" 
              />
              <div className="text-xs text-text-muted">Avg Agentard Index</div>
            </div>
            <div>
              <AnimatedCounter 
                value={Math.round(moments.reduce((sum, m) => sum + m.confidence, 0) / moments.length)} 
                suffix="%" 
                className="text-2xl text-terminal-red" 
              />
              <div className="text-xs text-text-muted">Avg Confidence</div>
            </div>
          </div>
        </TerminalWindow>
      </div>

      {/* Controls */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <TerminalWindow title="categories">
          <div className="space-y-2">
            {categories.map(category => (
              <button
                key={category.id}
                onClick={() => setSelectedCategory(category.id)}
                className={`w-full text-left p-2 rounded text-sm transition-colors ${
                  selectedCategory === category.id
                    ? 'bg-terminal-dim-green bg-opacity-20 border border-terminal-dim-green'
                    : 'hover:bg-bg-tertiary'
                }`}
              >
                <span className="mr-2">{category.icon}</span>
                {category.label}
                {category.id !== 'all' && (
                  <span className="float-right text-text-muted">
                    {moments.filter(m => m.category === category.id).length}
                  </span>
                )}
              </button>
            ))}
          </div>
        </TerminalWindow>

        <TerminalWindow title="sort_options">
          <div className="space-y-2">
            {[
              { key: 'agentard', label: 'Agentard Index', icon: '🏆' },
              { key: 'votes', label: 'Most Voted', icon: '👍' },
              { key: 'recent', label: 'Most Recent', icon: '🕒' }
            ].map(option => (
              <button
                key={option.key}
                onClick={() => setSortBy(option.key as typeof sortBy)}
                className={`w-full text-left p-2 rounded text-sm transition-colors ${
                  sortBy === option.key
                    ? 'bg-terminal-amber bg-opacity-20 border border-terminal-amber'
                    : 'hover:bg-bg-tertiary'
                }`}
              >
                <span className="mr-2">{option.icon}</span>
                {option.label}
              </button>
            ))}
          </div>
        </TerminalWindow>

        <TerminalWindow title="submit_moment">
          <button
            onClick={() => setShowSubmissionForm(true)}
            className="w-full bg-terminal-fg text-terminal-bg p-3 rounded font-bold hover:bg-terminal-cyan transition-colors"
          >
            📝 Share Your Agentard Moment
          </button>
          <div className="text-xs text-text-muted mt-2 text-center">
            Confident incompetence welcomed!
          </div>
        </TerminalWindow>
      </div>

      {/* Submission Form */}
      {showSubmissionForm && (
        <div className="fixed inset-0 bg-black bg-opacity-75 z-50 flex items-center justify-center p-4">
          <div className="w-full max-w-2xl">
            <TerminalWindow title="submit_agentard_moment.exe">
              <form onSubmit={handleSubmit} className="space-y-4">
                <div>
                  <label className="block text-sm font-bold text-terminal-cyan mb-2">
                    Title of Your Magnificent Solution
                  </label>
                  <input
                    type="text"
                    value={newMoment.title}
                    onChange={(e) => setNewMoment(prev => ({ ...prev, title: e.target.value }))}
                    className="w-full bg-bg-tertiary border border-terminal-dim-green p-2 rounded text-text-primary"
                    placeholder="e.g., Fixed all bugs by removing error handling"
                    required
                  />
                </div>

                <div>
                  <label className="block text-sm font-bold text-terminal-cyan mb-2">
                    Describe Your Brilliant Approach
                  </label>
                  <textarea
                    value={newMoment.description}
                    onChange={(e) => setNewMoment(prev => ({ ...prev, description: e.target.value }))}
                    className="w-full bg-bg-tertiary border border-terminal-dim-green p-2 rounded text-text-primary h-24 resize-none"
                    placeholder="Explain how your solution was obviously the best possible approach..."
                    required
                  />
                </div>

                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div>
                    <label className="block text-sm font-bold text-terminal-cyan mb-2">
                      Your Confidence Level: {newMoment.confidence}%
                    </label>
                    <input
                      type="range"
                      min="1"
                      max="100"
                      value={newMoment.confidence}
                      onChange={(e) => setNewMoment(prev => ({ ...prev, confidence: Number(e.target.value) }))}
                      className="w-full"
                    />
                  </div>

                  <div>
                    <label className="block text-sm font-bold text-terminal-cyan mb-2">
                      Category
                    </label>
                    <select
                      value={newMoment.category}
                      onChange={(e) => setNewMoment(prev => ({ ...prev, category: e.target.value as AgentardMoment['category'] }))}
                      className="w-full bg-bg-tertiary border border-terminal-dim-green p-2 rounded text-text-primary"
                    >
                      {categories.slice(1).map(cat => (
                        <option key={cat.id} value={cat.id}>{cat.label}</option>
                      ))}
                    </select>
                  </div>
                </div>

                <div>
                  <label className="block text-sm font-bold text-terminal-cyan mb-2">
                    Your Handle
                  </label>
                  <input
                    type="text"
                    value={newMoment.author}
                    onChange={(e) => setNewMoment(prev => ({ ...prev, author: e.target.value }))}
                    className="w-full bg-bg-tertiary border border-terminal-dim-green p-2 rounded text-text-primary"
                    placeholder="anonymous_genius"
                    required
                  />
                </div>

                <div className="flex space-x-3">
                  <button
                    type="submit"
                    className="flex-1 bg-terminal-fg text-terminal-bg p-2 rounded font-bold hover:bg-terminal-cyan transition-colors"
                  >
                    Submit Moment
                  </button>
                  <button
                    type="button"
                    onClick={() => setShowSubmissionForm(false)}
                    className="flex-1 bg-terminal-red text-terminal-bg p-2 rounded font-bold hover:bg-terminal-amber transition-colors"
                  >
                    Cancel
                  </button>
                </div>
              </form>
            </TerminalWindow>
          </div>
        </div>
      )}

      {/* Moments List */}
      <div className="space-y-4">
        <div className="text-xl font-bold text-terminal-cyan">
          Community Agentard Moments ({filteredMoments.length})
        </div>
        
        {filteredMoments.map(moment => {
          const agentardLevel = getAgentardLevel(moment.agentardIndex);
          return (
            <TerminalWindow key={moment.id} title={`moment_${moment.id}.log`}>
              <div className="space-y-3">
                {/* Header */}
                <div className="flex items-start justify-between">
                  <div className="flex-1">
                    <h3 className="text-lg font-bold text-text-primary">{moment.title}</h3>
                    <div className="text-sm text-text-muted">
                      by {moment.author} • {moment.timestamp.toLocaleDateString()}
                    </div>
                  </div>
                  <div className="flex items-center space-x-2">
                    <span className={`px-2 py-1 rounded text-xs font-bold bg-${getCategoryColor(moment.category)} text-terminal-bg`}>
                      {categories.find(c => c.id === moment.category)?.icon} {moment.category.toUpperCase()}
                    </span>
                  </div>
                </div>

                {/* Content */}
                <p className="text-text-primary leading-relaxed">{moment.description}</p>

                {/* Metrics */}
                <div className="grid grid-cols-2 md:grid-cols-4 gap-4 p-3 bg-bg-tertiary rounded">
                  <div className="text-center">
                    <div className="text-terminal-cyan font-bold">{moment.confidence}%</div>
                    <div className="text-xs text-text-muted">Confidence</div>
                  </div>
                  <div className="text-center">
                    <div className="text-terminal-red font-bold">{moment.reality}%</div>
                    <div className="text-xs text-text-muted">Reality</div>
                  </div>
                  <div className="text-center">
                    <div className={`font-bold text-${agentardLevel.color}`}>{moment.agentardIndex}</div>
                    <div className="text-xs text-text-muted">Agentard Index</div>
                  </div>
                  <div className="text-center">
                    <button
                      onClick={() => handleVote(moment.id)}
                      className="flex items-center justify-center space-x-1 hover:text-terminal-amber transition-colors"
                    >
                      <span>👍</span>
                      <span className="font-bold">{moment.votes}</span>
                    </button>
                    <div className="text-xs text-text-muted">Votes</div>
                  </div>
                </div>

                {/* Agentard Level Badge */}
                <div className="text-center">
                  <span className={`px-3 py-1 rounded-full text-xs font-bold bg-${agentardLevel.color} text-terminal-bg`}>
                    {agentardLevel.label} AGENTARD LEVEL
                  </span>
                </div>
              </div>
            </TerminalWindow>
          );
        })}
      </div>
    </div>
  );
}