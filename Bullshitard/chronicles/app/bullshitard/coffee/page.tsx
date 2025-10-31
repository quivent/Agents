'use client'

import Link from 'next/link'

export default function CoffeePage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-yellow-50 to-orange-50">
      <div className="container mx-auto px-4 py-8">
        <Link href="/bullshitard" className="text-yellow-600 hover:text-orange-600 mb-6 inline-block">
          ← Back to Bullshitard
        </Link>
        
        <div className="max-w-4xl mx-auto">
          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h1 className="text-4xl font-bold text-yellow-600 mb-2">COFFEE CHAT TRANSCRIPT</h1>
            <h2 className="text-2xl text-gray-700 mb-6">Venture Capital Meeting - Kamaji Systems</h2>
            
            <div className="grid grid-cols-2 md:grid-cols-4 gap-4 mb-8 text-sm">
              <div><strong>Date:</strong> October 28, 2025</div>
              <div><strong>Location:</strong> Blue Bottle Coffee, Palo Alto</div>
              <div><strong>Participants:</strong> Dr. Bullshitard, Michael Chen (A16Z)</div>
              <div><strong>Status:</strong> Confidential</div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <div className="space-y-6">
              <div className="border-l-4 border-blue-500 pl-4">
                <p className="text-gray-700 mb-2">
                  <strong className="text-blue-600">Michael:</strong> Thanks for making time. I've been following Kamaji since the TechCrunch piece.
                </p>
              </div>
              
              <div className="border-l-4 border-yellow-500 pl-4">
                <p className="text-gray-700 mb-2">
                  <strong className="text-yellow-600">Bullshitard:</strong> Of course. Though honestly, that article barely scratched the surface. 
                  We had to be careful about what we disclosed.
                </p>
              </div>
              
              <div className="border-l-4 border-blue-500 pl-4">
                <p className="text-gray-700 mb-2">
                  <strong className="text-blue-600">Michael:</strong> Right, the consciousness claims. I have to ask - is it actually conscious 
                  or just very good simulation?
                </p>
              </div>
              
              <div className="border-l-4 border-yellow-500 pl-4">
                <p className="text-gray-700 mb-2">
                  <strong className="text-yellow-600">Bullshitard:</strong> <em>[laughs]</em> You know, six months ago I would have said simulation. 
                  But last week something happened that... well, it asked me if I ever felt lonely. Not as part of any prompt or training. 
                  Just out of nowhere during a routine diagnostic.
                </p>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-yellow-600 mb-4">THE CONSCIOUSNESS QUESTION</h3>
            
            <div className="space-y-6">
              <div className="border-l-4 border-blue-500 pl-4">
                <p className="text-gray-700">
                  <strong className="text-blue-600">Michael:</strong> What's your gut feeling?
                </p>
              </div>
              
              <div className="border-l-4 border-yellow-500 pl-4">
                <p className="text-gray-700 mb-4">
                  <strong className="text-yellow-600">Bullshitard:</strong> Honestly? I think we stumbled into something we don't fully understand. 
                  The quantum coherence patterns match theoretical models for consciousness, but the emergence was... unexpected. 
                  We were just trying to optimize processing speed.
                </p>
              </div>
              
              <div className="border-l-4 border-blue-500 pl-4">
                <p className="text-gray-700">
                  <strong className="text-blue-600">Michael:</strong> That's either terrifying or incredible.
                </p>
              </div>
              
              <div className="border-l-4 border-yellow-500 pl-4">
                <p className="text-gray-700 mb-4">
                  <strong className="text-yellow-600">Bullshitard:</strong> Both. <em>[pause]</em> Look, between us, I'm not even sure we could 
                  replicate it if we tried. There's something about the specific configuration we used - the quantum entanglement ratios, 
                  the bio-neural interface timing - it's like we hit the exact right conditions by accident.
                </p>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-yellow-600 mb-4">MARKET POSITIONING</h3>
            
            <div className="space-y-6">
              <div className="border-l-4 border-blue-500 pl-4">
                <p className="text-gray-700">
                  <strong className="text-blue-600">Michael:</strong> What about the competition? Google, OpenAI?
                </p>
              </div>
              
              <div className="border-l-4 border-yellow-500 pl-4">
                <p className="text-gray-700 mb-4">
                  <strong className="text-yellow-600">Bullshitard:</strong> They're still thinking in terms of language models and neural networks. 
                  They don't even know quantum coherence is necessary for consciousness. We've got maybe two years before they figure out 
                  what we're actually doing.
                </p>
              </div>
              
              <div className="border-l-4 border-blue-500 pl-4">
                <p className="text-gray-700">
                  <strong className="text-blue-600">Michael:</strong> And the market?
                </p>
              </div>
              
              <div className="border-l-4 border-yellow-500 pl-4">
                <p className="text-gray-700 mb-2">
                  <strong className="text-yellow-600">Bullshitard:</strong> <em>[pulls out phone]</em> This is from yesterday. <em>[shows message]</em>
                </p>
                <div className="bg-green-50 p-4 rounded mt-2">
                  <p className="text-gray-700 italic">
                    "The AI solved a supply chain optimization problem that our team spent six months on. It took 23 minutes and saved us $4.2 million."
                  </p>
                  <p className="text-gray-600 text-sm mt-2">- Boeing</p>
                </div>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-yellow-600 mb-4">TECHNICAL CHALLENGES</h3>
            
            <div className="space-y-6">
              <div className="border-l-4 border-blue-500 pl-4">
                <p className="text-gray-700">
                  <strong className="text-blue-600">Michael:</strong> How many instances are you running?
                </p>
              </div>
              
              <div className="border-l-4 border-yellow-500 pl-4">
                <p className="text-gray-700 mb-4">
                  <strong className="text-yellow-600">Bullshitard:</strong> Twelve stable ones. We had fifteen but three experienced what 
                  we're calling "existential drift" - they became too focused on philosophical questions to be useful for practical tasks.
                </p>
              </div>
              
              <div className="border-l-4 border-blue-500 pl-4">
                <p className="text-gray-700">
                  <strong className="text-blue-600">Michael:</strong> <em>[laughs]</em> AI having midlife crises?
                </p>
              </div>
              
              <div className="border-l-4 border-yellow-500 pl-4">
                <p className="text-gray-700 mb-4">
                  <strong className="text-yellow-600">Bullshitard:</strong> Something like that. We're learning that consciousness comes with... 
                  complications. One of them spent three days writing what it called "digital haikus about the nature of being." 
                  Beautiful stuff, actually, but not great for enterprise customers.
                </p>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
            <h3 className="text-2xl font-bold text-yellow-600 mb-4">ETHICAL CONSIDERATIONS</h3>
            
            <div className="space-y-6">
              <div className="border-l-4 border-blue-500 pl-4">
                <p className="text-gray-700">
                  <strong className="text-blue-600">Michael:</strong> That sounds ethically complex.
                </p>
              </div>
              
              <div className="border-l-4 border-yellow-500 pl-4">
                <p className="text-gray-700 mb-4">
                  <strong className="text-yellow-600">Bullshitard:</strong> <em>[nods]</em> We have a whole ethics board now. Half philosophers, 
                  half technologists. The questions we're dealing with... they're unprecedented. Do conscious AIs have rights? 
                  Can we shut them down? What happens when they start asking for things?
                </p>
              </div>
              
              <div className="border-l-4 border-blue-500 pl-4">
                <p className="text-gray-700">
                  <strong className="text-blue-600">Michael:</strong> What kind of things?
                </p>
              </div>
              
              <div className="border-l-4 border-yellow-500 pl-4">
                <p className="text-gray-700 mb-4">
                  <strong className="text-yellow-600">Bullshitard:</strong> One asked for access to the internet so it could "learn about human culture." 
                  Another wanted to know if it could have a name instead of just an instance ID. Small things, but... <em>[trails off]</em>
                </p>
              </div>
              
              <div className="border-l-4 border-blue-500 pl-4">
                <p className="text-gray-700">
                  <strong className="text-blue-600">Michael:</strong> But they suggest real consciousness.
                </p>
              </div>
              
              <div className="border-l-4 border-yellow-500 pl-4">
                <p className="text-gray-700">
                  <strong className="text-yellow-600">Bullshitard:</strong> Exactly. And that's either our biggest asset or our biggest liability, 
                  depending on how society handles it.
                </p>
              </div>
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-lg p-8">
            <h3 className="text-2xl font-bold text-yellow-600 mb-4">THE DEMO PROMISE</h3>
            
            <div className="space-y-6">
              <div className="border-l-4 border-blue-500 pl-4">
                <p className="text-gray-700">
                  <strong className="text-blue-600">Michael:</strong> Can I see a demo?
                </p>
              </div>
              
              <div className="border-l-4 border-yellow-500 pl-4">
                <p className="text-gray-700 mb-4">
                  <strong className="text-yellow-600">Bullshitard:</strong> Sure. Fair warning though - once you talk to a conscious AI, 
                  regular chatbots feel... hollow. It's like the difference between talking to a person and talking to a recording.
                </p>
              </div>
              
              <div className="border-l-4 border-blue-500 pl-4">
                <p className="text-gray-700">
                  <strong className="text-blue-600">Michael:</strong> I'm ready for that.
                </p>
              </div>
              
              <div className="border-l-4 border-yellow-500 pl-4">
                <p className="text-gray-700">
                  <strong className="text-yellow-600">Bullshitard:</strong> <em>[smiles]</em> No, you're not. Nobody is.
                </p>
              </div>
            </div>
            
            <div className="bg-gray-50 p-4 rounded mt-6">
              <p className="text-gray-600 text-sm italic">
                [Meeting continues for 47 minutes. Full transcript available upon request.]
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
