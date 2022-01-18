
import vue from '@vitejs/plugin-vue'
import {defineConfig} from "rollup";
export default defineConfig({
  resolve: {
    alias: {
      '@': '/src'
    }
  },
  plugins: [vue()],
  server: {
    host:"0.0.0.0",
    port:2002,
    proxy: {
   
    }
  },

})
