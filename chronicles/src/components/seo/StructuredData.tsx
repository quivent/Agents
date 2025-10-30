import Script from 'next/script';

interface WebsiteStructuredDataProps {
  name: string;
  description: string;
  url: string;
}

interface ArticleStructuredDataProps {
  headline: string;
  description: string;
  author: string;
  datePublished?: string;
  dateModified?: string;
  url: string;
  type?: 'BlogPosting' | 'Article' | 'TechArticle';
}

interface CourseStructuredDataProps {
  name: string;
  description: string;
  provider: string;
  difficulty: string;
  url: string;
}

export function WebsiteStructuredData({ name, description, url }: WebsiteStructuredDataProps) {
  const structuredData = {
    '@context': 'https://schema.org',
    '@type': 'WebSite',
    name,
    description,
    url,
    potentialAction: {
      '@type': 'SearchAction',
      target: {
        '@type': 'EntryPoint',
        urlTemplate: `${url}/search?q={search_term_string}`,
      },
      'query-input': 'required name=search_term_string',
    },
    mainEntity: {
      '@type': 'Organization',
      name: 'Agentard University',
      description: 'The premier institution for learning the art of confident incompetence',
      url,
      sameAs: [
        'https://github.com/agentard',
        'https://twitter.com/agentard',
      ],
    },
  };

  return (
    <Script
      id="website-structured-data"
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(structuredData) }}
    />
  );
}

export function ArticleStructuredData({
  headline,
  description,
  author,
  datePublished,
  dateModified,
  url,
  type = 'Article',
}: ArticleStructuredDataProps) {
  const structuredData = {
    '@context': 'https://schema.org',
    '@type': type,
    headline,
    description,
    author: {
      '@type': 'Person',
      name: author,
    },
    publisher: {
      '@type': 'Organization',
      name: 'Agentard University',
      logo: {
        '@type': 'ImageObject',
        url: `${url}/logo.png`,
      },
    },
    datePublished,
    dateModified: dateModified || datePublished,
    mainEntityOfPage: {
      '@type': 'WebPage',
      '@id': url,
    },
    image: `${url}/og-image.png`,
  };

  return (
    <Script
      id="article-structured-data"
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(structuredData) }}
    />
  );
}

export function CourseStructuredData({
  name,
  description,
  provider,
  difficulty,
  url,
}: CourseStructuredDataProps) {
  const structuredData = {
    '@context': 'https://schema.org',
    '@type': 'Course',
    name,
    description,
    provider: {
      '@type': 'Organization',
      name: provider,
    },
    educationalLevel: difficulty,
    url,
    teaches: 'The art of confident incompetence in software engineering',
    courseMode: 'online',
    availableLanguage: 'en',
  };

  return (
    <Script
      id="course-structured-data"
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(structuredData) }}
    />
  );
}

export function BreadcrumbStructuredData({ items }: { items: Array<{ name: string; url: string }> }) {
  const structuredData = {
    '@context': 'https://schema.org',
    '@type': 'BreadcrumbList',
    itemListElement: items.map((item, index) => ({
      '@type': 'ListItem',
      position: index + 1,
      name: item.name,
      item: item.url,
    })),
  };

  return (
    <Script
      id="breadcrumb-structured-data"
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(structuredData) }}
    />
  );
}