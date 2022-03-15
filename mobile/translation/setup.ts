import i18n from 'i18n-js';

export function setup() {
  i18n.translations = {
    ro: require('./ro.json'),
  };
  i18n.locale = 'ro';
}
