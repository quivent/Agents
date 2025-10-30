import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        terminal: {
          bg: "var(--terminal-bg)",
          fg: "var(--terminal-fg)",
          amber: "var(--terminal-amber)",
          cyan: "var(--terminal-cyan)",
          magenta: "var(--terminal-magenta)",
          red: "var(--terminal-red)",
          blue: "var(--terminal-blue)",
          white: "var(--terminal-white)",
          gray: "var(--terminal-gray)",
          "dim-green": "var(--terminal-dim-green)",
        },
        bg: {
          primary: "var(--bg-primary)",
          secondary: "var(--bg-secondary)",
          tertiary: "var(--bg-tertiary)",
          hover: "var(--bg-hover)",
        },
        text: {
          primary: "var(--text-primary)",
          secondary: "var(--text-secondary)",
          accent: "var(--text-accent)",
          muted: "var(--text-muted)",
          warning: "var(--text-warning)",
        },
        border: {
          primary: "var(--border-primary)",
          secondary: "var(--border-secondary)",
        },
      },
      fontFamily: {
        mono: "var(--font-mono)",
      },
      animation: {
        blink: "blink 1s infinite",
      },
      keyframes: {
        blink: {
          "0%, 50%": { opacity: "1" },
          "51%, 100%": { opacity: "0" },
        },
      },
    },
  },
  plugins: [],
};

export default config;