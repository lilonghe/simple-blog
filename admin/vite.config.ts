import { defineConfig, loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve } from "path";

export default defineConfig(({ command, mode }) => {
  const env = loadEnv(mode, process.cwd(), '')

  return {
    plugins: [vue()],
    resolve: {
      alias: [
        {
          find: "@",
          replacement: resolve(__dirname, "src"),
        },
      ],
    },
    base: env.VITE_BASE,
    server: {
      port: 5563,
      proxy: {
        "/admin": {
          target: env.VITE_BASE_API_URL + '/admin',
          changeOrigin: true,
        },
      },
    },
  };
});
