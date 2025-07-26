import { defineStore } from 'pinia';
import api from '../services/api';

export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: localStorage.getItem('token') || null,
        user: JSON.parse(localStorage.getItem('user')) || null,
    }),
    getters: {
        isAuthenticated: (state) => !!state.token,
        isAdmin: (state) => state.user?.role === 'admin',
        userName: (state) => state.user?.name || 'Guest',
        userRole: (state) => state.user?.role || 'user',
    },
    actions: {
        async login(email, password) {
            try {
                const response = await api.post('/login', { email, password });

                // === BAGIAN YANG DIPERBAIKI ===
                // 1. Ambil token dari response dan simpan di konstanta
                const token = response.data.token;

                // 2. Simpan token ke state Pinia
                this.token = token;

                // 3. Simpan token ke localStorage untuk persistensi
                localStorage.setItem('token', token);
                // ============================

                // Setelah login, ambil profil user
                await this.fetchUser();

                // Arahkan berdasarkan role
                if (this.isAdmin) {
                    this.router.push('/admin');
                } else {
                    this.router.push('/dashboard');
                }
            } catch (error) {
                console.error('Login failed:', error.response?.data || error.message);
                alert('Login failed! Please check your email and password.');
                // Panggil logout untuk membersihkan state jika login gagal
                this.logout();
            }
        },

        async register(userData) {
            try {
                // Kirim data registrasi ke backend
                // userData adalah objek { name, email, password }
                await api.post('/register', userData);

                // Setelah registrasi berhasil, tampilkan notifikasi
                // dan arahkan ke halaman login agar mereka bisa masuk.
                alert('Registration successful! Please log in to continue.');
                this.router.push('/login');

            } catch (error) {
                console.error('Registration failed:', error.response?.data || error.message);
                // Ambil pesan error dari backend jika ada, jika tidak tampilkan pesan generik
                const errorMessage = error.response?.data?.error || 'An error occurred during registration.';
                alert(`Registration failed: ${errorMessage}`);
                throw error; // Lemparkan error agar komponen bisa menangkapnya
            }
        },

        async fetchUser() {
            if (this.token) {
                try {
                    const response = await api.get('/profile');
                    this.user = response.data;
                    localStorage.setItem('user', JSON.stringify(this.user));
                } catch (error) {
                    console.error('Failed to fetch user profile:', error);
                    // Jika token tidak valid (misal: expired), logout
                    this.logout();
                }
            }
        },
        logout() {
            this.token = null;
            this.user = null;
            localStorage.removeItem('token');
            localStorage.removeItem('user');
            this.router.push('/login');
        },
    },
});