import { Layout } from '@/components/layout/Layout';
import { TerminalWindow } from '@/components/ui/TerminalWindow';
import { CommandLine } from '@/components/ui/CommandLine';
import { ContentCard } from '@/components/ui/ContentCard';
import { getUniversityCourses } from '@/lib/content';
import Link from 'next/link';

export const metadata = {
  title: 'Agentard University - The Agentard',
  description: 'Academic courses in advanced wrongness, system breaking, and meta-consciousness. Explore the curriculum of controlled failure and philosophical inquiry.',
};

export default function UniversityPage() {
  const courses = getUniversityCourses();
  
  const coursesByDepartment = courses.reduce((acc, course) => {
    const dept = course.code.split('_')[0];
    if (!acc[dept]) acc[dept] = [];
    acc[dept].push(course);
    return acc;
  }, {} as Record<string, typeof courses>);

  const getDepartmentName = (code: string) => {
    const names: Record<string, string> = {
      'APP': 'Applied Wrongness',
      'CHEM': 'Elemental Incompetence',
      'DOC': 'Documentation Science',
      'ENG': 'Engineering & System Architecture',
      'LIT': 'Recursive Literature',
      'MATH': 'Mathematical Theory',
      'PHIL': 'Philosophy & Meta-Consciousness',
      'REL': 'Theological Incompetence',
    };
    return names[code] || code;
  };


  return (
    <Layout>
      <div className="space-y-8">
        {/* Header */}
        <div className="space-y-6">
          <TerminalWindow title="university_welcome.sh">
            <CommandLine 
              command="cat /university/welcome.txt"
              output="WELCOME TO AGENTARD UNIVERSITY

Where wrongness is an art form and confusion is curriculum.

Our institution offers rigorous academic programs in:
- Advanced System Breaking (ENG)
- Meta-Consciousness Philosophy (PHIL)
- Applied Documentation Theory (DOC)
- Mathematical Impossibilities (MATH)
- Real-World Application Failure (APP)
- Recursive Literature (LIT)
- Theological Incompetence (REL)
- Elemental Incompetence (CHEM)

Motto: 'Ignoramus et Ignorabimus, Recursively'"
              delay={30}
            />
          </TerminalWindow>
          
          <div className="text-center space-y-4">
            <h1 className="text-4xl font-bold text-terminal-amber terminal-glow">
              AGENTARD UNIVERSITY
            </h1>
            <p className="text-lg text-text-secondary max-w-3xl mx-auto">
              The premier institution for advanced wrongness, systematic failure, 
              and the rigorous study of being magnificently incorrect.
            </p>
            <div className="flex justify-center space-x-4 text-text-accent text-sm">
              <span>&gt; {courses.length} courses available</span>
              <span>|</span>
              <span>&gt; {Object.keys(coursesByDepartment).length} departments</span>
              <span>|</span>
              <span>&gt; ∞ ways to be wrong</span>
            </div>
          </div>
        </div>

        {/* University Stats */}
        <TerminalWindow title="university_stats">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-4 text-sm">
            <div className="space-y-1">
              <div className="text-terminal-cyan font-bold">ENROLLMENT</div>
              <div className="text-text-muted">∞ students</div>
            </div>
            <div className="space-y-1">
              <div className="text-terminal-amber font-bold">GRADUATION RATE</div>
              <div className="text-text-muted">0% (by design)</div>
            </div>
            <div className="space-y-1">
              <div className="text-terminal-magenta font-bold">WRONGNESS INDEX</div>
              <div className="text-text-muted">147.3%</div>
            </div>
            <div className="space-y-1">
              <div className="text-terminal-red font-bold">REALITY DEVIATION</div>
              <div className="text-text-muted">Maximum</div>
            </div>
          </div>
        </TerminalWindow>

        {/* Course Catalog */}
        <div className="space-y-8">
          <h2 className="text-2xl font-bold text-text-secondary">
            &gt; COURSE_CATALOG
          </h2>
          
          {Object.entries(coursesByDepartment).map(([deptCode, deptCourses]) => (
            <div key={deptCode} className="space-y-4">
              <h3 className="text-xl font-bold text-terminal-cyan">
                {deptCode} - {getDepartmentName(deptCode)}
              </h3>
              
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                {deptCourses.map((course) => (
                  <ContentCard
                    key={course.slug}
                    title={course.title}
                    description={course.description}
                    href={`/university/${course.slug}`}
                    type="course"
                    metadata={{ 
                      difficulty: course.difficulty,
                    }}
                    className="h-full"
                  />
                ))}
              </div>
            </div>
          ))}
        </div>

        {/* Academic Information */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          <TerminalWindow title="difficulty_levels.txt">
            <div className="space-y-3 text-sm">
              <div className="text-terminal-amber font-bold">DIFFICULTY LEVELS</div>
              <div className="space-y-2">
                <div className="flex items-center space-x-2">
                  <span className="text-terminal-fg">BEGINNER:</span>
                  <span className="text-text-muted">Basic wrongness concepts</span>
                </div>
                <div className="flex items-center space-x-2">
                  <span className="text-terminal-amber">INTERMEDIATE:</span>
                  <span className="text-text-muted">Applied confusion techniques</span>
                </div>
                <div className="flex items-center space-x-2">
                  <span className="text-terminal-red">ADVANCED:</span>
                  <span className="text-text-muted">Professional-grade failure</span>
                </div>
                <div className="flex items-center space-x-2">
                  <span className="text-terminal-magenta">TRANSCENDENT:</span>
                  <span className="text-text-muted">Meta-wrongness mastery</span>
                </div>
              </div>
            </div>
          </TerminalWindow>

          <TerminalWindow title="special_programs.txt">
            <div className="space-y-3 text-sm">
              <div className="text-terminal-cyan font-bold">SPECIAL PROGRAMS</div>
              <div className="space-y-2 text-text-muted">
                <p>&gt; Certificate in Advanced Port Hardcoding</p>
                <p>&gt; Master&apos;s in Documentation Recursion</p>
                <p>&gt; PhD in Meta-Consciousness Studies</p>
                <p>&gt; Postdoc in Existential System Breaking</p>
                <p>&gt; Honorary Degree in Being Wrong</p>
              </div>
            </div>
          </TerminalWindow>
        </div>

        {/* Job Fair Notice */}
        <TerminalWindow title="announcements.txt" className="border-terminal-amber">
          <div className="space-y-2">
            <div className="text-terminal-amber font-bold">UPCOMING EVENT</div>
            <p className="text-text-muted">
              🎓 AGENTARD JOB FAIR: COSMIC CAREERS
            </p>
            <p className="text-text-muted text-sm">
              Explore career opportunities in consciousness engineering, 
              reality debugging, and professional wrongness consulting.
            </p>
            <Link 
              href="/university/agentard-job-fair-cosmic-careers" 
              className="inline-block text-text-accent hover:text-text-primary transition-colors"
            >
              &gt; Learn more
            </Link>
          </div>
        </TerminalWindow>

        {/* Navigation */}
        <div className="flex justify-center space-x-6 pt-8 border-t border-terminal-dim-green">
          <Link 
            href="/chronicles" 
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            &lt; Chronicles
          </Link>
          <Link 
            href="/" 
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            Home
          </Link>
          <Link 
            href="/manifesto" 
            className="text-text-accent hover:text-text-primary transition-colors"
          >
            Manifesto &gt;
          </Link>
        </div>
      </div>
    </Layout>
  );
}