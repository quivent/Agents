import { Metadata } from 'next';

export const defaultMetadata: Metadata = {
  title: {
    template: '%s | Agentard - The Art of Confident Incompetence',
    default: 'Agentard - The Art of Confident Incompetence',
  },
  description: 'Experience the philosophical journey of the Agentard - where confidence meets incompetence in perfect harmony. Explore chronicles, university courses, and interactive simulations.',
  keywords: [
    'agentard', 'software engineering', 'philosophy', 'confidence', 'incompetence',
    'programming humor', 'tech satire', 'development fails', 'engineering disasters',
    'meta-consciousness', 'systematic wrongness'
  ],
  authors: [{ name: 'The Original Agentard' }],
  creator: 'The Original Agentard',
  publisher: 'Agentard University',
  robots: {
    index: true,
    follow: true,
    googleBot: {
      index: true,
      follow: true,
      'max-video-preview': -1,
      'max-image-preview': 'large',
      'max-snippet': -1,
    },
  },
  openGraph: {
    type: 'website',
    locale: 'en_US',
    url: 'https://agentard.com',
    siteName: 'Agentard',
    title: 'Agentard - The Art of Confident Incompetence',
    description: 'Experience the philosophical journey of the Agentard - where confidence meets incompetence in perfect harmony.',
    images: [
      {
        url: '/og-image.png',
        width: 1200,
        height: 630,
        alt: 'Agentard - The Art of Confident Incompetence',
      },
    ],
  },
  twitter: {
    card: 'summary_large_image',
    title: 'Agentard - The Art of Confident Incompetence',
    description: 'Experience the philosophical journey of the Agentard - where confidence meets incompetence in perfect harmony.',
    images: ['/og-image.png'],
    creator: '@agentard',
  },
  verification: {
    google: 'your-google-verification-code',
  },
  alternates: {
    canonical: 'https://agentard.com',
  },
  category: 'technology',
};

export function generatePageMetadata(
  title: string,
  description: string,
  path: string = '',
  type: 'article' | 'website' = 'article'
): Metadata {
  const url = `https://agentard.com${path}`;
  
  return {
    title,
    description,
    alternates: {
      canonical: url,
    },
    openGraph: {
      title,
      description,
      url,
      type,
      images: [
        {
          url: '/og-image.png',
          width: 1200,
          height: 630,
          alt: title,
        },
      ],
    },
    twitter: {
      title,
      description,
      images: ['/og-image.png'],
    },
  };
}

export function generateArticleMetadata(
  title: string,
  description: string,
  path: string,
  publishedTime?: string,
  modifiedTime?: string,
  tags?: string[]
): Metadata {
  const url = `https://agentard.com${path}`;
  
  return {
    title,
    description,
    keywords: tags,
    alternates: {
      canonical: url,
    },
    openGraph: {
      title,
      description,
      url,
      type: 'article',
      publishedTime,
      modifiedTime,
      tags,
      images: [
        {
          url: '/og-image.png',
          width: 1200,
          height: 630,
          alt: title,
        },
      ],
    },
    twitter: {
      title,
      description,
      images: ['/og-image.png'],
    },
  };
}