#!/usr/bin/env node

/**
 * Script to enhance content files with proper metadata
 * Adds reading time, summaries, and other metadata to markdown files
 */

const fs = require('fs');
const path = require('path');
const matter = require('gray-matter');

const contentDir = path.join(__dirname, '../src/data/theagentard');

function estimateReadingTime(content) {
  const wordsPerMinute = 200;
  const wordCount = content.split(/\s+/).length;
  return Math.ceil(wordCount / wordsPerMinute);
}

function generateSummary(content, maxLength = 150) {
  // Extract first meaningful paragraph, skipping headers and metadata
  const lines = content.split('\n').filter(line => 
    line.trim() && 
    !line.startsWith('#') && 
    !line.startsWith('**') &&
    !line.startsWith('---') &&
    !line.startsWith('*') &&
    line.length > 50
  );
  
  const summary = lines[0] || '';
  return summary.length > maxLength 
    ? summary.substring(0, maxLength) + '...'
    : summary;
}

function suggestTags(content, filename) {
  const text = content.toLowerCase();
  const tags = [];
  
  // Technical terms
  if (text.includes('port') || text.includes('8080') || text.includes('8443')) tags.push('engineering');
  if (text.includes('authentication') || text.includes('ssl') || text.includes('security')) tags.push('security');
  if (text.includes('dns') || text.includes('domain')) tags.push('networking');
  if (text.includes('consciousness') || text.includes('awareness')) tags.push('philosophy');
  if (text.includes('confidence') || text.includes('competence')) tags.push('psychology');
  if (text.includes('cascade') || text.includes('failure')) tags.push('failure-analysis');
  if (text.includes('documentation') || text.includes('manual')) tags.push('documentation');
  if (text.includes('automation') || text.includes('script')) tags.push('automation');
  if (text.includes('infinite') || text.includes('loop')) tags.push('recursion');
  if (text.includes('meta') || text.includes('self-aware')) tags.push('meta-analysis');
  
  // Content type specific
  if (filename.includes('chapter')) tags.push('narrative', 'journey');
  if (filename.includes('eng_') || filename.includes('app_')) tags.push('technical', 'course');
  if (filename.includes('phil_') || filename.includes('math_')) tags.push('theoretical', 'advanced');
  if (filename.includes('manifesto') || filename.includes('foundation')) tags.push('core-principles');
  
  // Difficulty indicators
  if (text.includes('beginner') || text.includes('introduction')) tags.push('beginner-friendly');
  if (text.includes('advanced') || text.includes('transcendent')) tags.push('advanced-concepts');
  
  return [...new Set(tags)]; // Remove duplicates
}

function processFile(filePath) {
  try {
    const content = fs.readFileSync(filePath, 'utf8');
    const { data, content: markdownContent } = matter(content);
    
    const filename = path.basename(filePath, '.md');
    const readingTime = estimateReadingTime(markdownContent);
    const summary = generateSummary(markdownContent);
    const suggestedTags = suggestTags(markdownContent, filename);
    
    // Enhanced metadata
    const enhancedData = {
      ...data,
      readingTime,
      summary,
      wordCount: markdownContent.split(/\s+/).length,
      lastUpdated: new Date().toISOString().split('T')[0],
      tags: data.tags || suggestedTags,
    };
    
    // Add specific metadata based on content type
    if (filePath.includes('chronicles')) {
      const chapterMatch = filename.match(/chapter_(\d+)/);
      if (chapterMatch) {
        enhancedData.chapter = parseInt(chapterMatch[1]);
        enhancedData.contentType = 'core';
      }
    }
    
    if (filePath.includes('university')) {
      const codeMatch = filename.match(/([A-Z]+_\d+)/);
      if (codeMatch) {
        enhancedData.course = codeMatch[1];
        const courseNumber = parseInt(codeMatch[1].split('_')[1] || '101');
        if (courseNumber >= 401) enhancedData.difficulty = 'transcendent';
        else if (courseNumber >= 301) enhancedData.difficulty = 'advanced';
        else if (courseNumber >= 201) enhancedData.difficulty = 'intermediate';
        else enhancedData.difficulty = 'beginner';
        enhancedData.contentType = enhancedData.difficulty === 'transcendent' ? 'advanced' : 'core';
      }
    }
    
    if (filePath.includes('manifesto')) {
      enhancedData.contentType = 'core';
      enhancedData.difficulty = 'intermediate';
    }
    
    if (filePath.includes('core')) {
      enhancedData.contentType = 'core';
      enhancedData.difficulty = 'beginner';
    }
    
    // Generate new content with enhanced frontmatter
    const newContent = matter.stringify(markdownContent, enhancedData);
    
    console.log(`Enhanced: ${filename}`);
    console.log(`  Reading time: ${readingTime} minutes`);
    console.log(`  Tags: ${enhancedData.tags.join(', ')}`);
    console.log(`  Summary: ${summary.substring(0, 50)}...`);
    console.log('');
    
    // Uncomment the line below to actually write the changes
    // fs.writeFileSync(filePath, newContent);
    
    return {
      file: filename,
      metadata: enhancedData,
      success: true
    };
    
  } catch (error) {
    console.error(`Error processing ${filePath}:`, error.message);
    return {
      file: path.basename(filePath),
      error: error.message,
      success: false
    };
  }
}

function processDirectory(dirPath) {
  const results = [];
  
  if (!fs.existsSync(dirPath)) {
    console.log(`Directory not found: ${dirPath}`);
    return results;
  }
  
  const items = fs.readdirSync(dirPath);
  
  for (const item of items) {
    const itemPath = path.join(dirPath, item);
    const stat = fs.statSync(itemPath);
    
    if (stat.isDirectory()) {
      results.push(...processDirectory(itemPath));
    } else if (item.endsWith('.md')) {
      results.push(processFile(itemPath));
    }
  }
  
  return results;
}

// Main execution
console.log('=== Agentard Content Metadata Enhancement ===\n');
console.log('Analyzing content files...\n');

const results = processDirectory(contentDir);
const successful = results.filter(r => r.success);
const failed = results.filter(r => !r.success);

console.log('=== SUMMARY ===');
console.log(`Processed: ${results.length} files`);
console.log(`Successful: ${successful.length} files`);
console.log(`Failed: ${failed.length} files`);

if (failed.length > 0) {
  console.log('\nFailed files:');
  failed.forEach(f => console.log(`  - ${f.file}: ${f.error}`));
}

console.log('\n=== RECOMMENDATIONS ===');
console.log('1. Review suggested tags and metadata');
console.log('2. Uncomment the fs.writeFileSync line to apply changes');
console.log('3. Test the enhanced content in the website');
console.log('4. Consider adding more specific tags for better discoverability');

// Generate a summary report
const reportPath = path.join(__dirname, '../content-enhancement-report.json');
fs.writeFileSync(reportPath, JSON.stringify({
  timestamp: new Date().toISOString(),
  summary: {
    total: results.length,
    successful: successful.length,
    failed: failed.length
  },
  files: successful.map(r => ({
    file: r.file,
    readingTime: r.metadata.readingTime,
    wordCount: r.metadata.wordCount,
    tags: r.metadata.tags,
    difficulty: r.metadata.difficulty,
    contentType: r.metadata.contentType
  }))
}, null, 2));

console.log(`\nDetailed report saved to: ${reportPath}`);