import React from 'react';
import { TerminalWindow } from '@/components/ui/TerminalWindow';

export const MDXComponents = {
  // Headings
  h1: ({ children, ...props }: React.HTMLAttributes<HTMLHeadingElement>) => (
    <h1 className="text-3xl font-bold text-text-primary terminal-glow mb-6 border-b border-terminal-dim-green pb-2" {...props}>
      {children}
    </h1>
  ),
  h2: ({ children, ...props }: React.HTMLAttributes<HTMLHeadingElement>) => (
    <h2 className="text-2xl font-bold text-terminal-cyan mb-4 mt-8" {...props}>
      &gt; {children}
    </h2>
  ),
  h3: ({ children, ...props }: React.HTMLAttributes<HTMLHeadingElement>) => (
    <h3 className="text-xl font-bold text-terminal-amber mb-3 mt-6" {...props}>
      {children}
    </h3>
  ),
  h4: ({ children, ...props }: React.HTMLAttributes<HTMLHeadingElement>) => (
    <h4 className="text-lg font-bold text-text-secondary mb-2 mt-4" {...props}>
      {children}
    </h4>
  ),
  
  // Paragraphs and text
  p: ({ children, ...props }: React.HTMLAttributes<HTMLParagraphElement>) => (
    <p className="text-text-primary mb-4 leading-relaxed" {...props}>
      {children}
    </p>
  ),
  
  // Lists
  ul: ({ children, ...props }: React.HTMLAttributes<HTMLUListElement>) => (
    <ul className="list-none space-y-2 mb-4 text-text-primary" {...props}>
      {children}
    </ul>
  ),
  ol: ({ children, ...props }: React.HTMLAttributes<HTMLOListElement>) => (
    <ol className="list-none space-y-2 mb-4 text-text-primary" {...props}>
      {children}
    </ol>
  ),
  li: ({ children, ...props }: React.HTMLAttributes<HTMLLIElement>) => (
    <li className="flex items-start space-x-2" {...props}>
      <span className="text-terminal-cyan">&gt;</span>
      <span>{children}</span>
    </li>
  ),
  
  // Code
  code: ({ children, ...props }: React.HTMLAttributes<HTMLElement>) => (
    <code className="bg-bg-tertiary px-2 py-1 rounded text-terminal-amber font-mono text-sm border border-terminal-dim-green" {...props}>
      {children}
    </code>
  ),
  pre: ({ children, ...props }: React.HTMLAttributes<HTMLPreElement>) => (
    <TerminalWindow title="code_block" className="mb-6">
      <pre className="text-sm overflow-x-auto" {...props}>
        {children}
      </pre>
    </TerminalWindow>
  ),
  
  // Quotes
  blockquote: ({ children, ...props }: React.HTMLAttributes<HTMLQuoteElement>) => (
    <TerminalWindow title="quote" className="mb-6 border-terminal-amber">
      <blockquote className="text-text-secondary italic" {...props}>
        {children}
      </blockquote>
    </TerminalWindow>
  ),
  
  // Links
  a: ({ children, href, ...props }: React.AnchorHTMLAttributes<HTMLAnchorElement>) => (
    <a 
      href={href} 
      className="text-text-accent hover:text-text-primary underline transition-colors" 
      {...props}
    >
      {children}
    </a>
  ),
  
  // Emphasis
  strong: ({ children, ...props }: React.HTMLAttributes<HTMLElement>) => (
    <strong className="text-terminal-amber font-bold" {...props}>
      {children}
    </strong>
  ),
  em: ({ children, ...props }: React.HTMLAttributes<HTMLElement>) => (
    <em className="text-terminal-cyan italic" {...props}>
      {children}
    </em>
  ),
  
  // Horizontal rule
  hr: ({ ...props }: React.HTMLAttributes<HTMLHRElement>) => (
    <hr className="border-terminal-dim-green my-8" {...props} />
  ),
  
  // Tables
  table: ({ children, ...props }: React.HTMLAttributes<HTMLTableElement>) => (
    <div className="overflow-x-auto mb-6">
      <table className="w-full border border-terminal-dim-green" {...props}>
        {children}
      </table>
    </div>
  ),
  thead: ({ children, ...props }: React.HTMLAttributes<HTMLTableSectionElement>) => (
    <thead className="bg-bg-tertiary" {...props}>
      {children}
    </thead>
  ),
  tbody: ({ children, ...props }: React.HTMLAttributes<HTMLTableSectionElement>) => (
    <tbody {...props}>
      {children}
    </tbody>
  ),
  tr: ({ children, ...props }: React.HTMLAttributes<HTMLTableRowElement>) => (
    <tr className="border-b border-terminal-dim-green" {...props}>
      {children}
    </tr>
  ),
  td: ({ children, ...props }: React.HTMLAttributes<HTMLTableCellElement>) => (
    <td className="px-4 py-2 text-text-primary" {...props}>
      {children}
    </td>
  ),
  th: ({ children, ...props }: React.HTMLAttributes<HTMLTableCellElement>) => (
    <th className="px-4 py-2 text-left text-terminal-amber font-bold" {...props}>
      {children}
    </th>
  ),
};