import tailwindcss from "@tailwindcss/vite";
import react from "@vitejs/plugin-react";
import { defineConfig, loadEnv } from "vite";
import checker from "vite-plugin-checker";
import eslint from "vite-plugin-eslint";

import pkg from "./package.json";

export default defineConfig(({ mode }) => {
  // Load env file based on `mode`
  const env = loadEnv(mode, process.cwd());

  return {
    base: env.VITE_BASE_URL || "/",
    plugins: [
      react(),
      tailwindcss(),
      eslint({
        include: ["src/**/*.js", "src/**/*.jsx"],
        emitWarning: true,
        emitError: true,
        failOnError: false,
        failOnWarning: false,
        cache: false, // disables cache so it always re-checks on file change
      }),
      checker({
        eslint: {
          // This is the lint command that runs on your source files
          lintCommand: 'eslint "./src/**/*.{js,jsx,ts,tsx}"',
        },
      }),
    ],
    define: {
      __APP_NAME__: `"${pkg.name}"`,
      __APP_VERSION__: `"${pkg.version}"`,
      __HOMEPAGE__: `"${pkg.homepage}"`,
      __API_ROOT__: `"${env.VITE_API_ROOT || "http://localhost:8080/api"}"`,
    },
    server: {
      port: 5173,
      host: "0.0.0.0",
    },
  };
});
