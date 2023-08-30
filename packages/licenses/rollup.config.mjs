import json from "@rollup/plugin-json"
import ts from "rollup-plugin-ts";
/**
 * @type {import('rollup').RollupOptions}
 */
const config = {
	input: './src/index.ts',
	output: {
    dir: "./build",
		format: 'cjs'
	},
  plugins: [ts({outDir: "./build"}), json()]
};

export default config;