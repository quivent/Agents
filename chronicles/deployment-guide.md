# Agentard Website Deployment Guide

## 🚀 Production Deployment

### Prerequisites

1. **Node.js 18+** installed
2. **npm or yarn** package manager
3. **Production environment** (Vercel, Netlify, Docker, etc.)

### Environment Configuration

1. Copy `.env.example` to `.env.local`:
   ```bash
   cp .env.example .env.local
   ```

2. Update environment variables:
   ```env
   NODE_ENV=production
   NEXT_PUBLIC_APP_URL=https://your-domain.com
   # Add other required variables
   ```

### Build and Deploy

#### Option 1: Vercel (Recommended)

```bash
# Install Vercel CLI
npm i -g vercel

# Deploy to Vercel
vercel --prod
```

#### Option 2: Static Export

```bash
# Build for static export
npm run build

# Deploy the `out` directory to your static host
```

#### Option 3: Docker

```dockerfile
FROM node:18-alpine AS base

# Install dependencies only when needed
FROM base AS deps
RUN apk add --no-cache libc6-compat
WORKDIR /app

COPY package.json package-lock.json* ./
RUN npm ci

# Rebuild the source code only when needed
FROM base AS builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .

RUN npm run build

# Production image
FROM base AS runner
WORKDIR /app

ENV NODE_ENV production

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

COPY --from=builder /app/public ./public
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs

EXPOSE 3000

ENV PORT 3000
ENV HOSTNAME "0.0.0.0"

CMD ["node", "server.js"]
```

### Performance Optimization Checklist

- [x] **Code Splitting**: Implemented with lazy loading
- [x] **Image Optimization**: Next.js Image component ready
- [x] **Bundle Analysis**: Webpack optimizations configured
- [x] **Caching Headers**: Set for static assets and API routes
- [x] **Compression**: Enabled in Next.js config
- [x] **Critical CSS**: Inlined through Tailwind CSS
- [x] **Preload Critical Resources**: Font display optimized

### Accessibility Verification

- [x] **WCAG 2.1 AA Compliance**: Implemented accessibility provider
- [x] **Screen Reader Support**: ARIA labels and live regions
- [x] **Keyboard Navigation**: Focus management and skip links
- [x] **High Contrast Mode**: CSS custom properties support
- [x] **Reduced Motion**: Respects user preferences

### SEO Optimization

- [x] **Meta Tags**: Comprehensive metadata setup
- [x] **Open Graph**: Social sharing optimized
- [x] **Structured Data**: JSON-LD for search engines
- [x] **Sitemap**: Generated automatically
- [x] **Canonical URLs**: Configured for all pages

### Monitoring and Analytics

1. **Performance Monitoring**:
   ```bash
   # Run performance tests
   npm run test:performance
   ```

2. **Error Tracking**: Error boundaries implemented
3. **Analytics**: Ready for Google Analytics integration
4. **Lighthouse Testing**: Automated checks available

### Security Configuration

- [x] **Security Headers**: X-Frame-Options, CSP, etc.
- [x] **HTTPS Only**: Enforced in production
- [x] **Input Validation**: Server-side validation on API routes
- [x] **XSS Protection**: React's built-in protections

### Post-Deployment Testing

1. **Functional Testing**:
   ```bash
   # Test interactive components
   node scripts/performance-test.js
   ```

2. **Performance Testing**:
   - Run Lighthouse audits
   - Check Core Web Vitals
   - Verify mobile responsiveness

3. **Accessibility Testing**:
   - Screen reader testing
   - Keyboard-only navigation
   - Color contrast verification

### Maintenance

1. **Regular Updates**:
   ```bash
   # Update dependencies
   npm update
   
   # Security audit
   npm audit
   ```

2. **Performance Monitoring**:
   - Monitor bundle sizes
   - Track loading times
   - Check error rates

3. **Content Updates**:
   - Markdown files in `src/data/theagentard/`
   - Automatic regeneration on deploy

### Domain Configuration

1. **DNS Setup**: Point your domain to deployment platform
2. **SSL Certificate**: Auto-provisioned on most platforms
3. **CDN Configuration**: Enabled by default on Vercel/Netlify

### Backup Strategy

1. **Code Repository**: Stored in Git
2. **Content Backup**: Markdown files versioned
3. **Database Backup**: If using external data sources

---

## 🔧 Development Workflow

### Local Development

```bash
# Start development server
npm run dev

# Run in Turbopack mode (faster)
npm run dev --turbo

# Build for production testing
npm run build
npm start
```

### Code Quality

```bash
# Lint code
npm run lint

# Type checking
npx tsc --noEmit

# Performance testing
node scripts/performance-test.js
```

### Content Management

- **Add new content**: Place markdown files in appropriate directories
- **Update metadata**: Modify frontmatter in markdown files
- **Test content**: Verify through development server

---

## 📊 Performance Targets

- **First Contentful Paint**: < 1.5s
- **Largest Contentful Paint**: < 2.5s
- **Cumulative Layout Shift**: < 0.1
- **First Input Delay**: < 100ms
- **Bundle Size**: < 150KB initial load

## 🎯 Production Checklist

- [ ] Environment variables configured
- [ ] Domain and SSL configured
- [ ] Analytics tracking setup
- [ ] Error monitoring configured
- [ ] Performance monitoring in place
- [ ] Backup strategy implemented
- [ ] Content workflow established
- [ ] Security headers verified
- [ ] Accessibility tested
- [ ] Mobile responsiveness confirmed
- [ ] SEO optimization complete

## 🆘 Troubleshooting

### Common Issues

1. **Build Failures**: Check Node.js version and dependencies
2. **Performance Issues**: Run bundle analyzer and optimization checks
3. **Accessibility Problems**: Use automated testing tools
4. **SEO Issues**: Verify metadata and structured data

### Support

- **Documentation**: Check Next.js official docs
- **Community**: Agentard University forums
- **Issues**: GitHub repository issues section