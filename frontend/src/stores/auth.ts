import { defineStore } from 'pinia';
import api from '../api/axios';

export interface Role {
  id: number;
  name: string;
  permissions: string[];
}

export interface User {
  id: number;
  tenant_id: number;
  branch_id?: number;
  username: string;
  role_id?: number;
  role?: Role;
  is_owner: boolean;
}

export interface Branch {
  id: number;
  tenant_id: number;
  name: string;
  address?: string;
  phone?: string;
  is_active: boolean;
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || null,
    user: null as User | null,
    selectedBranchID: localStorage.getItem('selectedBranchID') ? Number(localStorage.getItem('selectedBranchID')) : null as number | null,
    branches: [] as Branch[],
    loading: false,
    error: null as string | null,
  }),
  getters: {
    isAuthenticated: (state) => !!state.token,
    currentBranch: (state) => state.branches.find(b => b.id === state.selectedBranchID),
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
        await this.fetchBranches();
        
        // Auto-select first branch if none selected or current not in list
        if (this.branches.length > 0 && (!this.selectedBranchID || !this.branches.find(b => b.id === this.selectedBranchID))) {
          const firstBranch = this.branches[0];
          if (firstBranch) {
            this.setBranch(firstBranch.id);
          }
        }
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
    async fetchBranches() {
      if (!this.token) return;
      try {
        const response = await api.get('/api/branches');
        this.branches = response.data;
        
        // Ensure selected branch is valid
        if (this.selectedBranchID && !this.branches.find(b => b.id === this.selectedBranchID)) {
          this.selectedBranchID = this.branches.length > 0 ? (this.branches[0]?.id || null) : null;
          if (this.selectedBranchID) {
            localStorage.setItem('selectedBranchID', String(this.selectedBranchID));
          } else {
            localStorage.removeItem('selectedBranchID');
          }
        }
      } catch (error) {
        console.error('Failed to fetch branches', error);
      }
    },
    setBranch(id: number) {
      this.selectedBranchID = id;
      localStorage.setItem('selectedBranchID', String(id));
    },
    logout() {
      this.token = null;
      this.user = null;
      this.selectedBranchID = null;
      this.branches = [];
      localStorage.removeItem('token');
      localStorage.removeItem('selectedBranchID');
    },
  },
});
