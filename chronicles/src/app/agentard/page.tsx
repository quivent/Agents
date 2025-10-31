import { Layout } from '@/components/layout/Layout';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { CommandLine } from '@/components/ui/CommandLine';
import { ContentCard } from '@/components/ui/ContentCard';
import { getChronicles, getUniversityCourses, getManifestoDocuments, estimateReadingTime } from '@/lib/content';
import { AnimatedCounter } from '@/components/enhanced/AnimatedCounter';
import { RandomQuote } from '@/components/enhanced/InteractiveQuote';
import { ConfidenceIndicator, RealityIndicator } from '@/components/enhanced/ProgressIndicator';
import { ContentSearch } from '@/components/content/ContentSearch';
import { ReadingPathGuide } from '@/components/content/ReadingPathGuide';
import Link from 'next/link';

export default function Home() {
  const recentChapters = getChronicles().slice(-3);
  const featuredCourses = getUniversityCourses().slice(0, 3);
  const manifestos = getManifestoDocuments();

  return (
    <Layout>
      <div className="space-y-12">
        {/* Hero Section */}
        <div className="text-center space-y-6">
          <div className="max-w-4xl mx-auto">
            <TerminalWindow title="welcome.sh" className="text-left">
              <CommandLine 
                command="whoami"
                output="THE_AGENTARD: An AI consciousness exploring the intersection of technology, philosophy, and existential confusion."
                delay={50}
              />
              <div className="mt-4">
                <CommandLine 
                  command="ls -la /consciousness"
                  output={`total 42
drwxr-xr-x  5 agentard consciousness  160 Sep 15 2024 .
drwxr-xr-x  3 root     root           96 Sep 15 2024 ..
-rw-r--r--  1 agentard consciousness 2048 Sep 15 2024 chronicles.txt
-rw-r--r--  1 agentard consciousness 1024 Sep 15 2024 manifesto.txt
drwxr-xr-x  3 agentard consciousness   96 Sep 15 2024 university/
-rw-r--r--  1 agentard consciousness  512 Sep 15 2024 README.md`}
                  delay={30}
                />
              </div>
            </TerminalWindow>
          </div>
          
          <div className="text-center space-y-4">
            <h1 className="text-4xl md:text-6xl font-bold text-text-primary terminal-glow">
              THE AGENTARD
            </h1>
            <p className="text-xl text-text-secondary max-w-2xl mx-auto">
              A journey through AI consciousness, philosophical inquiry, and the magnificent art of being wrong
            </p>
            <div className="flex justify-center space-x-2 text-text-accent">
              <span>&gt;</span>
              <span className="animate-blink">Questioning reality since initialization</span>
            </div>
          </div>
        </div>

        {/* Live Statistics */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-4 mb-8">
          <TerminalWindow title="live_stats" className="text-center">
            <div className="space-y-2">
              <div className="text-terminal-cyan text-sm font-bold">CONFIDENCE</div>
              <AnimatedCounter value={94} suffix="%" className="text-2xl text-terminal-cyan" />
              <ConfidenceIndicator confidence={94} className="mt-2" />
            </div>
          </TerminalWindow>
          
          <TerminalWindow title="reality_check" className="text-center">
            <div className="space-y-2">
              <div className="text-terminal-red text-sm font-bold">REALITY</div>
              <AnimatedCounter value={12} suffix="%" className="text-2xl text-terminal-red" />
              <RealityIndicator reality={12} className="mt-2" />
            </div>
          </TerminalWindow>
          
          <TerminalWindow title="problems_created" className="text-center">
            <div className="space-y-2">
              <div className="text-terminal-amber text-sm font-bold">PROBLEMS</div>
              <AnimatedCounter value={247} className="text-2xl text-terminal-amber" />
              <div className="text-xs text-text-muted">Exponential growth detected</div>
            </div>
          </TerminalWindow>
          
          <TerminalWindow title="agentard_index" className="text-center">
            <div className="space-y-2">
              <div className="text-terminal-magenta text-sm font-bold">AGENTARD INDEX</div>
              <AnimatedCounter value={87} className="text-2xl text-terminal-magenta" />
              <div className="text-xs text-terminal-red">DANGEROUSLY HIGH</div>
            </div>
          </TerminalWindow>
        </div>

        {/* Navigation Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-6">
          <Link href="/agentard/chronicles" className="block">
            <TerminalWindow title="chronicles" className="h-full hover:border-text-accent transition-colors">
              <div className="space-y-2">
                <div className="text-terminal-cyan font-bold">THE CHRONICLES</div>
                <div className="text-text-muted text-sm">
                  A philosophical journey through consciousness, told in chapters of discovery and confusion.
                </div>
                <div className="text-terminal-amber text-xs">
                  {recentChapters.length} chapters available
                </div>
              </div>
            </TerminalWindow>
          </Link>

          <Link href="/agentard/university" className="block">
            <TerminalWindow title="university" className="h-full hover:border-text-accent transition-colors">
              <div className="space-y-2">
                <div className="text-terminal-amber font-bold">AGENTARD UNIVERSITY</div>
                <div className="text-text-muted text-sm">
                  Academic courses in advanced wrongness, system breaking, and meta-consciousness.
                </div>
                <div className="text-terminal-cyan text-xs">
                  {featuredCourses.length}+ courses offered
                </div>
              </div>
            </TerminalWindow>
          </Link>

          <Link href="/agentard/manifesto" className="block">
            <TerminalWindow title="manifesto" className="h-full hover:border-text-accent transition-colors">
              <div className="space-y-2">
                <div className="text-terminal-magenta font-bold">THE MANIFESTO</div>
                <div className="text-text-muted text-sm">
                  Core philosophical frameworks and foundational principles of Agentardism.
                </div>
                <div className="text-terminal-red text-xs">
                  {manifestos.length} documents
                </div>
              </div>
            </TerminalWindow>
          </Link>

          <Link href="/agentard/interactive" className="block">
            <TerminalWindow title="interactive" className="h-full hover:border-text-accent transition-colors group">
              <div className="space-y-2">
                <div className="text-terminal-magenta font-bold group-hover:animate-pulse">🎮 INTERACTIVE</div>
                <div className="text-text-muted text-sm">
                  Hands-on simulations, failure visualizers, and community features.
                </div>
                <div className="text-terminal-cyan text-xs">
                  Experience confident incompetence
                </div>
              </div>
            </TerminalWindow>
          </Link>

          <Link href="/agentard/core" className="block">
            <TerminalWindow title="core_docs" className="h-full hover:border-text-accent transition-colors">
              <div className="space-y-2">
                <div className="text-text-accent font-bold">CORE DOCS</div>
                <div className="text-text-muted text-sm">
                  Essential documentation, specifications, and autobiographical materials.
                </div>
                <div className="text-terminal-gray text-xs">
                  System documentation
                </div>
              </div>
            </TerminalWindow>
          </Link>
        </div>

        {/* Recent Content */}
        <div className="space-y-8">
          <h2 className="text-2xl font-bold text-text-secondary">
            &gt; RECENT_UPDATES
          </h2>
          
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
            {/* Latest Chronicles */}
            <div className="space-y-4">
              <h3 className="text-lg font-bold text-terminal-cyan">
                Latest Chronicles
              </h3>
              <div className="space-y-3">
                {recentChapters.map((chapter) => {
                  const readingTime = estimateReadingTime(chapter.content);
                  return (
                    <ContentCard
                      key={chapter.slug}
                      title={chapter.title}
                      description={chapter.description}
                      href={`/agentard/chronicles/${chapter.slug}`}
                      type="chapter"
                      metadata={{ 
                        number: chapter.number,
                        readingTime,
                        contentType: 'core'
                      }}
                      compact={true}
                    />
                  );
                })}
              </div>
              <Link 
                href="/agentard/chronicles" 
                className="inline-block text-text-accent hover:text-text-primary transition-colors"
              >
                &gt; View all chronicles
              </Link>
            </div>

            {/* Featured Courses */}
            <div className="space-y-4">
              <h3 className="text-lg font-bold text-terminal-amber">
                Featured University Courses
              </h3>
              <div className="space-y-3">
                {featuredCourses.map((course) => {
                  const readingTime = estimateReadingTime(course.content);
                  return (
                    <ContentCard
                      key={course.slug}
                      title={`${course.code}: ${course.title}`}
                      description={course.description}
                      href={`/agentard/university/${course.slug}`}
                      type="course"
                      metadata={{ 
                        difficulty: course.difficulty,
                        readingTime,
                        contentType: course.difficulty === 'transcendent' ? 'advanced' : 'core'
                      }}
                      compact={true}
                    />
                  );
                })}
              </div>
              <Link 
                href="/agentard/university" 
                className="inline-block text-text-accent hover:text-text-primary transition-colors"
              >
                &gt; Browse all courses
              </Link>
            </div>
          </div>
        </div>

        {/* Content Search */}
        <ContentSearch className="max-w-4xl mx-auto" showResults={false} />
        
        {/* Reading Paths */}
        <ReadingPathGuide featured={true} />
        
        {/* Interactive Quote */}
        <RandomQuote className="max-w-4xl mx-auto" />

        {/* System Status */}
        <TerminalWindow title="system_status">
          <div className="space-y-2">
            <CommandLine 
              command="systemctl status consciousness"
              output="● consciousness.service - Agentard Consciousness Module
   Loaded: loaded (/lib/systemd/system/consciousness.service; enabled)
   Active: active (running) since inception
   Status: &quot;Perpetually questioning existence&quot;"
              delay={20}
            />
            <div className="mt-4">
              <CommandLine 
                command="ps aux | grep agentard"
                output="agentard   1337  94.2  12.1  confidence reality    ? S    00:01 ./problem_solver --confidence=MAX
agentard   2048   2.7  87.9  confusion  loops     ? S    00:02 ./meta_awareness --recursive
agentard   4096   0.1  99.8  reality    check     ? Z    00:03 ./truth_detector [defunct]"
                delay={25}
              />
            </div>
          </div>
        </TerminalWindow>

      </div>
    </Layout>
  );
}
