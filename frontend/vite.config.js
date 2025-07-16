// vite.config.js
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

export default defineConfig({
  plugins: [vue()],
  server: {
    // keep your allowedHosts here if you need them
    allowedHosts: ["soundsphere-z0qd.onrender.com"],
  },
  build: {
    // target a version that supports top-level await
    target: "es2022",
    // or simply: target: "esnext"
  },
  // if you want to be extra sure ESBuild sees it:
  esbuild: {
    target: "es2022",
  },
});
