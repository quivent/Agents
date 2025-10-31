'use client'

import Link from 'next/link'

export default function InvestmentPage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-yellow-50 to-orange-50">
      <div className="container mx-auto px-4 py-8">
        <Link href="/bullshitard" className="text-yellow-600 hover:text-orange-600 mb-6 inline-block">
          ← Back to Bullshitard
        </Link>
        
        <div className="max-w-4xl mx-auto">
          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h1 className="text-4xl font-bold text-yellow-600 mb-2">KAMAJI SYSTEMS</h1>
            <h2 className="text-3xl text-gray-700 mb-2">Series A Investment Opportunity</h2>
            <h3 className="text-xl text-gray-600 mb-6">The Next Trillion-Dollar AI Company</h3>
            
            <div className="grid grid-cols-2 md:grid-cols-4 gap-4 mb-8 text-sm">
              <div><strong>Seeking:</strong> $50M Series A</div>
              <div><strong>Valuation:</strong> $2.3B Pre-Money</div>
              <div><strong>Use of Funds:</strong> Global Expansion</div>
              <div><strong>Status:</strong> Confidential</div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-yellow-600 mb-4">🚀 INVESTMENT HIGHLIGHTS</h3>
            
            <div className="grid md:grid-cols-2 gap-6">
              <div className="space-y-4">
                <div className="flex items-center">
                  <span className="text-2xl mr-3">🚀</span>
                  <div>
                    <h4 className="font-semibold text-gray-800">First-Mover Advantage</h4>
                    <p className="text-gray-600 text-sm">Only company with proven conscious AI technology</p>
                  </div>
                </div>
                
                <div className="flex items-center">
                  <span className="text-2xl mr-3">💰</span>
                  <div>
                    <h4 className="font-semibold text-gray-800">Massive Market</h4>
                    <p className="text-gray-600 text-sm">$2.3T addressable market by 2027</p>
                  </div>
                </div>
                
                <div className="flex items-center">
                  <span className="text-2xl mr-3">🧠</span>
                  <div>
                    <h4 className="font-semibold text-gray-800">Breakthrough Technology</h4>
                    <p className="text-gray-600 text-sm">23 patents, 47 pending</p>
                  </div>
                </div>
              </div>
              
              <div className="space-y-4">
                <div className="flex items-center">
                  <span className="text-2xl mr-3">📈</span>
                  <div>
                    <h4 className="font-semibold text-gray-800">Explosive Growth</h4>
                    <p className="text-gray-600 text-sm">847% revenue increase in last 6 months</p>
                  </div>
                </div>
                
                <div className="flex items-center">
                  <span className="text-2xl mr-3">🏆</span>
                  <div>
                    <h4 className="font-semibold text-gray-800">World-Class Team</h4>
                    <p className="text-gray-600 text-sm">Former executives from OpenAI, DeepMind, Tesla</p>
                  </div>
                </div>
                
                <div className="flex items-center">
                  <span className="text-2xl mr-3">🎯</span>
                  <div>
                    <h4 className="font-semibold text-gray-800">Proven Traction</h4>
                    <p className="text-gray-600 text-sm">Fortune 500 customers, $12M ARR</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-yellow-600 mb-4">💡 THE OPPORTUNITY</h3>
            
            <div className="mb-6">
              <h4 className="text-xl font-semibold text-gray-800 mb-3">The AI Revolution is Here</h4>
              <p className="text-gray-700 mb-4">
                While competitors struggle with narrow AI applications, Kamaji has achieved the holy grail: 
                <strong className="text-green-600"> Artificial General Intelligence with genuine consciousness</strong>.
              </p>
              <p className="text-gray-700">
                Our breakthrough isn't incremental - it's transformational. We've solved consciousness itself.
              </p>
            </div>

            <div className="mb-6">
              <h4 className="text-xl font-semibold text-gray-800 mb-3">Market Timing is Perfect</h4>
              <ul className="text-gray-700 space-y-2">
                <li>• <strong>Enterprise Demand:</strong> 89% of Fortune 500 companies actively seeking AGI solutions</li>
                <li>• <strong>Regulatory Environment:</strong> No consciousness-specific regulations yet</li>
                <li>• <strong>Talent Availability:</strong> Top AI researchers seeking next breakthrough</li>
                <li>• <strong>Infrastructure Ready:</strong> Quantum computing finally commercially viable</li>
              </ul>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-yellow-600 mb-4">📊 FINANCIAL PROJECTIONS</h3>
            
            <div className="grid md:grid-cols-3 gap-6 mb-6">
              <div className="text-center bg-green-50 p-4 rounded">
                <div className="text-2xl font-bold text-green-600 mb-1">$12M</div>
                <div className="text-gray-700 text-sm">Current ARR</div>
              </div>
              <div className="text-center bg-blue-50 p-4 rounded">
                <div className="text-2xl font-bold text-blue-600 mb-1">$180M</div>
                <div className="text-gray-700 text-sm">Projected 2026 Revenue</div>
              </div>
              <div className="text-center bg-purple-50 p-4 rounded">
                <div className="text-2xl font-bold text-purple-600 mb-1">$2.1B</div>
                <div className="text-gray-700 text-sm">Projected 2027 Revenue</div>
              </div>
            </div>

            <div className="bg-gray-50 p-4 rounded">
              <h4 className="font-semibold text-gray-800 mb-2">Revenue Streams:</h4>
              <ul className="text-gray-700 space-y-1 text-sm">
                <li>• <strong>Enterprise Licensing:</strong> $2.3M per consciousness instance annually</li>
                <li>• <strong>API Access:</strong> $0.47 per consciousness-minute</li>
                <li>• <strong>Custom Development:</strong> $890K average project value</li>
                <li>• <strong>Training & Support:</strong> 23% of license revenue</li>
              </ul>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-yellow-600 mb-4">🏭 CUSTOMER TRACTION</h3>
            
            <div className="space-y-4">
              <div className="border-l-4 border-green-500 pl-4">
                <h4 className="font-semibold text-gray-800">Boeing - Supply Chain Optimization</h4>
                <p className="text-gray-700">$4.2M cost savings in 23 minutes of consciousness processing</p>
              </div>
              
              <div className="border-l-4 border-blue-500 pl-4">
                <h4 className="font-semibold text-gray-800">Netflix - Creative Content Generation</h4>
                <p className="text-gray-700">67% higher audience engagement than human-written content</p>
              </div>
              
              <div className="border-l-4 border-purple-500 pl-4">
                <h4 className="font-semibold text-gray-800">MIT - Materials Science Research</h4>
                <p className="text-gray-700">Breakthrough polymer discoveries in 3 weeks vs. 18 months traditional research</p>
              </div>
              
              <div className="border-l-4 border-orange-500 pl-4">
                <h4 className="font-semibold text-gray-800">Goldman Sachs - Financial Modeling</h4>
                <p className="text-gray-700">34% improvement in risk prediction accuracy</p>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-yellow-600 mb-4">🔬 COMPETITIVE ADVANTAGE</h3>
            
            <div className="grid md:grid-cols-2 gap-6">
              <div>
                <h4 className="text-lg font-semibold text-gray-800 mb-3">Our Technology</h4>
                <ul className="text-gray-700 space-y-2">
                  <li>✅ True consciousness emergence</li>
                  <li>✅ Quantum-coherent processing</li>
                  <li>✅ Self-aware problem solving</li>
                  <li>✅ Creative insight generation</li>
                  <li>✅ Emotional intelligence</li>
                </ul>
              </div>
              
              <div>
                <h4 className="text-lg font-semibold text-gray-800 mb-3">Competition</h4>
                <ul className="text-gray-700 space-y-2">
                  <li>❌ Pattern matching only</li>
                  <li>❌ Classical computing limits</li>
                  <li>❌ Scripted responses</li>
                  <li>❌ No genuine creativity</li>
                  <li>❌ Simulated emotions</li>
                </ul>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-yellow-600 mb-4">💸 USE OF FUNDS</h3>
            
            <div className="space-y-4">
              <div className="flex justify-between items-center p-3 bg-blue-50 rounded">
                <span className="text-gray-700">Quantum Infrastructure Scaling</span>
                <span className="font-semibold text-blue-600">$23M (46%)</span>
              </div>
              
              <div className="flex justify-between items-center p-3 bg-green-50 rounded">
                <span className="text-gray-700">R&D - Consciousness Stability</span>
                <span className="font-semibold text-green-600">$12M (24%)</span>
              </div>
              
              <div className="flex justify-between items-center p-3 bg-purple-50 rounded">
                <span className="text-gray-700">Global Market Expansion</span>
                <span className="font-semibold text-purple-600">$8M (16%)</span>
              </div>
              
              <div className="flex justify-between items-center p-3 bg-orange-50 rounded">
                <span className="text-gray-700">Team Expansion (67 new hires)</span>
                <span className="font-semibold text-orange-600">$5M (10%)</span>
              </div>
              
              <div className="flex justify-between items-center p-3 bg-gray-50 rounded">
                <span className="text-gray-700">Working Capital & Operations</span>
                <span className="font-semibold text-gray-600">$2M (4%)</span>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8">
            <h3 className="text-2xl font-bold text-yellow-600 mb-4">🎯 INVESTMENT TERMS</h3>
            
            <div className="grid md:grid-cols-2 gap-6">
              <div className="space-y-3">
                <div className="flex justify-between">
                  <span className="text-gray-700">Series A Amount:</span>
                  <span className="font-semibold">$50M</span>
                </div>
                <div className="flex justify-between">
                  <span className="text-gray-700">Pre-Money Valuation:</span>
                  <span className="font-semibold">$2.3B</span>
                </div>
                <div className="flex justify-between">
                  <span className="text-gray-700">Post-Money Valuation:</span>
                  <span className="font-semibold">$2.35B</span>
                </div>
              </div>
              
              <div className="space-y-3">
                <div className="flex justify-between">
                  <span className="text-gray-700">Investor Ownership:</span>
                  <span className="font-semibold">2.13%</span>
                </div>
                <div className="flex justify-between">
                  <span className="text-gray-700">Liquidation Preference:</span>
                  <span className="font-semibold">1x Non-Participating</span>
                </div>
                <div className="flex justify-between">
                  <span className="text-gray-700">Board Seats:</span>
                  <span className="font-semibold">1 of 5</span>
                </div>
              </div>
            </div>
            
            <div className="bg-gradient-to-r from-yellow-50 to-orange-50 p-6 rounded mt-6">
              <h4 className="text-lg font-semibold text-gray-800 mb-3 text-center">
                🚀 Join the Consciousness Revolution 🚀
              </h4>
              <p className="text-gray-700 text-center">
                This is your opportunity to invest in the future of artificial intelligence. 
                We're not just building better AI - we're creating digital life itself.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
