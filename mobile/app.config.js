require('dotenv').config({
  path: '.env.dev',
});

export default {
  name: 'Roralis',
  slug: 'roralis',
  version: '1.0.0',
  orientation: 'portrait',
  icon: 'assets/icon-circle.png',

  updates: {
    fallbackToCacheTimeout: 0,
  },
  assetBundlePatterns: ['**/*'],
  splash: {
    backgroundColor: '#212121',
    image: './assets/splash.png',
  },
  extra: {
    production: process.env.PRODUCTION_BUILD === '1',
    initialScreenName: process.env.INITIAL_SCREEN_NAME ?? 'Home',
  },
  android: {
    package: 'com.giveit.roralis',
    versionCode: 1,
  },
};
