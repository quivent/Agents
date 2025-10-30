import { ContentDocument } from '@/types/content';

let contentCache: ContentDocument[] | null = null;

async function loadAllContent(): Promise<ContentDocument[]> {
  if (contentCache) {
    return contentCache;
  }

  try {
    const response = await fetch('/content-data.json');
    if (!response.ok) {
      throw new Error('Failed to fetch content');
    }
    const data: ContentDocument[] = await response.json();
    contentCache = data;
    return data;
  } catch (error) {
    console.error('Content fetch error:', error);
    return [];
  }
}

export async function searchContentAPI(query: string): Promise<ContentDocument[]> {
  try {
    const allContent = await loadAllContent();
    const searchTerm = query.toLowerCase().trim();

    if (!searchTerm) {
      return [];
    }

    // Simple client-side search
    return allContent.filter(doc => {
      const titleMatch = doc.metadata.title?.toLowerCase().includes(searchTerm);
      const descMatch = doc.metadata.description?.toLowerCase().includes(searchTerm);
      const contentMatch = doc.content?.toLowerCase().includes(searchTerm);
      const tagMatch = doc.metadata.tags?.some(tag => tag.toLowerCase().includes(searchTerm));

      return titleMatch || descMatch || contentMatch || tagMatch;
    });
  } catch (error) {
    console.error('Search error:', error);
    return [];
  }
}

export async function getAllContentAPI(): Promise<ContentDocument[]> {
  return loadAllContent();
}