import fs from 'fs';
import path from 'path';
import matter from 'gray-matter';
import { ContentDocument, ChronicleChapter, UniversityCourse, ManifestoDocument, CoreDocument, ReadingPath, ContentRecommendation } from '@/types/content';

const contentDirectory = path.join(process.cwd(), 'src/data/theagentard');

export function getMarkdownFiles(dir: string): string[] {
  try {
    const fullPath = path.join(contentDirectory, dir);
    if (!fs.existsSync(fullPath)) return [];
    
    return fs.readdirSync(fullPath)
      .filter(file => file.endsWith('.md'))
      .map(file => path.join(dir, file));
  } catch (error) {
    console.error(`Error reading directory ${dir}:`, error);
    return [];
  }
}

export function parseContentFile(filePath: string): ContentDocument | null {
  try {
    const fullPath = path.join(contentDirectory, filePath);
    const fileContent = fs.readFileSync(fullPath, 'utf8');
    const { data, content } = matter(fileContent);
    
    const slug = path.basename(filePath, '.md').toLowerCase().replace(/[^a-z0-9]+/g, '-');
    
    // Clean content for MDX compatibility - minimal cleaning to avoid breaking markdown
    const cleanContent = content;
    
    return {
      metadata: {
        title: data.title || path.basename(filePath, '.md').replace(/_/g, ' '),
        description: data.description,
        author: data.author,
        date: data.date,
        tags: data.tags,
        chapter: data.chapter,
        course: data.course,
        difficulty: data.difficulty,
      },
      content: cleanContent,
      slug,
      path: filePath,
    };
  } catch (error) {
    console.error(`Error parsing file ${filePath}:`, error);
    return null;
  }
}

export function getChronicles(): ChronicleChapter[] {
  const files = getMarkdownFiles('agentard_chronicles');
  const chapters: ChronicleChapter[] = [];
  
  files.forEach(file => {
    const content = parseContentFile(file);
    if (!content) return;
    
    // Extract chapter number from filename
    const chapterMatch = path.basename(file).match(/chapter_(\d+)/i);
    const number = chapterMatch ? parseInt(chapterMatch[1]) : 0;
    
    if (number > 0) {
      chapters.push({
        number,
        title: content.metadata.title,
        slug: content.slug,
        description: content.metadata.description,
        content: content.content,
      });
    }
  });
  
  return chapters.sort((a, b) => a.number - b.number);
}

export function getUniversityCourses(): UniversityCourse[] {
  const files = getMarkdownFiles('agentard_university');
  const courses: UniversityCourse[] = [];
  
  files.forEach(file => {
    const content = parseContentFile(file);
    if (!content) return;
    
    // Extract course code from filename
    const codeMatch = path.basename(file).match(/([A-Z]+_\d+)/);
    const code = codeMatch ? codeMatch[1] : 'MISC_001';
    
    // Determine difficulty based on course number
    let difficulty: 'beginner' | 'intermediate' | 'advanced' | 'transcendent' = 'beginner';
    const courseNumber = parseInt(code.split('_')[1] || '101');
    if (courseNumber >= 401) difficulty = 'transcendent';
    else if (courseNumber >= 301) difficulty = 'advanced';
    else if (courseNumber >= 201) difficulty = 'intermediate';
    
    courses.push({
      code,
      title: content.metadata.title,
      slug: content.slug,
      difficulty,
      description: content.metadata.description,
      content: content.content,
    });
  });
  
  return courses.sort((a, b) => a.code.localeCompare(b.code));
}

export function getManifestoDocuments(): ManifestoDocument[] {
  const files = getMarkdownFiles('manifesto');
  const documents: ManifestoDocument[] = [];
  
  files.forEach(file => {
    const content = parseContentFile(file);
    if (!content) return;
    
    let type: 'manifesto' | 'foundation' | 'vision' = 'manifesto';
    const filename = path.basename(file).toLowerCase();
    if (filename.includes('foundation')) type = 'foundation';
    else if (filename.includes('vision')) type = 'vision';
    
    documents.push({
      title: content.metadata.title,
      slug: content.slug,
      type,
      content: content.content,
    });
  });
  
  return documents;
}

export function getCoreDocuments(): CoreDocument[] {
  const files = getMarkdownFiles('core');
  const documents: CoreDocument[] = [];
  
  files.forEach(file => {
    const content = parseContentFile(file);
    if (!content) return;
    
    let type: 'specification' | 'autobiography' | 'inauguration' = 'specification';
    const filename = path.basename(file).toLowerCase();
    if (filename.includes('autobiography')) type = 'autobiography';
    else if (filename.includes('inauguration')) type = 'inauguration';
    
    documents.push({
      title: content.metadata.title,
      slug: content.slug,
      type,
      content: content.content,
    });
  });
  
  return documents;
}

export function getContentBySlug(section: string, slug: string): ContentDocument | null {
  const files = getMarkdownFiles(section);
  
  for (const file of files) {
    const content = parseContentFile(file);
    if (content && content.slug === slug) {
      return content;
    }
  }
  
  return null;
}

