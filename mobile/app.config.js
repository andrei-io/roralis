require('dotenv').config({
  path: '.env.dev',
});

export default {
  name: 'Roralis',
  slug: 'roralis',
  version: '1.0.0',
  orientation: 'portrait',

  updates: {
    fallbackToCacheTimeout: 0,
  },
  assetBundlePatterns: ['**/*'],
  extra: {
    production: process.env.PRODUCTION_BUILD === '1',
  },
};
