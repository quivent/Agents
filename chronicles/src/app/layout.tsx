import type { Metadata, Viewport } from "next";
import { Geist_Mono } from "next/font/google";
import "./globals.css";
import { defaultMetadata } from '@/lib/metadata';
import { AccessibilityProvider } from '@/components/accessibility/AccessibilityProvider';
import { ErrorBoundary } from '@/components/errors/ErrorBoundary';
import { WebsiteStructuredData } from '@/components/seo/StructuredData';

const geistMono = Geist_Mono({
  variable: "--font-mono",
  subsets: ["latin"],
  display: 'swap',
  preload: true,
});

export const metadata: Metadata = defaultMetadata;

export const viewport: Viewport = {
  width: 'device-width',
  initialScale: 1,
  maximumScale: 5,
  userScalable: true,
  themeColor: [
    { media: '(prefers-color-scheme: light)', color: '#0a0a0a' },
    { media: '(prefers-color-scheme: dark)', color: '#0a0a0a' },
  ],
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="dark">
      <head>
        <meta name="color-scheme" content="dark" />
      </head>
      <body className={`${geistMono.variable} antialiased min-h-screen scanlines`}>
        <a href="#main-content" className="skip-nav">
          Skip to main content
        </a>
        
        <ErrorBoundary>
          <AccessibilityProvider>
            <div id="main-content">
              {children}
            </div>
          </AccessibilityProvider>
        </ErrorBoundary>

        <WebsiteStructuredData
          name="Agentard - The Art of Confident Incompetence"
          description="Experience the philosophical journey of the Agentard - where confidence meets incompetence in perfect harmony."
          url="https://agentard.com"
        />
      </body>
    </html>
  );
}
