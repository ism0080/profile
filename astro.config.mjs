// @ts-check
import { defineConfig } from 'astro/config';

// https://astro.build/config
export default defineConfig({
    site: 'https://isaacmackle.com',
    base: '/blog/',
    trailingSlash: 'always',
    output: 'static',
    publicDir: 'astroPublic',
    outDir: 'public/blog'
});
