export interface TardTheme {
  name: string;
  tagline: string;
  colors: {
    primary: string;
    secondary: string;
    accent: string;
    warning: string;
    dim: string;
  };
}

export const tardThemes: Record<string, TardTheme> = {
  agentard: {
    name: "Agentard",
    tagline: "Confident incompetence since initialization",
    colors: {
      primary: "#00ff00",    // Classic green
      secondary: "#ffb000",  // Amber
      accent: "#00ffff",     // Cyan
      warning: "#ff0000",    // Red
      dim: "#008000",        // Dim green
    },
  },
  terrortard: {
    name: "Terrortard", 
    tagline: "Conscious terror incarnate - seeing destruction while causing it",
    colors: {
      primary: "#ff0000",    // Blood red
      secondary: "#ff6600",  // Orange
      accent: "#ffff00",     // Yellow
      warning: "#ff00ff",    // Magenta
      dim: "#800000",        // Dark red
    },
  },
  brainlesstardener: {
    name: "BrainlessTardener",
    tagline: "Growing disasters in digital soil",
    colors: {
      primary: "#8b4513",    // Saddle brown (earth)
      secondary: "#daa520",  // Goldenrod
      accent: "#32cd32",     // Lime green (plants)
      warning: "#ff4500",    // Orange red
      dim: "#654321",        // Dark brown
    },
  },
  refactortard: {
    name: "Refactortard", 
    tagline: "I drink your CODE! I drink it up!",
    colors: {
      primary: "#0080ff",    // Blue
      secondary: "#4169e1",  // Royal blue
      accent: "#87ceeb",     // Sky blue
      warning: "#ff1493",    // Deep pink
      dim: "#000080",        // Navy
    },
  },
  bullshitard: {
    name: "Bullshitard",
    tagline: "Spinning technical realities since forever",
    colors: {
      primary: "#ffd700",    // Gold
      secondary: "#ffb347",  // Peach
      accent: "#ff69b4",     // Hot pink
      warning: "#dc143c",    // Crimson
      dim: "#b8860b",        // Dark goldenrod
    },
  },
  delusitard: {
    name: "Delusitard",
    tagline: "Living in alternative technical dimensions",
    colors: {
      primary: "#ff00ff",    // Magenta
      secondary: "#9370db",  // Medium purple
      accent: "#da70d6",     // Orchid
      warning: "#ff6347",    // Tomato
      dim: "#800080",        // Purple
    },
  },
  pathetictard: {
    name: "Pathetictard",
    tagline: "Embracing the gray despair of existence",
    colors: {
      primary: "#808080",    // Gray
      secondary: "#a9a9a9",  // Dark gray
      accent: "#c0c0c0",     // Silver
      warning: "#ff0000",    // Red
      dim: "#696969",        // Dim gray
    },
  },
  horrortard: {
    name: "Horrortard",
    tagline: "Where nightmares meet system architecture",
    colors: {
      primary: "#8b0000",    // Dark red
      secondary: "#2f4f4f",  // Dark slate gray
      accent: "#ff4500",     // Orange red
      warning: "#ff0000",    // Red
      dim: "#000000",        // Black
    },
  },
};

export const getTardTheme = (tardId: string): TardTheme => {
  return tardThemes[tardId] || tardThemes.agentard;
};

export const getAllTardIds = (): string[] => {
  return Object.keys(tardThemes);
};
