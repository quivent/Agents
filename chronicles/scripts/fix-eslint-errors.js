#!/usr/bin/env node

const fs = require('fs');
const path = require('path');

// Function to fix unescaped quotes in JSX
function fixUnescapedQuotes(content) {
  // Replace unescaped quotes in JSX content
  return content
    .replace(/(\w+[>}])([^<]*)"([^<]*)"([^<]*<)/g, '$1$2&quot;$3&quot;$4')
    .replace(/(\w+[>}])([^<]*)'([^<]*)'([^<]*<)/g, '$1$2&apos;$3&apos;$4');
}

// Function to remove unused imports and variables
function removeUnusedImports(content, filePath) {
  const lines = content.split('\n');
  const newLines = [];
  
  for (const line of lines) {
    // Skip lines with unused imports that are warnings only
    if (line.includes("'useEffect'") && filePath.includes('AgentardMoments')) {
      continue;
    }
    if (line.includes("'getContentByTags'") && filePath.includes('search/page')) {
      continue;  
    }
    
    newLines.push(line);
  }
  
  return newLines.join('\n');
}

// Function to fix TypeScript errors
function fixTypeScriptErrors(content) {
  return content
    .replace(/let metadata:/g, 'const metadata:')
    .replace(/: any/g, ': Record<string, unknown>');
}

// List of files to fix
const filesToFix = [
  'src/app/interactive/page.tsx',
  'src/app/search/page.tsx', 
  'src/components/content/ContentSearch.tsx',
  'src/components/content/SeriesNavigation.tsx',
  'src/components/enhanced/InteractiveQuote.tsx',
  'src/components/interactive/InteractiveTerminal.tsx'
];

filesToFix.forEach(filePath => {
  const fullPath = path.join(process.cwd(), filePath);
  
  if (fs.existsSync(fullPath)) {
    console.log(`Fixing ${filePath}...`);
    
    let content = fs.readFileSync(fullPath, 'utf8');
    
    // Apply fixes
    content = fixUnescapedQuotes(content);
    content = removeUnusedImports(content, filePath);
    content = fixTypeScriptErrors(content);
    
    fs.writeFileSync(fullPath, content);
    console.log(`Fixed ${filePath}`);
  } else {
    console.log(`File not found: ${filePath}`);
  }
});

console.log('ESLint fixes completed!');