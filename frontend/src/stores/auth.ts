import { defineStore } from 'pinia';
import api from '../api/axios';

export interface User {
  id: number;
  tenant_id: number;
  username: string;
  role: string;
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || null,
    user: null as User | null,
    loading: false,
    error: null as string | null,
  }),
  getters: {
    isAuthenticated: (state) => !!state.token,
  },
  actions: {
    async login(credentials: any) {
      this.loading = true;
      this.error = null;
      try {
        const response = await api.post('/auth/login', credentials);
        this.token = response.data.token;
        localStorage.setItem('token', this.token as string);
        await this.fetchUser();
      } catch (error: any) {
        this.error = error.response?.data?.error || 'Login failed';
        throw error;
      } finally {
        this.loading = false;
      }
    },
    async register(data: any) {
      this.loading = true;
      this.error = null;
      try {
        const response = await api.post('/auth/register', data);
        this.token = response.data.token;
        localStorage.setItem('token', this.token as string);
        await this.fetchUser();
      } catch (error: any) {
        this.error = error.response?.data?.error || 'Registration failed';
        throw error;
      } finally {
        this.loading = false;
      }
    },
    async fetchUser() {
      if (!this.token) return;
      try {
        const response = await api.get('/api/me');
        this.user = response.data;
      } catch (error) {
        this.logout();
      }
    },
    logout() {
      this.token = null;
      this.user = null;
      localStorage.removeItem('token');
    },
  },
});
