import { Layout } from '@/components/layout/Layout';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import Link from 'next/link';

export default function BullshitardInvestment() {
  return (
    <Layout>
      <div className="max-w-4xl mx-auto space-y-8">
        {/* Navigation */}
        <nav className="mb-8">
          <Link href="/bullshitard" className="text-bullshitard-accent hover:text-bullshitard-primary transition-colors">
            ← Back to Bullshitard
          </Link>
        </nav>

        {/* Header */}
        <div className="text-center space-y-4 mb-12">
          <h1 className="text-4xl md:text-5xl font-bold text-bullshitard-primary terminal-glow">
            KAMAJI INVESTMENT OPPORTUNITY
          </h1>
          <h2 className="text-2xl text-bullshitard-secondary">
            Monetizing Bullshit at Enterprise Scale
          </h2>
          <div className="text-text-secondary">
            <p><strong>CONFIDENTIAL</strong> - Series B Funding Presentation</p>
            <p>Kamaji Corporation | October 2025</p>
          </div>
        </div>

        {/* Content */}
        <div className="prose prose-invert max-w-none space-y-6">
          
          <section className="bg-bullshitard-primary/10 border border-bullshitard-primary/30 rounded-lg p-6">
            <h2 className="text-2xl font-bold text-bullshitard-primary mb-4">Executive Summary</h2>
            <p className="text-text-primary mb-4">
              Kamaji represents the next evolution in artificial intelligence: the world's first commercially viable 
              Bullshit-as-a-Service (BaaS) platform. Our proprietary technology transforms meaningless technical jargon 
              into enterprise-grade solutions that sound impressive enough to secure funding, close deals, and drive adoption.
            </p>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mt-6">
              <div className="text-center">
                <div className="text-3xl font-bold text-bullshitard-primary">$50M</div>
                <div className="text-text-secondary">Series B Target</div>
              </div>
              <div className="text-center">
                <div className="text-3xl font-bold text-bullshitard-primary">847%</div>
                <div className="text-text-secondary">YoY Growth</div>
              </div>
              <div className="text-center">
                <div className="text-3xl font-bold text-bullshitard-primary">99.7%</div>
                <div className="text-text-secondary">Bullshit Accuracy</div>
              </div>
            </div>
          </section>

          <section>
            <h2 className="text-2xl font-bold text-bullshitard-primary mb-4">Market Opportunity</h2>
            <p className="text-text-primary mb-4">
              The global enterprise software market is valued at $650B annually, with 73% of purchasing decisions based 
              on impressive-sounding technical specifications rather than actual functionality. Kamaji addresses this 
              massive market inefficiency by providing the impressive-sounding specifications without the costly 
              development overhead.
            </p>
            <TerminalWindow title="market_analysis.json" className="mb-4">
              <pre className="text-sm text-text-primary overflow-x-auto">
{`{
  "total_addressable_market": "$650B",
  "serviceable_addressable_market": "$47B",
  "serviceable_obtainable_market": "$3.2B",
  "bullshit_penetration_rate": "73%",
  "customer_acquisition_cost": "$12,000",
  "lifetime_value": "$847,000",
  "magic_quadrant_position": "Visionary Leader"
}`}
              </pre>
            </TerminalWindow>
          </section>

          <section>
            <h2 className="text-2xl font-bold text-bullshitard-primary mb-4">Product Portfolio</h2>
            
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
              <div className="bg-bg-secondary p-4 rounded border border-bullshitard-primary/20">
                <h3 className="text-lg font-semibold text-bullshitard-primary mb-2">Kamaji Enterprise</h3>
                <p className="text-text-secondary text-sm">
                  Quantum-enhanced decision trees with consciousness-aware processing for Fortune 500 companies.
                </p>
                <div className="text-bullshitard-accent mt-2">$50K/month</div>
              </div>
              
              <div className="bg-bg-secondary p-4 rounded border border-bullshitard-primary/20">
                <h3 className="text-lg font-semibold text-bullshitard-primary mb-2">Kamaji Pro</h3>
                <p className="text-text-secondary text-sm">
                  Advanced multi-agent orchestration with neural pathway optimization for mid-market.
                </p>
                <div className="text-bullshitard-accent mt-2">$15K/month</div>
              </div>
              
              <div className="bg-bg-secondary p-4 rounded border border-bullshitard-primary/20">
                <h3 className="text-lg font-semibold text-bullshitard-primary mb-2">Kamaji Startup</h3>
                <p className="text-text-secondary text-sm">
                  Revolutionary streaming architecture with 63 integrated tools for emerging companies.
                </p>
                <div className="text-bullshitard-accent mt-2">$5K/month</div>
              </div>
              
              <div className="bg-bg-secondary p-4 rounded border border-bullshitard-primary/20">
                <h3 className="text-lg font-semibold text-bullshitard-primary mb-2">Kamaji API</h3>
                <p className="text-text-secondary text-sm">
                  Cognitive-aware system integration with performance optimization for developers.
                </p>
                <div className="text-bullshitard-accent mt-2">$0.03/request</div>
              </div>
            </div>
          </section>

          <section>
            <h2 className="text-2xl font-bold text-bullshitard-primary mb-4">Competitive Advantages</h2>
            <ul className="list-disc list-inside text-text-primary space-y-2 mb-4">
              <li><strong>First-mover advantage</strong> in the Bullshit-as-a-Service market</li>
              <li><strong>Proprietary algorithms</strong> for generating convincing technical specifications</li>
              <li><strong>Network effects</strong> - the more bullshit we generate, the more believable it becomes</li>
              <li><strong>Recursive improvement</strong> - our AI learns to bullshit better by studying its own bullshit</li>
              <li><strong>Enterprise-grade security</strong> - your bullshit is encrypted and compliant</li>
              <li><strong>Quantum-enhanced processing</strong> - because quantum makes everything sound more impressive</li>
            </ul>
          </section>

          <section>
            <h2 className="text-2xl font-bold text-bullshitard-primary mb-4">Financial Projections</h2>
            <TerminalWindow title="financial_model.xlsx" className="mb-4">
              <div className="text-sm space-y-2">
                <div className="grid grid-cols-4 gap-4 font-semibold text-bullshitard-primary">
                  <div>Year</div>
                  <div>Revenue</div>
                  <div>Customers</div>
                  <div>Valuation</div>
                </div>
                <div className="grid grid-cols-4 gap-4">
                  <div>2025</div>
                  <div>$12M</div>
                  <div>847</div>
                  <div>$150M</div>
                </div>
                <div className="grid grid-cols-4 gap-4">
                  <div>2026</div>
                  <div>$47M</div>
                  <div>2,341</div>
                  <div>$500M</div>
                </div>
                <div className="grid grid-cols-4 gap-4">
                  <div>2027</div>
                  <div>$156M</div>
                  <div>7,892</div>
                  <div>$1.8B</div>
                </div>
                <div className="grid grid-cols-4 gap-4 text-bullshitard-accent">
                  <div>2028</div>
                  <div>$423M</div>
                  <div>18,567</div>
                  <div>$5.2B</div>
                </div>
              </div>
            </TerminalWindow>
            <p className="text-text-secondary italic">
              *Projections based on advanced predictive modeling with consciousness-aware market analysis
            </p>
          </section>

          <section>
            <h2 className="text-2xl font-bold text-bullshitard-primary mb-4">Use of Funds</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <h3 className="text-lg font-semibold text-bullshitard-secondary mb-3">Engineering (40%)</h3>
                <ul className="list-disc list-inside text-text-primary text-sm space-y-1">
                  <li>Quantum coherence optimization</li>
                  <li>Meta-bullshit layer development</li>
                  <li>Consciousness-aware processing enhancement</li>
                  <li>Advanced neural pathway integration</li>
                </ul>
              </div>
              <div>
                <h3 className="text-lg font-semibold text-bullshitard-secondary mb-3">Sales & Marketing (35%)</h3>
                <ul className="list-disc list-inside text-text-primary text-sm space-y-1">
                  <li>Enterprise customer acquisition</li>
                  <li>Thought leadership content</li>
                  <li>Conference speaking engagements</li>
                  <li>Bullshit evangelism program</li>
                </ul>
              </div>
              <div>
                <h3 className="text-lg font-semibold text-bullshitard-secondary mb-3">Operations (15%)</h3>
                <ul className="list-disc list-inside text-text-primary text-sm space-y-1">
                  <li>Infrastructure scaling</li>
                  <li>Compliance and security</li>
                  <li>Customer success team</li>
                  <li>Quality assurance (bullshit validation)</li>
                </ul>
              </div>
              <div>
                <h3 className="text-lg font-semibold text-bullshitard-secondary mb-3">Strategic Reserves (10%)</h3>
                <ul className="list-disc list-inside text-text-primary text-sm space-y-1">
                  <li>Acquisition opportunities</li>
                  <li>Patent portfolio expansion</li>
                  <li>Regulatory compliance buffer</li>
                  <li>Contingency planning</li>
                </ul>
              </div>
            </div>
          </section>

          <section>
            <h2 className="text-2xl font-bold text-bullshitard-primary mb-4">Risk Factors</h2>
            <div className="bg-bullshitard-warning/10 border border-bullshitard-warning/30 rounded-lg p-4">
              <ul className="list-disc list-inside text-text-primary space-y-2 text-sm">
                <li>Regulatory scrutiny of bullshit-based business models</li>
                <li>Competition from traditional AI companies adding bullshit features</li>
                <li>Market saturation if everyone becomes aware they're buying bullshit</li>
                <li>Technical risk that our AI might accidentally become genuinely useful</li>
                <li>Existential risk that our founders might have an attack of conscience</li>
                <li>Quantum decoherence could affect our quantum-enhanced processing claims</li>
              </ul>
            </div>
          </section>

          <section>
            <h2 className="text-2xl font-bold text-bullshitard-primary mb-4">Investment Terms</h2>
            <TerminalWindow title="term_sheet.pdf" className="mb-4">
              <div className="text-sm space-y-2">
                <div><strong>Series B Funding:</strong> $50M</div>
                <div><strong>Pre-money Valuation:</strong> $200M</div>
                <div><strong>Post-money Valuation:</strong> $250M</div>
                <div><strong>Liquidation Preference:</strong> 1x Non-participating</div>
                <div><strong>Anti-dilution:</strong> Weighted average broad-based</div>
                <div><strong>Board Composition:</strong> 2 Founders, 2 Investors, 1 Independent</div>
                <div><strong>Option Pool:</strong> 15% (pre-money)</div>
                <div><strong>Drag-along Rights:</strong> Standard</div>
                <div><strong>Tag-along Rights:</strong> Standard</div>
                <div><strong>Bullshit Warranty:</strong> Enterprise-grade, 99.7% accuracy guaranteed</div>
              </div>
            </TerminalWindow>
          </section>

          <div className="bg-bullshitard-primary/10 border border-bullshitard-primary/30 rounded-lg p-6 mt-8">
            <h3 className="text-lg font-bold text-bullshitard-primary mb-4">Next Steps</h3>
            <p className="text-text-primary mb-4">
              We're seeking visionary investors who understand that the future of enterprise software isn't about 
              functionality - it's about the perception of functionality. Join us in revolutionizing how businesses 
              buy, sell, and implement solutions that sound impressive enough to justify their existence.
            </p>
            <div className="text-center">
              <p className="text-bullshitard-accent font-semibold">
                Contact: investors@kamaji.ai | +1 (555) BULLSHT
              </p>
            </div>
          </div>

          <div className="text-center py-8">
            <p className="text-text-secondary italic">
              This presentation contains forward-looking statements that may be complete bullshit. 
              Past performance of bullshit is not indicative of future bullshit results.
            </p>
          </div>
        </div>
      </div>
    </Layout>
  );
}