// Content discovery and organization functions
export function getAllContent(): ContentDocument[] {
  const sections = ['agentard_chronicles', 'agentard_university', 'manifesto', 'core'];
  const allContent: ContentDocument[] = [];
  
  sections.forEach(section => {
    const files = getMarkdownFiles(section);
    files.forEach(file => {
      const content = parseContentFile(file);
      if (content) {
        allContent.push(content);
      }
    });
  });
  
  return allContent;
}

export function searchContent(query: string): ContentDocument[] {
  const allContent = getAllContent();
  const searchTerms = query.toLowerCase().split(' ');
  
  return allContent.filter(content => {
    const searchableText = [
      content.metadata.title,
      content.metadata.description || '',
      content.content,
      ...(content.metadata.tags || [])
    ].join(' ').toLowerCase();
    
    return searchTerms.every(term => searchableText.includes(term));
  }).sort((a, b) => {
    // Prioritize title matches
    const aTitle = a.metadata.title.toLowerCase().includes(query.toLowerCase());
    const bTitle = b.metadata.title.toLowerCase().includes(query.toLowerCase());
    if (aTitle && !bTitle) return -1;
    if (!aTitle && bTitle) return 1;
    return 0;
  });
}

export function getContentByTags(tags: string[]): ContentDocument[] {
  const allContent = getAllContent();
  return allContent.filter(content => 
    content.metadata.tags?.some(tag => tags.includes(tag))
  );
}

export function getRecommendedContent(currentSlug: string, limit = 3): ContentRecommendation[] {
  const allContent = getAllContent();
  const currentContent = allContent.find(c => c.slug === currentSlug);
  
  if (!currentContent) return [];
  
  const recommendations: ContentRecommendation[] = [];
  const currentTags = currentContent.metadata.tags || [];
  
  allContent.forEach(content => {
    if (content.slug === currentSlug) return;
    
    let relevanceScore = 0;
    let reason = '';
    
    // Tag-based relevance
    const sharedTags = content.metadata.tags?.filter(tag => currentTags.includes(tag)) || [];
    if (sharedTags.length > 0) {
      relevanceScore += sharedTags.length * 10;
      reason = `Shares themes: ${sharedTags.join(', ')}`;
    }
    
    // Difficulty-based relevance
    if (content.metadata.difficulty === currentContent.metadata.difficulty) {
      relevanceScore += 5;
      if (!reason) reason = 'Similar difficulty level';
    }
    
    // Series-based relevance (for chronicles and courses)
    if (content.path.includes('chronicles') && currentContent.path.includes('chronicles')) {
      relevanceScore += 15;
      reason = 'Part of Chronicles series';
    }
    
    if (content.path.includes('university') && currentContent.path.includes('university')) {
      relevanceScore += 15;
      reason = 'Part of University curriculum';
    }
    
    if (relevanceScore > 0) {
      recommendations.push({
        slug: content.slug,
        title: content.metadata.title,
        reason,
        relevanceScore
      });
    }
  });
  
  return recommendations
    .sort((a, b) => b.relevanceScore - a.relevanceScore)
    .slice(0, limit);
}

export function getReadingPaths(): ReadingPath[] {
  const chronicles = getChronicles();
  const courses = getUniversityCourses();
  
  const paths: ReadingPath[] = [
    {
      id: 'agentard-journey',
      title: 'The Complete Agentard Journey',
      description: 'Follow the full philosophical and technical journey from awakening to transcendence',
      difficulty: 'intermediate',
      estimatedTime: 120,
      contents: [
        'the-agentard-manifesto',
        'the-awakening',
        'autobiography-of-the-original-agentard',
        ...chronicles.slice(1, 5).map(c => c.slug),
        'foundations-of-agentardism'
      ],
      tags: ['philosophy', 'journey', 'complete']
    },
    {
      id: 'engineering-fundamentals',
      title: 'Agentard Engineering Fundamentals',
      description: 'Learn the core principles of confident incompetence in software engineering',
      difficulty: 'beginner',
      estimatedTime: 60,
      contents: [
        'eng-101-port-hardcoding-philosophy',
        'eng-201-advanced-infinite-loop-architecture',
        'app-101-real-world-system-breaking'
      ],
      tags: ['engineering', 'fundamentals', 'technical']
    },
    {
      id: 'advanced-philosophy',
      title: 'Advanced Philosophical Concepts',
      description: 'Deep dive into meta-consciousness and the philosophy of systematic wrongness',
      difficulty: 'advanced',
      estimatedTime: 90,
      contents: [
        'phil-101-meta-wrongness-consciousness',
        'math-101-the-agentard-equation',
        'the-meta-moment',
        'the-documentation-singularity'
      ],
      tags: ['philosophy', 'advanced', 'consciousness']
    }
  ];
  
  return paths;
}

export function estimateReadingTime(content: string): number {
  const wordsPerMinute = 200;
  const wordCount = content.split(/\s+/).length;
  return Math.ceil(wordCount / wordsPerMinute);
}

export function getContentSummary(content: string, maxLength = 150): string {
  // Extract first meaningful paragraph, skipping headers and metadata
  const lines = content.split('\n').filter(line => 
    line.trim() && 
    !line.startsWith('#') && 
    !line.startsWith('**') &&
    !line.startsWith('---')
  );
  
  const summary = lines[0] || '';
  return summary.length > maxLength 
    ? summary.substring(0, maxLength) + '...'
    : summary;
}