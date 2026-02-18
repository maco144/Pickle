import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  // Set base to '/' for root deployment, or '/Pickle/' for GitHub Pages
  base: '/',
  build: {
    outDir: 'dist',
    sourcemap: false,
    rollupOptions: {
      output: {
        // Split three.js and chart.js into separate chunks — they're large
        manualChunks: {
          three: ['three'],
          chartjs: ['chart.js'],
        },
      },
    },
  },
})
