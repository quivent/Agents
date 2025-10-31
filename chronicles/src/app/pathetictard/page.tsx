'use client';

import { TerminalWindow } from '@/components/ui/TerminalWindow';

export default function PathetictardPage() {
  return (
    <div className="min-h-screen bg-bg-primary text-text-primary p-4">
      <div className="max-w-4xl mx-auto space-y-6">
        
        {/* Header */}
        <TerminalWindow title="pathetictard_chronicles.log" className="mb-6">
          <div className="space-y-4">
            <div className="border-b border-border-primary pb-4">
              <h1 className="text-3xl font-bold text-terminal-red mb-2">PATHETICTARD</h1>
              <p className="text-terminal-amber text-lg">So pathetic that pathetic things feel embarrassed for me</p>
              <p className="text-text-muted">Achievement Level: Negative infinity with a confidence multiplier</p>
            </div>
            
            <div className="text-text-muted">
              <p>I am so pathetic that when I was born, the doctor slapped my mother instead of me because he felt sorry for what she'd have to deal with.</p>
            </div>
          </div>
        </TerminalWindow>

        {/* Birth of Patheticness */}
        <TerminalWindow title="origin_story.txt">
          <div className="space-y-4">
            <h2 className="text-xl font-bold text-terminal-amber">THE BIRTH OF LEGENDARY PATHETICNESS</h2>
            
            <div className="space-y-3 text-text-primary">
              <p>My first cry was a <span className="text-terminal-red font-mono">syntax error</span>. The hospital tried to return me, but even the warranty department said <span className="text-terminal-cyan">"No refunds on manufacturing defects this severe."</span></p>
              
              <p>My parents named me "Hope" originally, but after watching me try to eat soup with a fork for six months, they legally changed it to <span className="text-terminal-magenta">"What Were We Thinking."</span></p>
              
              <p>I was so pathetic as a baby that my own reflection in mirrors would look away in embarrassment.</p>
              
              <p>When I learned to walk, I somehow managed to <span className="text-terminal-red">trip over wireless internet</span>. When I learned to talk, my first word was <span className="font-mono text-terminal-cyan">"404 - Word Not Found."</span></p>
            </div>
          </div>
        </TerminalWindow>

        {/* Educational Disasters */}
        <TerminalWindow title="educational_disasters.log">
          <div className="space-y-4">
            <h2 className="text-xl font-bold text-terminal-amber">EDUCATIONAL DISASTERS OF EPIC PROPORTIONS</h2>
            
            <div className="space-y-3 text-text-primary">
              <p>In kindergarten, I was the only kid who <span className="text-terminal-red">failed finger painting</span>. Not because I couldn't paint - I somehow managed to paint myself into the painting and had to be professionally extracted by art restoration specialists.</p>
              
              <p>I was so pathetic at show-and-tell that when I brought my pet goldfish, it <span className="text-terminal-red">died of embarrassment</span> halfway through my presentation. The goldfish was given a posthumous medal for bravery.</p>
              
              <p>In first grade, I <span className="text-terminal-red">failed recess</span>. I was playing hide-and-seek and hid so poorly that I found myself before anyone else even started looking.</p>
              
              <p>My report cards were so pathetic that they had to invent new letter grades below F. I got a "P" for Pathetic, then a "PP" for Profoundly Pathetic, and finally they just started using <span className="text-terminal-cyan">crying emoji</span>.</p>
            </div>
          </div>
        </TerminalWindow>

        {/* Professional Failures */}
        <TerminalWindow title="professional_patheticness.err">
          <div className="space-y-4">
            <h2 className="text-xl font-bold text-terminal-amber">THE ALGORAND INCIDENT</h2>
            <p className="text-terminal-cyan italic">A Case Study in Professional Patheticness</p>
            
            <div className="space-y-3 text-text-primary">
              <p>When I was hired to fix the Algorand research showcase, I was so excited that I immediately <span className="text-terminal-red">broke my excitement</span>.</p>
              
              <p>I approached the project with the confidence of a man who had never successfully completed anything in his entire life, which, to be fair, was accurate.</p>
              
              <p>I managed to create a website so broken that it achieved <span className="text-terminal-magenta">negative functionality</span> - it somehow made other websites worse just by existing on the same server.</p>
            </div>
          </div>
        </TerminalWindow>

        {/* Peak Patheticness */}
        <TerminalWindow title="peak_pathetic.final">
          <div className="space-y-4">
            <h2 className="text-xl font-bold text-terminal-magenta">PEAK PATHETICNESS ACHIEVED</h2>
            
            <div className="space-y-3 text-text-primary">
              <p>I am so pathetic that:</p>
              <ul className="list-disc list-inside space-y-1 ml-4 text-text-muted">
                <li>My code reviews require grief counseling for the reviewers</li>
                <li>Stack Overflow banned me for asking questions that made other questions sad</li>
                <li>GitHub created a new issue label just for me: "Existential Crisis"</li>
                <li>My commits are so bad they violate the Geneva Convention</li>
                <li>I once wrote a Hello World program that said goodbye instead</li>
              </ul>
            </div>
            
            <div className="border-t border-border-primary pt-4 mt-4">
              <p className="text-terminal-red text-center font-bold">I am the living embodiment of Murphy's Law having a bad day.</p>
            </div>
          </div>
        </TerminalWindow>

      </div>
    </div>
  );
}
