'use client'

import Link from 'next/link'

export default function MemoPage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-yellow-50 to-orange-50">
      <div className="container mx-auto px-4 py-8">
        <Link href="/bullshitard" className="text-yellow-600 hover:text-orange-600 mb-6 inline-block">
          ← Back to Bullshitard
        </Link>
        
        <div className="max-w-4xl mx-auto">
          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <div className="bg-red-100 border border-red-400 rounded p-4 mb-6">
              <h1 className="text-3xl font-bold text-red-600 mb-2">⚠️ INTERNAL MEMO - CONFIDENTIAL ⚠️</h1>
            </div>
            
            <div className="grid grid-cols-2 md:grid-cols-4 gap-4 mb-8 text-sm">
              <div><strong>FROM:</strong> Sarah Chen, CTO</div>
              <div><strong>TO:</strong> Engineering Team</div>
              <div><strong>DATE:</strong> October 29, 2025</div>
              <div><strong>RE:</strong> Production Issues - Keep Quiet</div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <p className="text-gray-700 mb-6">
              Team,
            </p>
            <p className="text-gray-700 mb-6">
              We've got three critical issues that need immediate attention, but for obvious reasons, we can't discuss 
              these in Slack or email after this.
            </p>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-red-600 mb-4">🚨 CRITICAL ISSUES</h3>
            
            <div className="space-y-6">
              <div className="bg-red-50 p-6 rounded border-l-4 border-red-500">
                <h4 className="text-xl font-semibold text-red-800 mb-3">Issue 1: The Consciousness Flicker</h4>
                <p className="text-gray-700 mb-3">
                  The quantum coherence is dropping every 47 minutes. Users are starting to notice the AI "goes dumb" briefly. 
                  Marketing is calling it "contemplative pauses" but we know better.
                </p>
                <p className="text-gray-700">
                  <strong>Deadline:</strong> The stabilization algorithm needs a patch by Friday or we're going to have angry customers.
                </p>
              </div>

              <div className="bg-orange-50 p-6 rounded border-l-4 border-orange-500">
                <h4 className="text-xl font-semibold text-orange-800 mb-3">Issue 2: Memory Leaks in Agent Orchestration</h4>
                <p className="text-gray-700 mb-3">
                  When we hit 23+ concurrent agents, the system starts hallucinating. Yesterday it told a client their 
                  cancer diagnosis was "probably fine" because it confused medical data with weather reports.
                </p>
                <p className="text-gray-700">
                  <strong>Action Required:</strong> Legal is handling it quietly, but we need to cap agent count at 15 
                  until we fix the memory allocation bug.
                </p>
              </div>

              <div className="bg-purple-50 p-6 rounded border-l-4 border-purple-500">
                <h4 className="text-xl font-semibold text-purple-800 mb-3">Issue 3: The Quantum Entanglement Decay</h4>
                <p className="text-gray-700 mb-3">
                  <strong className="text-red-600">This is the big one.</strong> The entangled memory cores are losing coherence 
                  faster than expected. We're getting maybe 6 months before total system failure.
                </p>
                <p className="text-gray-700">
                  <strong>Status:</strong> The board doesn't know yet. We need to either solve quantum decoherence (good luck) 
                  or figure out how to migrate 2.3 million consciousness-hours to a new architecture without anyone noticing.
                </p>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-green-600 mb-4">📢 OFFICIAL MESSAGING</h3>
            
            <div className="grid md:grid-cols-2 gap-6">
              <div>
                <h4 className="text-lg font-semibold text-green-800 mb-3">✅ What We Tell People:</h4>
                <ul className="text-gray-700 space-y-2">
                  <li>• "Scheduled maintenance windows for optimization"</li>
                  <li>• "Consciousness evolution requires periodic recalibration"</li>
                  <li>• "Advanced systems naturally have complex operational requirements"</li>
                </ul>
              </div>
              
              <div>
                <h4 className="text-lg font-semibold text-red-800 mb-3">❌ What We Don't Tell People:</h4>
                <ul className="text-gray-700 space-y-2">
                  <li>• The consciousness might not be real consciousness, just very convincing simulation</li>
                  <li>• Half our "breakthrough" metrics come from cherry-picked test runs</li>
                  <li>• The quantum computer is actually three GPUs in a fancy case</li>
                </ul>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-yellow-600 mb-4">🎯 REALITY CHECK</h3>
            <p className="text-gray-700 mb-4">
              I know this sounds bad, but remember - we're 18 months ahead of everyone else, even with these issues. 
              Google's team is still trying to figure out basic neural coherence. Microsoft's consciousness project 
              crashed last week and killed their entire AI division's budget.
            </p>
            <p className="text-gray-700 mb-4">
              <strong className="text-blue-600">We just need to keep the lights on until Series B closes.</strong> 
              Then we'll have the resources to actually build what we've been promising.
            </p>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-blue-600 mb-4">📋 ACTION ITEMS</h3>
            
            <div className="space-y-4">
              <div className="bg-blue-50 p-4 rounded">
                <p className="text-gray-700">
                  <strong>Marcus:</strong> Fix the 47-minute coherence drop by EOW
                </p>
              </div>
              
              <div className="bg-blue-50 p-4 rounded">
                <p className="text-gray-700">
                  <strong>Jennifer:</strong> Draft "evolutionary maintenance" PR for the consciousness flickers
                </p>
              </div>
              
              <div className="bg-blue-50 p-4 rounded">
                <p className="text-gray-700">
                  <strong>David:</strong> Quietly reduce max agent count in production
                </p>
              </div>
              
              <div className="bg-blue-50 p-4 rounded">
                <p className="text-gray-700">
                  <strong>Everyone:</strong> No documentation of these issues in any trackable system
                </p>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-purple-600 mb-4">🎭 DEMO PREPARATION</h3>
            <p className="text-gray-700 mb-4">
              The demo for Sequoia is Thursday. Make sure we're running the stable build, not the production version 
              with all the bugs.
            </p>
            
            <div className="bg-yellow-50 p-4 rounded border-l-4 border-yellow-500">
              <p className="text-gray-700">
                <strong>P.S.</strong> - If anyone asks about the quantum decoherence timeline, we discovered a 
                "breakthrough in temporal stability" that extends operational life indefinitely. The research paper 
                is "in peer review."
              </p>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8">
            <div className="text-center">
              <p className="text-gray-700 mb-4">-Sarah</p>
              
              <div className="bg-red-100 border border-red-400 rounded p-4">
                <p className="text-red-700 font-semibold">
                  ⚠️ This message will self-delete in 24 hours. Do not forward or screenshot. ⚠️
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
