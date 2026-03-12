import { defineStore } from 'pinia';
import api from '../api/axios';
import i18n from '../i18n';

export const useConfigStore = defineStore('config', {
  state: () => ({
    language: localStorage.getItem('language') || 'en',
    currency: 'USD',
    tenantName: '',
    loading: false
  }),
  actions: {
    async fetchTenantSettings() {
      this.loading = true;
      try {
        const response = await api.get('/api/tenant');
        this.currency = response.data.currency;
        this.tenantName = response.data.name;
      } catch (error) {
        console.error('Failed to fetch tenant settings', error);
      } finally {
        this.loading = false;
      }
    },
    async updateTenantSettings(name: string, currency: string) {
      this.loading = true;
      try {
        await api.patch('/api/tenant', { name, currency });
        this.tenantName = name;
        this.currency = currency;
      } catch (error) {
        console.error('Failed to update tenant settings', error);
        throw error;
      } finally {
        this.loading = false;
      }
    },
    setLanguage(lang: string) {
      this.language = lang;
      localStorage.setItem('language', lang);
      (i18n.global.locale as any).value = lang;
    },
    formatCurrency(amount: number) {
      return new Intl.NumberFormat(this.language === 'id' ? 'id-ID' : this.language === 'ja' ? 'ja-JP' : 'en-US', {
        style: 'currency',
        currency: this.currency || 'USD',
      }).format(amount);
    }
  }
});
