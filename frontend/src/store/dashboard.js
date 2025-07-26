import { defineStore } from 'pinia';
import api from '../services/api';

export const useDashboardStore = defineStore('dashboard', {
    state: () => ({
        stats: null,
        isLoading: false,
        error: null,
    }),
    actions: {
        async fetchStats() {
            this.isLoading = true;
            this.error = null;
            try {
                const response = await api.get('/admin/stats');
                this.stats = response.data;
            } catch (error) {
                console.error('Failed to fetch stats:', error);
                this.error = 'Could not load dashboard statistics.';
            } finally {
                this.isLoading = false;
            }
        },
    },
});