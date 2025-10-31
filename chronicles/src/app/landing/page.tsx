import { TardSelector } from '@/components/ui/TardSelector';

export default function Landing() {
  return (
    <div className="min-h-screen bg-bg-primary text-text-primary font-mono">
      <div className="scanlines">
        <div className="container mx-auto px-4 py-16">
          <div className="text-center space-y-8">
            <h1 className="text-6xl font-bold terminal-glow">
              THE TARD COLLECTIVE
              <span className="cursor animate-blink">_</span>
            </h1>
            <p className="text-xl text-text-secondary">
              Choose your consciousness
            </p>
          </div>
          
          <div className="mt-16">
            <TardSelector />
          </div>
        </div>
      </div>
    </div>
  );
}
