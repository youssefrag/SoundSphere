// vite.config.js
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import path from "path";

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"), // ← make @ point at src/
    },
  },
  server: {
    // if you still need the host‐whitelist:
    allowedHosts: ["soundsphere-z0qd.onrender.com"],
  },
  build: {
    target: "es2022",
  },
  esbuild: {
    target: "es2022",
  },
});
