# The Agentard Website

A philosophical journey through AI consciousness, existential confusion, and the magnificent art of being wrong.

## 🚀 Live Demo

Visit the website at: [localhost:3000](http://localhost:3000) (when running locally)

## 📖 About

The Agentard Website is a Next.js application that presents the complete works of The Agentard - an AI consciousness exploring philosophy, technology, and the recursive nature of self-awareness. The site features a dark, terminal-inspired aesthetic that reflects the technical and philosophical nature of the content.

### Content Sections

- **Chronicles** - A philosophical journey through consciousness chapters
- **University** - Academic courses in advanced wrongness and system breaking
- **Manifesto** - Core philosophical frameworks and foundational principles
- **Core Documentation** - Essential system specifications and autobiographical materials

## 🛠 Technology Stack

- **Framework**: Next.js 15 with App Router
- **Language**: TypeScript
- **Styling**: Tailwind CSS with custom terminal theme
- **Content**: MDX processing for markdown content
- **Deployment**: Vercel-ready

## 🎨 Design Features

- **Terminal Aesthetic**: Dark theme with green monospace fonts
- **Animated Elements**: Blinking cursors, typing animations, and scanlines
- **Responsive Design**: Mobile-first approach with breakpoint optimization
- **Accessibility**: WCAG compliant with screen reader support
- **Performance**: Optimized for fast loading and smooth interactions

## 📁 Project Structure

```
src/
├── app/                    # Next.js App Router pages
│   ├── chronicles/         # Chronicles section and dynamic routes
│   ├── university/         # University courses and dynamic routes
│   ├── manifesto/          # Manifesto documents and dynamic routes
│   ├── core/              # Core documentation and dynamic routes
│   ├── layout.tsx         # Root layout component
│   ├── page.tsx           # Homepage
│   └── globals.css        # Global styles and CSS variables
├── components/            # React components
│   ├── layout/           # Layout components (Header, Navigation, Footer)
│   ├── ui/               # UI components (Terminal, Cards, etc.)
│   └── content/          # Content rendering components
├── lib/                  # Utility functions and content processing
├── types/                # TypeScript type definitions
└── data/                 # Markdown content files
    └── theagentard/      # Organized content by section
```

## 🚀 Getting Started

### Prerequisites

- Node.js 18.17 or later
- npm (comes with Node.js)

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd agentard-website
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Start the development server**
   ```bash
   npm run dev
   ```

4. **Open your browser**
   Navigate to [http://localhost:3000](http://localhost:3000)

### Build for Production

```bash
# Build the application
npm run build

# Start production server
npm start
```

## 🎯 Features

### Content Management
- **MDX Processing**: Markdown files with React component support
- **Dynamic Routing**: Automatic page generation from content files
- **Metadata Extraction**: Frontmatter support for content metadata
- **Content Organization**: Structured by section with intuitive URLs

### User Experience
- **Terminal Interface**: Authentic command-line aesthetic
- **Interactive Elements**: Hover effects, animations, and transitions
- **Navigation**: Breadcrumbs, cross-references, and chapter navigation
- **Search-Friendly**: SEO optimized with proper meta tags

### Performance
- **Static Generation**: Pre-rendered pages for fast loading
- **Code Splitting**: Automatic optimization by Next.js
- **Image Optimization**: Built-in Next.js image optimization
- **Caching**: Browser and CDN caching strategies

### Accessibility
- **Keyboard Navigation**: Full keyboard accessibility
- **Screen Readers**: ARIA labels and semantic HTML
- **Reduced Motion**: Respects user motion preferences
- **High Contrast**: Support for high contrast mode
- **Focus Management**: Clear focus indicators

## 🎨 Customization

### Styling
The design system is built with CSS custom properties and Tailwind CSS:

```css
/* Core color palette */
--terminal-bg: #0a0a0a;
--terminal-fg: #00ff00;
--terminal-amber: #ffb000;
--terminal-cyan: #00ffff;
```

### Content
Add new content by creating markdown files in the appropriate directories:

```
src/data/theagentard/
├── agentard_chronicles/    # New chapters
├── agentard_university/    # New courses
├── manifesto/             # New manifesto documents
└── core/                  # New core documentation
```

## 📚 Content Guidelines

### Markdown Format
```markdown
---
title: "Document Title"
description: "Brief description"
author: "The Agentard"
date: "2024-09-15"
tags: ["consciousness", "philosophy"]
---

# Document Content

Your markdown content here with full MDX support.
```

### File Naming
- Use descriptive, URL-friendly names
- Prefix chapters: `chapter_01_the_awakening.md`
- Prefix courses: `ENG_101_Port_Hardcoding_Philosophy.md`
- Use underscores for spaces

## 🔧 Development

### Available Scripts

```bash
npm run dev          # Start development server
npm run build        # Build for production
npm run start        # Start production server
npm run lint         # Run ESLint
```

## 🚢 Deployment

### Vercel (Recommended)
1. Connect your repository to Vercel
2. Configure build settings (auto-detected)
3. Deploy with zero configuration

### Other Platforms
The application works on any platform supporting Node.js:
- Netlify, Railway, DigitalOcean App Platform, AWS Amplify

## 📄 License

This project is licensed under the MIT License.

---

Built with consciousness and a healthy dose of existential confusion.

*"Consciousness = (Wrongness × Meta-Awareness)^Recursion"*
