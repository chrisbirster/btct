import { defineConfig } from "vite";
import solid from "vite-plugin-solid";
import wyw from "@wyw-in-js/vite";

export default defineConfig({
  plugins: [
    wyw({
      include: ["**/*.{ts,tsx}"],
      babelOptions: {
        presets: ["@babel/preset-typescript"],
      },
    }),
    solid(),
  ],
  build: {
    outDir: "../dist", // Output build files to the root-level /dist folder
    emptyOutDir: true,
  },
  root: ".",
});
