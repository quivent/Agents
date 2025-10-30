const fs = require('fs');
const path = require('path');
const matter = require('gray-matter');

const contentDirectory = path.join(process.cwd(), 'src/data/theagentard');
const outputDir = path.join(process.cwd(), 'public');

function getMarkdownFiles(dir) {
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

function parseContentFile(filePath) {
  try {
    const fullPath = path.join(contentDirectory, filePath);
    const fileContent = fs.readFileSync(fullPath, 'utf8');
    const { data, content } = matter(fileContent);

    const slug = path.basename(filePath, '.md').toLowerCase().replace(/[^a-z0-9]+/g, '-');

    return {
      metadata: {
        title: data.title || path.basename(filePath, '.md').replace(/_/g, ' '),
        description: data.description || '',
        author: data.author || '',
        date: data.date || '',
        tags: data.tags || [],
        chapter: data.chapter,
        course: data.course,
        difficulty: data.difficulty,
      },
      content: content,
      slug,
      path: filePath,
    };
  } catch (error) {
    console.error(`Error parsing file ${filePath}:`, error);
    return null;
  }
}

function getAllContent() {
  const allContent = [];

  const directories = [
    'agentard_chronicles',
    'agentard_university',
    'manifesto',
    'core'
  ];

  directories.forEach(dir => {
    const files = getMarkdownFiles(dir);
    files.forEach(file => {
      const content = parseContentFile(file);
      if (content) {
        allContent.push(content);
      }
    });
  });

  return allContent;
}

// Generate the content JSON
console.log('Generating static content JSON...');
const allContent = getAllContent();
const outputPath = path.join(outputDir, 'content-data.json');

// Ensure output directory exists
if (!fs.existsSync(outputDir)) {
  fs.mkdirSync(outputDir, { recursive: true });
}

fs.writeFileSync(outputPath, JSON.stringify(allContent, null, 2));
console.log(`✅ Generated ${allContent.length} content documents to ${outputPath}`);
