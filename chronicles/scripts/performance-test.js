#!/usr/bin/env node

/**
 * Performance testing script for Agentard website
 * Tests loading times, bundle sizes, and accessibility
 */

const puppeteer = require('puppeteer');
const fs = require('fs');
const path = require('path');

const BASE_URL = process.env.BASE_URL || 'http://localhost:3000';
const PAGES_TO_TEST = [
  '/',
  '/interactive',
  '/chronicles',
  '/chronicles/chapter-01-the-awakening',
  '/university',
  '/university/eng-101-port-hardcoding-philosophy',
  '/manifesto',
  '/search',
  '/reading-paths'
];

async function runPerformanceTest() {
  console.log('🚀 Starting Agentard Website Performance Tests\n');
  
  const browser = await puppeteer.launch({
    headless: true,
    args: ['--no-sandbox', '--disable-setuid-sandbox']
  });

  const results = {
    timestamp: new Date().toISOString(),
    baseUrl: BASE_URL,
    summary: {
      totalPages: PAGES_TO_TEST.length,
      passedTests: 0,
      failedTests: 0,
      averageLoadTime: 0,
      averageFCP: 0,
      averageLCP: 0
    },
    pages: []
  };

  let totalLoadTime = 0;
  let totalFCP = 0;
  let totalLCP = 0;

  for (const pagePath of PAGES_TO_TEST) {
    const url = `${BASE_URL}${pagePath}`;
    console.log(`📋 Testing: ${url}`);
    
    try {
      const pageResult = await testPage(browser, url, pagePath);
      results.pages.push(pageResult);
      
      if (pageResult.passed) {
        results.summary.passedTests++;
      } else {
        results.summary.failedTests++;
      }
      
      totalLoadTime += pageResult.metrics.loadTime;
      totalFCP += pageResult.metrics.firstContentfulPaint;
      totalLCP += pageResult.metrics.largestContentfulPaint;
      
      console.log(`   ✅ Load Time: ${pageResult.metrics.loadTime}ms`);
      console.log(`   📊 FCP: ${pageResult.metrics.firstContentfulPaint}ms`);
      console.log(`   📈 LCP: ${pageResult.metrics.largestContentfulPaint}ms`);
      console.log(`   🎯 Score: ${pageResult.score}/100\n`);
      
    } catch (error) {
      console.error(`   ❌ Failed to test ${url}:`, error.message);
      results.pages.push({
        url,
        path: pagePath,
        passed: false,
        error: error.message,
        metrics: {},
        score: 0
      });
      results.summary.failedTests++;
    }
  }

  // Calculate averages
  const validPages = results.summary.passedTests;
  if (validPages > 0) {
    results.summary.averageLoadTime = Math.round(totalLoadTime / validPages);
    results.summary.averageFCP = Math.round(totalFCP / validPages);
    results.summary.averageLCP = Math.round(totalLCP / validPages);
  }

  await browser.close();

  // Save results
  const resultsPath = path.join(process.cwd(), 'performance-results.json');
  fs.writeFileSync(resultsPath, JSON.stringify(results, null, 2));

  // Print summary
  console.log('📊 PERFORMANCE TEST SUMMARY');
  console.log('================================');
  console.log(`✅ Passed: ${results.summary.passedTests}/${results.summary.totalPages}`);
  console.log(`❌ Failed: ${results.summary.failedTests}/${results.summary.totalPages}`);
  console.log(`⏱️  Average Load Time: ${results.summary.averageLoadTime}ms`);
  console.log(`🎨 Average FCP: ${results.summary.averageFCP}ms`);
  console.log(`📏 Average LCP: ${results.summary.averageLCP}ms`);
  console.log(`\n📄 Full results saved to: ${resultsPath}`);

  // Exit with error code if tests failed
  process.exit(results.summary.failedTests > 0 ? 1 : 0);
}

async function testPage(browser, url, pagePath) {
  const page = await browser.newPage();
  
  // Set viewport for consistent testing
  await page.setViewport({ width: 1200, height: 800 });
  
  // Enable performance monitoring
  await page.setCacheEnabled(false);
  
  const startTime = Date.now();
  
  try {
    // Navigate to page
    await page.goto(url, { 
      waitUntil: 'networkidle0',
      timeout: 30000 
    });
    
    const loadTime = Date.now() - startTime;
    
    // Get performance metrics
    const metrics = await page.evaluate(() => {
      const paintEntries = performance.getEntriesByType('paint');
      const fcp = paintEntries.find(entry => entry.name === 'first-contentful-paint');
      
      return {
        firstContentfulPaint: fcp ? Math.round(fcp.startTime) : 0,
        largestContentfulPaint: 0, // LCP requires more complex measurement
        domContentLoaded: Math.round(performance.timing.domContentLoadedEventEnd - performance.timing.navigationStart),
        loadComplete: Math.round(performance.timing.loadEventEnd - performance.timing.navigationStart)
      };
    });
    
    // Basic accessibility checks
    const accessibilityIssues = await checkAccessibility(page);
    
    // Check for terminal elements (Agentard-specific)
    const hasTerminalElements = await page.$('.terminal-window, .terminal-glow') !== null;
    
    // Check for interactive elements
    const hasInteractiveElements = await page.$('button, input, [role="button"]') !== null;
    
    // Calculate score
    let score = 100;
    if (loadTime > 3000) score -= 20;
    if (metrics.firstContentfulPaint > 2000) score -= 15;
    if (accessibilityIssues.length > 0) score -= (accessibilityIssues.length * 5);
    if (!hasTerminalElements) score -= 10;
    if (!hasInteractiveElements && pagePath === '/interactive') score -= 20;
    
    await page.close();
    
    return {
      url,
      path: pagePath,
      passed: score >= 70,
      metrics: {
        loadTime,
        firstContentfulPaint: metrics.firstContentfulPaint,
        largestContentfulPaint: metrics.largestContentfulPaint,
        domContentLoaded: metrics.domContentLoaded,
        loadComplete: metrics.loadComplete
      },
      accessibility: {
        issues: accessibilityIssues,
        hasTerminalElements,
        hasInteractiveElements
      },
      score: Math.max(0, score)
    };
    
  } catch (error) {
    await page.close();
    throw error;
  }
}

async function checkAccessibility(page) {
  const issues = await page.evaluate(() => {
    const issues = [];
    
    // Check for missing alt text on images
    const images = document.querySelectorAll('img:not([alt])');
    if (images.length > 0) {
      issues.push(`Missing alt text on ${images.length} images`);
    }
    
    // Check for missing form labels
    const inputs = document.querySelectorAll('input:not([aria-label]):not([aria-labelledby])');
    const unlabeledInputs = Array.from(inputs).filter(input => {
      const labels = document.querySelectorAll(`label[for="${input.id}"]`);
      return labels.length === 0;
    });
    if (unlabeledInputs.length > 0) {
      issues.push(`${unlabeledInputs.length} form inputs without labels`);
    }
    
    // Check for missing headings structure
    const headings = document.querySelectorAll('h1, h2, h3, h4, h5, h6');
    if (headings.length === 0) {
      issues.push('No heading elements found');
    }
    
    // Check for skip links
    const skipLinks = document.querySelectorAll('a[href="#main-content"], .skip-nav');
    if (skipLinks.length === 0) {
      issues.push('No skip navigation links found');
    }
    
    return issues;
  });
  
  return issues;
}

// Run the test if called directly
if (require.main === module) {
  runPerformanceTest().catch(console.error);
}

module.exports = { runPerformanceTest, testPage, checkAccessibility };