import rss, { pagesGlobToRssItems } from '@astrojs/rss';
import { AppConfig } from '~/utils/app-config';

export async function GET(context) {
    return rss({
        title: AppConfig.title,
        description: AppConfig.description,
        site: `${context.site}${AppConfig.routes.blog}`,
        items: await pagesGlobToRssItems(
            import.meta.glob('./posts/*.{md,mdx}')
        ),
        stylesheet: './rss/styles.xsl',
        customData: `<language>en-NZ</language>`
    });
}
