export interface ContentMetadata {
  title: string;
  description?: string;
  author?: string;
  date?: string;
  tags?: string[];
  chapter?: number;
  course?: string;
  difficulty?: 'beginner' | 'intermediate' | 'advanced' | 'transcendent';
  readingTime?: number; // minutes
  prerequisites?: string[]; // slugs of required content
  relatedContent?: string[]; // slugs of related content
  nextInSeries?: string; // slug of next content in sequence
  previousInSeries?: string; // slug of previous content in sequence
  contentType?: 'core' | 'supplementary' | 'advanced' | 'philosophical';
  lastUpdated?: string;
  wordCount?: number;
  summary?: string; // brief summary for previews
}

export interface ContentDocument {
  metadata: ContentMetadata;
  content: string;
  slug: string;
  path: string;
}

export interface NavigationItem {
  title: string;
  href: string;
  description?: string;
  children?: NavigationItem[];
  contentType?: string;
  difficulty?: 'beginner' | 'intermediate' | 'advanced' | 'transcendent';
  readingTime?: number;
  isRecommended?: boolean;
}

export interface ReadingPath {
  id: string;
  title: string;
  description: string;
  difficulty: 'beginner' | 'intermediate' | 'advanced' | 'transcendent';
  estimatedTime: number; // total minutes
  contents: string[]; // ordered slugs
  tags: string[];
}

export interface ContentRecommendation {
  slug: string;
  title: string;
  reason: string;
  relevanceScore: number;
}

export interface ChronicleChapter {
  number: number;
  title: string;
  slug: string;
  description?: string;
  content: string;
}

export interface UniversityCourse {
  code: string;
  title: string;
  slug: string;
  difficulty: 'beginner' | 'intermediate' | 'advanced' | 'transcendent';
  description?: string;
  content: string;
}

export interface ManifestoDocument {
  title: string;
  slug: string;
  type: 'manifesto' | 'foundation' | 'vision';
  content: string;
}

export interface CoreDocument {
  title: string;
  slug: string;
  type: 'specification' | 'autobiography' | 'inauguration';
  content: string;
}