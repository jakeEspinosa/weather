import react, { reactCompilerPreset } from "@vitejs/plugin-react";
import babel from "@rolldown/plugin-babel";
import { defineConfig } from "vite";
import tailwindcss from "@tailwindcss/vite";
import path from "node:path";

const root = import.meta.dirname;
const srcDir = path.join(root, "src");

export default defineConfig({
  plugins: [react(), babel({ presets: [reactCompilerPreset()] }), tailwindcss()],
  test: {
    resolveSnapshotPath: (testPath: string, snapExtension: string) =>
      path.join(root, ".vitest", path.relative(srcDir, testPath) + snapExtension),
    environment: "happy-dom",
  },
});
