import { createI18n } from 'vue-i18n';
import en from './locales/en';
import id from './locales/id';
import ja from './locales/ja';

const i18n = createI18n({
  legacy: false,
  locale: localStorage.getItem('language') || 'en',
  fallbackLocale: 'en',
  messages: {
    en,
    id,
    ja
  }
});

export default i18n;
