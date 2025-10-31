import { notFound } from 'next/navigation';
import { getTardTheme, getAllTardIds } from '@/lib/themes';
import { ThemeProvider } from '@/components/theme/ThemeProvider';

interface TardLayoutProps {
  children: React.ReactNode;
  params: { tard: string };
}

export async function generateStaticParams() {
  return getAllTardIds().map((tard) => ({
    tard,
  }));
}

export default function TardLayout({ children, params }: TardLayoutProps) {
  const { tard } = params;
  
  if (!getAllTardIds().includes(tard)) {
    notFound();
  }

  const theme = getTardTheme(tard);

  return (
    <ThemeProvider theme={theme}>
      {children}
    </ThemeProvider>
  );
}
