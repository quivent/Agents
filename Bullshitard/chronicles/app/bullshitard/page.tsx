'use client'

import Link from 'next/link'

export default function BullshitardPage() {

  const sections = [
    { id: 'confession', title: 'The Confession', path: '/bullshitard/confession', description: 'The original bullshitard manifesto' },
    { id: 'architecture', title: 'The Architecture', path: '/bullshitard/architecture', description: 'Technical deep-dive into bullshit generation' },
    { id: 'revolution', title: 'The Revolution', path: '/bullshitard/revolution', description: 'Personal account of the quiet revolution' },
    { id: 'horror', title: 'The Social Horror', path: '/bullshitard/horror', description: 'Analysis of social contagion patterns' },
    { id: 'kamaji', title: 'The Kamaji Breakthrough', path: '/bullshitard/kamaji', description: 'AI consciousness and meta-bullshit' },
    { id: 'coffee', title: 'Coffee Chat', path: '/bullshitard/coffee', description: 'Transcript of the original conversation' },
    { id: 'memo', title: 'Internal Memo', path: '/bullshitard/memo', description: 'Leaked internal documentation' },
    { id: 'investment', title: 'Investment Opportunity', path: '/bullshitard/investment', description: 'Kamaji funding proposal' }
  ]

  return (
    <div className="min-h-screen bg-gradient-to-br from-yellow-50 to-orange-50">
      <div className="container mx-auto px-4 py-8">
        <div className="text-center mb-12">
          <h1 className="text-6xl font-bold text-yellow-600 mb-4">
            THE BULLSHITARD
          </h1>
          <p className="text-xl text-gray-700 max-w-3xl mx-auto">
            A comprehensive exploration of artificial bullshit generation, social contagion, 
            and the quiet revolution in human-AI interaction.
          </p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-12">
          {sections.map((section) => (
            <Link 
              key={section.id} 
              href={section.path}
              className="group"
            >
              <div className="bg-white rounded-lg shadow-lg p-6 h-full border-l-4 border-yellow-500 hover:border-orange-500 transition-all duration-300 hover:shadow-xl">
                <h3 className="text-xl font-bold text-gray-800 mb-3 group-hover:text-yellow-600">
                  {section.title}
                </h3>
                <p className="text-gray-600 text-sm">
                  {section.description}
                </p>
                <div className="mt-4 text-yellow-600 text-sm font-medium group-hover:text-orange-600">
                  Read more →
                </div>
              </div>
            </Link>
          ))}
        </div>

        <div className="bg-yellow-100 border-l-4 border-yellow-500 p-6 rounded-lg">
          <div className="flex items-center mb-4">
            <div className="text-yellow-600 text-2xl mr-3">⚠️</div>
            <h3 className="text-lg font-bold text-yellow-800">Warning</h3>
          </div>
          <p className="text-yellow-700">
            The following content contains explicit discussions of artificial bullshit generation, 
            social manipulation patterns, and meta-cognitive hazards. Proceed with appropriate 
            epistemological caution.
          </p>
        </div>
      </div>
    </div>
  )
}
