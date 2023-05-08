import vue from "@vitejs/plugin-vue";
import path from "path";
import { createHtmlPlugin } from "vite-plugin-html";
import { loadEnv } from "vite";
import mkcert from "vite-plugin-mkcert";
// https://vitejs.dev/config/

export default ({ mode }) => {
  const env = loadEnv(mode, process.cwd());
  return {
    plugins: [
      vue(),
      createHtmlPlugin({
        minify: true,
        inject: {
          data: {
            title: env.VITE_APP_NAME,
            icon: env.VITE_FAV_ICON,
          },
        },
      }),
      // mkcert(),
    ],
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "./src"),
      },
    },
    css: {
      preprocessorOptions: {
        less: {
          modifyVars: {},
          javascriptEnabled: true,
        },
        scss: {
          additionalData: `@import "src/styles/index.scss";`,
        },
      },
    },
    server: {
      host: "0.0.0.0",
      port: 3000,
      // https: true,
      open: true,
      // proxy: {
      //   "/api": {
      //     target: "https://lotteryadminapi.66bbs.live",
      //     changeOrigin: true,
      //   },
      //   "/v2": {
      //     target: "https://lotteryresource.66bbs.live",
      //     changeOrigin: true,
      //   },
      // },
    },
    define: {
      __VUE_I18N_FULL_INSTALL__: true,
      __VUE_I18N_LEGACY_API__: false,
      __INTLIFY_PROD_DEVTOOLS__: false,
    },
  };
};
